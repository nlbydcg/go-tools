package test_common

import (
	"fmt"
	"reflect"
	"testing"
)

func AssertEqError(t *testing.T, got, want interface{}) error {
	if reflect.DeepEqual(got, want) {
		return fmt.Errorf("not eq !got:%#v ,want:%#v", got, want)
	}
	return nil
}
