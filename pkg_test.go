package main

import (
	"os"
	"path/filepath"
	"testing"
)

func Helper() *TodoList {
	return &TodoList{
		List: []Todo{},
	}
}

func TestAdd(t *testing.T) {
	todo := Helper()
	todo.Add("write test")

	if len(todo.List) != 1 {
		t.Errorf("want %v and got %v", 1, len(todo.List))
	}
	if todo.List[0].Task != "write test" {
		t.Errorf("want %v and got %v", "write test", todo.List[0].Task)
	}

}

func TestUpdate(t *testing.T) {
	todo := Helper()
	todo.Add("task 1")
	todo.Add("task 2")
	todo.Update(1, "new task")

	if todo.List[1].Task != "new task" {
		t.Errorf("want %v and got %v", "new task", todo.List[1].Task)
	}

}

func TestCheck(t *testing.T) {
	todo := Helper()
	todo.Add("task 1")
	todo.Add("task 2")

	todo.Check(0)
	todo.Check(1)
	if todo.List[0].Check != "Done" {
		t.Errorf("want %v got %v", "Done", todo.List[0].Check)

	}
	todo.UnCheck(1)

	if todo.List[1].Check != "" {
		t.Errorf("want %v got %v", "", todo.List[1].Check)
	}
}

func TestDelete(t *testing.T) {
	todo := Helper()
	todo.Add("task 1")
	todo.Add("task 2")
	todo.Add("task 3")

	todo.Delete(2)
	if len(todo.List) != 2 {
		t.Errorf("want %v got %v", 2, len(todo.List))
	}
	if todo.List[1].Id != 2 {
		t.Errorf("want %v got %v", 2, todo.List[1].Id)
	}
}

func TestLoadAndSave(t *testing.T) {
	dir := t.TempDir()
	file := filepath.Join(dir, "db.json")

	data := []byte(`[
		{"id":1,"task":"from test","check":"Done"}
	]`)
	err := os.WriteFile(file, data, 0664)
	if err != nil {
		t.Fatal(err)
	}
	todo := &TodoList{
		File: file,
	}
	err = todo.Load()
	if err != nil {

		t.Fatal(err)
	}
	if len(todo.List) != 1 {
		t.Errorf("want %v got %v", 1, len(todo.List))
	}
	if todo.List[0].Task != "from test" {
		t.Errorf("want %v got %v ", "from test", todo.List[0].Task)

	}
}
