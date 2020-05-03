# Go (Golang): The Complete Bootcamp

[udemy] course  on the `go` language.
Contains progress of my work in the course and the exercises...

## Exercises and content

All exrcises and course contents are available in the instructors `learngo` [repository]

## Additional stuff...

> Going through the `go` documentation from command line

```bash
$ go doc -src strconv FormatBool
package strconv // import "strconv"

// FormatBool returns "true" or "false" according to the value of b.
func FormatBool(b bool) string {
	if b {
		return "true"
	}
	return "false"
}
```
> using the godoc tool

start  a http server and access the documentation in browser:

```sh
godoc -http=:8080
```
- string length
	The len built-in function returns the length of v, according to its type. For `string` it always returns the number of bytes.

	For getting the number of characters in `string`, use `utf8.RuneCountInString`
	`string literals` are utf-8 encoded objects. If the string contains non english characters it's better to use the `utf8.RuneCountInString` for getting the string length.


[repository]: <github.com/inancgumus/learngo>
[udemy]: <www.udemy.com/course/learn-go-the-complete-bootcamp-course-golang>
