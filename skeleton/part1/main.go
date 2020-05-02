// Solution to part 1 of the Whispering Gophers code lab.
// This program reads from standard input and writes JSON-encoded messages to
// standard output. For example, this input line:
//	Hello!
// Produces this output:
//	{"Body":"Hello!"}
//
package main

import (
	"bufio"
	"encoding/json"
	"log"
	"os"
)

type Message struct {
	Body string
}

func main() {
	s := bufio.NewScanner(os.Stdin)   // TODO: Create a new bufio.Scanner reading from the standard input.
	enc := json.NewEncoder(os.Stdout) // TODO: Create a new json.Encoder writing into the standard output.
	for s.Scan() /* TODO: Iterate over every line in the scanner */ {
		m := Message{Body: s.Text()} // TODO: Create a new message with the read text.
		err := enc.Encode(m)
		if err != nil {
			log.Fatal(err)
		} // TODO: Encode the message, and check for errors!
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	} // TODO: Check for a scan error.
}
