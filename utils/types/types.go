package types

//! Debricated

import (
	"encoding/json"
)

type Nullable[T any] struct {
	Value  T
	IsNull bool
}

func (n Nullable[T]) MarshalJSON() ([]byte, error) {
	if n.IsNull {
		return json.Marshal(nil)
	}
	return json.Marshal(n.Value)
}

func (n *Nullable[T]) UnmarshalJSON(data []byte) error {
	var v *T
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	if v != nil {
		n.Value = *v
		n.IsNull = false
	} else {
		var none T
		n.Value = none
		n.IsNull = true
	}
	return nil
}
