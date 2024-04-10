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

package types_test

import (
	"testing"

	"github.com/berachain/beacon-kit/mod/da/types"
	primitives "github.com/berachain/beacon-kit/mod/primitives"
	"github.com/berachain/beacon-kit/mod/primitives/kzg"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEmptySidecarMarshalling(t *testing.T) {
	// Create an empty BlobSidecar
	sidecar := types.BlobSidecar{
		Index:             0,
		Blob:              kzg.Blob{},
		BeaconBlockHeader: &primitives.BeaconBlockHeader{},
		InclusionProof: [][]byte{
			[]byte("00000000000000000000000000000001"),
			[]byte("00000000000000000000000000000002"),
			[]byte("00000000000000000000000000000003"),
			[]byte("00000000000000000000000000000004"),
			[]byte("00000000000000000000000000000005"),
			[]byte("00000000000000000000000000000006"),
			[]byte("00000000000000000000000000000007"),
			[]byte("00000000000000000000000000000008"),
		},
	}

	// Marshal the empty sidecar
	marshalled, err := sidecar.MarshalSSZ()
	require.NoError(
		t,
		err,
		"Marshalling empty sidecar should not produce an error",
	)
	require.NotNil(
		t,
		marshalled,
		"Marshalling empty sidecar should produce a result",
	)

	// Unmarshal the empty sidecar
	unmarshalled := types.BlobSidecar{}
	err = unmarshalled.UnmarshalSSZ(marshalled)
	require.NoError(
		t,
		err,
		"Unmarshalling empty sidecar should not produce an error",
	)

	// Compare the original and unmarshalled empty sidecars
	assert.Equal(
		t,
		sidecar,
		unmarshalled,
		"The original and unmarshalled empty sidecars should be equal",
	)
}
func TestValidateBlockRoots(t *testing.T) {
	// Create a sample BlobSidecar with valid roots
	validSidecar := types.BlobSidecar{
		Index: 0,
		Blob:  kzg.Blob{},
		BeaconBlockHeader: &primitives.BeaconBlockHeader{
			StateRoot: [32]byte{1},
			BodyRoot:  [32]byte{2},
		},
		InclusionProof: [][]byte{
			[]byte("00000000000000000000000000000001"),
			[]byte("00000000000000000000000000000002"),
			[]byte("00000000000000000000000000000003"),
			[]byte("00000000000000000000000000000004"),
			[]byte("00000000000000000000000000000005"),
			[]byte("00000000000000000000000000000006"),
			[]byte("00000000000000000000000000000007"),
			[]byte("00000000000000000000000000000008"),
		},
	}

	// Validate the sidecar with valid roots
	sidecars := types.BlobSidecars{
		Sidecars: []*types.BlobSidecar{&validSidecar},
	}
	err := sidecars.ValidateBlockRoots()
	require.NoError(
		t,
		err,
		"Validating sidecar with valid roots should not produce an error",
	)

	// Create a sample BlobSidecar with invalid roots
	differentBlockRootSidecar := types.BlobSidecar{
		Index: 0,
		Blob:  kzg.Blob{},
		BeaconBlockHeader: &primitives.BeaconBlockHeader{
			StateRoot: [32]byte{},
			BodyRoot:  [32]byte{},
		},
		InclusionProof: [][]byte{
			[]byte("00000000000000000000000000000001"),
			[]byte("00000000000000000000000000000002"),
			[]byte("00000000000000000000000000000003"),
			[]byte("00000000000000000000000000000004"),
			[]byte("00000000000000000000000000000005"),
			[]byte("00000000000000000000000000000006"),
			[]byte("00000000000000000000000000000007"),
			[]byte("00000000000000000000000000000008"),
		},
	}
	// Validate the sidecar with invalid roots
	sidecarsInvalid := types.BlobSidecars{
		Sidecars: []*types.BlobSidecar{
			&validSidecar,
			&differentBlockRootSidecar,
		},
	}
	err = sidecarsInvalid.ValidateBlockRoots()
	require.Error(
		t,
		err,
		"Validating sidecar with invalid roots should produce an error",
	)
}