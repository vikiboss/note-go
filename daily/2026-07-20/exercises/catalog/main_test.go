package main

import "testing"

func TestCatalogSearchIsSortedAndDefensive(t *testing.T) {
	var c Catalog
	inputTags := []string{"language"}
	for _, item := range []Item{
		{Name: "Zulu Go", Stock: 2, Tags: inputTags},
		{Name: "Alpha Go", Stock: 1},
		{Name: "No stock Go", Stock: 0},
	} {
		if err := c.Add(item); err != nil {
			t.Fatal(err)
		}
	}
	inputTags[0] = "changed"

	got := c.Search("go", 1)
	if len(got) != 2 || got[0].Name != "Alpha Go" || got[1].Name != "Zulu Go" {
		t.Fatalf("unexpected result: %#v", got)
	}
	got[1].Tags[0] = "mutated"
	if c.Items()[2].Tags[0] != "language" {
		t.Fatal("caller mutated catalog through returned slice")
	}
}

func TestCatalogRejectsInvalidItem(t *testing.T) {
	tests := []Item{{}, {Name: "book", Stock: -1}}
	for _, item := range tests {
		if err := new(Catalog).Add(item); err == nil {
			t.Fatalf("Add(%#v) returned nil error", item)
		}
	}
}
