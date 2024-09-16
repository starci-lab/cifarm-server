package collections_common

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"reflect"

	"github.com/heroiclabs/nakama-common/api"
	"github.com/heroiclabs/nakama-common/runtime"
)

func ToValue[TValue any](
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	object *api.StorageObject,
) (*TValue, error) {
	if object == nil {
		errMsg := "object cannot be nil"
		logger.Error(errMsg)
		return nil, errors.New(errMsg)
	}
	var data TValue
	err := json.Unmarshal([]byte(object.Value), &data)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	value := reflect.ValueOf(&data).Elem()
	keyField := value.FieldByName("Key")

	if keyField.IsValid() && keyField.CanSet() {
		if keyField.Kind() == reflect.String {
			keyField.SetString(object.Key)
		} else {
			errMsg := "key field is not of type string"
			logger.Error(errMsg)
			return nil, errors.New(errMsg)
		}
	} else {
		errMsg := "key field is not found or cannot be set"
		logger.Error(errMsg)
		return nil, errors.New(errMsg)
	}

	return &data, nil
}

func ToValues[TValue any](
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	objects *api.StorageObjects,
) ([]*TValue, error) {
	var values []*TValue
	for _, object := range objects.Objects {
		var data TValue
		err := json.Unmarshal([]byte(object.Value), &data)
		if err != nil {
			logger.Error(err.Error())
			return nil, err
		}

		value := reflect.ValueOf(&data).Elem()
		keyField := value.FieldByName("Key")

		if keyField.IsValid() && keyField.CanSet() {
			if keyField.Kind() == reflect.String {
				keyField.SetString(object.Key)
			} else {
				errMsg := "key field is not of type string"
				logger.Error(errMsg)
				return nil, errors.New(errMsg)
			}
		} else {
			errMsg := "key field is not found or cannot be set"
			logger.Error(errMsg)
			return nil, errors.New(errMsg)
		}

		values = append(values, &data)
	}

	return values, nil
}

func ToValues2[TValue any](
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	objects []*api.StorageObject,
) ([]*TValue, error) {
	var values []*TValue
	for _, object := range objects {
		var data TValue
		err := json.Unmarshal([]byte(object.Value), &data)
		if err != nil {
			logger.Error(err.Error())
			return nil, err
		}

		value := reflect.ValueOf(&data).Elem()
		keyField := value.FieldByName("Key")

		if keyField.IsValid() && keyField.CanSet() {
			if keyField.Kind() == reflect.String {
				keyField.SetString(object.Key)
			} else {
				errMsg := "key field is not of type string"
				logger.Error(errMsg)
				return nil, errors.New(errMsg)
			}
		} else {
			errMsg := "key field is not found or cannot be set"
			logger.Error(errMsg)
			return nil, errors.New(errMsg)
		}

		values = append(values, &data)
	}

	return values, nil
}
