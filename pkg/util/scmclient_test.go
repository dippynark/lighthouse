package util_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/jenkins-x/lighthouse/pkg/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// This test ensures that we can retrieve the Git token from the environment
func TestGetSCMToken(t *testing.T) {
	os.Setenv("GIT_TOKEN", "mytokenfromenvvar")

	token, err := util.GetSCMToken("github")
	require.NoError(t, err, "failed to get SCM token")
	assert.Equal(t, "mytokenfromenvvar", token, "failed to get expected SCM token: %s", token)
}

// This test ensures that we can retrieve the Git token from the filesystem
func TestGetSCMTokenPath(t *testing.T) {
	err := os.Unsetenv("GIT_TOKEN")
	require.NoError(t, err, "failed to unset environment variable GIT_TOKEN")

	os.Setenv("GIT_TOKEN_PATH", filepath.Join("test_data", "secret_dir", "git-token"))

	token, err := util.GetSCMToken("github")
	require.NoError(t, err, "failed to get SCM token")
	assert.Equal(t, "mytokenfrompath", token, "failed to get expected SCM token: %s", token)
}

// This test ensures that GIT_TOKEN takes priority over GIT_TOKEN_PATH for backwards compatibility
func TestGetSCMTokenPathFallback(t *testing.T) {
	os.Setenv("GIT_TOKEN", "mytokenfromenvvar")
	os.Setenv("GIT_TOKEN_PATH", filepath.Join("test_data", "secret_dir", "git-token"))

	token, err := util.GetSCMToken("github")
	require.NoError(t, err, "failed to get SCM token")
	assert.Equal(t, "mytokenfromenvvar", token, "failed to get expected SCM token: %s", token)
}

// This test ensures that we receive an error if path GIT_TOKEN_PATH does not exist
func TestGetSCMTokenPathMissingError(t *testing.T) {
	err := os.Unsetenv("GIT_TOKEN")
	require.NoError(t, err, "failed to unset environment variable GIT_TOKEN")

	os.Setenv("GIT_TOKEN_PATH", filepath.Join("test_data", "secret_dir", "does-not-exist"))

	_, err = util.GetSCMToken("github")
	require.Error(t, err)
}

// This test ensures that we receive an error if neither GIT_TOKEN or GIT_TOKEN_PATH are set
func TestGetSCMTokenUnsetError(t *testing.T) {
	err := os.Unsetenv("GIT_TOKEN")
	require.NoError(t, err, "failed to unset environment variable GIT_TOKEN")

	err = os.Unsetenv("GIT_TOKEN_PATH")
	require.NoError(t, err, "failed to unset environment variable GIT_TOKEN_PATH")

	_, err = util.GetSCMToken("github")
	require.Error(t, err)
}
