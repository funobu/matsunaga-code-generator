package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/dave/jennifer/jen"
)

// CommandName define what kind of generate command.
type CommandName string

// below is the defined generate command.
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
	Value string       `json:"value"`
	Query Query        `json:"sql"`
}

type Commands interface {
	CreateVariable(ctx context.Context, input *jen.File, args Args) (*jen.File, error)
	HashPassword(ctx context.Context, input *jen.File, args Args) (*jen.File, error)
	DatabaseConnect(ctx context.Context, input *jen.File, args Args) (*jen.File, error)
	DatabaseQuery(ctx context.Context, input *jen.File, args Args) (*jen.File, error)
}

type generateCode struct {
	ID string
}

func NewGenerateCode(id string) Commands {
	return &generateCode{
		ID: id,
	}
}

func (gc *generateCode) CreateVariable(ctx context.Context, input *jen.File, args Args) (file *jen.File, err error) {
	input.Func().Id("main").Params().Block(
		jen.Id(args.Name).Op(":=").Lit(args.Value),
	)
	return
}

func (gc *generateCode) HashPassword(ctx context.Context, input *jen.File, args Args) (file *jen.File, err error) {
	return
}

func (gc *generateCode) DatabaseConnect(ctx context.Context, input *jen.File, args Args) (file *jen.File, err error) {
	return
}

func (gc *generateCode) DatabaseQuery(ctx context.Context, input *jen.File, args Args) (file *jen.File, err error) {
	return
}

func (gc *generateCode) Add(ctx context.Context, input *jen.File, command *Command) (file *jen.File, err error) {
	switch command.Name {
	case CreateVariable:
		return gc.CreateVariable(ctx, input, command.Args)
	case HashPassword:
		return gc.HashPassword(ctx, input, command.Args)
	case DatabaseConnect:
		return gc.DatabaseConnect(ctx, input, command.Args)
	case DatabaseQuery:
		return gc.DatabaseQuery(ctx, input, command.Args)
	default:
		return nil, nil
	}

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
