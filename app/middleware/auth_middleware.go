// Package middleware contains middleware for authentication and user context management
package middleware

import (
	"clinic-vet-api/app/modules/account/auth/token/service"
	repositoryimpl "clinic-vet-api/app/modules/account/user/infrastructure/repository"
	"clinic-vet-api/app/modules/core/domain/entity/user"
	"clinic-vet-api/app/modules/core/domain/valueobject"
	"clinic-vet-api/app/modules/core/repository"
	s "clinic-vet-api/app/modules/core/service"

	"clinic-vet-api/app/shared/mapper"
	"clinic-vet-api/app/shared/response"
	"clinic-vet-api/sqlc"
	"context"
	"strconv"

	autherror "clinic-vet-api/app/shared/error/auth"

	"github.com/gin-gonic/gin"
)

type UserContext struct {
	UserID      uint
	Email       string
	PhoneNumber string
	Role        string
	CustomerID  uint
	EmployeeID  uint
}

func UserToUserContext(user user.User) *UserContext {
	userCTX := &UserContext{
		UserID:      user.ID().Value(),
		Email:       user.Email().String(),
		PhoneNumber: user.PhoneNumber().String(),
		Role:        string(user.Role()),
	}

	if user.IsEmployee() {
		userCTX.EmployeeID = user.EmployeeID().Value()
	} else if user.IsCustomer() {
		userCTX.CustomerID = user.CustomerID().Value()
	}

	return userCTX
}

type AuthMiddleware struct {
	jwtService s.JWTService
	userRepo   repository.UserRepository
}

func NewAuthMiddleware(jwtSecret string, queries *sqlc.Queries) *AuthMiddleware {
	jwtService := service.NewJWTService(jwtSecret)
	userRepo := repositoryimpl.NewSqlcUserRepository(queries, mapper.NewSqlcFieldMapper())
	return &AuthMiddleware{
		jwtService: jwtService,
		userRepo:   userRepo,
	}
}

func (am *AuthMiddleware) Authenticate() gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			response.Unauthorized(c, autherror.UnauthorizedError("auth token required"))
			c.Abort()
			return
		}

		token, err := am.jwtService.ExtractToken(authHeader)
		if err != nil {
			response.Unauthorized(c, autherror.UnauthorizedError(err.Error()))
			c.Abort()
			return
		}

		c.Set("jwtToken", token)

		claim, err := am.jwtService.ValidateToken(token)
		if err != nil {
			response.Unauthorized(c, autherror.UnauthorizedError(err.Error()))
			c.Abort()
			return
		}

		idSTR := claim.UserID
		idUInt, err := strconv.ParseUint(idSTR, 10, 0)
		if err != nil {
			response.ServerError(c, err)
			c.Abort()
			return
		}

		user, err := am.userRepo.FindByID(c.Request.Context(), valueobject.NewUserID(uint(idUInt)))
		if err != nil {
			response.ApplicationError(c, err)
			return
		}

		c.Set("user", UserToUserContext(user))
		c.Set("userID", user.ID().Value())
		c.Set("userEmail", user.Email().String())
		c.Set("userRole", user.Role().String())

		c.Next()
	})
}

func (am *AuthMiddleware) OptionalAuth() gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.Next()
			return
		}

		token, err := am.jwtService.ExtractToken(authHeader)
		if err != nil {
			c.Next()
			return
		}

		c.Set("jwtToken", token)

		claims, err := am.jwtService.ValidateToken(token)
		if err != nil {
			c.Next()
			return
		}

		idUInt, err := strconv.ParseUint(claims.UserID, 10, 0)
		if err != nil {
			response.ServerError(c, err)
			c.Abort()
			return
		}

		user, err := am.userRepo.FindByID(context.Background(), valueobject.NewUserID(uint(idUInt)))
		if err != nil {
			response.ApplicationError(c, err)
			return
		}

		c.Set("user", UserToUserContext(user))
		c.Set("userID", user.ID().Value())
		c.Set("userEmail", user.Email().String())
		c.Set("userRole", user.Role().String())

		c.Next()
	})
}

// GetUserFromContext obtiene el contexto completo del usuario
func GetUserFromContext(c *gin.Context) (*UserContext, bool) {
	user, exists := c.Get("user")
	if !exists {
		return nil, false
	}
	return user.(*UserContext), true
}

// GetUserIDFromContext obtiene solo el ID del usuario
func GetUserIDFromContext(c *gin.Context) (valueobject.UserID, bool) {
	userID, exists := c.Get("userID")
	if !exists {
		return valueobject.UserID{}, false
	}

	idUint := userID.(uint)
	return valueobject.NewUserID(idUint), true
}

// GetUserEmailFromContext obtiene solo el email del usuario
func GetUserEmailFromContext(c *gin.Context) (string, bool) {
	email, exists := c.Get("userEmail")
	if !exists {
		return "", false
	}
	return email.(string), true
}

// GetUserRolesFromContext obtiene solo los roles del usuario
func GetUserRolesFromContext(c *gin.Context) (string, bool) {
	roles, exists := c.Get("userRole")
	if !exists {
		return "", false
	}
	return roles.(string), true
}

// HasRole verifica si el usuario tiene un rol específico
func HasRole(c *gin.Context, role string) bool {
	userRole, exists := GetUserRolesFromContext(c)
	if !exists {
		return false
	}

	if userRole == role {
		return true
	}

	return false
}

func (am *AuthMiddleware) RequireAnyRole(roles ...string) gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		userRole, exists := GetUserRolesFromContext(c)
		if !exists {
			response.Unauthorized(c, autherror.UnauthorizedError("user role not found in context"))
			c.Abort()
			return
		}

		for _, role := range roles {
			if userRole == role {
				c.Next()
				return
			}
		}

		response.Forbidden(c, autherror.ForbiddenError("insufficient permissions", userRole))
		c.Abort()
	})
}
