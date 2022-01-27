package service

import (
	"github.com/SaravananPitchaimuthu/pub-api/domain"
	"github.com/SaravananPitchaimuthu/pub-api/dto"
	"github.com/SaravananPitchaimuthu/pub-api/errs"
)

type AccountService interface {
	NewAccount(request dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError)
}

type defaultAccountService struct {
	repo domain.AccountRepository
}

func NewAccountService(repo domain.AccountRepositoryDb) defaultAccountService {
	return defaultAccountService{repo: repo}
}

func (s defaultAccountService) NewAccount(req dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}
	account := domain.NewAccount(req.CustomerId, req.AccountType, req.Amount)
	if newAccount, err := s.repo.Save(account); err != nil {
		return nil, err
	} else {
		return newAccount.ToNewAccountResponseDto(), nil
	}

}
