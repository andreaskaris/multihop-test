package main

import (
	"encoding/json"
	"fmt"

	"github.com/vishvananda/netlink"
)

func main() {
	filter := &netlink.Route{Dst: nil}
	mask := netlink.RT_FILTER_DST
	// filter.LinkIndex = 2
	// mask = mask | netlink.RT_FILTER_OIF

	fmt.Println("IPv4")
	family := 2
	routes, _ := netlink.RouteListFiltered(family, filter, mask)
	rb, _ := json.Marshal(routes)
	fmt.Println()
	fmt.Println("result: ", string(rb))

	fmt.Println("===================================")

	fmt.Println("IPv6")
	family = 10
	routes, _ = netlink.RouteListFiltered(family, filter, mask)
	rb, _ = json.Marshal(routes)
	fmt.Println()
	fmt.Println("result: ", string(rb))
}
