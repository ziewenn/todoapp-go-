package main

import (
	"flag"
	"strconv"
	"strings"
)

type cmdFlags struct {
	Add    string
	Del    int
	Edit   string
	Toggle int
	List   bool
}

func NewCmdFlags() *cmdFlags {
	cf := cmdFlags{}

	flag.StringVar(&cf.Add, "add", "", "Add a new todo item")
	flag.StringVar(&cf.Edit, "edit", "", "Edit a todo item by format: <index>:<new title>")
	flag.IntVar(&cf.Del, "del", -1, "Delete a todo item by index")
	flag.IntVar(&cf.Toggle, "toggle", -1, "Toggle a todo item by index")
	flag.BoolVar(&cf.List, "list", false, "List all todo items")
	flag.Parse()

	return &cf
}

func (cf *cmdFlags) Execute(todos *Todos) {
	switch {
	case cf.List:
		todos.print()

	case cf.Add != "":
		todos.add(cf.Add)

	case cf.Del >= 0:
		if err := todos.delete(cf.Del); err != nil {
			println("Delete error:", err.Error())
			return
		}

	case cf.Toggle >= 0:
		if err := todos.marked(cf.Toggle); err != nil {
			println("Toggle error:", err.Error())
			return
		}

	case cf.Edit != "":
		index, title := parseEditFlag(cf.Edit)
		if index < 0 || index >= len(*todos) {
			println("Invalid index for edit")
			return
		}
		if err := todos.edit(index, title); err != nil {
			println("Edit error:", err.Error())
			return
		}

	default:
		println("No valid command provided. Use -h for help.")
	}
}

func parseEditFlag(edit string) (int, string) {
	parts := strings.SplitN(edit, ":", 2)
	if len(parts) != 2 {
		return -1, ""
	}
	index, err := strconv.Atoi(parts[0])
	if err != nil {
		return -1, ""
	}
	return index, parts[1]
}
