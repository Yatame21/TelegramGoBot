package Tg

const msgHelp = `This bot stores your links and returns them upon request.

To save a link: simply send it to me.
To retrieve a link: use the /rnd command to receive a random page from your saved list.

Please note: once a link is shown to you, it will be permanently removed from your list to ensure timely reading.
`

const msgHello = "Hello there! \n\n" + msgHelp

const (
	msgUnknownCommand = "Unknown command"
	msgNoSavedPages   = "You have no saved pages"
	msgSaved          = "Saved!"
	msgAlreadyExists  = "This page is already in your list"
)
