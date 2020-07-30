// Package token provides token stuff for db.
package token

import (
	"context"
	"fmt"

	"github.com/go-snart/snart/db"
	"github.com/go-snart/snart/logs"
)

const _p = "token"

// Log is the logger for token.
var Debug, Info, Warn = logs.Loggers(_p)

// Tokens returns a list of suitable tokens.
func Tokens(ctx context.Context, d *db.DB) []string {
	Debug.Println("enter->env")

	allToks := []string(nil)

	toks, err := EnvTokens()
	if err != nil {
		err = fmt.Errorf("env tok: %w", err)
		Warn.Println(err)
	} else {
		allToks = append(allToks, toks...)
	}

	Debug.Println("env->select")

	toks, err = SelectTokens(ctx, d)
	if err != nil {
		err = fmt.Errorf("select tok: %w", err)
		Warn.Println(err)
	} else {
		allToks = append(allToks, toks...)
	}

	Debug.Println("select->stdin")

	if len(allToks) == 0 {
		toks, err = StdinTokens()
		if err != nil {
			err = fmt.Errorf("stdin tok: %w", err)
			Warn.Println(err)
		} else {
			Debug.Println("stdin->insert")

			InsertTokens(ctx, d, toks)

			Debug.Println("insert->exit")

			allToks = append(allToks, toks...)
		}
	}

	Debug.Println("stdin->exit")

	return allToks
}