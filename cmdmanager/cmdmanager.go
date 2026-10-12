package cmdmanager

import "fmt"

type CMDManager struct{}

func (cm CMDManager) ReadLines() ([]string, error) {
	fmt.Println("Please provide the prices (0 to terminate). Confirm entry to each line by pressing enter.")

	lines := []string{}

	for {
		fmt.Print("Price: ")
		var line string
		fmt.Scan(&line)

		if line == "0" {
			break
		}

		lines = append(lines, line)
	}

	return lines, nil
}

func (cm CMDManager) WriteResults(data any) error {
	fmt.Println(data)
	return nil
}
