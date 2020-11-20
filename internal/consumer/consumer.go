/*
  Copyright (C) 2019 - 2021 MWSOFT
  This program is free software: you can redistribute it and/or modify
  it under the terms of the GNU General Public License as published by
  the Free Software Foundation, either version 3 of the License, or
  (at your option) any later version.
  This program is distributed in the hope that it will be useful,
  but WITHOUT ANY WARRANTY; without even the implied warranty of
  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
  GNU General Public License for more details.
  You should have received a copy of the GNU General Public License
  along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/
package consumer

import (
	"time"

	"github.com/segmentio/kafka-go"

	"github.com/superhero-match/consumer-delete-messages/internal/config"
)

// Consumer holds Kafka consumer related data.
type Consumer struct {
	Consumer *kafka.Reader
}

// NewConsumer configures Kafka consumer that consumes from configured topic.
func NewConsumer(cfg *config.Config) (*Consumer, error) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:       cfg.Consumer.Brokers,
		Topic:         cfg.Consumer.Topic,
		GroupID:       cfg.Consumer.GroupID,
		QueueCapacity: int(1),
		MaxWait:       time.Second,
	})

	return &Consumer{
		Consumer: r,
	}, nil
}
