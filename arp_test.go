package arp

import "testing"

func TestARP(t *testing.T) {
	list, err := GetEntries()
	if err != nil {
		t.Error(err)
	}
	if len(list) <= 0 {
		t.Errorf("Invalid length (got %d expected >0)", len(list))
	}
}
