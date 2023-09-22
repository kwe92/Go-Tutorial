package main

import "fmt"

// TODO: Add a "String() string" method to IPAddr.
// keys: loopback, google | values: 127,0,0,1 && 8,8,8,8 | example {8,8,8,8} should print 8.8.8.8

func main() {

	hosts := map[string]IpAddr{
		"loopback": {127, 0, 0, 1},
		"google":   {8, 8, 8, 8},
	}

	for k, v := range hosts {
		fmt.Printf("\nKey: %v | Value: %v\n", k, v)
	}

}

type IpAddr [4]byte

func (ip IpAddr) String() string {
	var str string

	for i, ele := range ip {
		if i == 0 {
			str += fmt.Sprint(ele)
		}
		if i > 0 {
			str += fmt.Sprintf(".%v", ele)

		}
	}
	return str
}
