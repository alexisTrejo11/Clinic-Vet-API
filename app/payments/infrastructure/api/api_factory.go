package paymentAPI

import (
	paymentCmd "github.com/alexisTrejo11/Clinic-Vet-API/app/payments/application/command"
	paymentQuery "github.com/alexisTrejo11/Clinic-Vet-API/app/payments/application/queries"
	paymentDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/payments/domain"
	paymentController "github.com/alexisTrejo11/Clinic-Vet-API/app/payments/infrastructure/api/controller"
	paymentRoutes "github.com/alexisTrejo11/Clinic-Vet-API/app/payments/infrastructure/api/routes"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// PaymentAPI is the main factory for payment API
type PaymentAPI struct {
	validator   *validator.Validate
	paymentRepo paymentDomain.PaymentRepository
	commandBuss *paymentCmd.CommandBus
	queryBus    *paymentQuery.QueryBus
}

// NewPaymentAPI creates a new payment API factory
func NewPaymentAPI(
	validator *validator.Validate,
	paymentRepo paymentDomain.PaymentRepository,
) *PaymentAPI {
	commandBuss := paymentCmd.NewPaymentCommandBus(paymentRepo)
	queryBuss := paymentQuery.NewQueryBus(paymentRepo)

	return &PaymentAPI{
		validator:   validator,
		paymentRepo: paymentRepo,
		commandBuss: &commandBuss,
		queryBus:    &queryBuss,
	}
}

// PaymentControllers holds all payment controllers
type PaymentControllers struct {
	Admin   *paymentController.AdminPaymentController
	Client  *paymentController.ClientPaymentController
	Query   *paymentController.PaymentQueryController
	Command *paymentController.PaymentController
}

// CreateControllers creates all payment controllers
func (f *PaymentAPI) CreateControllers() *PaymentControllers {

	queryController := paymentController.NewPaymentQueryController(
		f.validator,
		f.paymentRepo,
		*f.queryBus,
	)

	commandController := paymentController.NewPaymentController(
		f.validator,
		*f.commandBuss,
	)

	clientController := paymentController.NewClientPaymentController(
		f.validator,
		f.paymentRepo,
		queryController,
	)

	adminController := paymentController.NewAdminPaymentController(
		f.validator,
		queryController,
		commandController,
	)

	return &PaymentControllers{
		Admin:   adminController,
		Client:  clientController,
		Query:   queryController,
		Command: commandController,
	}
}

// RegisterAllRoutes registers all payment routes
func (f *PaymentAPI) RegisterAllRoutes(router *gin.Engine) {
	controllers := f.CreateControllers()

	paymentRoutes.RegisterAdminPaymentRoutes(router, controllers.Admin)
	paymentRoutes.RegisterClientPaymentRoutes(router, controllers.Client)
}

// Validate validates the factory configuration
func (f *PaymentAPI) Validate() error {
	if f.validator == nil {
		return paymentDomain.NewPaymentError("INVALID_API_CONFIG", "validator cannot be nil", 0, "")
	}

	if f.paymentRepo == nil {
		return paymentDomain.NewPaymentError("INVALID_API_CONFIG", "payment repository cannot be nil", 0, "")
	}

	if f.commandBuss == nil {
		return paymentDomain.NewPaymentError("INVALID_API_CONFIG", "command bus cannot be nil", 0, "")

	}

	if f.queryBus == nil {
		return paymentDomain.NewPaymentError("INVALID_API_CONFIG", "query bus cannot be nil", 0, "")
	}

	return nil
}
