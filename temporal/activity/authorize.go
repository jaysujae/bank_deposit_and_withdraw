package activity

import (
	"context"
	"errors"
	"strconv"

	"github.com/google/uuid"

	"encore.app/temporal"
	tb_types "github.com/tigerbeetledb/tigerbeetle-go/pkg/types"
)

const bank = "1"

func (a *Activities) CreateAuthorizeTransfer(ctx context.Context, p temporal.AuthorizeParams) (string, error) {
	transferID := uuid.NewString()[:5]
	_, err := a.Client.CreateTransfers([]tb_types.Transfer{
		{
			ID:              uint128(transferID),
			DebitAccountID:  uint128(strconv.Itoa(p.ID)),
			CreditAccountID: uint128(bank),
			Ledger:          1,
			Code:            1,
			Amount:          uint64(p.Amount),
			Flags: tb_types.TransferFlags{Pending: true}.ToUint16(),
		},
	})	
	if err != nil {
		return "", errors.New("test")
	}
	return transferID, nil
}

func (a *Activities) CancelAuthorize(ctx context.Context, pendingID string) (error) {
	// when authorize is cancelled, bank gets reserved credit and customer loses credit
	_, err := a.Client.CreateTransfers([]tb_types.Transfer{
		{
			ID:              uint128(uuid.NewString()[:5]),
			PendingID:  uint128(pendingID),
			Flags: tb_types.TransferFlags{PostPendingTransfer: true}.ToUint16(),
		},
	})
	if err != nil {
		return errors.New("test")
	}
	return nil
}

