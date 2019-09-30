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
	"context"
	"errors"
	"testing"
)

func TestNew(t *testing.T) {
	t.Run("New generates error", func(t *testing.T) {
		msg := "error message"
		var e error = NewError(msg, true)

		if e.Error() != msg {
			t.Errorf("expected [%s], got[%s]", msg, e.Error())
		}
	})

	t.Run("Nuts Error allocation", func(t *testing.T) {
		msg := "error message"
		var e error = NewError(msg, true)
		var e2 error = NewError(msg, true)

		if e == e2 {
			t.Error("expected allocations to be different")
		}
	})
}

func TestErrorf(t *testing.T) {
	t.Run("Errorf generates error with correct string arg", func(t *testing.T) {
		format := "error message, cause: %s"
		cause := "catastrophic failure"
		expected := "error message, cause: catastrophic failure"

		e := Errorf(format, false, cause)

		if e.Error() != expected {
			t.Errorf("expected [%s], got[%s]", expected, e.Error())
		}
	})

	t.Run("Errorf generates error with correct error arg", func(t *testing.T) {
		format := "error message, cause: %w"
		cause := errors.New("catastrophic failure")
		expected := "error message, cause: catastrophic failure"

		e := Errorf(format, false, cause)

		if e.Error() != expected {
			t.Errorf("expected [%s], got[%s]", expected, e.Error())
		}
	})
}

func TestNutsError_As(t *testing.T) {
	t.Run("Errorf generates error compatible with errors.As()", func(t *testing.T) {
		format := "error message, cause: %w"
		cause := errors.New("catastrophic failure")

		e := Errorf(format, false, cause)

		var nutsError *NutsEventError

		if !errors.As(e, &nutsError) {
			t.Error("expected NutsEventError to be able to be type casted")
		}
	})
}

func TestNutsError_UnWrap(t *testing.T) {
	t.Run("Errorf generates error that can be unwrapped", func(t *testing.T) {
		format := "error message, cause: %w"
		cause := errors.New("catastrophic failure")

		e := Errorf(format, false, cause)
		var nutsError *NutsEventError
		errors.As(e, &nutsError)

		if nutsError.UnWrap().Error() != "catastrophic failure" {
			t.Error("expected Unwrapped error to equals [catastrophic failure]")
		}
	})
}

func TestNutsError_Is(t *testing.T) {
	t.Run("Errorf generates error compatible with errors.Is()", func(t *testing.T) {
		format := "error message, cause: %w"
		cause := errors.New("catastrophic failure")

		e := Errorf(format, false, cause)

		if !errors.Is(e, cause) {
			t.Errorf("expected [%s] to be a [%s]", e.Error(), cause.Error())
		}
	})

	t.Run("Is() is compatible with errors.Is()", func(t *testing.T) {
		format := "error message, cause: %w"
		cause := errors.New("catastrophic failure")

		e := Errorf(format, false, cause)
		var nutsError *NutsEventError
		errors.As(e, &nutsError)

		if !nutsError.Is(cause) {
			t.Errorf("expected [%s] to be a [%s]", e.Error(), cause.Error())
		}
	})
}

type errTimeout struct{}
func (e *errTimeout) Error() string { return "error" }
func (e *errTimeout) Timeout() bool { return true }
func (e *errTimeout) Temporary() bool { return true }

func TestWrap(t *testing.T) {
	t.Run("Wrap returns recoverable error for Timeout error", func(t *testing.T) {
		e := Wrap(&errTimeout{})

		if !e.Recoverable() {
			t.Error("Expected timeout error to be recoverable")
		}
	})

	t.Run("Wrap returns recoverable error for context deadline", func(t *testing.T) {
		e := Wrap(context.DeadlineExceeded)

		if !e.Recoverable() {
			t.Error("Expected deadline error to be recoverable")
		}
	})

	t.Run("Wrap returns non-recoverable error by default", func(t *testing.T) {
		e := Wrap(errors.New("error"))

		if e.Recoverable() {
			t.Error("Expected normal error to be non-recoverable")
		}
	})
}
