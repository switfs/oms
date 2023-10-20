package main

import (
	"fmt"

	"github.com/switfs/oms/pkg/ipip"
)

func GetCityUsers(ip string) (local ipip.Location, err error) {
	city, err := ipip.NewCity("./mydata4vipday2.datx") // For City Level IP Database
	if err != nil {
		return ipip.Location{}, err
	}

	loc, err := city.FindLocation(ip)
	if err != nil {
		return ipip.Location{}, err
	}
	local = loc

	return local, nil
}

func main() {
	ipaddr := "119.163.28.110"
	ipctiy, err := GetCityUsers(ipaddr)
	if err != nil {
		return
	}

	fmt.Println(string(ipctiy.ToJSON()))

}
