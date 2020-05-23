
# Go (Golang): The Complete Bootcamp

The excellent [udemy] course on the `golang` development bootcamp.

Contains progress of my work in the course and the exercises...

## Exercises and content

All exercises and course contents are available in the instructors [learngo] repository
  
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

+ _nil_ value:
  `nil` is a predeclared identifier in `go` similar to other predeclared identifiers like `true`, `false`, `float64` etc. It can be used anywhere without any special packages.
  Analogically it's similar to the below in other languages
  - `java: null`
  - `python: None`
  - `javascript: null` ...

	All `nil` means is that there does not exist any particular value, but `nil` itself is the value. It's a special value which says there is no value.

+ **Composite Types**:
	The following are the _composite types_ in `go`. They can store multiple values at the same time.
	- _Arrays_
	- _Slices_
	- _String Internals_
	- _Maps_
	- _Structs_
  Except _structs_ all the above types support index based element lookup.

- **Arrays** and **Slices**
	`Arrays` are a collection of elements (indexable) of fixed length. Length of an array is fixed and is determined at the compile time. Arrays cannot `grow` or `shrink`. `Go` being a statically typed language the size of the array specification is fixed during the compile time and cannot be changed thereafter. 
	representation: `var variable_name [length]element_type`
	eg: `var colors [7]string`
	Also worth to note is that the length and element types of an array are inseparable parts of its type. The type of the colors array above is `[7]string` and not just `string`.

	Once an array is defined, it's length is fixed and always gives the same value regardless of whether the array is populated with values or not. This is because, once the array is defined all the places will be populated with their default zero values as per the array's type.

	`Slices` are similar to `Arrays`, but they are of dynamic length and can either grow or shrink as per demand.

- **Keyed Elements**
	Each key of the array corresponds to an _index_ of the array.
	Here is an example in which we are using indexes inside the element list of the array literal.
```go
items = [3]float64 {
	0: 1.5,				// key = 0 is the index of 1st element
	1: 2.6,				// key = 1 is the index of 2nd element
	2: 3.2,				// key = 2 is the index of 3rd element
}
```
The keyed elements can now occur in any order as below.
```go
items = [3]float64 {
	2: 3.2,				// key = 2 is the index of 3rd element
	0: 1.5,				// key = 0 is the index of 1st element
	1: 2.6,				// key = 1 is the index of 2nd element
}
```

- **Slices**
	Slices are the dynamic arrays of `go`. New elements can be added or existing elements removed from a slice during runtime.
	Example: `var values []int`
	The length or the size of the slice is left empty as that is dynamically updated during runtime unlike an array where the value is fixed during compile time.
	The zero value of a slice is `nil`.

	+ Appending the slices:
	For appending two slices, always use the _ellipsis_ (`...`) operator as below, otherwise it will through an error.

	```go
	a := []int{1, 2, 3}
	b := []int{4, 5}
	c := append(a, b...)
	```

	The _ellipsis_ operator sends the values of the slice as additional arguments to the variadic function append, there by apending the additional elements of the slice explicitly.

	_note:_ appending slices will always result in a new slice; we have to save it over our original slice, or result of append would be lost.

	_example_: here `x` will remain as is and `y` will get new values as shown below

	```go
	x := []int{6, 7, 8}
	y := append(x, []int{11, 12, 13}...)

	fmt.Println(x)
	fmt.Println(y)

	// this will print the below output
	// [6 7 8]
	// [6 7 8 11 12 13]
	```

	_A String is a `byte slice` behind the scenes._

	+ Slicing an array, slice or a string
	
	representation:
	```go
	sliceable[start_position : end_position]
	```
	where,
	`start_position` = *value from which slicing can be done* or can be considered as start slicing from the `index` position; so it can start from `0`.
	`end_position` = *slice up till this value* or can be considered as slice till `(index + 1)`; so it should start from `1`.

	_note_ : `slice`  expressions return a new slice derived from the original slice.

	A `nil` slice does not have a backing array, so both it's length as well as capacity are `zero`.

	* A view of the slice header: The slice header is a tiny datastructure containing the below metadata information.

	_Pointer_

	_Length_
	
	_Capacity_

	* `make` function allows to create a slice with preallocated backing array of given capacity.
	
	```go
	make([]Type, length, capacity)
	
	make([]int, 5)

	s := make([]int, 0, 5)
	```


[learngo]: https://github.com/inancgumus/learngo

[udemy]: https://udemy.com/course/learn-go-the-complete-bootcamp-course-golang