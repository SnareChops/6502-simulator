package sim

import (
	"io/ioutil"
	"os"
)

func Start(path string) {
	data := Load(path)

	Run(data)
}

// Load loads a file and returns it's contents
func Load(path string) []byte {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	return data
}

func Run(data []byte) {

}

func Tick() {

}
