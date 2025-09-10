package paymentAPI

import (
	domainerr "github.com/alexisTrejo11/Clinic-Vet-API/app/core/error"
	repository "github.com/alexisTrejo11/Clinic-Vet-API/app/core/repositories"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/payments/infrastructure/api/controller"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/payments/infrastructure/api/routes"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/payments/infrastructure/bus"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type PaymentAPIConfig struct {
	Router      *gin.Engine
	Validator   *validator.Validate
	PaymentRepo repository.PaymentRepository
}

type PaymentAPIComponents struct {
	Repository  repository.PaymentRepository
	CommandBus  *bus.PaymentCommandBus
	QueryBus    *bus.PaymentQueryBus
	Controllers *PaymentControllers
}

type PaymentControllers struct {
	Admin   *controller.AdminPaymentController
	Client  *controller.ClientPaymentController
	Query   *controller.PaymentQueryController
	Command *controller.PaymentController
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

	commandBus := bus.NewPaymentCommandBus(f.config.PaymentRepo)
	queryBus := bus.NewQueryBus(f.config.PaymentRepo)

	controllers := f.createControllers(commandBus, queryBus)

	f.registerRoutes(controllers)

	f.components = &PaymentAPIComponents{
		Repository:  f.config.PaymentRepo,
		CommandBus:  commandBus,
		QueryBus:    queryBus,
		Controllers: controllers,
	}

	f.isBuilt = true
	return nil
}

func (f *PaymentAPIBuilder) createControllers(commandBus *bus.PaymentCommandBus, queryBus *bus.PaymentQueryBus) *PaymentControllers {
	queryController := controller.NewPaymentQueryController(
		f.config.Validator,
		queryBus,
	)

	commandController := controller.NewPaymentController(
		f.config.Validator,
		commandBus,
	)

	clientController := controller.NewClientPaymentController(
		f.config.Validator,
		f.config.PaymentRepo,
		queryController,
	)

	adminController := controller.NewAdminPaymentController(
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
	routes.RegisterAdminPaymentRoutes(f.config.Router, controllers.Admin)
	routes.RegisterClientPaymentRoutes(f.config.Router, controllers.Client)
}

func (f *PaymentAPIBuilder) validateConfig() error {
	if f.config == nil {
		return domainerr.NewPaymentError("INVALID_API_CONFIG", "configuration cannot be nil", 0, "")
	}
	if f.config.Router == nil {
		return domainerr.NewPaymentError("INVALID_API_CONFIG", "router cannot be nil", 0, "")
	}
	if f.config.Validator == nil {
		return domainerr.NewPaymentError("INVALID_API_CONFIG", "validator cannot be nil", 0, "")
	}
	if f.config.PaymentRepo == nil {
		return domainerr.NewPaymentError("INVALID_API_CONFIG", "payment repository cannot be nil", 0, "")
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
