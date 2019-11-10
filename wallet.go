// Copyright © 2019 Weald Technology Trading
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

// Package scratch provides an in-memory store.  It is ephemeral, and should not be used where accounts are required to be stored.
package scratch

import (
	"fmt"

	types "github.com/wealdtech/go-eth2-wallet-types"
)

// StoreWallet stores wallet-level data.
func (s *Store) StoreWallet(wallet types.Wallet, data []byte) error {
	s.wallets[wallet.Name()] = data
	if _, exists := s.accounts[wallet.Name()]; !exists {
		s.accounts[wallet.Name()] = make(map[string][]byte)
	}
	return nil
}

// RetrieveWallet retrieves wallet-level data.  It will fail if it cannot retrieve the data.
func (s *Store) RetrieveWallet(walletName string) ([]byte, error) {
	data, exists := s.wallets[walletName]
	if !exists {
		return nil, fmt.Errorf("no wallet at %s", walletName)
	}
	return data, nil
}

// RetrieveWallets retrieves wallet-level data for all wallets.
func (s *Store) RetrieveWallets() <-chan []byte {
	ch := make(chan []byte, 1024)
	go func() {
		for _, data := range s.wallets {
			ch <- data
		}
		close(ch)
	}()
	return ch
}
