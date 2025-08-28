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

type PaymentAPIConfig struct {
	Router      *gin.Engine
	Validator   *validator.Validate
	PaymentRepo paymentDomain.PaymentRepository
}

type PaymentAPIComponents struct {
	Repository  paymentDomain.PaymentRepository
	CommandBus  *paymentCmd.CommandBus
	QueryBus    *paymentQuery.QueryBus
	Controllers *PaymentControllers
}

type PaymentControllers struct {
	Admin   *paymentController.AdminPaymentController
	Client  *paymentController.ClientPaymentController
	Query   *paymentController.PaymentQueryController
	Command *paymentController.PaymentController
}

type PaymentAPIBuilder struct {
	config     *PaymentAPIConfig
	components *PaymentAPIComponents
	isBuilt    bool
}

func NewPaymentAPIBuilder(config *PaymentAPIConfig) *PaymentAPIBuilder {
	return &PaymentAPIBuilder{
		config:  config,
		isBuilt: false,
	}
}

func (f *PaymentAPIBuilder) Build() error {
	if f.isBuilt {
		return nil
	}

	if err := f.validateConfig(); err != nil {
		return err
	}

	commandBus := paymentCmd.NewPaymentCommandBus(f.config.PaymentRepo)
	queryBus := paymentQuery.NewQueryBus(f.config.PaymentRepo)

	controllers := f.createControllers(&commandBus, &queryBus)

	f.registerRoutes(controllers)

	f.components = &PaymentAPIComponents{
		Repository:  f.config.PaymentRepo,
		CommandBus:  &commandBus,
		QueryBus:    &queryBus,
		Controllers: controllers,
	}

	f.isBuilt = true
	return nil
}

func (f *PaymentAPIBuilder) createControllers(commandBus *paymentCmd.CommandBus, queryBus *paymentQuery.QueryBus) *PaymentControllers {
	queryController := paymentController.NewPaymentQueryController(
		f.config.Validator,
		f.config.PaymentRepo,
		*queryBus,
	)

	commandController := paymentController.NewPaymentController(
		f.config.Validator,
		*commandBus,
	)

	clientController := paymentController.NewClientPaymentController(
		f.config.Validator,
		f.config.PaymentRepo,
		queryController,
	)

	adminController := paymentController.NewAdminPaymentController(
		f.config.Validator,
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

func (f *PaymentAPIBuilder) registerRoutes(controllers *PaymentControllers) {
	paymentRoutes.RegisterAdminPaymentRoutes(f.config.Router, controllers.Admin)
	paymentRoutes.RegisterClientPaymentRoutes(f.config.Router, controllers.Client)
}

func (f *PaymentAPIBuilder) validateConfig() error {
	if f.config == nil {
		return paymentDomain.NewPaymentError("INVALID_API_CONFIG", "configuration cannot be nil", 0, "")
	}
	if f.config.Router == nil {
		return paymentDomain.NewPaymentError("INVALID_API_CONFIG", "router cannot be nil", 0, "")
	}
	if f.config.Validator == nil {
		return paymentDomain.NewPaymentError("INVALID_API_CONFIG", "validator cannot be nil", 0, "")
	}
	if f.config.PaymentRepo == nil {
		return paymentDomain.NewPaymentError("INVALID_API_CONFIG", "payment repository cannot be nil", 0, "")
	}
	return nil
}

func (f *PaymentAPIBuilder) GetComponents() (*PaymentAPIComponents, error) {
	if !f.isBuilt {
		if err := f.Build(); err != nil {
			return nil, err
		}
	}
	return f.components, nil
}

func (f *PaymentAPIBuilder) GetControllers() (*PaymentControllers, error) {
	components, err := f.GetComponents()
	if err != nil {
		return nil, err
	}
	return components.Controllers, nil
}
