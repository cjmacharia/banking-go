package service

import (
	"github.com/cjmash/banking/domain"
)

type AccountService interface {
	CreateAccount(account domain.Account) error
}

type AccountS struct {
	acc domain.AccountRepository
}

func (a AccountS) CreateAccount(account domain.Account) error {
	return a.acc.CreateAccount(account)
}

func NewAccountService(repository domain.AccountRepository) AccountS {
	return AccountS{repository}
}
