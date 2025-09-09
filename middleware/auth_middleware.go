package middleware

import (
	"context"
	"strconv"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/entity/user"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
	repository "github.com/alexisTrejo11/Clinic-Vet-API/app/core/repositories"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/auth/application/jwt"
	autherror "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/error/auth"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/response"
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
	return &UserContext{
		UserID:      user.ID().Value(),
		Email:       user.Email().String(),
		PhoneNumber: user.PhoneNumber().String(),
		Role:        string(user.Role()),
		CustomerID:  0,
		EmployeeID:  0,
	}
}

type AuthMiddleware struct {
	jwtService jwt.JWTService
	userRepo   repository.UserRepository
}

func NewAuthMiddleware(jwtService jwt.JWTService, userRepo repository.UserRepository) *AuthMiddleware {
	return &AuthMiddleware{
		jwtService: jwtService,
		userRepo:   userRepo,
	}
}

func (am *AuthMiddleware) Authenticate() gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		authHeader := c.GetHeader("Auhtorization")
		if authHeader == "" {
			response.Unauthorized(c, autherror.UnauthorizedError("auth token required"))
			c.Abort()
			return
		}

		idSTR, err := am.jwtService.ExtractToken(authHeader)
		if err != nil {
			response.Unauthorized(c, autherror.UnauthorizedError(err.Error()))
		}

		ctx := context.Background()

		idINT, err := strconv.Atoi(idSTR)
		if err != nil {
			response.ServerError(c, err)
			c.Abort()
			return
		}

		userID, err := valueobject.NewUserID(idINT)
		if err != nil {
			response.ServerError(c, err)
			c.Abort()
			return
		}

		user, err := am.userRepo.GetByID(ctx, userID)
		if err != nil {
			response.ApplicationError(c, err)
		}

		c.Set("user", UserToUserContext(user))
		c.Set("userID", user.ID)
		c.Set("userEmail", user.Email)
		c.Set("userRoles", user.Role())

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

		claims, err := am.jwtService.ValidateToken(token)
		if err != nil {
			c.Next()
			return
		}

		idINT, err := strconv.Atoi(claims.ID)
		if err != nil {
			response.ServerError(c, err)
			c.Abort()
			return
		}

		userID, err := valueobject.NewUserID(idINT)
		if err != nil {
			response.ServerError(c, err)
			c.Abort()
			return
		}

		user, err := am.userRepo.GetByID(context.Background(), userID)
		if err != nil {
			response.ApplicationError(c, err)
		}

		c.Set("user", UserToUserContext(user))
		c.Set("userID", user.ID)
		c.Set("userEmail", user.Email)
		c.Set("userRoles", user.Role())

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
func GetUserIDFromContext(c *gin.Context) (int, bool) {
	userID, exists := c.Get("userID")
	if !exists {
		return 0, false
	}
	return userID.(int), true
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
	roles, exists := c.Get("userRoles")
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
