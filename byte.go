package sqlnull

import (
	"database/sql"
	"encoding/json"
	"reflect"
)

type NullByte sql.NullByte

func (n NullByte) MarshalJSON() ([]byte, error) {
	if !n.Valid {
		return json.Marshal(nil)
	}
	return json.Marshal(n.Byte)
}

func (n *NullByte) UnmarshalJSON(data []byte) error {
	var target *byte
	if err := json.Unmarshal(data, &target); err != nil {
		return err
	}

	n.Valid = target != nil
	if n.Valid {
		n.Byte = *target
	} else {
		n.Byte = 0x00
	}
	return nil
}

func (n *NullByte) Scan(src any) error {
	var sqln sql.NullByte
	if err := sqln.Scan(src); err != nil {
		return err
	}
	n.Byte = sqln.Byte
	n.Valid = reflect.TypeOf(src) != nil
	return nil
}
