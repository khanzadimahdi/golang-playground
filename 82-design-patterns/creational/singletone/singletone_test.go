package singletone

import "testing"

func TestGetInstance(t *testing.T) {
	firstInstance := GetInstance()
	secondInstance := GetInstance()

	if firstInstance == nil {
		t.Error("expects a singletone instance but receives nil")
	}

	if firstInstance != secondInstance {
		t.Error("expects same instances but received different instances")
	}
}
