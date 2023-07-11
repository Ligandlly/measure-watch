package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

const recordFile = "record.json"

type Record struct {
	SysTime   time.Time `json:"sysTime"`
	WatchTime time.Time `json:"watchTime"`
}

func main() {
	if len(os.Args) == 2 && (os.Args[1] == "rm" || os.Args[1] == "remove") {
		err := os.Remove(recordFile)
		if err != nil {
			panic(err)
		}
		return
	}
	fmt.Println("Press enter")
	fmt.Scanln()
	sysNow := time.Now()

	var minutes int
	fmt.Println("Watch minuts: ")
	fmt.Scanf("%d", &minutes)
	watchNow := time.Date(sysNow.Year(), sysNow.Month(), sysNow.Day(), sysNow.Hour(), minutes, 0, 0, sysNow.Location())

	record := Record{SysTime: sysNow, WatchTime: watchNow}

	if _, err := os.Stat(recordFile); err != nil {

		sysNowByte, _ := json.Marshal(record)
		file, _ := os.Create(recordFile)
		defer file.Close()
		file.Write(sysNowByte)
	} else {
		var pastRecord Record
		recordByte, err := os.ReadFile(recordFile)
		if err != nil {
			panic(err)
		}
		// os.Remove(recordFile)
		json.Unmarshal(recordByte, &pastRecord)

		sysDuration := record.SysTime.Sub(pastRecord.SysTime)
		watchDuration := record.WatchTime.Sub(pastRecord.WatchTime)

		rlt := (-watchDuration.Microseconds() + sysDuration.Microseconds()) * (24 * time.Hour.Microseconds() / sysDuration.Microseconds())
		fmt.Println("sysDuration", sysDuration)
		fmt.Println("watchDuration", watchDuration)
		fmt.Println(float64(rlt)/1000000, "seconds pre day")
	}
}
