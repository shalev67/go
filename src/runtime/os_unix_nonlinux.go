// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build unix && !linux

package runtime

// sigFromUser reports whether the signal was sent because of a call
// to kill.
func (c *sigctxt) sigFromUser() bool {
	return c.sigcode() == _SI_USER
}
