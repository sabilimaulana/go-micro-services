package data

import "testing"

func TestChecksValidaton(t *testing.T) {
	p := &Product{
		Name:  "Tea",
		Price: 2.00,
		SKU:   "abc-def-asw",
	}

	err := p.Validate()
	if err != nil {
		t.Fatal(err)
	}
}
