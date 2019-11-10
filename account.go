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

package scratch

import (
	types "github.com/wealdtech/go-eth2-wallet-types"
)

// StoreAccount stores an account.  It will fail if it cannot store the data.
// Note this will overwrite an existing account with the same ID.  It will not, however, allow multiple accounts with the same
// name to co-exist in the same wallet.
func (s *Store) StoreAccount(wallet types.Wallet, account types.Account, data []byte) error {
	s.accounts[wallet.Name()][account.Name()] = data
	return nil
}

// RetrieveAccount retrieves account-level data.  It will fail if it cannot retrieve the data.
func (s *Store) RetrieveAccount(wallet types.Wallet, name string) ([]byte, error) {
	return s.accounts[wallet.Name()][name], nil
}

// RetrieveAccounts retrieves all account-level data for a wallet.
func (s *Store) RetrieveAccounts(wallet types.Wallet) <-chan []byte {
	ch := make(chan []byte, 1024)
	go func() {
		for _, account := range s.accounts[wallet.Name()] {
			ch <- account
		}
		close(ch)
	}()
	return ch
}
