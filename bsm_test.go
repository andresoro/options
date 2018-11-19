package bsm_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/andresoro/options"
)

// BSMTest test for option price
func TestBSM(t *testing.T) {

	//Test option on SPY with following data:
	// Current price: 268.97
	// Exp Date: 11/30/18
	// Strike Price: 270
	// Vol: 19.3
	// Call price: 3.21-3.24 (spread)
	today := time.Now().Local()
	date := time.Date(2018, time.November, 30, 0, 0, 0, 0, time.UTC)
	opt := bsm.New(268.97, 270, 0.193, true, "SPY", date)

	price := opt.Value()
	fmt.Println(price)
	if price != 3.21 {
		t.Log(today)
		t.Log(date)
		t.Logf("Call Price: %f", price)
		t.Error("Option returns wrong value")
	}

}
