package cmd

import (
	"fmt"
	"strconv"
)

func argsToIds(args []string) []int {
	var ids []int
	for _, arg := range args {
		id, err := strconv.Atoi(arg)
		if err != nil {
			fmt.Println("Failed to parse argument:", arg)
		} else {
			ids = append(ids, id)
		}
	}
	return ids
}
