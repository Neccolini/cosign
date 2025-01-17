// Copyright 2022 The Sigstore Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cosign

import "fmt"

var (
	// NoMatchingAttestations
	ErrNoMatchingAttestationsMessage = "no matching attestations"
	ErrNoMatchingAttestationsType    = "NoMatchingAttestations"

	// NoMatchingSignatures
	ErrNoMatchingSignaturesType    = "NoMatchingSignatures"
	ErrNoMatchingSignaturesMessage = "no matching signatures"
)

// VerificationError is the type of Go error that is used by cosign to surface
// errors actually related to verification (vs. transient, misconfiguration,
// transport, or authentication related issues).
type VerificationError struct {
	errorType string
	message   string
}

// NewVerificationError constructs a new VerificationError in a manner similar
// to fmt.Errorf
func NewVerificationError(msg string, args ...interface{}) error {
	return &VerificationError{
		message: fmt.Sprintf(msg, args...),
	}
}

// Assert that we implement error at build time.
var _ error = (*VerificationError)(nil)

// Error implements error
func (ve *VerificationError) Error() string {
	return ve.message
}

// Error implements error
func (ve *VerificationError) ErrorType() string {
	return ve.errorType
}

func (ve *VerificationError) SetErrorType(errorType string) {
	ve.errorType = errorType
}
