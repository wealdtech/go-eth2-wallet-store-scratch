// Copyright 2019 - 2023 Weald Technology Trading.
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
	"sync"

	"github.com/google/uuid"
	wtypes "github.com/wealdtech/go-eth2-wallet-types/v2"
)

// Store is the store for the wallet.
type Store struct {
	wallets       map[string][]byte
	walletMu      sync.RWMutex
	accounts      map[string]map[string][]byte
	accountMu     sync.RWMutex
	accountIndex  map[uuid.UUID][]byte
	batchMu       sync.RWMutex
	walletBatches map[uuid.UUID][]byte
}

// New creates a new scratch store.
func New() wtypes.Store {
	return &Store{
		wallets:       make(map[string][]byte),
		accounts:      make(map[string]map[string][]byte),
		accountIndex:  make(map[uuid.UUID][]byte),
		walletBatches: make(map[uuid.UUID][]byte),
	}
}

// Name returns the name of this store.
func (s *Store) Name() string {
	return "scratch"
}
