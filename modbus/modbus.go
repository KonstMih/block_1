package modbus

import (
	"encoding/binary"
	"fmt"
	"math"
	"strconv"
	"time"
)

func Ieee_754(response []byte) float32 {

	var s int32
	var m int32
	var m1 float32

	slovo := int32(binary.BigEndian.Uint32(response))

	if (slovo >> 31) == 0 {
		s = 1
	} else {
		s = -1
	}

	e := (slovo >> 23) & 0xff

	if e != 0 {
		m = ((slovo & 0x7fffff) | 0x800000)
	} else {
		m = (slovo & 0x7fffff) << 1
	}

	m1 = float32(m) * float32(math.Pow(2, -23))

	return float32(s) * m1 * float32(math.Pow(2, float64(e)-127))
}

func Ascii_temp(response []byte) []float32 {

	var temp float64
	var mass_temp []float32

	temp, _ = strconv.ParseFloat(string(response), 32)
	mass_temp = append(mass_temp, float32(temp))

	return mass_temp

}

func Rtu_temp(response []byte) []float32 {

	var mass_temp []float32

	for i := 0; i < len(response); i += 4 {
		var m []byte = response[i : i+4]
		var temp float32 = Ieee_754(m)
		mass_temp = append(mass_temp, temp)

	}
	return mass_temp
}

func Today() (string, string) {
	today := time.Now()
	hours, min, sec := today.Clock()
	year, month, day := today.Date()

	t := fmt.Sprintf("%d:%d:%d", hours, min, sec)
	d := fmt.Sprintf("%d:%d:%d", year, int(month), day)

	return t, d
}
