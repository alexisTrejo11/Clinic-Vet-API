package authUsecase

import authDto "github.com/alexisTrejo11/Clinic-Vet-API/app/auth/application/dtos"

type SignUpUseCase struct {
}

func (uc *SignUpUseCase) Execute(requestSignup authDto.RequestSignup) error {
	return nil
}
