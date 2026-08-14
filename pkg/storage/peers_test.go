package storage

import (
	"context"
	"fmt"
	"testing"

	"github.com/gotd/td/telegram/peers"

	"github.com/iyear/tdl/pkg/kv"
)

type wrappedNotFoundKV struct{}

func (wrappedNotFoundKV) Get(string) ([]byte, error) {
	return nil, fmt.Errorf("wrapped lookup failure: %w", kv.ErrNotFound)
}

func (wrappedNotFoundKV) Set(string, []byte) error {
	return nil
}

func (wrappedNotFoundKV) Delete(string) error {
	return nil
}

func TestPeersStorageTreatsWrappedNotFoundAsMissing(t *testing.T) {
	ctx := context.Background()
	store := NewPeers(wrappedNotFoundKV{})

	if _, found, err := store.Find(ctx, peers.Key{}); err != nil || found {
		t.Fatalf("Find() = found %v, err %v; want missing peer without error", found, err)
	}

	if _, _, found, err := store.FindPhone(ctx, "12345"); err != nil || found {
		t.Fatalf("FindPhone() = found %v, err %v; want missing phone without error", found, err)
	}

	if hash, err := store.GetContactsHash(ctx); err != nil || hash != 0 {
		t.Fatalf("GetContactsHash() = hash %d, err %v; want zero hash without error", hash, err)
	}
}
