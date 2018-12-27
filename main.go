package main

import (
	"flag"
	"io/ioutil"
	"log"
	"strconv"
	"strings"

	"github.com/google/tcpproxy"
	"gopkg.in/yaml.v2"
)

type HostConfig struct {
	Hostname string
	Port_set []struct {
		Type       string
		PortNumber int `yaml:"port_number"`
	}
}

var Config []HostConfig

func init() {
	var configFilePath string
	flag.StringVar(&configFilePath, "c", "config.yml", "Path of Abghand config file.")
	flag.Parse()

	configFileContent, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		log.Fatal(err)
	}

	Config = make([]HostConfig, 0)
	if err := yaml.Unmarshal(configFileContent, &Config); err != nil {
		log.Fatal(err)
	}
}

func main() {
	var server tcpproxy.Proxy
	for _, host := range Config {
		if len(host.Port_set) == 0 {
			server.AddHTTPHostRoute(":80", host.Hostname, tcpproxy.To(host.Hostname+":80"))
			server.AddSNIRoute(":443", host.Hostname, tcpproxy.To(host.Hostname+":443"))
		} else {
			for _, port := range host.Port_set {
				switch strings.ToLower(port.Type) {
				case "http":
					server.AddHTTPHostRoute(
						":"+strconv.Itoa(port.PortNumber),
						host.Hostname,
						tcpproxy.To(host.Hostname+":"+strconv.Itoa(port.PortNumber)))
				case "https":
					server.AddSNIRoute(":"+strconv.Itoa(port.PortNumber),
						host.Hostname,
						tcpproxy.To(host.Hostname+":"+strconv.Itoa(port.PortNumber)))
				}
			}
		}
	}
	log.Println("Starting Abghand ...")
	log.Fatal(server.Run())
}
