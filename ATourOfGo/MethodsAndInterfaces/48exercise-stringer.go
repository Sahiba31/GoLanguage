package main

import "fmt"

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.

func (ip IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", ip[0], ip[1], ip[2], ip[3])
}

// func (ip IPAddr) String() string {
// 	tmp := make([]interface{}, len(ip))
// 	for i, val := range ip {
// 		tmp[i] = val
// 	}
// 	return fmt.Sprintf("%d.%d.%d.%d", tmp...)
// }

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
