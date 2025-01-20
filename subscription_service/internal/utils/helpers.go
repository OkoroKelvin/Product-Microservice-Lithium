package utils

import "encoding/json"

func UnmarshalToType[T any](data any, valueType *T) error {
	var (
		err       error
		dataBytes []byte
	)

	switch data := data.(type) {
	case string:
		dataBytes = []byte(data)
	case []byte:
		dataBytes = data
	default:
		dataBytes, err = json.Marshal(data)
		if err != nil {
			return err
		}
	}
	err = json.Unmarshal(dataBytes, valueType)
	if err != nil {
		return err
	}
	return nil
}
