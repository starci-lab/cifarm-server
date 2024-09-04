package rpcs

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/heroiclabs/nakama-common/runtime"
)

type HealthcheckResponse struct {
	Status string `json:"status"`
}

func HealthcheckRpc(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, payload string) (string, error) {
	response := &HealthcheckResponse{Status: "ok"}

	out, err := json.Marshal(response)
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	return string(out), nil
}
