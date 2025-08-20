package paymentAPI

import (
	paymentDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/payments/domain"
	paymentRoutes "github.com/alexisTrejo11/Clinic-Vet-API/app/payments/infrastructure/api/routes"
	paymentRepo "github.com/alexisTrejo11/Clinic-Vet-API/app/payments/infrastructure/persistence"
	"github.com/alexisTrejo11/Clinic-Vet-API/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type PaymentAPIConfig struct {
	EnableAdminRoutes   bool
	EnableClientRoutes  bool
	EnableQueryRoutes   bool
	EnableCommandRoutes bool
	RoutePrefix         string
	EnableValidation    bool
	EnableMetrics       bool
	EnableAuditing      bool
}

// DefaultPaymentAPIConfig returns default API configuration
func DefaultPaymentAPIConfig() *PaymentAPIConfig {
	return &PaymentAPIConfig{
		EnableAdminRoutes:   true,
		EnableClientRoutes:  true,
		EnableQueryRoutes:   false,
		EnableCommandRoutes: false,
		RoutePrefix:         "api/v2",
		EnableValidation:    true,
		EnableMetrics:       false,
		EnableAuditing:      false,
	}
}

// PaymentAPIBuilder provides a fluent interface for building payment API
type PaymentAPIBuilder struct {
	config      *PaymentAPIConfig
	validator   *validator.Validate
	paymentRepo paymentDomain.PaymentRepository
	router      *gin.Engine
}

// NewPaymentAPIBuilder creates a new payment API builder
func NewPaymentAPIBuilder() *PaymentAPIBuilder {
	return &PaymentAPIBuilder{
		config: DefaultPaymentAPIConfig(),
	}
}

// WithConfig sets the API configuration
func (b *PaymentAPIBuilder) WithConfig(config *PaymentAPIConfig) *PaymentAPIBuilder {
	if config != nil {
		b.config = config
	}
	return b
}

// WithValidator sets the validator
func (b *PaymentAPIBuilder) WithValidator(validator *validator.Validate) *PaymentAPIBuilder {
	b.validator = validator
	return b
}

// WithPaymentRepository sets the payment repository
func (b *PaymentAPIBuilder) WithPaymentRepository(queries *sqlc.Queries) *PaymentAPIBuilder {
	b.paymentRepo = paymentRepo.NewSQLCPaymentRepository(queries)
	return b
}

// WithRouter sets the gin router
func (b *PaymentAPIBuilder) WithRouter(router *gin.Engine) *PaymentAPIBuilder {
	b.router = router
	return b
}

// Build builds and configures the payment API
func (b *PaymentAPIBuilder) Build() (*PaymentAPI, error) {
	// Validate required dependencies
	if b.validator == nil {
		return nil, paymentDomain.NewPaymentError("MISSING_DEPENDENCY", "validator is required", 0, "")
	}

	if b.paymentRepo == nil {
		return nil, paymentDomain.NewPaymentError("MISSING_DEPENDENCY", "payment repository is required", 0, "")
	}

	// Create factory
	factory := NewPaymentAPI(b.validator, b.paymentRepo)

	// Validate factory
	if err := factory.Validate(); err != nil {
		return nil, err
	}

	// Register routes if router is provided
	if b.router != nil {
		b.registerRoutes(factory)
	}

	return factory, nil
}

// registerRoutes registers the appropriate routes based on configuration
func (b *PaymentAPIBuilder) registerRoutes(factory *PaymentAPI) {
	controllers := factory.CreateControllers()

	if b.config.EnableAdminRoutes {
		paymentRoutes.RegisterAdminPaymentRoutes(b.router, controllers.Admin)
	}

	if b.config.EnableClientRoutes {
		paymentRoutes.RegisterClientPaymentRoutes(b.router, controllers.Client)
	}

	if b.config.EnableQueryRoutes {
		paymentRoutes.RegisterPaymentQueryRoutes(b.router, controllers.Query)
	}

	if b.config.EnableCommandRoutes {
		paymentRoutes.RegisterPaymentCommandRoutes(b.router, controllers.Command)
	}
}

func SetupPaymentAPI(router *gin.Engine, validator *validator.Validate, queries *sqlc.Queries) (*PaymentAPI, error) {

	return NewPaymentAPIBuilder().
		WithRouter(router).
		WithValidator(validator).
		WithPaymentRepository(queries).
		WithConfig(&PaymentAPIConfig{
			EnableAdminRoutes:   true,
			EnableClientRoutes:  true,
			EnableQueryRoutes:   true,
			EnableCommandRoutes: true,
			EnableValidation:    true,
		}).
		Build()
}
