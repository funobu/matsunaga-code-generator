package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type CommandName string

const (
	CreateVariable  = CommandName("create variable")
	HashPassword    = CommandName("hash password")
	DatabaseConnect = CommandName("database connect")
	DatabaseQuery   = CommandName("database query")
)

type Command struct {
	Name CommandName `json:"command"`
	Args `json:"args"`
}

type QueryType string

const (
	Select = QueryType("select")
	Insert = QueryType("insert")
	Update = QueryType("update")
	Delete = QueryType("delete")
)

type Query struct {
	Table string    `json:"table"`
	Type  QueryType `json:"type"`
}

type VariableType string

const (
	String = VariableType("string")
)

type Args struct {
	Name  string       `json:"name"`
	Type  VariableType `json:"type"`
	Query Query        `json:"sql"`
}

func main() {
	commands := make([]*Command, 0)

	input, err := ioutil.ReadFile("./mock/sample1.json")
	if err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal(input, &commands); err != nil {
		log.Fatal(err)
	}

	fmt.Println(len(commands))

	for _, command := range commands {
		fmt.Printf("%s: args -> %v\n", command.Name, command.Args)
	}

}
