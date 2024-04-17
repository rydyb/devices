package metzconnect

import (
	"os"
	"testing"

	"github.com/goburrow/modbus"
)

var address = os.Getenv("ADDRESS")

func TestEWIO2_AnalogInput(t *testing.T) {
	ewio2 := EWIO2{modbus.TCPClient(address)}

	for i := 1; i <= 3; i++ {
		voltage, err := ewio2.AnalogInput(i)
		if err != nil {
			t.Fatalf("failed to read analog input: %s", err)
		}

		if voltage < 0.0 || voltage > 10.0 {
			t.Errorf("voltage out of range: %f", voltage)
		}
	}
}