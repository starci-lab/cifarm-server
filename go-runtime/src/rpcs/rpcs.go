package rpcs

import (
	rpcs_assets "cifarm-server/src/rpcs/assets"
	rpcs_auth "cifarm-server/src/rpcs/auth"
	rpcs_claim "cifarm-server/src/rpcs/claim"
	rpcs_community "cifarm-server/src/rpcs/community"
	rpcs_farming "cifarm-server/src/rpcs/farming"
	rpcs_miscellaneous "cifarm-server/src/rpcs/miscellaneous"
	rpcs_nfts "cifarm-server/src/rpcs/nfts"
	rpcs_placement "cifarm-server/src/rpcs/placement"
	rpcs_profile "cifarm-server/src/rpcs/profile"
	rpcs_shop "cifarm-server/src/rpcs/shop"
	rpcs_tests "cifarm-server/src/rpcs/tests"
	"context"
	"database/sql"

	"github.com/heroiclabs/nakama-common/runtime"
)

func Initialize(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, initializer runtime.Initializer) error {
	err := initializer.RegisterRpc("go_healthcheck", HealthcheckRpc)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	err = rpcs_auth.Initialize(ctx, logger, db, nk, initializer)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	err = rpcs_farming.Initialize(ctx, logger, db, nk, initializer)
	if err != nil {
		return err
	}

	err = rpcs_shop.Initialize(ctx, logger, db, nk, initializer)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	err = rpcs_claim.Initialize(ctx, logger, db, nk, initializer)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	err = rpcs_nfts.Initialize(ctx, logger, db, nk, initializer)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	err = rpcs_community.Initialize(ctx, logger, db, nk, initializer)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	err = rpcs_tests.Initialize(ctx, logger, db, nk, initializer)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	err = rpcs_assets.Initialize(ctx, logger, db, nk, initializer)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	err = rpcs_placement.Initialize(ctx, logger, db, nk, initializer)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	err = rpcs_miscellaneous.Initialize(ctx, logger, db, nk, initializer)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	err = rpcs_profile.Initialize(ctx, logger, db, nk, initializer)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	return nil
}
