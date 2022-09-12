package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
)

func main() {
	// import file and check that it exists without error
	file, err := os.Open("Section002.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	// Create a scanner for the file
	scanner := bufio.NewScanner(file)

	// Create an array to store all of the student names, then iterate through the scanner and add each student to the array
	var student_array []string

	student_count := 0

	for scanner.Scan() {
		line := scanner.Text()
		student_array = append(student_array, line)
		student_count++
	}

	// get the team size from user input and then calculate the number of teams needed
	var teamSize int

	fmt.Print("What is the max group size?: ")
	fmt.Scan(&teamSize)

	numGroups := student_count/teamSize + 1

	// Create an array of arrays and add an empty array for each group
	var groups_array [][]string

	i := 0

	for i < numGroups {
		var group []string
		groups_array = append(groups_array, group)
		i++
	}

	// for each student, randomly select a group to add them to, if that group is already maxed out then pick a new group
	i = 0

	for i < student_count {
		groupIndex := rand.Intn(numGroups)
		for len(groups_array[groupIndex]) >= teamSize {
			groupIndex = rand.Intn(numGroups)
		}
		groups_array[groupIndex] = append(groups_array[groupIndex], student_array[i])
		i++
	}

	// print the results
	i = 0

	for i < len(groups_array) {
		fmt.Println("--- Group", i+1, "---", "( Size:", len(groups_array[i]), ")")
		n := 0
		for n < len(groups_array[i]) {
			fmt.Println(groups_array[i][n])
			n++
		}
		fmt.Println()
		i++
	}
}
