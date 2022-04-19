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
	ChangeVariable  = CommandName("change variable")
	JudgeTarget     = CommandName("judge target")
	HashPassword    = CommandName("hash password")
	DatabaseConnect = CommandName("database connect")
	DatabaseQuery   = CommandName("database query")
)

// command is the struct of input command.
type Command struct {
	Name CommandName `json:"command"`
	Args `json:"args"`
}

// QueryType is the type of SQL Query CRUD.
type QueryType string

const (
	Select = QueryType("select")
	Insert = QueryType("insert")
	Update = QueryType("update")
	Delete = QueryType("delete")
)

// Query is the struct of SQL Query.
type Query struct {
	Table  string    `json:"table"`
	Type   QueryType `json:"type"`
	Notion string    `json:"notion"`
	Update string    `json:"update"`
}

// VariableType is the type of golang variable.
type VariableType string

const (
	String = VariableType("string")
)

// Args is the struct of input command args.
type Args struct {
	Name  string       `json:"name"`
	Type  VariableType `json:"type"`
	Value string       `json:"value"`
	Query Query        `json:"sql"`
}

// Commands is the interface of input commands.
type Commands interface {
	CreateVariable(ctx context.Context, output *jen.File, args Args) (*jen.File, error)
	HashPassword(ctx context.Context, output *jen.File, args Args) (*jen.File, error)
	DatabaseConnect(ctx context.Context, output *jen.File, args Args) (*jen.File, error)
	DatabaseQuery(ctx context.Context, output *jen.File, args Args) (*jen.File, error)
}

// generateCode is the struct of Commands implements.
type generateCode struct {
	ID string
}

// NewGenerateCode is the function that makes generateCode instance.
func NewGenerateCode(id string) Commands {
	return &generateCode{
		ID: id,
	}
}

func (gc *generateCode) CreateVariable(ctx context.Context, output *jen.File, args Args) (file *jen.File, err error) {
	output.Func().Id("main").Params().Block(
		jen.Id(args.Name).Op(":=").Lit(args.Value),
	)
	return
}

func (gc *generateCode) HashPassword(ctx context.Context, output *jen.File, args Args) (file *jen.File, err error) {
	return
}

func (gc *generateCode) DatabaseConnect(ctx context.Context, output *jen.File, args Args) (file *jen.File, err error) {
	return
}

func (gc *generateCode) DatabaseQuery(ctx context.Context, output *jen.File, args Args) (file *jen.File, err error) {
	return
}

// Add method write generated code to the target file.
func (gc *generateCode) Add(ctx context.Context, output *jen.File, command *Command) (file *jen.File, err error) {
	switch command.Name {
	case CreateVariable:
		return gc.CreateVariable(ctx, output, command.Args)
	case HashPassword:
		return gc.HashPassword(ctx, output, command.Args)
	case DatabaseConnect:
		return gc.DatabaseConnect(ctx, output, command.Args)
	case DatabaseQuery:
		return gc.DatabaseQuery(ctx, output, command.Args)
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

	generateCommandLines := make([]*jen.Statement, 0)

	generateCommandLines = append(generateCommandLines, jen.Qual("fmt", "Println").Call(jen.Lit("Hello, world")))
	generateCommandLines = append(generateCommandLines, jen.Qual("fmt", "Println").Call(jen.Lit("This is a Code Generator for Matsunaga Project.")))
	generateCommandLines = append(generateCommandLines, jen.Qual("fmt", "Println").Call(jen.Lit("By JSON file, code will be created.")))
	generateCommandLines = append(generateCommandLines, jen.Qual("fmt", "Println").Call(jen.Lit("Please use this library!!!")))

	// test
	f := jen.NewFile("main")
	f.Func().Id("main").Params().BlockFunc(func(g *jen.Group) {
		for i := range generateCommandLines {
			g.Add(generateCommandLines[i])
		}
	})
	fmt.Printf("%#v", f)

}
