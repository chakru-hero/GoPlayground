package main
import "fmt"
func main(){
    inputString := "";
    fmt.Print("Enter a string : ");
    fmt.Scan(&inputString);
    fmt.Println(findNonRepeatingSubString(inputString));


}

func findNonRepeatingSubString(str string) int {
    m := make(map[byte]bool)
    left := 0
    size := 0

    for right := 0; right < len(str); right++ {
        for m[str[right]] {
            delete(m, str[left])
            left++
        }

        m[str[right]] = true

        if len(m) > size {
            size = len(m)
        }
    }

    return size
}

