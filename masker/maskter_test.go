package masker

import (
	"testing"

	"github.com/nlbydcg/go-tools/test_common"
)

func TestMaskerMobile(t *testing.T) {
	t.Run("masker mobile", func(t *testing.T) {
		got := Mobile("13888882242")
		want := "138****2242"
		test_common.AssertEqError(t, got, want)
	})
}

func TestMaskerEmail(t *testing.T) {
	t.Run("masker email", func(t *testing.T) {
		got := Email("594617354@qq.com")
		want := "59461****@qq.com"
		test_common.AssertEqError(t, got, want)
	})
}
