package main

import (
	"block_1/devices"
	"block_1/modbus"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"time"
)

func main() {

	conn, err := net.Dial("tcp", "192.168.127.254:951")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	f, err := os.OpenFile("data.json", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		fmt.Println(err)
	}

	defer f.Close()

	for {

		var data = make(map[string]interface{})

		t, d := modbus.Today()

		data["T_1time"] = t
		data["Day"] = d

		for _, i := range devices.Devices {

			var dev = i.Get_device()

			if _, err = conn.Write(dev.Request); err != nil {
				fmt.Println(err)
				return
			}

			buff := make([]byte, 100)

			t := time.NewTimer(500 * time.Millisecond)
			<-t.C

			_, err := conn.Read(buff)
			if err != nil {
				fmt.Println(err)
				return
			}

			var mass_temp = i.Get_data()

			for j := 0; j < int(dev.Qan_chanels); j++ {
				data[dev.Chanels[j]] = mass_temp[j]
			}
		}

		data_js, _ := json.Marshal(data)

		f.WriteString(string(data_js) + "\n")
	}
}
