package db

import (
	"fmt"
	"strings"
)

// StdinConfigs gets a config from input on the command line.
func StdinConfigs() ([]string, error) {
	fmt.Print("enter your postgres config(s): ")

	confs := ""

	_, err := fmt.Scanln(&confs)
	if err != nil {
		err = fmt.Errorf("scanln confs: %w", err)

		Warn.Println(err)

		return nil, err
	}

	return strings.Split(confs, ":"), nil
}