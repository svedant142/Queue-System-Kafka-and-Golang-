package entity

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type URLS []string

func (m URLS) Value() (driver.Value, error) {
	if len(m) == 0 {
		return nil, nil
	}
	
	j, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	return driver.Value([]byte(j)), nil
}

func (m *URLS) Scan(src interface{}) error {
	var source []byte
	var _m []string

	switch src.(type) {
	case []uint8:
		source = []byte(src.([]uint8))
	case nil:
		return nil
	default:
		return errors.New("incompatible type for URLS")
	}
	err := json.Unmarshal(source, &_m)
	if err != nil {
		return err
	}
	*m = URLS(_m)
	return nil
}
