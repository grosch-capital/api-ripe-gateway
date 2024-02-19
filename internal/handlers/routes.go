package handlers

import (
	"fmt"
	"net"
	"time"

	"golang.org/x/net/ipv4"
)

func Traceroute(destination string) ([]string, error) {
	var hops []string

	ipAddr, err := net.ResolveIPAddr("ip", destination)
	if err != nil {
		return nil, err
	}

	conn, err := net.Dial("ip:icmp", ipAddr.String())
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	err = conn.SetDeadline(time.Now().Add(5 * time.Second))
	if err != nil {
		return nil, err
	}

	for ttl := 1; ttl <= 30; ttl++ {

		ipv4Conn := ipv4.NewConn(conn)
		err = ipv4Conn.SetTTL(ttl)
		if err != nil {
			return nil, err
		}

		start := time.Now()
		_, err = conn.Write([]byte("HELLO"))
		if err != nil {
			return nil, err
		}

		// buffer := make([]byte, 1024)
		// n, err := conn.Read(buffer)
		if err != nil {
			return nil, err
		}

		duration := time.Since(start)
		hop := fmt.Sprintf("%d. %s (%s) %s", ttl, ipAddr.String(), conn.RemoteAddr().String(), duration)
		hops = append(hops, hop)

		if ipAddr.String() == conn.RemoteAddr().String() {
			break
		}
	}

	return hops, nil
}
