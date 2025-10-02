// Package api provides the API builder for the payments module.
package api

import (
	"clinic-vet-api/app/middleware"
	domainerr "clinic-vet-api/app/modules/core/error"
	"clinic-vet-api/app/modules/core/repository"
	"clinic-vet-api/app/modules/payment/infrastructure/bus"
	repositoryimpl "clinic-vet-api/app/modules/payment/infrastructure/repository"
	"clinic-vet-api/app/modules/payment/presentation/controller"
	"clinic-vet-api/app/modules/payment/presentation/routes"
	"clinic-vet-api/sqlc"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type PaymentAPIConfig struct {
	Router         *gin.RouterGroup
	Validator      *validator.Validate
	AuthMiddleware *middleware.AuthMiddleware
	Queries        *sqlc.Queries
}

type PaymentAPIComponents struct {
	Repository  repository.PaymentRepository
	Bus         *bus.PaymentBus
	Controllers *PaymentControllers
}

type PaymentControllers struct {
	Admin  *controller.AdminPaymentController
	Client *controller.ClientPaymentController
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

	repository := repositoryimpl.NewSqlcPaymentRepository(f.config.Queries)

	paymentBus := bus.NewPaymentBus(repository)
	controllers := f.createControllers(*paymentBus)

	f.registerRoutes(controllers)

	f.components = &PaymentAPIComponents{
		Repository:  repository,
		Bus:         paymentBus,
		Controllers: controllers,
	}

	f.isBuilt = true
	return nil
}

func (f *PaymentAPIBuilder) createControllers(paymentBus bus.PaymentBus) *PaymentControllers {
	controllerOperations := controller.NewPaymentControllerOperations(
		f.config.Validator,
		&paymentBus,
	)

	clientController := controller.NewClientPaymentController(
		f.config.Validator,
		controllerOperations,
	)

	adminController := controller.NewAdminPaymentController(
		f.config.Validator,
		controllerOperations,
	)

	return &PaymentControllers{
		Admin:  adminController,
		Client: clientController,
	}
}

func (f *PaymentAPIBuilder) registerRoutes(controllers *PaymentControllers) {
	paymentRoutes := routes.NewPaymentRoutes(controllers.Admin, controllers.Client)
	paymentRoutes.RegisterAdminPaymentRoutes(f.config.Router, f.config.AuthMiddleware)
	paymentRoutes.RegisterClientRoutes(f.config.Router, f.config.AuthMiddleware)
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
	if f.config.AuthMiddleware == nil {
		return domainerr.NewPaymentError("INVALID_API_CONFIG", "auth middleware cannot be nil", 0, "")
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
