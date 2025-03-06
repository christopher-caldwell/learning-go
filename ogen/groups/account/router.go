package account

import (
	"context"

	petstore "oapigen/petstore"
)

type AccountService struct {
	accounts map[int64]petstore.Account
}

var _ petstore.AccountHandler = (*AccountService)(nil)

func (a *AccountService) GetAccountById(ctx context.Context, params petstore.GetAccountByIdParams) (petstore.GetAccountByIdRes, error) {
	account, ok := a.accounts[params.AccountId]
	if !ok {
		return &petstore.GetAccountByIdNotFound{}, nil
	}
	return &account, nil
}

var Service = AccountService{
	accounts: map[int64]petstore.Account{
		1: {
			ID:   1,
			Name: "Jim",
		},
	},
}
