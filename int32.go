package sqlnull

import (
	"database/sql"
	"encoding/json"
	"reflect"
)

type NullInt32 sql.NullInt32

func (n NullInt32) MarshalJSON() ([]byte, error) {
	if !n.Valid {
		return json.Marshal(nil)
	}
	return json.Marshal(n.Int32)
}

func (n *NullInt32) UnmarshalJSON(data []byte) error {
	var target *int32
	if err := json.Unmarshal(data, &target); err != nil {
		return err
	}

	n.Valid = target != nil
	if n.Valid {
		n.Int32 = *target
	} else {
		n.Int32 = 0
	}
	return nil
}

func (n *NullInt32) Scan(src any) error {
	var sqln sql.NullInt32
	if err := sqln.Scan(src); err != nil {
		return err
	}
	n.Int32 = sqln.Int32
	n.Valid = reflect.TypeOf(src) != nil
	return nil
}
