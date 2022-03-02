package hello

import "testing"

func TestHello(t *testing.T ){

	expect :="Hello, world."

	if got := Hello(); got != expect {
		t.Errorf("expecting %s got %s",expect,got)
	}
	
}