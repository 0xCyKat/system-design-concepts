package main

import (
	"fmt"
	"sd_concepts/unique_id_generator/id_generator"
)

func main() {
	id, err := id_generator.NextID()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(id)
}
