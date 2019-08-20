package main

import (
	"errors"
	"fmt"
	"gotest/test/test1"
	"gotest/test/test2"
	"reflect"
)

var (
	TTs = &testtasks{}
)

func main() {
	err := Start()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, v := range TTs.Tasks {
		v.StartTask()
	}
}

func Start() (err error) {
	defer func() {
		if err1 := recover(); err1 != nil {
			err = err1.(error)
		}
	}()
	TTs.Tasks = make([]task, 0)
	register(test1.InitTest1())
	register(test2.InitTest2())
	return nil
}

func register(t task, err error) {
	if err != nil {
		panic(errors.New(reflect.TypeOf(t).String() + err.Error()))
	}
	TTs.Tasks = append(TTs.Tasks, t)
}

type task interface {
	StartTask()
}

type testtasks struct {
	Tasks []task
}
