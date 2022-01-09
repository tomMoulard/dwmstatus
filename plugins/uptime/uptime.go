package uptime

import (
	"dwmstatus/modules"
	"fmt"
	"log"
	"time"

	"github.com/shirou/gopsutil/v3/host"
)

func init() {
	c := make(chan string)
	go handleUptime(c)
	modules.Modules = append(modules.Modules, c)
}

func handleUptime(c chan string) {
	for {
		c <- getUptime()
		time.Sleep(time.Second)
	}
}

func getUptime() string {
	u, err := host.Uptime()
	if err != nil {
		log.Panic(err)
	}

	hours := 0
	mins := u / 60

	for mins >= 60 {
		mins -= 60
		hours += 1
	}

	return fmt.Sprintf("Up: %dh %dm", hours, mins)
}
