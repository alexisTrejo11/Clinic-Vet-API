package domainerr

import (
	"context"
	"strconv"
)

func customerNotFoundError(ctx context.Context, id int) error {
	return EntityNotFoundError(ctx, "customer", strconv.Itoa(int(id)), "Customer Find")
}

func PhoneConflictError(ctx context.Context) error {
	return ConflictError(ctx, "phone number", "Phone Number Already Taken", "unique phone validation")
}
