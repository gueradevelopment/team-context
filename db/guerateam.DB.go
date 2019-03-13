package db

import (
	"errors"

	"github.com/gueradevelopment/team-context/models"
)

// GuerateamDB - Guerateam model database accessor
type GuerateamDB struct{}

var (
	guerateamItems = make(map[string]models.Guerateam)
)

// Get - retrieves a single resource
func (db *GuerateamDB) Get(id string, c chan Result) {
	defer close(c)

	result := Result{}
	item, ok := guerateamItems[id]
	if ok {
		result.Result = item
		result.Err = nil
	} else {
		result.Err = errors.New("No result")
	}

	c <- result
}

// GetAll - retrieves all resources
func (db *GuerateamDB) GetAll(c chan ResultArray, where map[string][]string) {
	defer close(c)
	result := ResultArray{}
	var arr = []Model{}
	for _, v := range guerateamItems {
		arr = append(arr, v)
	}
	result.Result = arr
	c <- result
}

// Add - creates a resource
func (db *GuerateamDB) Add(item models.Guerateam, c chan Result) {
	defer close(c)

	result := Result{}
	if _, found := guerateamItems[item.ID]; found {
		result.Err = errors.New("Duplicated ID")
	} else {
		guerateamItems[item.ID] = item
		result.Result = item
	}

	c <- result
}

// Edit - updates a resource
func (db *GuerateamDB) Edit(item models.Guerateam, c chan Result) {
	defer close(c)
	result := Result{}

	if _, found := guerateamItems[item.ID]; !found {
		result.Err = errors.New("No such ID")
	} else {
		guerateamItems[item.ID] = item
		result.Result = item
	}

	c <- result
}

// Delete - deletes a resource
func (db *GuerateamDB) Delete(id string, c chan Result) {
	defer close(c)
	result := Result{}

	if item, found := guerateamItems[id]; !found {
		result.Err = errors.New("No such ID")
	} else {
		result.Result = item
		delete(guerateamItems, id)
	}

	c <- result
}
