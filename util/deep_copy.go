package util

import (
	"fmt"
	"github.com/jinzhu/copier"
)

func DeepCopy[T any](object *T) (*T, error) {
	other := new(T)
	err := copier.Copy(other, object)
	if err != nil {
		return nil, fmt.Errorf("cannot copy the data: %w", err)
	}
	return other, nil
}
