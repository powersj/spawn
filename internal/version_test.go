package internal

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestConst(t *testing.T) {
	require.Equal(t, "spawn", AppName)
}

func TestUnsetVar(t *testing.T) {
	require.Equal(t, "unset", Version)
	require.Equal(t, "unset", Branch)
	require.Equal(t, "unset", Commit)
}

func TestUserAgent(t *testing.T) {
	require.Contains(t, UserAgent(), "spawn/unset Go/")
}

func TestAppVersion(t *testing.T) {
	require.Contains(t, AppVersion(), "spawn version unset (git: unset@unset) (go:")
}
