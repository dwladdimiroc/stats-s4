package files

import (
	//		"fmt"
	"bytes"
	"io/ioutil"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func WriteOutput(component string, output string) {
	var buffer bytes.Buffer

	pathFile := "logs/" + component + ".log"
	errWrite := ioutil.WriteFile(pathFile, []byte(output), 0644)
	check(errWrite)

	buffer.Reset()

}
