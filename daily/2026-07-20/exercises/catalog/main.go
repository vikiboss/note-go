package main

import (
	"fmt"
	"sort"
	"strings"
)

type Item struct {
	Name  string
	Stock int
	Tags  []string
}

type Catalog struct {
	items []Item
}

func (c *Catalog) Add(item Item) error {
	item.Name = strings.TrimSpace(item.Name)
	if item.Name == "" {
		return fmt.Errorf("name is required")
	}
	if item.Stock < 0 {
		return fmt.Errorf("stock must be non-negative")
	}
	item.Tags = append([]string(nil), item.Tags...)
	c.items = append(c.items, item)
	return nil
}

func (c *Catalog) Search(query string, minStock int) []Item {
	query = strings.ToLower(strings.TrimSpace(query))
	var result []Item
	for _, item := range c.items {
		if item.Stock < minStock || !strings.Contains(strings.ToLower(item.Name), query) {
			continue
		}
		copyItem := item
		copyItem.Tags = append([]string(nil), item.Tags...)
		result = append(result, copyItem)
	}
	sort.Slice(result, func(i, j int) bool { return result[i].Name < result[j].Name })
	return result
}

func (c *Catalog) Items() []Item {
	return c.Search("", 0)
}

func main() {
	var c Catalog
	_ = c.Add(Item{Name: "Go 语言设计与实现", Stock: 3, Tags: []string{"go", "systems"}})
	_ = c.Add(Item{Name: "Concurrency in Go", Stock: 1, Tags: []string{"go", "concurrency"}})
	for _, item := range c.Search("go", 1) {
		fmt.Printf("%s (stock=%d)\n", item.Name, item.Stock)
	}
}
