package good

import "testing"

func TestConfigure(t *testing.T) {
	Configure("./app.yml", false)
}
