// программа выведет [3 2 3]

package main
 
import (
  	"fmt"
)
 
func main() {
	var s = []string{"1", "2", "3"} // len=3 cap=3 [1 2 3]
	modifySlice(s)
	fmt.Println(s)
}
 
func modifySlice(i []string) {
	i[0] = "3" // len=3 cap=3 [3 2 3]
	i = append(i, "4") // len=4 cap=6 [3 2 3 4] 
	i[1] = "5" // len=4 cap=6 [3 5 3 4]
	i = append(i, "6") // len=5 cap=6 [3 5 3 4 6]
}