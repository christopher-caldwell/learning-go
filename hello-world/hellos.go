package main

import "fmt"

func my_hellos(names []string) (map[string]string, error) {
	messages_map := make(map[string]string)

	for index, name := range names {
		message, error := hello(name)

		fmt.Print(index)

		if error != nil {
			return nil, error
		}

		messages_map[name] = message
	}

	return messages_map, nil
}
