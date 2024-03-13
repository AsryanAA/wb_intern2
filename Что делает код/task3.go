// программа выведет nil, false потому что интерфейс это структура из ссылки на конкретный тип и его значение (значение nil, но тип *fs.PathError)

package main
 
import (
    "fmt"
    "os"
)
 
func Foo() error {
    var err *os.PathError = nil
    return err
}
 
func main() {
    err := Foo()
	// var err interface{} // для того чтобы получить true
    fmt.Println(err)
    fmt.Println(err == nil)
	fmt.Printf("type %T, value %v", err, err)
}