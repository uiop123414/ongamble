package utils

import (
	"chatgpt-article/internal/models"
	"encoding/json"

	"github.com/xeipuuv/gojsonschema"
)

const RequestSizeLimit = 2 * 1024 * 1024 // 2 kBs

func ReadJSON(payload []byte, loader gojsonschema.JSONLoader, data interface{}) error {
	err := validateJSON(loader, payload)
	if err != nil {
		return err
	}

	err = json.Unmarshal(payload, data)
	if err != nil {
		return err
	}

	return nil
}

func validateJSON(loader gojsonschema.JSONLoader, data []byte) error {
	payloadLoader := gojsonschema.NewBytesLoader(data)

	result, err := gojsonschema.Validate(loader, payloadLoader)
	if err != nil {
		return err
	}

	if !result.Valid() {
		return models.ErrJSONNotValid
	}

	return nil
}