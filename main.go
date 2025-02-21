package main

import (
	"flag"
	"fmt"
	pcnm "go_converter/network/routes/pair" 
	"net/http"
	"os"
)

var client = http.Client{}

func main() {
	convertFlagSet := flag.NewFlagSet("convert", flag.ExitOnError)
	var cnvFromValue string
	var cnvToValue string
	var amountToConvert float64
	convertFlagSet.StringVar(&cnvFromValue, "from", "USD", "The currency to convert from")
	convertFlagSet.StringVar(&cnvToValue, "to", "EUR", "The currency to convert to")
	convertFlagSet.Float64Var(&amountToConvert, "a", 1, "Amount of currency to convert")

	listFlag := flag.String("list", "", "")

	flag.Parse()

	if len(os.Args) < 2 {
        fmt.Println("expected 'convert' or 'list' subcommands")
        os.Exit(1)
    }

	switch os.Args[1] {
	case "convert":
		convertFlagSet.Parse(os.Args[2:])
		convertFlagSet.Parse(os.Args[4:])
		convertFlagSet.Parse(os.Args[6:])

		result, err := pcnm.MakeRequest(
			pcnm.PairCurrencyNM{Client: client}, 
			pcnm.Request{
				Amount: amountToConvert, 
				From: cnvFromValue,
				To: cnvToValue,
		})
		if err != nil {
			fmt.Println(err)
			return 
		}

		fmt.Println(result.Result)
	case "list":
		fmt.Println("list requested", listFlag)
	default:
		fmt.Println("Unknown command")
		os.Exit(1)
	}
}