package models

import "bubble/dao"

// Todo Model
type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

// Todo 增删改查

// CreateATodo 创建Todo
func CreateATodo(todo *Todo) (err error) {
	err = dao.DB.Create(&todo).Error
	return err
}

// GetTodoList 查询所有的Todo
func GetTodoList() (todos []*Todo, err error) {
	err = dao.DB.Find(&todos).Error //查询所欲数据
	if err != nil {
		return nil, err
	}
	return
}

// GetATodoByID 查询一个Todo
func GetATodoByID(id string) (todo *Todo, err error) {
	todo = new(Todo)
	err = dao.DB.Where("id  = ?", id).First(todo).Error
	if err != nil {
		return nil, err
	}
	return
}

// UpdateTodo 修改数据
func UpdateTodo(todo *Todo) (err error) {
	err = dao.DB.Save(&todo).Error
	return err
}

// DeleteTodo 删除数据
func DeleteTodo(id string) (err error) {
	err = dao.DB.Where("id = ?", id).Delete(Todo{}).Error
	return err
}
