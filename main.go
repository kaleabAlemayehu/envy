package main

import (
	"fmt"
	"log"
	"os"
)

type Envy struct {
	Logger *log.Logger
}

func NewEnvy(l *log.Logger) *Envy {
	return &Envy{l}
}
func main() {
	l := log.New(os.Stdout, "EnvyLog#_ ", log.LstdFlags)
	envy := NewEnvy(l)
	envy.Config(".env")
}

func (e *Envy) Config(location ...string) {
	for _, value := range location {
		data, err := os.ReadFile(value)
		if err != nil {
			e.Logger.Fatalln(err.Error())
		}

		_, _, err = e.Parse(data)
		if err != nil {
			e.Logger.Fatalln(err.Error())
		}
	}
}
func (e *Envy) Parse(data []byte) (string, string, error) {
	var i int
	var chunkPoint []int = []int{}
	for i < len(data) {
		if data[i] == '\n' {
			chunkPoint = append(chunkPoint, i-1)
			fmt.Printf("%v on %d\n", data[i], i)
			i++
		}
		fmt.Printf("%v on %d\n", string(data[i]), i)
		i++
	}
	lastpoint := 0
	for _, point := range chunkPoint {
		fmt.Println("chunk", string(data[lastpoint:point]))
		lastpoint = point
	}
	return "", "", nil
}
