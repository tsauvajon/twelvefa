package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/tsauvajon/twelvefa/calc"
	"google.golang.org/grpc"
)

/*
first draft of the API:

cli help                 // displays a help message

cli connect $ADDRESS     // connects to the server, opens the interactive mode:
> Add 2 3                // 5
> Max 7 2                // 7
> NthPrimes 12 9658 789  // 37 100801 6047
> exit

*/

func init() {
	// print all logs to stdout for streaming, using Docker logs...
	log.SetOutput(os.Stdout)
}

type calcli struct {
	client *client
}

func main() {
	argsWithoutProg := os.Args[1:]

	cmd := argsWithoutProg[0]

	// two choices here:
	// cli connect address => open interactive mode
	// anything else => display the help menu and quit
	if len(argsWithoutProg) != 2 || cmd != "connect" {
		help()
		return
	}

	cli := &calcli{
		client: &client{
			address: argsWithoutProg[1],
		},
	}

	conn, err := grpc.Dial(cli.client.address,
		grpc.WithBlock(),                // check we can reach the server
		grpc.WithTimeout(5*time.Second), // but not forever
		grpc.WithInsecure(),
	)

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	cli.client.calc = calc.NewCalcClient(conn)

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("calc> ")

		// read next line
		scanner.Scan()
		input := scanner.Text()

		args := strings.Split(input, " ")

		switch args[0] {
		case "exit":
			return
		case "":
			// if the user just wants to add some whitespace
			break
		case "max":
			if len(args) != 3 {
				fmt.Printf("expected 2 parameters\n")
				break
			}
			a, err := strconv.ParseInt(args[1], 10, 64)
			if err != nil {
				fmt.Printf("%s is not an int64\n", args[1])
				break
			}
			b, err := strconv.ParseInt(args[2], 10, 64)
			if err != nil {
				fmt.Printf("%s is not an int64\n", args[2])
				break
			}
			cli.max(a, b)
		case "add":
			if len(args) != 3 {
				fmt.Printf("expected 2 parameters\n")
				break
			}
			a, err := strconv.ParseInt(args[1], 10, 64)
			if err != nil {
				fmt.Printf("%s is not an int64\n", args[1])
				break
			}
			b, err := strconv.ParseInt(args[2], 10, 64)
			if err != nil {
				fmt.Printf("%s is not an int64\n", args[2])
				break
			}
			cli.add(a, b)
		case "np":
			if len(args) == 1 {
				fmt.Printf("expected at least 1 parameter\n")
				break
			}

			params := []uint64{}

			for i := 1; i < len(args); i++ {
				n, err := strconv.ParseUint(args[i], 10, 64)
				if err != nil {
					fmt.Printf("%s is not an int64\n", args[1])
					break
				}
				params = append(params, n)
			}
			cli.np(params)
		default:
			cli.help()
		}
	}
}

func help() {
	helpMsg := `cli connect address    Open a connection to the calc service
cli help                Displays this message`

	fmt.Println(helpMsg)
}
