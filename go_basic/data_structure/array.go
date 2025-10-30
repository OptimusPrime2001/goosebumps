package main

import "fmt"

type DispayOption struct {
	enabled     bool
	selectedIds []string
	displayMode string
}

func TestArray() {
	// var (
	// 	students = [3]string{"Trung", "HaMi", "Hoat"}
	// 	fordeer  = [...]int{3, 343}
	// )
	var displaySettings = [2][2]DispayOption{
		{
			{
				enabled:     true,
				selectedIds: []string{"1", "2"},
				displayMode: "list",
			},
			{
				enabled:     true,
				selectedIds: []string{"3", "4"},
				displayMode: "list",
			},
		},
		{
			{
				enabled:     true,
				selectedIds: []string{"5", "6"},
				displayMode: "list",
			},
			{
				enabled:     true,
				selectedIds: []string{"7", "8"},
				displayMode: "list",
			},
		},
	}
	for _, displayOption := range displaySettings {
		for _, option := range displayOption {
			fmt.Println(option.selectedIds)
		}
	}
	// fmt.Println(students, fordeer[1])
}
