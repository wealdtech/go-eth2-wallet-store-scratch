// Copyright 2023 Weald Technology Trading.
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

package scratch

import (
	"context"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

// StoreBatch stores wallet batch data.  It will fail if it cannot store the data.
func (s *Store) StoreBatch(_ context.Context, walletID uuid.UUID, _ string, data []byte) error {
	s.batchMu.Lock()
	s.walletBatches[walletID] = data
	s.batchMu.Unlock()

	return nil
}

// RetrieveBatch retrieves the batch of accounts for a given wallet.
func (s *Store) RetrieveBatch(_ context.Context, walletID uuid.UUID) ([]byte, error) {
	s.batchMu.RLock()
	batch, exists := s.walletBatches[walletID]
	s.batchMu.RUnlock()

	if !exists {
		return nil, errors.New("no batch")
	}

	return batch, nil
}
