package auth

import (
	"context"
	"database/sql"

	"github.com/heroiclabs/nakama-common/runtime"
)

func Initialize(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	initializer runtime.Initializer,
) error {
	err := initializer.RegisterBeforeAuthenticateCustom(BeforeAuthenticate)
	if err != nil {
		return err
	}

	err = initializer.RegisterAfterAuthenticateCustom(AfterAuthenticate)
	if err != nil {
		return err
	}
	return nil
}
