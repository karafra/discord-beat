package main

import (
	beater2 "DiscordBeat/beater"
	"github.com/elastic/beats/libbeat/beat"
	"os"
)

func main() {
	err := beat.Run("discord-beat", "", beater2.New)
	if err != nil {
		os.Exit(1)
	}
}
