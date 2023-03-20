package activity

import (
	"context"
	"errors"

	"encore.app/temporal"
	"github.com/google/uuid"
	tb_types "github.com/tigerbeetledb/tigerbeetle-go/pkg/types"
)

func (a *Activities) Present(ctx context.Context, pendingID string) (temporal.Account, error) {
	_, err := a.Client.CreateTransfers([]tb_types.Transfer{
		{
			ID:              uint128(uuid.NewString()[:5]),
			PendingID:  uint128(pendingID),
			Flags: tb_types.TransferFlags{PostPendingTransfer: true}.ToUint16(),
		},
	})
	if err != nil {
		return temporal.Account{}, errors.New("test")
	}
	return temporal.Account{
	}, nil
}


