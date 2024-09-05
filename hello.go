package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func introduction() {
	var version float32 = 1.11
	fmt.Println("Version", version)
}

func showMenu() {
	fmt.Println("1 - Start Monitoring")
	fmt.Println("2 - Show Logs")
	fmt.Println("3 - Exit")
}

func getCommand() int {
	var commandRead int
	fmt.Scan(&commandRead)
	return commandRead
}

func monitor() {
	sites := readFromFile()

	for i, site := range sites {
		fmt.Println("Testing site", i)
		resp, err := http.Get(site)

		if err != nil {
			fmt.Println("Error: ", err)
		}

		if resp.StatusCode == 200 {
			fmt.Println(site, "is online")
			writeLog(site, true)
		} else {
			fmt.Println(site, "is ***offline***")
			writeLog(site, false)
		}
	}
}

func readFromFile() []string {

	var sites []string

	arquive, err := os.Open("D:/Demos/Go/src/hello/sites.txt")

	if err != nil {
		fmt.Println("Error: ", err)
	}

	reader := bufio.NewReader(arquive)
	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		sites = append(sites, line)
		if err == io.EOF {
			break
		}
	}
	arquive.Close()
	return sites
}

func writeLog(site string, status bool) {
	arquive, err := os.OpenFile("D:/Demos/Go/src/hello/log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("ERROR: ", err)
	}
	arquive.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - online:" + strconv.FormatBool(status) + "\n")
	arquive.Close()
}

func printLogs() {
	arquive, err := os.ReadFile("D:/Demos/Go/src/hello/log.txt")
	if err != nil {
		fmt.Println("ERROR: ", err)
	}

	fmt.Println(string(arquive))
}

func main() {
	introduction()

	for {
		showMenu()

		command := getCommand()

		switch command {
		case 1:
			monitor()
		case 2:
			printLogs()
		case 3:
			fmt.Println("Bye...")
			os.Exit(0)
		default:
			fmt.Println("Unvalid command")
			os.Exit(-1)
		}
	}
}
