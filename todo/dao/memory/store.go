package memory

import (
	"day4/todo/dao"
	"day4/todo/model"
)

var nextID int = 1

type Store struct {
	idMap map[int]model.ToDo
}

func New() dao.ToDoDAO {
	return &Store{idMap: make(map[int]model.ToDo)}
}

func (store *Store) Add(todo model.ToDo) (id int, err error) {
	if todo.ID > 0 {
		id = todo.ID
	} else {
		id = nextID
		nextID = nextID + 1
	}

	todo.ID = id
	store.idMap[todo.ID] = todo
	return
}

func (store *Store) FindByID(ID int) (todo model.ToDo, err error) {
	todo = store.idMap[ID]
	return
}

func (store *Store) FindAll() (todos []model.ToDo, err error) {

	todos = make([]model.ToDo, len(store.idMap))
	for key, vlaue := range store.idMap {
		todos[key-1] = vlaue
	}
	return
}

func (store *Store) DeleteById(id int) (err error) {
	delete(store.idMap, id)
	return
}

func (store *Store) UpdateById(id int, todo model.ToDo) (err error) {
	_, ok := store.idMap[id]

	if ok {
		store.idMap[id] = todo
	}
	return
}
