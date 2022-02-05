// package greeting_test ...
package greeting_test

import (
	"testing"

	"api/libs/api/greeting"
)

func TestGreeting(t *testing.T) {
	result := greeting.Greeting("works")
	if result != "Greeting works" {
		t.Error("Expected Greeting to append 'works'")
	}
}
