/*
 * Nuts go
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
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/nuts-foundation/nuts-go-core/mock"
	"github.com/stretchr/testify/assert"
)

func TestRegisterEngine(t *testing.T) {
	t.Run("adds an engine to the list", func(t *testing.T) {
		ctl := EngineControl{
			Engines: []*Engine{},
		}
		ctl.registerEngine(&Engine{})

		if len(ctl.Engines) != 1 {
			t.Errorf("Expected 1 registered engine, Got %d", len(ctl.Engines))
		}
	})
}

func TestNewStatusEngine_Routes(t *testing.T) {
	t.Run("Registers a single route for listing all engines", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		echo := mock.NewMockEchoRouter(ctrl)

		echo.EXPECT().GET("/status/diagnostics", gomock.Any())
		echo.EXPECT().GET("/status", gomock.Any())

		NewStatusEngine().Routes(echo)
	})
}

func TestNewStatusEngine_Cmd(t *testing.T) {
	t.Run("Cmd returns a cobra command", func(t *testing.T) {
		e := NewStatusEngine().Cmd
		if e.Name() != "engineStatus" {
			t.Errorf("Expected a command with name engineStatus, Got %s", e.Name())
		}
	})
}

func TestNewStatusEngine_Diagnostics(t *testing.T) {
	RegisterEngine(NewStatusEngine())
	RegisterEngine(NewLoggerEngine())

	t.Run("Diagnostics returns engine list", func(t *testing.T) {
		ds := NewStatusEngine().Diagnostics()
		assert.Len(t, ds, 1)
		assert.Equal(t, "Registered engines", ds[0].Name())
		assert.Equal(t, "Status,Logging", ds[0].String())
	})

	t.Run("ListAllEngines renders json output of list of engines", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		echo := mock.NewMockContext(ctrl)

		echo.EXPECT().String(http.StatusOK, "Registered engines: Status,Logging\nLogger verbosity: ")

		diagnosticsOverview(echo)
	})
}

func TestStatusOK(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	echo := mock.NewMockContext(ctrl)

	echo.EXPECT().String(http.StatusOK, "OK")

	StatusOK(echo)
}
