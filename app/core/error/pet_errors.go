package domainerr

import "context"

func PetNotFoundErr(ctx context.Context, petID string) error {
	return EntityNotFoundError(ctx, "pet", petID, "Pet finding")
}
