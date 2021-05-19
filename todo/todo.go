package todo

import (
	uuid "github.com/nu7hatch/gouuid"
)

type Todo struct {
	Id string `json:"id"`
	Title string `json:"title"`
	Done bool `json:"done"`
}

var list []*Todo

func Create(todo *Todo) (*Todo) {
	id, _ := uuid.NewV4()
	todo.Id = id.String()
	todo.Done = false
	list = append(list, todo)
	return todo
}

func Read(ids ...string) ([]*Todo) {
	if len(ids) > 0 {
		var temp []*Todo
		for _, item := range list {
			for _, id := range ids {
				if item.Id == id {
					temp = append(temp, item)
				}
			}
		}
		return temp;
	}
	return list
}

func Update(todo *Todo) (*Todo) {
	for _, t := range list {
		if t.Id == todo.Id {
			t.Done = todo.Done
			return t
		}
	}
	return nil
}