package main

import "fmt"

type customError struct {
	msg string
}

// реализуем интерфейс Error
func (e *customError) Error() string {
	return e.msg
}

func test1() *customError {
	{
		// do something
	}
	// return &customError{msg: "text"}
	return nil
}

func main() {
	var err error
	err = test1()
	// значение интерфейса nil тогда и только тогда и значение и тип nil
	fmt.Printf("type %T, value %v \n", err, err) //у нас тип != nil
	if err != nil {
		println("error")
		return
	}
	println("ok")
}
