package main

import (
	"fmt"
	"os"
	"os/exec"
	"pokedex/internal/locations"
)

// commandClear clears the terminal screen
func commandClear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

// commandExit exits the program
func commandExit() {
	os.Exit(0)
}

// commandMap fetch 20 location areas
func commandMap(cfg *locations.LocationPaginate, next bool) {
	var pageUrl string
	if cfg == nil {
		fmt.Println("error fetching locations areas: no config params")
		return
	}
	if next {
		if cfg.NextPage != nil {
			pageUrl = *cfg.NextPage
		} else {
			fmt.Println("Reached end, no more pages to fetch")
			return
		}
	} else {
		if cfg.PrevPage != nil {
			pageUrl = *cfg.PrevPage
		} else {
			fmt.Println("Already at start, no more areas to fetch")
			return
		}
	}

	areas, err := locations.GetBatchAreas(pageUrl, cfg)
	if err != nil {
		fmt.Println(err)
	}
	for _, area := range areas {
		fmt.Println(area)
	}
}
