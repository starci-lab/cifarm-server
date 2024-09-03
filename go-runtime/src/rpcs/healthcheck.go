package rpcs

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/heroiclabs/nakama-common/runtime"
)

type HealthcheckResponse struct {
	Status bool `json:"status"`
}

func HealthcheckRpc(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, payload string) (string, error) {
	response := &HealthcheckResponse{Status: true}

	out, err := json.Marshal(response)
	if err != nil {
		return "", runtime.NewError("Cannot marshal type", 13)
	}

	return string(out), nil
}
