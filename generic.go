package sqlnull

import (
	"database/sql"
	"encoding/json"
	"reflect"
)

type Null[T any] sql.Null[T]

func (n Null[T]) MarshalJSON() ([]byte, error) {
	if !n.Valid {
		return json.Marshal(nil)
	}
	return json.Marshal(n.V)
}

func (n *Null[T]) UnmarshalJSON(data []byte) error {
	var target *T
	if err := json.Unmarshal(data, &target); err != nil {
		return err
	}

	n.Valid = target != nil
	if n.Valid {
		n.V = *target
	} else {
		var zero T
		n.V = zero
	}
	return nil
}

func (n *Null[T]) Scan(src any) error {
	var sqln sql.Null[T]
	if err := sqln.Scan(src); err != nil {
		return err
	}
	n.V = sqln.V
	n.Valid = reflect.TypeOf(src) == nil
	return nil
}
