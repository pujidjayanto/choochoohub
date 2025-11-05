package jsonb

import (
	"database/sql/driver"
	"fmt"

	"github.com/bytedance/sonic"
)

// JSONB stands for json binary or json better
// https://gorm.io/docs/data_types.html#Scanner-x2F-Valuer
type JSON map[string]interface{}

// Scan scan value into jsonb.JSON, implements sql.Scanner interface
func (j *JSON) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("type assertion to []byte failed")
	}

	return sonic.Unmarshal(b, &j)
}

// Value return json value, implement driver.Valuer interface
func (j JSON) Value() (driver.Value, error) {
	if len(j) == 0 {
		return nil, nil
	}

	return sonic.Marshal(j)
}

// Marshal marshals v to JSON using sonic
func Marshal(v interface{}) ([]byte, error) {
	return sonic.Marshal(v)
}

// Unmarshal unmarshals JSON data into v using sonic
func Unmarshal(data []byte, v interface{}) error {
	return sonic.Unmarshal(data, v)
}

// MarshalString marshals v to JSON string using sonic
func MarshalString(v interface{}) (string, error) {
	return sonic.MarshalString(v)
}

// UnmarshalString unmarshals JSON string into v using sonic
func UnmarshalString(data string, v interface{}) error {
	return sonic.UnmarshalString(data, v)
}

// MarshalIndent marshals v to indented JSON for pretty printing
func MarshalIndent(v interface{}, prefix, indent string) ([]byte, error) {
	return sonic.ConfigDefault.MarshalIndent(v, prefix, indent)
}

// Valid checks if data is valid JSON
func Valid(data []byte) bool {
	return sonic.Valid(data)
}
