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
package reader

import (
	"go.uber.org/zap"

	"github.com/superhero-match/consumer-delete-messages/internal/cache"
	"github.com/superhero-match/consumer-delete-messages/internal/config"
	"github.com/superhero-match/consumer-delete-messages/internal/consumer"
)

const timeFormat = "2006-01-02T15:04:05"

// Reader holds all the data relevant.
type Reader struct {
	Consumer          *consumer.Consumer
	Cache             cache.Cache
	Logger            *zap.Logger
	TimeFormat        string
	MessagesKeyFormat string
}

// NewReader configures Reader.
func NewReader(cfg *config.Config) (r *Reader, err error) {
	ch, err := cache.NewCache(cfg)
	if err != nil {
		return nil, err
	}

	cs, err := consumer.NewConsumer(cfg)
	if err != nil {
		return nil, err
	}

	logger, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}

	defer logger.Sync()

	return &Reader{
		Consumer:          cs,
		Cache:             ch,
		Logger:            logger,
		TimeFormat:        timeFormat,
		MessagesKeyFormat: cfg.Cache.MessagesKeyFormat,
	}, nil
}
