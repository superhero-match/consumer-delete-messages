/*
  Copyright (C) 2019 - 2020 MWSOFT
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
	"github.com/superhero-match/consumer-delete-messages/internal/cache"
	"github.com/superhero-match/consumer-delete-messages/internal/config"
	"github.com/superhero-match/consumer-delete-messages/internal/consumer"
)

// Reader holds all the data relevant.
type Reader struct {
	Consumer *consumer.Consumer
	Cache    *cache.Cache
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

	return &Reader{
		Consumer: cs,
		Cache:    ch,
	}, nil
}
