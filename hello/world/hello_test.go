package world

import "testing"

func TestPrintSecEle(t *testing.T) {
	ll := []string{"a", "b", "c"}
	output := PrintSecELe(ll)
	if output != "b" {
		t.Error("output should be b, is", output)
	}
}
