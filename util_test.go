package gocybos

import "testing"

func TestIsUserAnAdmin(t *testing.T) {
	v, err := IsUserAnAdmin()
	if err != nil {
		panic(err)
	}
	t.Logf("IsAdmin: %v", v)
}
