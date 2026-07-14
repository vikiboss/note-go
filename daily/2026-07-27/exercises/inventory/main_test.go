package main

import (
	"errors"
	"testing"
)

func TestInventoryRoundTrip(t *testing.T) {
	inv := NewInventory()
	for _, item := range []Item{{SKU: "b", Quantity: 2}, {SKU: "a", Quantity: 1}} {
		if err := inv.Put(item); err != nil {
			t.Fatal(err)
		}
	}
	data, err := inv.Export()
	if err != nil {
		t.Fatal(err)
	}
	clone := NewInventory()
	if err := clone.Import(data); err != nil {
		t.Fatal(err)
	}
	got, err := clone.Get("a")
	if err != nil || got.Quantity != 1 {
		t.Fatalf("Get(a) = %#v, %v", got, err)
	}
}

func TestInventoryErrorsRemainInspectable(t *testing.T) {
	inv := NewInventory()
	_, err := inv.Get("missing")
	if !errors.Is(err, ErrNotFound) {
		t.Fatalf("error = %v, want ErrNotFound", err)
	}
	if err := inv.Import([]byte(`[{"sku":"x","quantity":-1}]`)); err == nil {
		t.Fatal("Import accepted negative quantity")
	}
}
