package main

import (
	"fmt"
)

func commandExplore(cfg *config, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("please provide a location to explore")
	}
	locationName := args[0]

	// location, err := cfg.pokeapiClient.GetLocation(locationName)
	// if err != nil {
	// 	return err
	// }

	fmt.Printf("Exploring location: %s\n", locationName)
	return nil
}
