package array

import (
	"fmt"
	"testing"

	"github.com/nlbydcg/go-tools/test_common"
)

func TestCheckHas(t *testing.T) {
	t.Run("tools array check_has ", func(t *testing.T) {
		got := CheckHas([]int64{1, 2, 3, 4, 5}, int64(2))
		want := true
		test_common.AssertEqError(t, got, want)
	})

	t.Run("tools array remove duplicate ", func(t *testing.T) {
		got := RemoveDuplicate([]int{1, 2, 3, 4, 5, 6, 1, 2})
		want := []int{1, 2, 3, 4, 5, 6}
		test_common.AssertEqError(t, got, want)
	})

	t.Run("tools array get duplicate", func(t *testing.T) {
		got := GetDuplicate([]int{1, 2, 3, 4, 5, 6, 1, 2, 3, 4})
		want := []int{1, 2, 3, 4}
		test_common.AssertEqError(t, got, want)
	})

	t.Run("tools array to strings", func(t *testing.T) {
		type T struct {
			Id int64
		}
		list := []T{{Id: 1}, {Id: 2}, {Id: 3}}

		got := ToStrings(list)
		fmt.Printf("%#v\n", got)
	})
}
