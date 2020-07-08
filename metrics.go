/*
 * Nuts go core
 * Copyright (C) 2019 Nuts community
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <https://www.gnu.org/licenses/>.
 *
 */

package core

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const nutsMetricsPrefix = "nuts_"

//NewMetricsEngine creates a new Engine for exposing prometheus metrics via http.
//Metrics are exposed on /metrics, by default the GoCollector and ProcessCollector are enabled.
func NewMetricsEngine() *Engine {
	return &Engine{
		Name:      "Metrics",
		Configure: configure,
		Routes: func(router EchoRouter) {
			router.GET("/metrics", echo.WrapHandler(promhttp.Handler()))
		},
	}
}

var nutsRegistry prometheus.Registerer

// Registerer returns a prometheus register with the engine name added as prefix. The nuts_ prefix has also been added.
func Registerer(engine *Engine) prometheus.Registerer {
	var regex = regexp.MustCompile(`\s+`)
	var engineName = regex.ReplaceAllString(engine.Name, "_")
	engineName = fmt.Sprintf("%s_", strings.ToLower(engineName))

	return prometheus.WrapRegistererWithPrefix(engineName, nutsRegistry)
}

func configure() error {
	nutsRegistry = prometheus.WrapRegistererWithPrefix(nutsMetricsPrefix, prometheus.NewRegistry())
	nutsRegistry.MustRegister(prometheus.NewGoCollector())
	nutsRegistry.MustRegister(prometheus.NewProcessCollector(prometheus.ProcessCollectorOpts{}))

	return nil
}
