Что выведет программа? Объяснить вывод программы. Объяснить как работают defer’ы и их порядок вызовов.

```go
package main

import (
	"fmt"
)


func test() (x int) {
	defer func() {
		x++
	}()
	x = 1
	return
}


func anotherTest() int {
	var x int
	defer func() {
		x++
	}()
	x = 1
	return x
}


func main() {
	fmt.Println(test())
	fmt.Println(anotherTest())
}
```

Ответ:
```
2
1
Defer может читать и изменять возвращаемую именнованную переменную,
поэтому в первом случае вернется 2. 
Во втором же случае вернётся 1, так как defer изменит локальную переменную, а не значение которое вернётся 

```