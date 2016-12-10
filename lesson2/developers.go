package main

import (
	"fmt"
	"strings"
)

type developer struct {
	manager    string
	department string
	team       string
	floor      int
}

var developers = map[string]developer{
	"James": developer{
		manager:    "Rudolf",
		department: "Development",
		team:       "DAS",
		floor:      5,
	},
	"Kamil": developer{
		manager:    "Rudolf",
		department: "Development",
		team:       "DAS",
		floor:      5,
	},
	"Chris": developer{
		manager:    "Charlie",
		department: "Development",
		team:       "Product",
		floor:      5,
	},
	"Andrew": developer{
		manager:    "Charlie",
		department: "Development",
		team:       "Product",
		floor:      5,
	},
	"Gijs": developer{
		manager:    "Frazer",
		department: "Project Management",
		team:       "Product",
		floor:      5,
	},
	"Sorin": developer{
		manager:    "Edina",
		department: "Development",
		team:       "SmartShop",
		floor:      5,
	},
	"Bomany": developer{
		manager:    "Simon",
		department: "Devops",
		team:       "SmartShop",
		floor:      5,
	},
	"Peter": developer{
		manager:    "David",
		department: "Development",
		team:       "Web",
		floor:      4,
	},
}

func main() {
	fmt.Println("Employees who work in development on the 5th floor", developersOnTheFifthFloor())
	fmt.Println("Employees who Charlie manages", developersCharlieManages())
	fmt.Println("Employees who work in product", employeesWhoWorkInProduct())
}

func developersCharlieManages() []string {
	charliesTeam := make([]string, 0)
	for name, details := range developers {
		if strings.ToLower(details.manager) == "charlie" {
			charliesTeam = append(charliesTeam, name)
		}
	}
	return charliesTeam
}

func developersOnTheFifthFloor() []string {
	developersOnFifthFloor := make([]string, 0)
	for name, details := range developers {
		if strings.ToLower(details.department) == "development" && details.floor == 5 {
			developersOnFifthFloor = append(developersOnFifthFloor, name)
		}
	}
	return developersOnFifthFloor
}

func employeesWhoWorkInProduct() []string {
	employeesWhoWorkInProduct := make([]string, 0)
	for name, details := range developers {
		if strings.ToLower(details.team) == "product" {
			employeesWhoWorkInProduct = append(employeesWhoWorkInProduct, name)
		}
	}
	return employeesWhoWorkInProduct
}
