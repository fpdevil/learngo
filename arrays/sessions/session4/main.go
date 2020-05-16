package main

import "fmt"

func main() {
	// keyed elements
	values := [...]float64{
		0: 0.5,
		5: 2.5,
		1.75,
		2: 1.5,
	}
	fmt.Printf("%#v\n", values)

	// crypto currency exchange ratios
	// without keyed indexes
	// rates := [...]float64{
	// 	120.5, // wanchain
	// 	25.5,  // ethereum
	// }

	// with keyed indexes
	// const (
	// 	ETH = iota // initialize to 0
	// 	WAN        // auto intialized to 1 now
	// )

	// rates := [...]float64{
	// 	WAN: 120.5, // wanchain
	// 	ETH: 25.5,  // ethereum
	// }

	// additional currencies...
	const (
		ETH = 9 - iota // initialize to 9
		WAN            // auto intialized to 8 now
		ICX            // auto initialized to 7 now
	)

	rates := [...]float64{
		WAN: 120.5, // wanchain
		ETH: 25.5,  // ethereum
		ICX: 20,
	}

	fmt.Printf("1 BTC is %g ETH\n", rates[ETH])
	fmt.Printf("1 BTC is %g WAN\n", rates[WAN])
	fmt.Printf("Crypto Currency Rates: %#v\n", rates)
}
