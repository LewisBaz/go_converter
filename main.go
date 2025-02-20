package main

import (
	"flag"
	"fmt"
	"go_converter/network"
	"os"
)

func main() {

	convertFlagSet := flag.NewFlagSet("cnv", flag.ExitOnError)
	var cnvFromValue string
	var cnvToValue string
	var amountToConvert float64
	convertFlagSet.StringVar(&cnvFromValue, "from", "USD", "The currency to convert from")
	convertFlagSet.StringVar(&cnvToValue, "to", "EUR", "The currency to convert to")
	convertFlagSet.Float64Var(&amountToConvert, "a", 1, "Amount of currency to convert")

	if len(os.Args) < 2 {
        fmt.Println("expected 'cnv' or 'test' subcommands")
        os.Exit(1)
    }

	switch os.Args[1] {
	case "cnv":
		convertFlagSet.Parse(os.Args[2:])
		convertFlagSet.Parse(os.Args[4:])
		convertFlagSet.Parse(os.Args[6:])

		request := network.Request{
			Amount: amountToConvert, 
			From: cnvFromValue,
			To: cnvToValue,
		}

		result, err := network.MakeRequest(request)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(result.Result)
	default:
		fmt.Println("Unknown command")
		os.Exit(1)
	}
}