// программа выведет [77 78 79] потому что срез включает в себя начальный аргумент и исключает конечный

package main

import (
	"fmt"
)

func main() {
	a := [5]int{76, 77, 78, 79, 80}
	var b []int = a[1:4]
	fmt.Println(b)
}
