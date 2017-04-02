package exists

import "testing"

func TestExists(t *testing.T) {
	actual := true
	expected := true

	actual = File("exists_test.go")
	expected = true
	if actual != expected {
		t.Errorf("get %v\nwant %v", actual, expected)
	}

	actual = File("doesnotexists.go")
	expected = false
	if actual != expected {
		t.Errorf("get %v\nwant %v", actual, expected)
	}
}
