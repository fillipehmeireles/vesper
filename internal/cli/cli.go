package cli

import (
	"flag"
	"errors"
	"fmt"
)

type CliArgs struct{
	NodeName string
	NodePort int
	Headless bool
}

func Usage() {
	fmt.Println("Flags:")
	fmt.Println("--name\t\t Node name.")
	fmt.Println("--headless\t\t set headless mode.")
}

func GetArgs() (CliArgs, error) {
	name := flag.String("name", "", "Node name")
	port := flag.Int("port", 0, "Node Port")
	headless := flag.Bool("headless", true, "Headless mode")
	flag.Parse()

	if *name == "" {
		Usage()
		return CliArgs{}, errors.New("Please provide the node name (--name/-n).")
	}

	if *port == 0 {
		Usage()
		return CliArgs{}, errors.New("Please provide the node port (--port/-p).")
	}

	return CliArgs{
		NodeName: *name,
		NodePort: *port,
		Headless: *headless,
	}, nil
}


