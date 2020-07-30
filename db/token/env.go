package token

import (
	"errors"
	"os"
	"strings"
)

// EnvName is the name of the environment variable used to load tokens.
const EnvName = "SNART_TOKENS"

// ErrEnvUnset occurs when the EnvName environment variable is not set.
var ErrEnvUnset = errors.New(EnvName + " is not set")

// EnvTokens returns the tokens listed in the EnvName environment variable.
func EnvTokens() ([]string, error) {
	toks, ok := os.LookupEnv(EnvName)
	if !ok {
		return nil, ErrEnvUnset
	}

	return strings.Split(toks, ":"), nil
}