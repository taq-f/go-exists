package exists

import "testing"

func TestExists(t *testing.T) {
	actual := true
	expected := true

	actual = Exists("exists_test.go")
	expected = true
	if actual != expected {
		t.Errorf("get %v\nwant %v", actual, expected)
	}

	actual = Exists("doesnotexists.go")
	expected = false
	if actual != expected {
		t.Errorf("get %v\nwant %v", actual, expected)
	}
}
