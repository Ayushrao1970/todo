package api

import (
	"encoding/json"
	"fmt"
	"os"
	"slices"
	"strings"
	"time"
)

type TodoList struct {
	List []Todo `json:"list"`
	File string
}

type Todo struct {
	Id    int       `json:"id"`
	Task  string    `json:"task"`
	Time  time.Time `json:"time,omitempty"`
	Check string    `json:"check"`
}

func (t *TodoList) Check(i int) {
	i = i - 1
	if t.List[i].Check == "Pending" || t.List[i].Check == "" {
		t.List[i].Check = "Done"
		t.List[i].Time = time.Now()
	}

	t.Save()
}
func (t *TodoList) UnCheck(i int) {
	i = i - 1
	if t.List[i].Check == "Done" || t.List[i].Check == "" {
		t.List[i].Check = "Pending"
		t.List[i].Time = time.Now()
	}

	t.Save()
}

// get
func (t *TodoList) Get() {
	fmt.Printf("%-4s %-45s %-25s %-10s\n", "ID", "TASK", "CREATED", "CHECK")
	fmt.Println(strings.Repeat("-", 100))

	for _, item := range t.List {
		fmt.Printf(
			"%-4d %-45s %-25s %-10s\n",
			item.Id,
			truncate(item.Task, 45),
			item.Time.Format("02 Jan 2006, 03:04 PM"),
			item.Check,
		)
	}
}

func (t *TodoList) Add(task string) {
	p := Todo{
		Id:    len(t.List) + 1,
		Task:  task,
		Time:  time.Now(),
		Check: "Pending",
	}
	t.List = append(t.List, p)
	t.Save()

}
func (t *TodoList) Delete(i int) {
	i = i - 1
	t.List = slices.Delete(t.List, i, i+1)
	for j := range len(t.List) {
		t.List[j].Id = j + 1

	}
	t.Save()

}
func (t *TodoList) Update(i int, task string) {
	i = i - 1
	t.List[i].Task = task
	t.List[i].Time = time.Now()
	t.Save()

}

func (t *TodoList) Save() {
	b, _ := json.MarshalIndent(t.List, "", "")

	os.WriteFile(t.File, b, 0644) // we create the file here

}

func (t *TodoList) Load() error {

	data, err := os.ReadFile(t.File)

	if err != nil {
		if os.IsNotExist(err) { // if file doesn't exist then we create the todolist
			t.List = []Todo{}
			return nil
		}
		return err
	}

	if len(data) == 0 { //if len of data from db.json is empty we create the todolist
		t.List = []Todo{}
		return nil
	}

	return json.Unmarshal(data, &t.List) //if none of the case from above then unmarshal the data byte to list struct
}

// func (t Todo) String() string {
// 	return fmt.Sprintf(
// 		"[%d] %s\n    Created: %s  Check:%s",
// 		t.Id,
// 		t.Task,
// 		t.Time.Format("02 Jan 2006, 03:04 PM"),
// 		t.Check,
// 	)
// }

func truncate(s string, max int) string {
	if len(s) <= max {
		return s
	}
	return s[:max-3] + "..."
}
