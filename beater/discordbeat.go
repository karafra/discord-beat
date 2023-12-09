package beater

import (
	"fmt"
	"time"

	"github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/logp"
	"github.com/elastic/beats/libbeat/publisher"
	"github.com/karafra/discord-beat/config"
)

type DiscordBeat struct {
	done   chan struct{}
	config config.Config
	client publisher.Client
}

// New Creates beater
//
//goland:noinspection GoUnusedParameter
func New(b *beat.Beat, cfg *common.Config) (beat.Beater, error) {
	providedConfig := config.DefaultConfig
	if err := cfg.Unpack(&providedConfig); err != nil {
		return nil, fmt.Errorf("error reading providedConfig file: %v", err)
	}

	bt := &DiscordBeat{
		done:   make(chan struct{}),
		config: providedConfig,
	}
	return bt, nil
}

func (bt *DiscordBeat) Run(b *beat.Beat) error {
	logp.Info("DiscordBeat is running! Hit CTRL-C to stop it.")

	bt.client = b.Publisher.Connect()
	ticker := time.NewTicker(bt.config.Period)
	counter := 1
	for {
		select {
		case <-bt.done:
			return nil
		case <-ticker.C:
		}

		event := common.MapStr{
			"@timestamp": common.Time(time.Now()),
			"type":       b.Name,
			"counter":    counter,
		}
		bt.client.PublishEvent(event)
		logp.Info("Event sent")
		counter++
	}
}

func (bt *DiscordBeat) Stop() {
	err := bt.client.Close()
	if err != nil {
		logp.Err("Error while closing beat.", err)
		return
	}
	close(bt.done)
}
