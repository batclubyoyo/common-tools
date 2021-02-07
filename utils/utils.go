package utils

import (
	"errors"
	"fmt"
	"net"
)

/*
AbortAppRun AbortAppRun
*/
func AbortAppRun(err error) {
	if err != nil {
		panic(err)
	}
}

/*
StopAppRun StopAppRun
*/
func StopAppRun(msg string) {
	panic(msg)
}

/*
HandleError HandleError
*/
func HandleError(err error) {
	if err == nil {
		return
	}
	fmt.Printf("error:%s\n", err.Error())
}

/*
Percentage Percentage
*/
func Percentage(denominator, numerator int) (delta float64) {
	if denominator <= 0 {
		delta = 0
	}
	delta = (float64(numerator) / float64(denominator)) * 100
	return
}

/*
IPs get local ips, include eth, ens(wifi), dockers
*/
func IPs() (map[string]string, error) {
	ips := make(map[string]string)

	interfaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}

	for _, i := range interfaces {
		byName, err := net.InterfaceByName(i.Name)
		if err != nil {
			return nil, err
		}
		addresses, err := byName.Addrs()
		if err != nil {
			return nil, err
		}
		for _, v := range addresses {
			if ipnet, ok := v.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
				if ipnet.IP.To4() != nil {
					ips[byName.Name] = ipnet.IP.String()
				}
			}
		}
	}

	if len(ips) < 1 {
		return nil, errors.New("No Net Interface")
	}

	return ips, nil
}
