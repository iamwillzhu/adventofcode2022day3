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
	Items             string
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
			Items:             ruckSackItems,
			FirstCompartment:  Compartment(firstCompartment),
			SecondCompartment: Compartment(secondCompartment),
		}

		ruckSackList = append(ruckSackList, ruckSack)
	}

	return
}

const ItemTypePriorityOrder = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func getItemTypeInBothCompartments(firstCompartment Compartment, secondCompartment Compartment) string {
	firstCompartmentItemTypeMap := make(map[string]bool)
	secondCompartmentItemTypeMap := make(map[string]bool)

	for _, itemType := range firstCompartment {
		firstCompartmentItemTypeMap[string(itemType)] = true
	}

	for _, itemType := range secondCompartment {
		secondCompartmentItemTypeMap[string(itemType)] = true
	}

	for itemType, _ := range firstCompartmentItemTypeMap {
		if _, ok := secondCompartmentItemTypeMap[itemType]; ok {
			return itemType
		}
	}

	return ""
}

func getItemTypeInRuckSacks(ruckSackOne *RuckSack, ruckSackTwo *RuckSack, ruckSackThree *RuckSack) string {
	ruckSackOneItemTypeMap := make(map[string]bool)
	ruckSackTwoItemTypeMap := make(map[string]bool)
	ruckSackThreeItemTypeMap := make(map[string]bool)

	for _, itemType := range ruckSackOne.Items {
		ruckSackOneItemTypeMap[string(itemType)] = true
	}

	for _, itemType := range ruckSackTwo.Items {
		ruckSackTwoItemTypeMap[string(itemType)] = true
	}

	for _, itemType := range ruckSackThree.Items {
		ruckSackThreeItemTypeMap[string(itemType)] = true
	}

	for itemType, _ := range ruckSackOneItemTypeMap {
		_, inRuckSackTwo := ruckSackTwoItemTypeMap[itemType]
		_, inRuckSackThree := ruckSackThreeItemTypeMap[itemType]

		if inRuckSackTwo && inRuckSackThree {
			return itemType
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
	sumOfItemTypeInRuckSacksPriority := 0

	for index, ruckSack := range ruckSackList {
		itemTypeInBothCompartments := getItemTypeInBothCompartments(ruckSack.FirstCompartment, ruckSack.SecondCompartment)

		if itemTypeInBothCompartments == "" {
			log.Panicf("No item type in both compartments of rucksack %d", index)
		}

		sumOfItemTypeInBothCompartmentsPriority += getItemTypePriority(itemTypeInBothCompartments)
		fmt.Printf("rucksack :%d, first compartment: %s, second compartment: %s, item type in both compartments: %s\n", index+1, ruckSack.FirstCompartment, ruckSack.SecondCompartment, itemTypeInBothCompartments)

		if index%3 == 2 && index >= 2 {
			itemTypeInRuckSacks := getItemTypeInRuckSacks(ruckSackList[index-2], ruckSackList[index-1], ruckSackList[index])
			fmt.Printf("rucksacks %d to %d badge item type is %s\n\n", index-1, index+1, itemTypeInRuckSacks)

			sumOfItemTypeInRuckSacksPriority += getItemTypePriority(itemTypeInRuckSacks)
		}
	}

	fmt.Printf("The sum of the priorities of the item types in both compartments for all rucksacks is %d\n", sumOfItemTypeInBothCompartmentsPriority)
	fmt.Printf("The sum of the priorities of the item types that are badges for rucksacks in groups of three is %d\n", sumOfItemTypeInRuckSacksPriority)
}
