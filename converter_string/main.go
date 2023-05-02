package main

import (
	"fmt"
	"strconv"
)

type HexConverter interface {
	Convert(hexString string) (decString string)
}

type hexString struct {
	hexString string
}

func getPow(length int) int64 {
	var pow int64
	pow = 1
	for i := 0; i < length-1; i++ {
		pow *= 16
	}
	return pow
}
func (h hexString) Convert(hexString string) (decString string) {
	runeString := []rune(hexString)
	var result uint64
	for idx, value := range runeString {
		decInt, err := strconv.ParseInt(string(value), 16, 10)
		decInt = decInt * getPow(len(runeString)-idx)
		result += uint64(decInt)
		if err != nil {
			fmt.Println(err)
		}

	}
	decString = strconv.Itoa(int(result))
	return decString
}

func main() {
	str := ""
	_, err := fmt.Scan(&str)
	if err != nil {
		fmt.Println(err)
	}
	hexString := hexString{hexString: str}
	decString := hexString.Convert(hexString.hexString)
	fmt.Println(decString)

}