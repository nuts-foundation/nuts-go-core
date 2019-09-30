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
	"errors"
	"fmt"
	"net"
)

// Error is the interface that extends the default error interface
type Error interface {
	error

	// Recoverable indicates if an action which resulted in the error can be retried
	Recoverable() bool
}

// NutsEventError is the main implementation adding a recoverable field to an error.
// This field will tell if the error is definitive or can be retried at a later moment.
type NutsEventError struct {
	err         error
	recoverable bool
}

// NewError is a convenience method for creating a simple error
func NewError(msg string, recoverable bool) Error {
	return &NutsEventError{
		err:         errors.New(msg),
		recoverable: recoverable,
	}
}

// Errorf creates a new NutsEventError with given format and values
func Errorf(format string, recoverable bool, a ...interface{}) Error {
	return &NutsEventError{
		err:         fmt.Errorf(format, a...),
		recoverable: recoverable,
	}
}

// Wrap tries to identify the error and sets recoverable
func Wrap(err error) Error {
	var recoverable bool

	// net.Error interface
	var netError net.Error
	if errors.As(err, &netError) {
		recoverable = true
	}

	// json.SyntaxError is not recoverable

	return Errorf("%w", recoverable, err)
}

func (ne *NutsEventError) Error() string {
	return ne.err.Error()
}

// Recoverable indicates if an action which resulted in the error can be retried
func (ne *NutsEventError) Recoverable() bool {
	return ne.recoverable
}

// Is is a wrapper for errors.Is()
func (ne *NutsEventError) Is(target error) bool {
	return errors.Is(ne.err, target)
}

// UnWrap is needed for NutsEventError to be UnWrapped
func (ne NutsEventError) UnWrap() error {
	return errors.Unwrap(ne.err)
}
