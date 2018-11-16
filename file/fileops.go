package file

import (
	"fmt"
	"io/ioutil"
)

func SaveToFile(filename string, input string) error {	
	return ioutil.WriteFile(filename , []byte(input), 0775)
}


func ReadFile(filename string) string {
	output, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error : ", err)
	}
	return string(output)
}