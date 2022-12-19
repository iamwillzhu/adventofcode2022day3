package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

type Compartment string

type RuckSack struct {
	FirstCompartment  Compartment
	SecondCompartment Compartment
}

func getRuckSackList(reader io.Reader) (ruckSackList []*RuckSack) {
	ruckSackList = make([]*RuckSack, 0)

	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		ruckSackItems := scanner.Text()
		numberItems := len(ruckSackItems)
		midIndex := numberItems / 2

		firstCompartment := ruckSackItems[:midIndex]
		secondCompartment := ruckSackItems[midIndex:]

		ruckSack := &RuckSack{
			FirstCompartment:  Compartment(firstCompartment),
			SecondCompartment: Compartment(secondCompartment),
		}

		ruckSackList = append(ruckSackList, ruckSack)
	}

	return
}

const ItemTypePriorityOrder = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func getItemTypeInBothCompartments(firstCompartment Compartment, secondCompartment Compartment) string {
	firstCompartmentCharacterMap := make(map[string]bool)
	secondCompartmentCharacterMap := make(map[string]bool)

	for _, character := range firstCompartment {
		firstCompartmentCharacterMap[string(character)] = true
	}

	for _, character := range secondCompartment {
		secondCompartmentCharacterMap[string(character)] = true
	}

	for character, _ := range firstCompartmentCharacterMap {
		if _, ok := secondCompartmentCharacterMap[character]; ok {
			return character
		}
	}

	return ""
}

func getItemTypePriority(itemType string) int {
	for index, character := range ItemTypePriorityOrder {
		if string(character) == itemType {
			return index + 1
		}
	}

	return -1
}

func main() {
	file, err := os.Open("/home/ec2-user/go/src/github.com/iamwillzhu/adventofcode2022day3/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	ruckSackList := getRuckSackList(file)
	sumOfItemTypeInBothCompartmentsPriority := 0

	for index, ruckSack := range ruckSackList {
		itemTypeInBothCompartments := getItemTypeInBothCompartments(ruckSack.FirstCompartment, ruckSack.SecondCompartment)

		if itemTypeInBothCompartments == "" {
			log.Fatalf("No item type in both compartments of rucksack %d", index)
			continue
		}

		sumOfItemTypeInBothCompartmentsPriority += getItemTypePriority(itemTypeInBothCompartments)
		fmt.Printf("rucksack :%d, first compartment: %s, second compartment: %s, item type in both compartments: %s\n", index, ruckSack.FirstCompartment, ruckSack.SecondCompartment, itemTypeInBothCompartments)
	}

	fmt.Printf("The sum of the priorities of the above item types is %d\n", sumOfItemTypeInBothCompartmentsPriority)
}
