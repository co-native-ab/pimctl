// Original file is located at: https://github.com/Azure/azure-sdk-for-go/blob/sdk/resourcemanager/authorization/armauthorization/v3.0.0-beta.2/sdk/internal/uuid/uuid.go
//
// Original file header:
// ---
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// ---
//
// Original license:
// ---
// The MIT License (MIT)

// Copyright (c) Microsoft Corporation.

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.
// ---

package azurerm

import (
	"crypto/rand"
	"errors"
	"fmt"
	"strconv"
)

// The UUID reserved variants.
const (
	reservedRFC4122 byte = 0x40
)

// A uuid representation compliant with specification in RFC4122 document.
type uuid [16]byte

// New returns a new UUID using the RFC4122 algorithm.
func newUUID() (uuid, error) {
	u := uuid{}
	// Set all bits to pseudo-random values.
	// NOTE: this takes a process-wide lock
	_, err := rand.Read(u[:])
	if err != nil {
		return u, err
	}
	u[8] = (u[8] | reservedRFC4122) & 0x7F // u.setVariant(ReservedRFC4122)

	var version byte = 4
	u[6] = (u[6] & 0xF) | (version << 4) // u.setVersion(4)
	return u, nil
}

// String returns the UUID in "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx" format.
func (u uuid) String() string {
	return fmt.Sprintf("%x-%x-%x-%x-%x", u[0:4], u[4:6], u[6:8], u[8:10], u[10:])
}

// Parse parses a string formatted as "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"
// or "{xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx}" into a UUID.
func parseUUID(s string) (uuid, error) {
	var uuid uuid
	// ensure format
	switch len(s) {
	case 36:
		// xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
	case 38:
		// {xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx}
		s = s[1:37]
	default:
		return uuid, errors.New("invalid UUID format")
	}
	if s[8] != '-' || s[13] != '-' || s[18] != '-' || s[23] != '-' {
		return uuid, errors.New("invalid UUID format")
	}
	// parse chunks
	for i, x := range [16]int{
		0, 2, 4, 6,
		9, 11,
		14, 16,
		19, 21,
		24, 26, 28, 30, 32, 34} {
		b, err := strconv.ParseUint(s[x:x+2], 16, 8)
		if err != nil {
			return uuid, fmt.Errorf("invalid UUID format: %s", err)
		}
		uuid[i] = byte(b)
	}
	return uuid, nil
}
