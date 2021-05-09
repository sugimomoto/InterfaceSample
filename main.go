package main

import (
	"fmt"
)

func main() {

	s := &StudentImpl{}

	Show(s)

	mock := &StudentMock{}

	Show(mock)

}

func Show(s Student) {
	s.Name()
	s.Age()
}

type Student interface {
	Name()
	Age()
}

type StudentImpl struct{}

func (s *StudentImpl) Name() {
	fmt.Println("Taro")
}

func (s *StudentImpl) Age() {
	fmt.Println(15)
}

type StudentMock struct{}

func (s *StudentMock) Name() {
	fmt.Println("Mock")
}

func (s *StudentMock) Age() {
	fmt.Println(100)
}

type Tasks interface {
	GetTasks() (tasks []Task, err error)
	GetTask(id int) (task Task, err error)
}

type Task struct {
	Id    int
	Title string
}
