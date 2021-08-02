package logger

import (
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
)

func TestInit(t *testing.T) {
	Log = logrus.New()
	LogError = logrus.New()
}

func TestInitLogger(t *testing.T) {
	require := require.New(t)

	err := initLogger(Log, "../../logs/info.log")
	require.NoError(err)
	err = initLogger(LogError, "../../logs/error.log")
	require.NoError(err)
}
