package main
import (
    "fmt"
    "strconv"
)

func fizzbuzz(num int) string {
    if num%15 == 0 {
        return("FizzBuzz")
    } else if num%3 == 0 {
        return("Fizz")
    } else if num%5 == 0 {
        return("Buzz")
    } 
    return(strconv.Itoa(num))
}

func main() {  
    for i := 1; i <= 100; i++ {
        fmt.Println(fizzbuzz(i))
    }
}