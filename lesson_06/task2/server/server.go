package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
)

func main(){
	fmt.Println("Server...")
	ln, err := net.Listen("tcp", ":8081")
	if err != nil{
		log.Fatal()
	}
	conn, _ := ln.Accept()
	for {
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Println("Message:", message)
		r := []rune(message)
		r = r[:len(r)-1]
		message = string(r)
		var newMessage string
		i, err := strconv.Atoi(message)
		if err == nil{
			newMessage = strconv.Itoa(i*2)
		} else {
			newMessage = strings.ToUpper(message)
		}
		_, err = conn.Write([]byte(newMessage + "\n"))
		if err != nil{
			log.Print("Error")
		}
	}
}
