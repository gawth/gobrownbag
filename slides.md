class: middle, center
#GoLang BrownBag

By Alan Gawthorpe

---
class: middle, center
# Personal GoLang History
---
## CSharpe

```cs
public void DoFizzBuzz()
{
    for (int i = 1; i <= 100; i++)
    {
        bool fizz = i % 3 == 0;
        bool buzz = i % 5 == 0;
        if (fizz && buzz)
            Console.WriteLine ("FizzBuzz");
        else if (fizz)
            Console.WriteLine ("Fizz");
        else if (buzz)
            Console.WriteLine ("Buzz");
        else
            Console.WriteLine (i);
    }
}
```

---
## Python 

```python
def fizzbuzz(n):

    if n % 3 == 0 and n % 5 == 0:
        return 'FizzBuzz'
    elif n % 3 == 0:
        return 'Fizz'
    elif n % 5 == 0:
        return 'Buzz'
    else:
        return str(n)

print "\n".join(fizzbuzz(n) for n in xrange(1, 21))
```
---
## Javascript
```js
function fizzbuzz(i) {
  if (i % 5 === 0 && i % 3 === 0) {
    return 'FizzBuzz';
  } else if (i % 3 === 0) {
    return 'Fizz';
  } else if (i % 5 === 0) {
    return 'Buzz';
  } else {
    return i
  }
};
for (var i = 1; i <= 20; i++) {
  console.log(fizzbuzz(i) + '\n')
}
```

---
## GoLang

```go
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
```

---
# So What

---
# Google Background

![](Rob-pike-oscon.jpg) Rob Pike - Worked at Bell labs.  Co-creator of UTF-8?! 

![](robert.jpg) Robert Griesemer - some Java stuff 

![](220px-Ken_n_dennis.jpg) Ken Thompson - designed and implemented Unix and B (predecessor of C) 

![](Plan9bunnysmblack.jpg) ![](Golang_gopher.png) 

---
# Google Background

![](Rob-pike-oscon.jpg) Rob Pike - Worked at Bell labs.  Co-creator of UTF-8?! 

![](robert.jpg) Robert Griesemer - some Java stuff 

![](220px-Ken_n_dennis.jpg) Ken Thompson - designed and implemented Unix and B (predecessor of C) 

![](Plan9bunnysmblack.jpg) ![](Golang_gopher.png) Renée French

---
# Context

* C++, Python, Java
* Single repository
* Complex build system

---
# Key Features
Too many to list but here is a few I thought worth mentioning:
* Structure of working directories, etc
* Compilation is very picky - imports, documentation, unused variables
* No inheritence
* Interface implementation is implicit
* Gargage Collection
* Error handling
* Go routines and channels

---
# Interfaces

```go
type geometry interface {
    area() float64
    perim() float64
}

type rect struct {
    width, height float64
}
func (r rect) area() float64 {
    return r.width * r.height
}
func (r rect) perim() float64 {
    return 2*r.width + 2*r.height
}

func measure(g geometry) {
    fmt.Println(g)
    fmt.Println(g.area())
    fmt.Println(g.perim())
}
func main() {
    r := rect{width: 3, height: 4}
    
    measure(r)
}
```

---
exclude: true
# Tool Bloat

![](vs2015-fizzbuzz-solution.png)

---
#Other Examples

##Erlang
```erlang
fizzbuzz() ->
	F = fun(N) when N rem 15 == 0 -> "FizzBuzz";
		(N) when N rem 3 == 0  -> "Fizz";
		(N) when N rem 5 == 0  -> "Buzz";
		(N) -> integer_to_list(N)
	end,
[F(N)++"\n" || N <- lists:seq(1,100)]
```

##JS - Paul Irish
```js
for (var i = 1; i <= 100; i++) {
  var f = i % 3 == 0, b = i % 5 == 0;
  console.log(f ? b ? "FizzBuzz" : "Fizz" : b ? "Buzz" : i);
}
```

---
### Chapter 1.2

Some content

---
## And so on...