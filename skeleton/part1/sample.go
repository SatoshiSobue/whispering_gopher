package main

import (
	"fmt"
	"log"
	"net"
	"reflect"
)

func main() {
	c, err := net.Dial("tcp", "www.google.com:80")
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(reflect.TypeOf(c))
	fmt.Fprintln(c, "GET /")
	fmt.Println(reflect.TypeOf(c))
	// io.Copy(os.Stdout, c)
	// fmt.Fprintln(os.Stdout, c)
	c.Close()
}
