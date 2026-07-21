package Tg

import "TestGOBot/clients/telegram"

type Dispatcher struct {
	tg     *telegram.Client
	offset int
	// storage
}

// func New
