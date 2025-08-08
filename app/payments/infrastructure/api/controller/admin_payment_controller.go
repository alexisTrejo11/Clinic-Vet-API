package paymentController

import "github.com/gin-gonic/gin"

type AdminPaymentController struct{}

func NewAdminPaymentController() *AdminPaymentController {
	return &AdminPaymentController{}
}

func (c *AdminPaymentController) SearchPayments(ctx *gin.Context) {
	// Implementation for searching payments
}

func (c *AdminPaymentController) CreatePayment(ctx *gin.Context) {
	// Implementation for creating a payment
}

func (c *AdminPaymentController) GetPayment(ctx *gin.Context) {
	// Implementation for retrieving a payment
}

func (c *AdminPaymentController) UpdatePayment(ctx *gin.Context) {
	// Implementation for updating a payment
}

func (c *AdminPaymentController) DeletePayment(ctx *gin.Context) {
	// Implementation for deleting a payment
}
