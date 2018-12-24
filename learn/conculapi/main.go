package main

import (
	"fmt"
	"net"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/hashicorp/consul/api"
)

func main() {
	cl, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		fmt.Println(err)
	}
	c := cl.KV()
	c.Put(&api.KVPair{
		Key:   "name",
		Value: []byte("lyl"),
	}, nil)
	v, _, err := c.Get("name", nil)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(v.Key + ":" + string(v.Value))
	}

	agent := cl.Agent()
	err = agent.ServiceRegister(&api.AgentServiceRegistration{
		ID:      "lyl.consulapi",
		Name:    "lyl.consulapi",
		Tags:    []string{"learn", "api"},
		Address: "10.112.132.22",
		Port:    110,
		Check: &api.AgentServiceCheck{
			DeregisterCriticalServiceAfter: "1m",
			Interval:                       "10s",
			HTTP:                           "http://localhost:5000/health",
		},
	})
	if err != nil {
		fmt.Println(err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/health", func(http.ResponseWriter, *http.Request) {})
	http.ListenAndServe("localhost:5000", r)
}

func IPs() []string {
	ifaces, err := net.Interfaces()
	if err != nil {
		return nil
	}

	var ipAddrs []string

	for _, i := range ifaces {
		addrs, err := i.Addrs()
		if err != nil {
			continue
		}

		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}

			if ip == nil {
				continue
			}

			ip = ip.To4()
			if ip == nil {
				continue
			}

			ipAddrs = append(ipAddrs, ip.String())
		}
	}
	return ipAddrs
}
