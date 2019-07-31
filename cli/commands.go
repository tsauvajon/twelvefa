package main

import (
	"fmt"
	"strings"
)

func (cli *calcli) help() {
	helpMsg := `add a b               Add two integers
max a b               Return the max of two integers
np a [b] [c]          Get the nth prime numbers
exit                  Quit the client`

	fmt.Println(helpMsg)
}

func (cli *calcli) max(a, b int64) {
	res, err := cli.client.Max(a, b)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(res.Max)
}

func (cli *calcli) add(a, b int64) {
	res, err := cli.client.Add(a, b)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(res.Sum)
}

func (cli *calcli) np(n []uint64) {
	res, err := cli.client.NthPrimes(n)

	if err != nil {
		fmt.Println(err)
		return
	}

	str := fmt.Sprint(res.NthPrimes)
	trimmed := strings.Trim(str, "[]")

	fmt.Println(trimmed)
}
