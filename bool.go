package sqlnull

import (
	"database/sql"
	"encoding/json"
	"reflect"
)

type NullBool sql.NullBool

func (n NullBool) MarshalJSON() ([]byte, error) {
	if !n.Valid {
		return json.Marshal(nil)
	}
	return json.Marshal(n.Bool)
}

func (n *NullBool) UnmarshalJSON(data []byte) error {
	var target *bool
	if err := json.Unmarshal(data, &target); err != nil {
		return err
	}

	n.Valid = target != nil
	if n.Valid {
		n.Bool = *target
	} else {
		n.Bool = false
	}
	return nil
}

func (n *NullBool) Scan(src any) error {
	var sqln sql.NullBool
	if err := sqln.Scan(src); err != nil {
		return err
	}
	n.Bool = sqln.Bool
	n.Valid = reflect.TypeOf(src) != nil
	return nil
}
