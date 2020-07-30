// Package admin handles bot-wide administration permissions.
package admin

import (
	"context"
	"fmt"

	"github.com/go-snart/snart/db"
	"github.com/go-snart/snart/logs"
	"github.com/go-snart/snart/route"
)

const _p = "admin"

var Debug, Info, Warn = logs.Loggers(_p)

// Table builds the table of admins.
func Table(ctx context.Context, d *db.DB) {
	const (
		e = `CREATE TABLE IF NOT EXISTS admin(
			id TEXT PRIMARY KEY UNIQUE
		)`
	)

	_, err := d.Conn(&ctx).Exec(ctx, e)
	if err != nil {
		err = fmt.Errorf("exec %#q: %w", e, err)

		Warn.Println(err)

		return
	}
}

// IsAdmin checks if the author has bot-wide admin privileges.
func IsAdmin(d *db.DB) route.Okay {
	return func(c *route.Ctx) bool {
		for _, admin := range List(c, d) {
			if c.Message.Author.ID == admin {
				return true
			}
		}

		app, err := c.Session.Application("@me")
		if err != nil {
			err = fmt.Errorf("app @me: %w", err)
			Warn.Println(err)

			return false
		}

		if app.Owner != nil && c.Message.Author.ID == app.Owner.ID {
			return true
		}

		if app.Team != nil {
			if c.Message.Author.ID == app.Team.OwnerID {
				return true
			}

			for _, member := range app.Team.Members {
				if c.Message.Author.ID == member.User.ID {
					return true
				}
			}
		}

		return false
	}
}

// List returns a list of known admin IDs from the database.
func List(ctx context.Context, d *db.DB) []string {
	Table(ctx, d)

	const q = `SELECT id FROM admin`

	rows, err := d.Conn(&ctx).Query(ctx, q)
	if err != nil {
		err = fmt.Errorf("query %#q: %w", q, err)

		Warn.Println(err)

		return nil
	}
	defer rows.Close()

	list := []string(nil)

	for rows.Next() {
		admin := ""

		err = rows.Scan(&admin)
		if err != nil {
			err = fmt.Errorf("scan admin: %w", err)

			return nil
		}

		list = append(list, admin)
	}

	err = rows.Err()
	if err != nil {
		err = fmt.Errorf("rows err: %w", err)

		return nil
	}

	return list
}