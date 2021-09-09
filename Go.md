# Go Guide
This document describes followings:
* IDE
* Go Extension for VS Code
* Linter tool
* Naming
* Getter and Setter
* Mixed Caps
* Interface Naming
* Go Cheat Sheet
* Useful Links including `go style guide`


For the go style guideline, please read `go style` in the `Useful Links`

## IDE
Visual Studio Code is the IDE we are using in this SDK.
You can download it [here](https://code.visualstudio.com/download)

## Go Extension for VS Code
Go extension for VS Code should be installed.
It allows to to choose a linter tool for the project.

###Linter tool
Use golangci-lint. You can set this up  by going to settings (CTRL+,), typing go.lintTool and selecting it in the dropdownlist.
This extension would take care of go lang source code format.

## Naming
*By convention, packages are given lower case, single-word names;there should be no need for underscores or mixedCaps. 
*Another convention is that the package name is the base name of its source directory; 
*The importer of a package will use the name to refer to its contents, so exported names in the package can use that fact to avoid repetition. For example, bufio.reader not bufReader

## Getter and Setter
For Getter, do not use `Get` prefix, but do use `Set` postfix for the setter

## Mixed Caps
If you are using multi-words for variables, func, or methods, instead of using `-`, use upper camel case or lower camel case based on needs.

## interface Naming
By convention, one-method interfaces are named by the method name plus an -er suffix or similar modification to construct an agent noun: Reader, Writer, Formatter, CloseNotifier etc.


## Go Cheat Sheet

#### Variables
declaration
```
var msg string
msg = "Hello"
```

Shortcut of above (Infers type)
```
msg := "Hello"
```

#### Constants
Constants can be character, string, boolean, or numeric values.
```
const Phi = 1.618
```

#### Strings
```
str := "Hello"
str := `Multiline
string`
```
Strings are of type string.

#### Pointers
```
func main () {
  b := *getPointer()
  fmt.Println("Value is", b)
}
 
func getPointer () (myPointer *int) {
  a := 234
  return &a
}
 
a := new(int)
*a = 234
```
Pointers point to a memory location of a variable. Go is fully garbage-collected.

#### Numbers
Typical types
```
num := 3          // int
num := 3.         // float64
num := 3 + 4i     // complex128
num := byte('a')  // byte (alias for uint8)
```
Other types
```
var u uint = 7        // uint (unsigned)
var p float32 = 22.7  // 32-bit float
```
#### Type conversions
```
i := 2
f := float64(i)
u := uint(i)
```
####
```
// var numbers [5]int
numbers := [...]int{0, 0, 0, 0, 0}
```
Arrays have a fixed size.

#### Slices
```
slice := []int{2, 3, 4}
slice := []byte("Hello")
```
Slices have a dynamic size, unlike arrays.

### Flow control
#### Conditional
```
if day == "sunday" || day == "saturday" {
  rest()
} else if day == "monday" && isTired() {
  groan()
} else {
  work()
}
```
#### Statements in if
```
if _, err := doThing(); err != nil {
  fmt.Println("Uh oh")
}
```
A condition in an if statement can be preceded with a statement before a ;. Variables declared by the statement are only in scope until the end of the if.

#### Switch
```
switch day {
  case "sunday":
    // cases don't "fall through" by default!
    fallthrough

  case "saturday":
    rest()

  default:
    work()
}
```
#### For loop
```
for count := 0; count <= 10; count++ {
  fmt.Println("My counter is at", count)
}
```
#### For-Range loop
```
entry := []string{"Jack","John","Jones"}
for i, val := range entry {
  fmt.Printf("At position %d, the character %s is present\n", i, val)
}
```
#### While loop
```
n := 0
x := 42
for n != x {
  n := guess()
}
```
### Functions
#### Lambdas
```
myfunc := func() bool {
  return x > 10000
}
```
Functions are first class objects.

#### Multiple return types
```
a, b := getMessage()
func getMessage() (a string, b string) {
  return "Hello", "World"
}
```
#### Named return values
```
func split(sum int) (x, y int) {
  x = sum * 4 / 9
  y = sum - x
  return
}
```

By defining the return value names in the signature, a return (no args) will return variables with those names.

### Packages
#### Importing
```
import "fmt"
import "math/rand"
import (
  "fmt"        // gives fmt.Println
  "math/rand"  // gives rand.Intn
)
```
Both are the same.

#### Aliases
```
import r "math/rand"
 
r.Intn()
Exporting names
func Hello () {
  ···
}
```
Exported names begin with capital letters.

#### Packages
package hello
Every package file has to start with package.

### Concurrency
#### Goroutines
```
func main() {
  // A "channel"
  ch := make(chan string)

  // Start concurrent routines
  go push("Moe", ch)
  go push("Larry", ch)
  go push("Curly", ch)

  // Read 3 results
  // (Since our goroutines are concurrent,
  // the order isn't guaranteed!)
  fmt.Println(<-ch, <-ch, <-ch)
}
 
func push(name string, ch chan string) {
  msg := "Hey, " + name
  ch <- msg
}
```
Channels are concurrency-safe communication objects, used in goroutines.

#### Buffered channels
```
ch := make(chan int, 2)
ch <- 1
ch <- 2
ch <- 3
```
// fatal error:
// all goroutines are asleep - deadlock!
 
Buffered channels limit the amount of messages it can keep.

#### Closing channels
Closes a channel
```
ch <- 1
ch <- 2
ch <- 3
close(ch)
``` 
#### Iterates across a channel until its closed
```
for i := range ch {
  ···
}
 
//Closed if ok == false
v, ok := <- ch
```
* WaitGroup
```
import "sync"

func main() {
  var wg sync.WaitGroup
  
  for _, item := range itemList {
    // Increment WaitGroup Counter
    wg.Add(1)
    go doOperation(item)
  }
  // Wait for goroutines to finish
  wg.Wait()
  
}

func doOperation(item string) {
  defer wg.Done()
  // do operation on item
  // ...
}
```
A WaitGroup waits for a collection of goroutines to finish. The main goroutine calls Add to set the number of goroutines to wait for. The goroutine calls wg.Done() when it finishes. See: WaitGroup

### Error control
#### Defer
```
func main() {
  defer fmt.Println("Done")
  fmt.Println("Working...")
}
```
Defers running a function until the surrounding function returns. The arguments are evaluated immediately, but the function call is not ran until later.

#### Deferring functions
```
func main() {
  defer func() {
    fmt.Println("Done")
  }()
  fmt.Println("Working...")
}
``` 
Lambdas are better suited for defer blocks.
```
func main() {
  var d = int64(0)
  defer func(d *int64) {
    fmt.Printf("& %v Unix Sec\n", *d)
  }(&d)
  fmt.Print("Done ")
  d = time.Now().Unix()
}
``` 
The defer func uses current value of d, unless we use a pointer to get final value at end of main.

### Structs
#### Defining
```
type Vertex struct {
  X int
  Y int
}
 
func main() {
  v := Vertex{1, 2}
  v.X = 4
  fmt.Println(v.X, v.Y)
}
```

#### Literals
```
v := Vertex{X: 1, Y: 2}
// Field names can be omitted
v := Vertex{1, 2}
// Y is implicit
v := Vertex{X: 1}
```
You can also put field names.

#### Pointers to structs
```
v := &Vertex{1, 2}
v.X = 2
```
Doing v.X is the same as doing (*v).X, when v is a pointer.

### Methods
#### Receivers
```
type Vertex struct {
  X, Y float64
}
func (v Vertex) Abs() float64 {
  return math.Sqrt(v.X * v.X + v.Y * v.Y)
}
 
v := Vertex{1, 2}
v.Abs()
```
There are no classes, but you can define functions with receivers.

#### Mutation
```
func (v *Vertex) Scale(f float64) {
  v.X = v.X * f
  v.Y = v.Y * f
}
 
v := Vertex{6, 12}
v.Scale(0.5)
```
// `v` is updated
By defining your receiver as a pointer (*Vertex), you can do mutations.

### Interfaces
#### A basic interface
```
type Shape interface {
  Area() float64
  Perimeter() float64
}
```
#### Struct
```
type Rectangle struct {
  Length, Width float64
}
```
Struct Rectangle implicitly implements interface Shape by implementing all of its methods.

#### Methods
```
func (r Rectangle) Area() float64 {
  return r.Length * r.Width
}

func (r Rectangle) Perimeter() float64 {
  return 2 * (r.Length + r.Width)
}
```
The methods defined in Shape are implemented in Rectangle.

#### Interface example
```
func main() {
  var r Shape = Rectangle{Length: 3, Width: 4}
  fmt.Printf("Type of r: %T, Area: %v, Perimeter: %v.", r, r.Area(), r.Perimeter())
}
```

## Useful Links
* [A tour of Go](https://tour.golang.org)
* [Golang wiki](https://github.com/golang/go/wiki)
* [Effective Go](https://golang.org/doc/effective_go.html)
* [Go by Example](https://gobyexample.com)
* [Awesome Go](https://awesome-go.com)
* [JustForFunc Youtube](yhttps://www.youtube.com/channel/UC_BzFbxG2za3bp5NRRRXJSw)
* [Style Guide](https://github.com/golang/go/wiki/CodeReviewComments)
* [Edgecast Go Style](https://gitlab.edgecastcdn.net/edgecast/docs/developer-guide/-/blob/master/style/go/go.md)

