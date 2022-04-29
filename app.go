package main

import (
	"io"
	"log"
	"net"
)

func main() {
	local, err := net.Listen("tcp", *cfgLocalAddr)
	if err != nil {
		log.Fatalln(err)
	}
	for {
		conn, err := local.Accept()
		if err != nil {
			log.Println(err)
		}
		go HandleConn(conn)
	}
}

func HandleConn(conn net.Conn) {
	log.Println("client", conn.RemoteAddr(), "connected!")
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			log.Println(err)
		}
		log.Println("client", conn.RemoteAddr(), "closed!")
	}(conn)

	end, err := net.Dial("tcp", *cfgEndpointAddr)
	if err != nil {
		log.Println(err)
		return
	}
	go func(conn, end net.Conn) {
		log.Println("endpoint", end.LocalAddr(), "connected!")
		defer func(end net.Conn) {
			err := end.Close()
			if err != nil {
				log.Println(end)
			}
			log.Println("endpoint", end.LocalAddr(), "closed!")
		}(end)
		_, err = io.Copy(end, conn)
		if err != nil {
			log.Println(err)
		}
	}(conn, end)

	_, err = io.Copy(conn, end)
	if err != nil {
		log.Println(err)
	}
}
