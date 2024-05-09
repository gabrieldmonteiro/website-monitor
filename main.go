package main

import (
	"fmt"
	"os"
	"web-monitor/functions"
)

func main() {

	fmt.Println("╦ ╦┌─┐┌┐   ╔╦╗┌─┐┌┐┌┬┌┬┐┌─┐┬─┐")
	fmt.Println("║║║├┤ ├┴┐  ║║║│ │││││ │ │ │├┬┘")
	fmt.Println("╚╩╝└─┘└─┘  ╩ ╩└─┘┘└┘┴ ┴ └─┘┴└─")

	for {
		functions.ShowMenu()

		command := functions.GetInput()

		switch command {
		case 1:
			functions.StartMonitore()
		case 2:
			fmt.Println("Showing Logs...")
			functions.ShowLogs()
		case 0:
			fmt.Println("Exiting...")
			os.Exit(0)
		default:
			fmt.Println("IDK this command... Exiting...")
			os.Exit(-1)
		}
	}
}
