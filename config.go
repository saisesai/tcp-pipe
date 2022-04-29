package main

import (
	"flag"
	"log"
)

var cfgLocalAddr = flag.String("local", "0.0.0.0:13001", "local listen address")
var cfgEndpointAddr = flag.String("end", "0.0.0.0:13000", "endpoint address")
var cfgDialRetryInterval = flag.Int("retry", 1, "retry interval when dial to endpoint failed in seconds")

func init() {
	flag.Parse()
	log.Println("local listen address:", *cfgLocalAddr)
	log.Println("endpoint address:", *cfgEndpointAddr)
	log.Println("using \"-help\" to see usage!")
}
