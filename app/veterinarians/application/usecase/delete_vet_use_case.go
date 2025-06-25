package vetUsecase

import (
	"context"
)

type DeleteVetUseCase struct {
}

func (uc *DeleteVetUseCase) Execute(ctx context.Context, vetId uint) error {
	return nil
}
