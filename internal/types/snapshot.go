// Copyright 2026 Erst Users
// SPDX-License-Identifier: Apache-2.0

package types

// StateSnapshot mirrors the Rust simulator snapshot payload.
// Field order is kept consistent with the Rust struct for serde/bincode parity.
type StateSnapshot struct {
	LedgerEntries    map[string]string `json:"ledger_entries"`
	Timestamp        uint64            `json:"timestamp"`
	InstructionIndex uint32            `json:"instruction_index"`
	Events           []string          `json:"events"`
}

// LedgerDelta mirrors the Rust state::StateDiff JSON shape used by the UI.
type LedgerDelta struct {
	NewKeys      []string `json:"new_keys"`
	ModifiedKeys []string `json:"modified_keys"`
	DeletedKeys  []string `json:"deleted_keys"`
}
