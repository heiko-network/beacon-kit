// SPDX-License-Identifier: MIT
//
// Copyright (c) 2024 Berachain Foundation
//
// Permission is hereby granted, free of charge, to any person
// obtaining a copy of this software and associated documentation
// files (the "Software"), to deal in the Software without
// restriction, including without limitation the rights to use,
// copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the
// Software is furnished to do so, subject to the following
// conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES
// OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT
// HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY,
// WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
// FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
// OTHER DEALINGS IN THE SOFTWARE.

package kzg_test

import (
	"testing"

	"github.com/berachain/beacon-kit/crypto/kzg"
)

func TestConvertCommitmentToVersionedHash(t *testing.T) {
	commitment := [48]byte{}
	copy(commitment[:], []byte("test commitment"))
	// Assuming BlobCommitmentVersion is a byte value
	expectedPrefix := kzg.BlobCommitmentVersion

	hash := kzg.ConvertCommitmentToVersionedHash(commitment)
	if hash[0] != expectedPrefix {
		t.Errorf(
			"expected first byte of hash to be %v, got %v",
			expectedPrefix,
			hash[0],
		)
	}

	if len(hash) != 32 {
		t.Errorf("expected hash length to be 32, got %d", len(hash))
	}
}

func TestConvertCommitmentsToVersionedHashes(t *testing.T) {
	commitments := make([][48]byte, 2)
	copy(commitments[0][:], "commitment 1")
	copy(commitments[1][:], "commitment 2")
	hashes := kzg.ConvertCommitmentsToVersionedHashes(commitments)

	if len(hashes) != len(commitments) {
		t.Errorf("expected %d hashes, got %d", len(commitments), len(hashes))
	}

	for i, hash := range hashes {
		if hash[0] != kzg.BlobCommitmentVersion {
			t.Errorf(
				"expected first byte of hash %d to be %v, got %v",
				i,
				kzg.BlobCommitmentVersion,
				hash[0],
			)
		}
	}
}