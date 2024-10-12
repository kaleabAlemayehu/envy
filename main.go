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
	l := log.New(os.Stdout, "EnvyLog>", log.LstdFlags)
	envy := NewEnvy(l)
	envy.Config(".env", ".production.env")
}

func (e *Envy) Config(location ...string) {
	for _, value := range location {
		data, err := os.ReadFile(value)
		if err != nil {
			e.Logger.Fatalln(err.Error())
		}
		fmt.Println(string(data))
	}
}
