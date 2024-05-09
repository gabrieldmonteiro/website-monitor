package functions

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const loops = 24
const delay = 3600

func ShowMenu() {
	fmt.Println("1- Start Monitoring")
	fmt.Println("2- Show Logs")
	fmt.Println("0- Exit")
}

func GetInput() int {
	var command int
	fmt.Scan(&command)
	fmt.Println("")
	fmt.Println("Command Selected >>>", command)
	fmt.Println("")

	return command
}

func StartMonitore() {
	fmt.Println("Monitoring...")
	sites := GetWebSites()

	for i := 0; i < loops; i++ {
		for i, site := range sites {
			fmt.Println("Site", i, ":", site)
			TestSite(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}

	fmt.Println("")
}

func TestSite(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Error:", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println(site, "is ONLINE!")
		WriteLog(site, true)
	} else {
		fmt.Println(site, "ERROR! Status Code:", resp.StatusCode)
		WriteLog(site, false)
	}
}

func GetWebSites() []string {
	var sites []string
	file, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("ERROR:", err)
	}

	reader := bufio.NewReader(file)
	for {
		row, err := reader.ReadString('\n')
		row = strings.TrimSpace(row)

		sites = append(sites, row)

		if err == io.EOF {
			break
		}

	}

	file.Close()
	return sites
}

func WriteLog(site string, status bool) {

	file, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
	}

	file.WriteString(time.Now().Format("01/02/2006 15:04:05") + " - " + site + " - Online: " + strconv.FormatBool(status) + "\n")

	file.Close()
}

func ShowLogs() {

	file, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(file))

}
