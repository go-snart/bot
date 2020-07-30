package route

import (
	"fmt"

	dg "github.com/bwmarrin/discordgo"
)

// Okay is a function which checks if a Ctx should be used.
type Okay func(*Ctx) bool

// Any is a logical OR of Okays.
func Any(chs ...Okay) Okay {
	return func(c *Ctx) bool {
		for _, ch := range chs {
			if ch(c) {
				return true
			}
		}

		return false
	}
}

// All is a logical AND of Okays.
func All(chs ...Okay) Okay {
	return func(c *Ctx) bool {
		for _, ch := range chs {
			if !ch(c) {
				return false
			}
		}

		return true
	}
}

// False is an Okay that always returns false.
var False Okay = func(*Ctx) bool {
	return false
}

// True is an Okay that always returns true.
var True Okay = func(*Ctx) bool {
	return true
}

// GuildAdmin is an Okay that checks if the user has administrator privileges on the guild.
var GuildAdmin Okay = func(c *Ctx) bool {
	const _f = "GuildAdmin"

	perm, err := c.Session.UserChannelPermissions(c.Message.Author.ID, c.Message.ChannelID)
	if err != nil {
		err = fmt.Errorf("perm: %w", err)
		Log.Warn(_f, err)

		return false
	}

	return perm&(dg.PermissionAdministrator|
		dg.PermissionManageServer) > 0
}
