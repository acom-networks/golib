package acom

import (
	"fmt"
	"net"
	"net/http"
	"net/url"
	"time"
)

var timeout = time.Duration(2 * time.Second)

func dialTimeout(network, addr string) (net.Conn, error) {
	return net.DialTimeout(network, addr, timeout)
}

func SendSms(uri, msg string) (*http.Response, error) {
	transport := http.Transport{
		Dial: dialTimeout,
	}

	client := http.Client{
		Transport: &transport,
	}

	ourl := fmt.Sprintf("%s%s", uri, url.QueryEscape(msg))
	resp, err := client.Get(ourl)
	return resp, err
}
