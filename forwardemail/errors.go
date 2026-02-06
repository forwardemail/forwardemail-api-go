// Copyright Forward Email 2026
// SPDX-License-Identifier: MIT

package forwardemail

import "errors"

// ErrRequestFailure is returned when an API request fails due to a non-2xx status code.
var ErrRequestFailure = errors.New("failed to complete request")
