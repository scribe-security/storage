// --------------------------------------------------------------
// SPDX-FileCopyrightText: Copyright © 2024 The Protobom Authors
// SPDX-FileType: SOURCE
// SPDX-License-Identifier: Apache-2.0
// --------------------------------------------------------------
package ent

import (
	"fmt"

	"github.com/protobom/protobom/pkg/sbom"
	"github.com/protobom/protobom/pkg/storage"

	"github.com/protobom/storage/internal/backends/ent"
	"github.com/protobom/storage/internal/backends/ent/hashesentry"
)

type HashesBackend Backend[map[sbom.HashAlgorithm]string]

var _ storage.Backend[map[sbom.HashAlgorithm]string] = (*HashesBackend)(nil)

func (backend *HashesBackend) Store(hashes *map[sbom.HashAlgorithm]string, _opts *storage.StoreOptions) error {
	for alg, content := range *hashes {
		err := backend.client.HashesEntry.Create().
			SetHashAlgorithmType(hashesentry.HashAlgorithmType(alg.String())).
			SetHashData(content).
			OnConflict().
			Ignore().
			Exec(backend.ctx)

		if err != nil && !ent.IsConstraintError(err) {
			return fmt.Errorf("failed creating ent.HashesEntry: %w", err)
		}
	}

	return nil
}

func (backend *HashesBackend) Retrieve(
	_id string,
	_opts *storage.RetrieveOptions,
) (*map[sbom.HashAlgorithm]string, error) {
	return nil, nil
}
