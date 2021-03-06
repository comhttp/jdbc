package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/comhttp/jdbc"
)

func check(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s\n", err.Error())
		os.Exit(1)
	}
}

func main() {
	endpoint := flag.String("endpoint", "http://localhost:4338", "Address:port to connect to")
	auth := flag.String("auth", "", "Optional Authorization string (for stulbe)")
	command := flag.String("command", "", "Command to run (supported: kget/kset)")
	key := flag.String("key", "", "Key to run command on")
	data := flag.String("data", "", "Optional data argument for commands that require it")
	flag.Parse()

	if *command == "" {
		check(fmt.Errorf("must specify a valid -command"))
	}
	if *key == "" {
		check(fmt.Errorf("must specify a valid -key"))
	}

	headers := http.Header{}
	if *auth != "" {
		headers.Add("Authorization", "Bearer "+*auth)
	}

	client, err := jdbc.NewClient(*endpoint, jdbc.ClientOptions{Headers: headers})
	check(err)

	switch strings.ToLower(*command) {
	case "kget":
		str, err := client.GetKey(*key)
		check(err)
		fmt.Println(str)
	case "kset":
		check(client.SetKey(*key, *data))
	default:
		check(fmt.Errorf("unknown command \"%s\"", *command))
	}

}
