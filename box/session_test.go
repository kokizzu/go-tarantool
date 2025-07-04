package box_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/tarantool/go-tarantool/v2/box"
	th "github.com/tarantool/go-tarantool/v2/test_helpers"
)

func TestBox_Session(t *testing.T) {
	b := box.New(th.Ptr(th.NewMockDoer(t)))
	require.NotNil(t, b.Session())
}

func TestNewSessionSuRequest(t *testing.T) {
	_, err := box.NewSessionSuRequest("admin")
	require.NoError(t, err)
}
