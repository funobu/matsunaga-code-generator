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
	createVariable(ctx context.Context, args Args) error
	hashPassword(ctx context.Context, args Args) error
	databaseConnect(ctx context.Context, args Args) error
	databaseQuery(ctx context.Context, args Args) error
	Add(ctx context.Context, command *Command) error
	Generate(ctx context.Context) error
}

// generateCode is the struct of Commands implements.
type generateCode struct {
	ID           string
	CommandLines []*jen.Statement
}

// NewGenerateCode is the function that makes generateCode instance.
func NewGenerateCode(id string) Commands {
	return &generateCode{
		ID:           id,
		CommandLines: make([]*jen.Statement, 0),
	}
}

func (gc *generateCode) createVariable(ctx context.Context, args Args) (err error) {
	gc.CommandLines = append(gc.CommandLines, jen.Id(args.Name).Op(":=").Lit(args.Value))
	return
}

func (gc *generateCode) hashPassword(ctx context.Context, args Args) (err error) {
	gc.CommandLines = append(gc.CommandLines, jen.Id(args.Name).Op(":=").Lit(args.Value))
	return
}

func (gc *generateCode) databaseConnect(ctx context.Context, args Args) (err error) {
	gc.CommandLines = append(gc.CommandLines, jen.Id(args.Name).Op(":=").Lit(args.Value))
	return
}

func (gc *generateCode) databaseQuery(ctx context.Context, args Args) (err error) {
	gc.CommandLines = append(gc.CommandLines, jen.Id(args.Name).Op(":=").Lit(args.Value))
	return
}

// Add method write generated code to the target file.
func (gc *generateCode) Add(ctx context.Context, command *Command) (err error) {
	switch command.Name {
	case CreateVariable:
		return gc.createVariable(ctx, command.Args)
	case HashPassword:
		return gc.hashPassword(ctx, command.Args)
	case DatabaseConnect:
		return gc.databaseConnect(ctx, command.Args)
	case DatabaseQuery:
		return gc.databaseQuery(ctx, command.Args)
	default:
		return nil
	}

}

func (gc *generateCode) Generate(ctx context.Context) (err error) {
	f := jen.NewFile("main")
	f.Func().Id("main").Params().BlockFunc(func(g *jen.Group) {
		for i := range gc.CommandLines {
			g.Add(gc.CommandLines[i])
		}
	})
	fmt.Printf("%#v", f)

	return
}

func main() {
	ctx := context.Background()
	commands := make([]*Command, 0)

	input, err := ioutil.ReadFile("./mock/sample1.json")
	if err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal(input, &commands); err != nil {
		log.Fatal(err)
	}

	fmt.Println(len(commands))

	gc := NewGenerateCode("111")

	for _, command := range commands {
		fmt.Printf("%s: args -> %v\n", command.Name, command.Args)
		gc.Add(ctx, command)
	}

	gc.Generate(ctx)

}
