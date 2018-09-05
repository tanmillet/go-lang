package src

import "github.com/pkg/errors"

func Add(a int, b int) (ret int, err error) {
	if a < 10 || b < 0 {
		err = errors.New("this is test error")
		return
	}
	return a + b, nil
}
