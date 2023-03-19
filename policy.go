package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {

	byte, err := os.ReadFile("redflags.json")
	if err != nil {
		log.Fatal("error: ", err)
	}

	var data map[string]interface{}
	err = json.Unmarshal(byte, &data)
	if err != nil {
		log.Fatal("error: ", err)
	}

	ref := make(map[interface{}]interface{})

	items := data["items"].([]interface{})
	for _, item := range items {
		var id = item.(map[string]interface{})["ref"]
		ref[id] = ""
	}

	// part 2 --------------------------------------

	byte1, err := os.ReadFile("processes.json")
	if err != nil {
		log.Fatal("error: ", err)
	}
	var data2 map[string]interface{}

	err = json.Unmarshal(byte1, &data2)
	if err != nil {
		log.Fatal("error: ", err)
	}

	items2 := data2["items"].([]interface{})
	for _, item := range items2 {
		for key := range ref {
			// fmt.Println("I'm refernce", key)
			itemm := item.(map[string]interface{})["id"]
			// fmt.Println("I'm id", itemm)
			if key == itemm {
				cgroup := item.(map[string]interface{})["cgroup"]
				ref[key] = cgroup
			}
		}
	}

	// part 3-----------------------------------

	policymap := map[interface{}]int{}
	counter := 1

	for _, policy := range ref {
		if _, ok := policymap[policy]; ok {
			counter += 1
			policymap[policy] = counter
		} else {
			policymap[policy] = counter
		}
	}
	for policies, count := range policymap {
		fmt.Println(policies, count)
		fmt.Println()
	}
}
