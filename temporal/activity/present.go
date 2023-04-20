package activity

import (
	"context"
	"errors"

	"encore.app/temporal"
	"github.com/google/uuid"
	tb_types "github.com/tigerbeetledb/tigerbeetle-go/pkg/types"
)

func (a *Activities) Present(ctx context.Context, pendingID string) (temporal.Account, error) {
	_, err := a.Client.LookupTransfers([]tb_types.Uint128{uint128(pendingID)})
	if err != nil {
		return temporal.Account{}, errors.New("test")
	}
	// cancel pending transfer - money is set.
	_, err = a.Client.CreateTransfers([]tb_types.Transfer{
		{
			ID:              uint128(uuid.NewString()[:5]),
			PendingID:  uint128(pendingID),
			Flags: tb_types.TransferFlags{VoidPendingTransfer: true}.ToUint16(),
		},
	})
	if err != nil {
		return temporal.Account{}, errors.New("test")
	}
	return temporal.Account{}, nil
}


