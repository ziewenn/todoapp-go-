package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/aquasecurity/table"
)

type Todo struct {
	Title       string
	Done        bool
	CreatedAt   time.Time
	CompletedAt *time.Time
}

type Todos []Todo

func (todos *Todos) add(title string) {
	now := time.Now()
	todo := Todo{
		Title:     title,
		Done:      false,
		CreatedAt: now,
		CompletedAt: nil,
	}
	*todos = append(*todos, todo)
}

func (todos *Todos) validate(index int) error {
	if index < 0 || index >= len(*todos) {
		err := errors.New("Invalid index")
		fmt.Println(err)
		return err
	}
	return nil
}

func (todos *Todos) delete(index int) error {
	if err := todos.validate(index); err != nil {
		return err
	}
	t := *todos
	*todos = append(t[:index], t[index+1:]...)
	return nil
}

func (todos *Todos) marked(index int) error {
	if err := todos.validate(index); err != nil {
		return err
	}

	t := *todos
	isCompleted := t[index].Done

	if !isCompleted {
		completionTime := time.Now()
		t[index].CompletedAt = &completionTime
	} else {
		t[index].CompletedAt = nil
	}

	t[index].Done = !isCompleted
	return nil
}

func (todos *Todos) edit(index int, title string) error {
	if err := todos.validate(index); err != nil {
		return err
	}
	(*todos)[index].Title = title
	return nil
}

func (todos *Todos) print() {
	table := table.New(os.Stdout)
	table.SetRowLines(false)
	table.SetHeaders("#", "Title", "Completed", "Created At", "Completed At")

	for index, t := range *todos {
		completed := "❌"
		completedAt := ""

		if t.Done {
			completed = "☑️"
			if t.CompletedAt != nil {
				completedAt = t.CompletedAt.Format(time.RFC1123)
			}
		}

		createdAt := ""
		if !t.CreatedAt.IsZero() {
			createdAt = t.CreatedAt.Format(time.RFC1123)
		}

		table.AddRow(strconv.Itoa(index), t.Title, completed, createdAt, completedAt)
	}

	table.Render()
}

