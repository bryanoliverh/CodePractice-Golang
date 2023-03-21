package main

import "fmt"

type Person struct {
    name     string
    age      int
    location string
    next     *Person
}

func main() {
    // Create a map of people with their ID as key and a pointer to their Person object as value
    people := make(map[int]*Person)

    // Add people to the linked list and map
    addPerson(&people, "John", 17, "New York")
    addPerson(&people, "Jane", 22, "Los Angeles")
    addPerson(&people, "Bob", 30, "Chicago")

    // Print all the people
    printPeople(&people)

    // Print the name and age of people above the age of 2
    printPeopleAboveAge(&people, 2)
}

// Add a new person to the linked list and map
func addPerson(people *map[int]*Person, name string, age int, location string) {
    newPerson := &Person{
        name:     name,
        age:      age,
        location: location,
    }

    // Find the next available ID for the person
    id := 1
    for _, exists := (*people)[id]; exists; _, exists = (*people)[id] {
        id++
    }

    // Add the person to the linked list
    if len(*people) == 0 {
        (*people)[id] = newPerson
    } else {
        // Find the last person in the linked list
        lastPerson := (*people)[1]
        for lastPerson.next != nil {
            lastPerson = lastPerson.next
        }

        // Add the new person to the end of the linked list
        lastPerson.next = newPerson
    }

    // Add the person to the map with their ID as key
    (*people)[id] = newPerson
}

// Print all the people in the linked list and map
func printPeople(people *map[int]*Person) {
	fmt.Printf("Name of All Registered People\n")
    for id, person := range *people {
        fmt.Printf("ID: %d, Name: %s, Age: %d, Location: %s\n", id, person.name, person.age, person.location)
    }
}

// Print the name and age of people above the given age
func printPeopleAboveAge(people *map[int]*Person, age int) {
    for _, person := range *people {
		fmt.Printf("Name of People Above 20 yo:")
        if person.age > age {
            fmt.Printf("Name: %s, Age: %d\n", person.name, person.age)
        }
    }
}
