package internals

import "fmt"

func PrintList(templates []string) {
	fmt.Println(Header.Render("Available Templates:"))
	for _, t := range templates {
		fmt.Println(Success.Render(" • ") + t)
	}
}

func PrintCreate(template, dest string) {
	fmt.Println(Success.Render("✔ Project created successfully!\n\n"))
	fmt.Println(InfoBox.Render(fmt.Sprintf("Template: %s\nPath: %s", template, dest)))
}
