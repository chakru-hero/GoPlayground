package basicFuncs

import "fmt"

func SwapTwoNumbers(){
    x :=0;
    y :=0;
    fmt.Print("Enter a value for x: ");
    fmt.Scanln(&x);
    fmt.Print("Enter a value for y : ");
    fmt.Scan(&y);
    fmt.Println("The values before swap : ");
    fmt.Println("x = ", x , " y = ", y);
    swap(&x,&y);
fmt.Println("The values after swap : ");
    
    fmt.Print("x = ", x , " y = ", y);
}

func swap(x *int, y *int){
    *y,*x = *x,*y;
}

func NonRepeatingSubString(){
    
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

