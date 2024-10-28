package sqlnull

import (
	"database/sql"
	"encoding/json"
	"reflect"
)

type NullInt16 sql.NullInt16

func (n NullInt16) MarshalJSON() ([]byte, error) {
	if !n.Valid {
		return json.Marshal(nil)
	}
	return json.Marshal(n.Int16)
}

func (n *NullInt16) UnmarshalJSON(data []byte) error {
	var target *int16
	if err := json.Unmarshal(data, &target); err != nil {
		return err
	}

	n.Valid = target != nil
	if n.Valid {
		n.Int16 = *target
	} else {
		n.Int16 = 0
	}
	return nil
}

func (n *NullInt16) Scan(src any) error {
	var sqln sql.NullInt16
	if err := sqln.Scan(src); err != nil {
		return err
	}
	n.Int16 = sqln.Int16
	n.Valid = reflect.TypeOf(src) != nil
	return nil
}
