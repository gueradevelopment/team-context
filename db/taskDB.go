package db

import (
	"errors"

	"github.com/gueradevelopment/team-context/models"
)

// TaskDB - Task model database accessor
type TaskDB struct{}

var (
	items = make(map[string]models.Task)
)

// Get - retrieves a single resource
func (db *TaskDB) Get(id string, c chan Result) {
	defer close(c)
	result := Result{}

	item, ok := checklistItems[id]
	if ok {
		result.Result = item
		result.Err = nil
	} else {
		result.Err = errors.New("No result")
	}

	c <- result
}

// GetAll - retrieves all resources
func (db *TaskDB) GetAll(c chan ResultArray, resources map[string][]string) {
	defer close(c)
	result := ResultArray{}
	var arr = []Model{}
	var boardID string
	if resources["boardId"] != nil {
		boardID = resources["boardId"][0]
	}
	var checklistID string
	if resources["checklistId"] != nil {
		checklistID = resources["checklistId"][0]
	}
	for _, v := range items {
		if checklistID != "" && v.ChecklistID == checklistID {
			arr = append(arr, v)
		}
		if checklistID == "" {
			if boardID != "" && v.BoardID == boardID {
				arr = append(arr, v)
			}
			if boardID == "" {
				arr = append(arr, v)
			}
		}
	}
	result.Result = arr
	c <- result
}

// Add - creates a resource
func (db *TaskDB) Add(item models.Task, c chan Result) {
	defer close(c)
	result := Result{}
	if items[item.ID] == (models.Task{}) {
		items[item.ID] = item
		result.Result = item
	} else {
		result.Err = errors.New("Duplicated ID")
	}
	c <- result
}

// Edit - updates a resource
func (db *TaskDB) Edit(item models.Task, c chan Result) {
	defer close(c)
	result := Result{}
	if items[item.ID] == (models.Task{}) {
		result.Err = errors.New("No such ID")
	} else {
		items[item.ID] = item
		result.Result = item
	}
	c <- result
}

// Delete - deletes a resource
func (db *TaskDB) Delete(id string, c chan Result) {
	defer close(c)
	result := Result{}
	item := items[id]
	if item == (models.Task{}) {
		result.Err = errors.New("No such ID")
	} else {
		result.Result = item
		delete(items, id)
	}
	c <- result
}
