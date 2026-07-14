package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"sort"
)

var ErrNotFound = errors.New("item not found")

type Item struct {
	SKU      string `json:"sku"`
	Quantity int    `json:"quantity"`
}

type Inventory struct {
	items map[string]Item
}

func NewInventory() *Inventory {
	return &Inventory{items: make(map[string]Item)}
}

func (i *Inventory) Put(item Item) error {
	if item.SKU == "" {
		return fmt.Errorf("put item: sku is required")
	}
	if item.Quantity < 0 {
		return fmt.Errorf("put %q: quantity must be non-negative", item.SKU)
	}
	i.items[item.SKU] = item
	return nil
}

func (i *Inventory) Get(sku string) (Item, error) {
	item, ok := i.items[sku]
	if !ok {
		return Item{}, fmt.Errorf("get %q: %w", sku, ErrNotFound)
	}
	return item, nil
}

func (i *Inventory) UpdateQuantity(sku string, quantity int) error {
	if quantity < 0 {
		return fmt.Errorf("update %q: quantity must be non-negative", sku)
	}
	item, err := i.Get(sku)
	if err != nil {
		return fmt.Errorf("update quantity: %w", err)
	}
	item.Quantity = quantity
	i.items[sku] = item
	return nil
}

func (i *Inventory) Delete(sku string) error {
	if _, err := i.Get(sku); err != nil {
		return fmt.Errorf("delete item: %w", err)
	}
	delete(i.items, sku)
	return nil
}

func (i *Inventory) Export() ([]byte, error) {
	items := make([]Item, 0, len(i.items))
	for _, item := range i.items {
		items = append(items, item)
	}
	sort.Slice(items, func(a, b int) bool { return items[a].SKU < items[b].SKU })
	return json.MarshalIndent(items, "", "  ")
}

func (i *Inventory) Import(data []byte) error {
	var items []Item
	if err := json.Unmarshal(data, &items); err != nil {
		return fmt.Errorf("import inventory: %w", err)
	}
	next := make(map[string]Item, len(items))
	for _, item := range items {
		if item.SKU == "" || item.Quantity < 0 {
			return fmt.Errorf("import inventory: invalid item %#v", item)
		}
		if _, exists := next[item.SKU]; exists {
			return fmt.Errorf("import inventory: duplicate sku %q", item.SKU)
		}
		next[item.SKU] = item
	}
	i.items = next
	return nil
}

func main() {
	inv := NewInventory()
	_ = inv.Put(Item{SKU: "book-go", Quantity: 4})
	data, _ := inv.Export()
	fmt.Println(string(data))
}
