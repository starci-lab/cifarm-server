package collections_tools

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/heroiclabs/nakama-common/runtime"
)

type WriteManyParams struct {
	Tools []Tool `json:"tools"`
}

func WriteMany(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	params WriteManyParams,
) error {
	var writes []*runtime.StorageWrite
	for _, tool := range params.Tools {
		key := tool.Key
		tool.Key = ""
		value, err := json.Marshal(tool)
		if err != nil {
			continue
		}

		write := &runtime.StorageWrite{
			Collection:      COLLECTION_NAME,
			Key:             key,
			Value:           string(value),
			PermissionRead:  2,
			PermissionWrite: 0,
		}
		writes = append(writes, write)
	}

	_, err := nk.StorageWrite(ctx, writes)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	return nil
}
