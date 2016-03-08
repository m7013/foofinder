package foofinder

import (
	"testing"

	"github.com/bmizerany/assert"
)

func TestIsItFoo(t *testing.T) {
	word := "foo"
	foo, err := IsItFoo(word)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, true, foo)
}
