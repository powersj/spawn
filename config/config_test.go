package config

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestNewConfig(t *testing.T) {
	type testcase struct {
		name     string
		filename string
		cfg      Config
	}
	testcases := []testcase{
		{
			"Basic testcase",
			"testcases/basic.toml",
			Config{
				Agent: agentConfig{
					Interval: 5 * time.Second,
				},
			},
		},
		{
			"Interval set",
			"testcases/interval.toml",
			Config{
				Agent: agentConfig{
					Interval: 10 * time.Second,
				},
			},
		},
	}
	for _, tc := range testcases {
		c, err := NewConfig(tc.filename)
		require.NoError(t, err)
		require.Equal(t, &tc.cfg, c)
	}
}

func TestNewConfigError(t *testing.T) {
	type testcase struct {
		name     string
		filename string
		errorMsg string
	}
	testcases := []testcase{
		{
			"Filename does not exist",
			"fileDoesNotExist.toml",
			"no such file or directory",
		},
		{
			"File not readable",
			"testcases/base64.toml",
			"expected value but found '=' instead",
		},
		{
			"Invalid TOML file",
			"testcases/invalid.toml",
			"expected value but found '\\n' instead",
		},
	}
	for _, tc := range testcases {
		c, err := NewConfig(tc.filename)
		require.Nil(t, c)
		require.Error(t, err)
		require.ErrorContains(t, err, tc.errorMsg)
	}
}
