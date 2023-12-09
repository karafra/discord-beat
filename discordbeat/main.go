package main

import (
	"os"

	"github.com/elastic/beats/libbeat/beat"

	"github.com/karafra/discord-beat/discordbeat/beater"
)

func main() {
	err := beat.Run("discordbeat", "", beater.New)
	if err != nil {
		os.Exit(1)
	}
}
