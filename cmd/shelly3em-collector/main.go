package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/r10r/shelly-prometheus/pkg/devices"
)

func getNodeStatus(nodeURL string) (*devices.Shelly3EM, error) {
	resp, err := http.Get(nodeURL)
	if err != nil {
		return nil, err
	}
	defer func() {
		if resp.Body != nil {
			io.Copy(io.Discard, resp.Body)
			resp.Body.Close()
		}
	}()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected response status: %d %s", resp.StatusCode, resp.Status)
	}

	data, err := ioutil.ReadAll(resp.Body)
	node := &devices.Shelly3EM{}
	err = json.Unmarshal(data, node)
	if err != nil {
		return nil, err
	}
	return node, nil
}

func main() {
	var address string
	flag.StringVar(&address, "address", "http://127.0.0.1", "node address")
	flag.Parse()

	u, err := url.Parse(address)
	if err != nil {
		panic(err)
	}
	u.Path = "/status"
	statusURL := u.String()

	ticker := time.NewTicker(time.Second)
	for {
		select {
		case <-ticker.C:
			node, err := getNodeStatus(statusURL)
			if err != nil {
				log.Printf("%s", err)
				continue
			}
			fmt.Printf("%#v\n", node)
			fmt.Println(node.IP)
			for i, em := range node.Emeters {
				fmt.Printf("%d %.2f (%.2fV %.2fA %.2f%%)\n", i, em.Power, em.Voltage, em.Current, em.PowerFactor)
			}
		}
	}
}
