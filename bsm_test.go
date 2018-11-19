package bsm

import (
	"fmt"
	"testing"
	"time"
)

// BSMTest test for option price
func TestBSM(t *testing.T) {

	//Test option on AAPL with following data:
	// Current price: 186.28
	// Exp Date: 11/30/18
	// Strike Price: 190
	// Vol: 34.4
	// Call price: 2.90-2.93 (spread)
	date := time.Date(2018, time.November, 30, 0, 0, 0, 0, time.UTC)
	opt := New(186.28, 190.0, 34.4, true, "AAPL", date)

	price := opt.Value()
	fmt.Println(price)
	if price != 2.90 {
		t.Logf("Call Price: %f", price)
		t.Error("Option returns wrong value")
	}

}
