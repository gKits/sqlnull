package sqlnull

import (
	"database/sql"
	"encoding/json"
	"reflect"
)

type NullString sql.NullString

func (n NullString) MarshalJSON() ([]byte, error) {
	if !n.Valid {
		return json.Marshal(nil)
	}
	return json.Marshal(n.String)
}

func (n *NullString) UnmarshalJSON(data []byte) error {
	var target *string
	if err := json.Unmarshal(data, &target); err != nil {
		return err
	}

	n.Valid = target != nil
	if n.Valid {
		n.String = *target
	} else {
		n.String = ""
	}
	return nil
}

func (n *NullString) Scan(src any) error {
	var sqln sql.NullString
	if err := sqln.Scan(src); err != nil {
		return err
	}
	n.String = sqln.String
	n.Valid = reflect.TypeOf(src) != nil
	return nil
}

func (n *NullString) IsZero() bool {
	return !n.Valid
}
