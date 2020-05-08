
# Go (Golang): The Complete Bootcamp
  
[udemy] course on the `go` language.

Contains progress of my work in the course and the exercises...

## Exercises and content


All exercises and course contents are available in the instructors `learngo`  [repository]
  

## Additional stuff...

> Going through the `go` documentation from command line. We can use the `go doc` utility to read the `golang` documentation.

```bash
$ go doc -src strconv FormatBool

package strconv // import "strconv"
// FormatBool returns "true" or "false" according to the value of b.
func FormatBool(b bool) string {

if b {
	return  "true"
}
	return  "false"
}
```

> using the `godoc` utility

start a _http server_ and access the documentation in browser:

```sh
godoc -http=:8080
```
### Generic Information

+ **Strings**
	-  Getting string length
		The len built-in function returns the length of v, according to its type. For `string` it always returns the number of bytes.
	- For getting the number of characters in `string`, use `utf8.RuneCountInString`.

`string literals` are utf-8 encoded objects. If the string contains non english characters it's better to use the `utf8.RuneCountInString` for getting the string length.


+ **Declarations and Types**

+ A constant without a name is just a literal. So, constant should always be a named constant. _Constants will always be evaluated during compile time instead of runtime_. It helps to catch errors in the ealy stages.

+ A variable is an exact opposite of the constant and it gets executed during runtime. Also functions calls get executed during runtime.

+ _Typeless constants_ are the ones which do not have a type until one is needed. Examples as below

	```go
	const  max  =  101.08
	```

	Basic literals are also typed constants.

	```go
	3.14

	"Howdy"
	```

+ **IOTA**
	It's a Built-in constant number generator

`Iota` is a built-in constant generator which generates ever increasing numbers. Example

```go
const (
a  =  iota
b
c
d
)

fmt.Println(a, b, c, d)
// this will print 0, 1, 2, 3
// each subsequent variables after the first follow from the
// previous value of iota
```


[repository]: <github.com/inancgumus/learngo>

[udemy]: <www.udemy.com/course/learn-go-the-complete-bootcamp-course-golang>