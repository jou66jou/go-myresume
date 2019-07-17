package test2

import (
	"errors"
	"fmt"
)

type test2 struct {
	name string
	nike string
}

func InitTest2() (*test2, error) {
	// t := &test2{"joy", "jy"}
	// return t, nil
	return nil, errors.New("tmpStr")

}

func (t *test2) StartTask() {
	fmt.Println(t.name)
}
