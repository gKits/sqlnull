package sqlnull

import (
	"database/sql"
	"encoding/json"
	"reflect"
	"time"
)

type NullTime sql.NullTime

func (n NullTime) MarshalJSON() ([]byte, error) {
	if !n.Valid {
		return json.Marshal(nil)
	}
	return json.Marshal(n.Time)
}

func (n *NullTime) UnmarshalJSON(data []byte) error {
	var target *time.Time
	if err := json.Unmarshal(data, &target); err != nil {
		return err
	}

	n.Valid = target != nil
	if n.Valid {
		n.Time = *target
	} else {
		n.Time = time.Time{}
	}
	return nil
}

func (n *NullTime) Scan(src any) error {
	var sqln sql.NullTime
	if err := sqln.Scan(src); err != nil {
		return err
	}
	n.Time = sqln.Time
	n.Valid = reflect.TypeOf(src) != nil
	return nil
}

func (n *NullTime) IsZero() bool {
	return !n.Valid
}
