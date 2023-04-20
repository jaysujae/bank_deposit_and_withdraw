package activity

import (
	"context"
	"errors"
	"strconv"

	"encore.app/temporal"
	tb_types "github.com/tigerbeetledb/tigerbeetle-go/pkg/types"
)

func (a *Activities) GetAccount(ctx context.Context, accountNumber int) (temporal.Account, error) {
	accounts, err := a.Client.LookupAccounts([]tb_types.Uint128{uint128(strconv.Itoa(accountNumber))})
	if err != nil {
		return temporal.Account{}, err
	}
	if len(accounts) == 0{
		return temporal.Account{}, errors.New("account doesn't exist")
	}
	
    account := accounts[0]
	return temporal.Account{
		ID: accountNumber,
		Available: int(account.CreditsPosted - account.DebitsPosted - account.DebitsPending + account.CreditsPending),
		Reserved: int(account.DebitsPending - account.CreditsPending),
	}, nil
}

func uint128(value string) tb_types.Uint128 {
	x, err := tb_types.HexStringToUint128(value)
	if err != nil {
		panic(err)
	}
	return x
}