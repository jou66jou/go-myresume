package test1

import "fmt"

type test1 struct {
	name string
}

func InitTest1() (*test1, error) {
	t := &test1{"jase"}
	return t, nil

}

func (t *test1) StartTask() {
	fmt.Println(t.name)
}
