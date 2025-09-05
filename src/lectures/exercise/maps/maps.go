//--Summary:
//  Create a program to display server status. The server statuses are
//  defined as constants, and the servers are represented as strings
//  in the `servers` slice.
//
//--Requirements:
//* Create a function to print server status displaying:
//  - number of servers
//  - number of servers for each status (Online, Offline, Maintenance, Retired)
//* Create a map using the server names as the key and the server status
//  as the value
//* Set all of the server statuses to `Online` when creating the map
//* After creating the map, perform the following actions:
//  - call display server info function
//  - change server status of `darkstar` to `Retired`
//  - change server status of `aiur` to `Offline`
//  - call display server info function
//  - change server status of all servers to `Maintenance`
//  - call display server info function

package main

import "fmt"

const (
	Online      = 0
	Offline     = 1
	Maintenance = 2
	Retired     = 3
)

type Hosts map[string]int

func main() {
	servers := []string{"darkstar", "aiur", "omicron", "w359", "baseline"}
	hosts := make(Hosts)

	for _, s := range servers {
		hosts[s] = Online
	}
	hosts.PrintStatus()
	hosts.CountByStatus()
}

func (h Hosts) PrintStatus() {
	fmt.Println("Server status:")
	for k, v := range h {
		fmt.Printf("%s status : %s \n", k, intToString(v))
	}

}

func (h Hosts) Count() int {
	return len(h)
}

func (h Hosts) CountByStatus() {
	statusMap := make(map[string]int)

	for _, s := range h {
		fmt.Print(s)
		code := intToString(s)
		if count, exists := statusMap[code]; exists {
			statusMap[code] = count + 1
		} else {
			statusMap[code] = 1
		}
	}

	fmt.Print(statusMap)
	fmt.Println("Server status:")
	for k, v := range statusMap {
		fmt.Printf("%s : %v \n", k, v)
	}
	fmt.Println("-------------")
}

func intToString(a int) string {
	switch a {
	case Online:
		return "Online"
	case Offline:
		return "Offline"
	case Maintenance:
		return "Maintainance"
	case Retired:
		return "Retired"
	default:
		return "Not implemented"
	}

}
