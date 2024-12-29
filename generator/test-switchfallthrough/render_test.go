package testswitchfallthrough

import (
	"context"
	"github.com/google/go-cmp/cmp"
	"strings"
	"testing"
)

// Need to add a parser and node type for it,
// and a validator to make sure it only appears as a child of the case statement,
// and that there are no nodes between the fallthrough and the next case node.

var input = 0

const expected = `<p>hey</p>`

func TestRender(t *testing.T) {
	w := new(strings.Builder)
	err := example(input).Render(context.Background(), w)
	if err != nil {
		t.Errorf("failed to render: %v", err)
	}
	if diff := cmp.Diff(expected, w.String()); diff != "" {
		t.Error(diff)
	}
}
