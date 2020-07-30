package bot

import "time"

// WaitReady loops until the Bot has a valid Session, and then calls the Bot's DB's WaitReady method.
func (b *Bot) WaitReady() {
	const _f = "(*Bot).WaitReady"

	for {
		Log.Debug(_f, "wait for session")

		if b.Session.State.User != nil {
			break
		}

		time.Sleep(time.Second / 10)
	}
}
