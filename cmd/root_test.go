package cmd

import (
	"path/filepath"
	"testing"

	"github.com/spf13/viper"
)

// TestInitConfigEnvBinding guards issue #61: OAuth2 credentials must be
// resolvable from CNTB_-prefixed environment variables. The config keys are
// hyphenated (e.g. "oauth2-clientid"), so without an env key replacer viper
// looks for the un-settable name "CNTB_OAUTH2-CLIENTID" and the credentials
// come back empty ("clientId is empty").
func TestInitConfigEnvBinding(t *testing.T) {
	origCfg := cfgFile
	t.Cleanup(func() {
		cfgFile = origCfg
		viper.Reset()
	})

	viper.Reset()
	// Point at a non-existent file so no real ~/.cntb.yaml is read.
	cfgFile = filepath.Join(t.TempDir(), "does-not-exist.yaml")

	t.Setenv("CNTB_OAUTH2_CLIENTID", "env-client-id")          // single-word suffix
	t.Setenv("CNTB_OAUTH2_CLIENT_SECRET", "env-client-secret") // multi-hyphen key
	t.Setenv("CNTB_OAUTH2_USER", "env-user")

	initConfig()

	for _, c := range []struct{ name, got, want string }{
		{"oauth2-clientid", oauth2ClientId, "env-client-id"},
		{"oauth2-client-secret", oauth2ClientSecret, "env-client-secret"},
		{"oauth2-user", oauth2User, "env-user"},
	} {
		if c.got != c.want {
			t.Errorf("%s from env = %q, want %q", c.name, c.got, c.want)
		}
	}
}
