package cli

import (
	"flag"
	"errors"
	"fmt"
)


func Usage() {
	fmt.Println("Flags:")
	fmt.Println("--name\t\t Node name.")
}

func GeNameFlag() (string, error) {
	name := flag.String("name", "", "Node name")
	flag.Parse()
	if *name == "" {
		Usage()
		return "", errors.New("Please provide the node name (--name/-n).")
	}

	return *name, nil
}
