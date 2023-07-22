# Go learn

## Resources

### Get going

Make a repo with the [GitHub CLI](https://cli.github.com/):

```sh
gh repo create golearn \
  --clone \
  --description "Resources from learning Go" \
  --disable-issues \
  --disable-wiki \
  --gitignore "Go" \
  --license "CC0-1.0" \
  --public
```

Go through these resources:

- [A Tour of Go](https://go.dev/tour/welcome/1)
- [Learn Go with Tests](https://github.com/quii/learn-go-with-tests)

### Go deeper

- [Go docs](https://go.dev/doc/) (note that the Go docs don't have built-in search, despite the fact that Go comes from Google)
- [Effective Go](https://go.dev/doc/effective_go) (I guess this is like the more comprehensive reference manual?)
- [Go references](https://go.dev/doc/#references) (ah, okay, _this_ is the more comprehensive reference manual)
  - [Go language specification](https://go.dev/ref/spec)
  - [Go modules reference](https://go.dev/ref/mod)
- [Go 1 release notes](https://go.dev/doc/go1)
- [Go GitHub Wiki](https://github.com/golang/go/wiki)
  - [Table-driven tests](https://github.com/golang/go/wiki/TableDrivenTests)
- [Awesome Go](https://github.com/avelino/awesome-go)
- [Go Time podcast](https://changelog.com/gotime)
- [Domain-Driven Design with Go](https://www.packtpub.com/product/domain-driven-design-with-golang/9781804613450) (found via [Go Time 273](https://changelog.com/gotime/273))

## Syntax

### Names

[Effective Go](https://go.dev/doc/effective_go#names) has general info on names. Key points:

> By convention, packages are given lower case, single-word names; there should be no need for underscores or mixedCaps...
> Another convention is that the package name is the base name of its source directory; the package in `src/encoding/base64` is imported as `"encoding/base64"` but has name `base64`, not `encoding_base64` and not `encodingBase64`...
> Finally, the convention in Go is to use `MixedCaps` or `mixedCaps` rather than underscores to write multiword names.

The comments on not using underscores apply to _package names_, not _file names_. [Test files](https://go.dev/doc/code#Testing) are written with names ending in `_test.go`.

[A Tour of Go](https://go.dev/tour/basics/3) explains exported names:

> In Go, a name is exported if it begins with a capital letter. For example, `Pizza` is an exported name, as is `Pi`, which is exported from the `math` package.

Go files can be written with names out-of-order, and Go is able to handle this when determining [order of evaluation](https://go.dev/ref/spec#Order_of_evaluation) of the expressions. [Python resolves names](https://docs.python.org/3/reference/executionmodel.html#resolution-of-names) approximately top-to-bottom (from the top of a file to the bottom), so in Python, the order of names is more important.

### Command-line applications

As described in the ["Get started with Go" tutorial](https://go.dev/doc/tutorial/getting-started) and [Go language spec](https://go.dev/ref/spec#Program_execution), `func main()` is used to determine behavior when running the program from the command-line, like `if __name__ == "__main__":` [in Python](https://docs.python.org/3/tutorial/modules.html#executing-modules-as-scripts).

### Iteration

[`for` is Go's `while`](https://go.dev/tour/flowcontrol/3):

```go
package main

import "fmt"

func main() {
	sum := 1
	for sum < 10 {
		sum += sum
	}
	fmt.Println(sum)
}
```

An infinite loop is simply `for`:

```go
package main

func main() {
	for {
	}
}
```

### Strings and runes

Strings must be double-quoted. A single-quoted character is called a [rune](https://go.dev/doc/go1#rune) and has its own data type.

### Arrays and slices

Go's arrays are fixed-length. Variable-length arrays are called "[slices](https://go.dev/doc/effective_go#slices)" and, [like "package" and "module,"](#go-packages-and-modules) the term "slice" is used differently in Go than it may be elsewhere. [In Python, a "slice" is a part of a list](https://docs.python.org/3/library/stdtypes.html#sequence-types-list-tuple-range), whereas in Go a slice is the list itself. Slices can be sliced.

### Structs

[A Tour of Go](https://go.dev/tour/moretypes/2) says, "A struct is a collection of fields. Struct fields are accessed using a dot. Struct fields can be accessed through a struct pointer." Hmm, not very helpful. The "[How to Write Go Code](https://go.dev/doc/code)" docs are not particularly helpful either - the example code introduces a struct without explaining what it is. The [Go language spec section on struct types](https://go.dev/ref/spec#Struct_types) is esoteric and doesn't get us much further. So let's just look at an example from the [structs chapter of "Learn Go with Tests"](https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/structs-methods-and-interfaces) (in [`shapes.go`](./structs/shapes.go) in this repo).

```go
type Rectangle struct {
	Width  float64
	Height float64
}
```

### Methods

[The Go language spec section on method declarations](https://go.dev/ref/spec#Method_declarations) says, "A method is a function with a receiver." The "receiver" can be any type (built-in types or structs). When methods are called, they must be called on an instance of the type. In this sense, it seems like Go methods are somewhat like [Python class](https://docs.python.org/3/tutorial/classes.html) instance methods, which accept a class instance (`self`) as their receiver.

The receiver variable is conventionally named with the first letter of the type, like `r Rectangle`.

Methods are not nested or indented inside their receiver types, so the `func` syntax itself is the only syntactic indication that a method is associated with a receiver.

Adding methods to the struct example would look like this:

```go
type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}
```

### Interfaces

[Interfaces](https://go.dev/ref/spec#Interface_types) are somewhat like [abstract base classes in Python](https://docs.python.org/3/glossary.html#term-abstract-base-class). One difference from other programming languages is that interface resolution is implicit in Go. In the example below, `Circle` is automatically a `Shape` because it implements an `Area()` method that returns a `float64`.

```go
import "math"

type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}
```

### Pointers

Go copies values when they are passed in to functions or methods. Pointers can be used to access values directly instead of copying them. Pointers are indicated with an asterisk (`*`) before the name of the value. Within functions that accept pointers as arguments, the pointers are automatically "dereferenced," meaning that the value at the memory address is identified rather than the memory address itself. See the example from the pointers chapter of "Learn Go with Tests" in [`wallet.go`](./pointers/wallet.go).

```go
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}
```

### Formatting

[Go has a built-in formatter](https://go.dev/doc/effective_go#formatting) `gofmt` that formats indentations with tabs and does not enforce a maximum line length.

## Troubleshooting

### Go workspaces

Initially, I had a `tour.go` file from working on [A Tour of Go](https://go.dev/tour/welcome/1) and then I started [Learn Go with Tests](https://github.com/quii/learn-go-with-tests). I thought I might put the files into subdirectories to keep them separated.

- `golearn/`
  - `helloworld/`
    - `hello_test.go`
    - `hello.go`
  - `tour/`
    - `tour.go`

<details><summary>The contents of each file looked like this <em>(expand)</em>.</summary>

`helloworld/hello_test.go`

```go
package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello()
	want := "Hello, world"
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

```

`helloworld/hello.go`

```go
package main

import "fmt"

func Hello() string {
	return "Hello, world"
}

func main() {
	fmt.Println(Hello())
}

```

`tour/tour.go`

```go
package main

import (
	"fmt"
	"time"
)

func add(x, y int) int {
	return x + y
}

func addAndPrint(x, y int) string {
	return fmt.Sprintf("The sum of %d and %d is %d", x, y, add(x, y))
}

func main() {
	var x, y int = 42, 13
	fmt.Println("Welcome to the playground!")
	fmt.Println("The time is", time.Now())
	fmt.Println(addAndPrint(x, y))
}

```

</details>

With this directory configuration, I can run each module like `go run helloworld/hello.go` or `go run tour/tour.go`, but I can't run tests.

If I try to run `go test helloworld/hello_test.go`, I see an error:

```text
golearn on  main [?⇡] via 🐹 v1.20.4
❯ go test helloworld/hello_test.go
# command-line-arguments [command-line-arguments.test]
helloworld/hello_test.go:6:9: undefined: Hello
FAIL    command-line-arguments [build failed]
FAIL
```

If I try to change into the `helloworld` directory and run `go test`, I see a completely different error:

```text
golearn on  main [?⇡] via 🐹 v1.20.4
❯ go test
go: go.mod file not found in current directory or any parent directory; see 'go help modules'
```

[The "Hello, World" chapter of Learn Go with Tests](https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/hello-world) suggested running `go mod init`.

After running `go mod init` in the subfolder and then hovering over the module or package name in a Go file with the [VSCode Go extension](https://open-vsx.org/extension/golang/Go) enabled (extension ID `golang.go`), I see an error:

> gopls was not able to find modules in your workspace.
> When outside of GOPATH, gopls needs to know which modules you are working on.
> You can fix this by opening your workspace to a folder inside a Go module, or
> by using a go.work file to specify multiple modules.
> See the documentation for more information on setting up your workspace:
> https://github.com/golang/tools/blob/master/gopls/doc/workspace.md. `go list`

<img src="https://github.com/br3ndonland/golearn/assets/26674818/937a5350-cc66-4d09-bfac-7f1d6c117632" alt="Screenshot of VSCode showing gopls error" width="50%" />

The [`gopls` docs](https://github.com/golang/tools/blob/5974258e689a4f8a93448a0d181737afa4506e3f/gopls/doc/workspace.md) (the Go language server) say:

> #### Go workspaces (Go 1.18+)
>
> Starting with Go 1.18, the `go` command has native support for multi-module
> workspaces, via [`go.work`](https://go.dev/ref/mod#workspaces) files. These
> files are recognized by gopls starting with `gopls@v0.8.0`.
>
> The easiest way to work on multiple modules in Go 1.18 and later is therefore
> to create a `go.work` file containing the modules you wish to work on, and set
> your workspace root to the directory containing the `go.work` file.
>
> For example, suppose this repo is checked out into the `$WORK/tools` directory.
> We can work on both `golang.org/x/tools` and `golang.org/x/tools/gopls`
> simultaneously by creating a `go.work` file using `go work init`, followed by
> `go work use MODULE_DIRECTORIES...` to add directories containing `go.mod` files to the
> workspace:
>
> ```sh
> cd $WORK
> go work init
> go work use ./tools/ ./tools/gopls/
> ```
>
> ...followed by opening the `$WORK` directory in our editor.

The [Go docs on workspaces](https://go.dev/doc/tutorial/workspaces) provide further details.

At this point, I ran:

```sh
golearn on  main [?⇡] via 🐹 v1.20.4
❯ go work init

golearn on  main [?⇡] via 🐹 v1.20.4
❯ go work use helloworld tour
```

The `go.work` file came out like this:

```go.work
go 1.20

use (
	./helloworld
	./tour
)
```

Note that the [GitHub Go `.gitignore` template](https://github.com/github/gitignore/blob/4488915eec0b3a45b5c63ead28f286819c0917de/Go.gitignore) is configured to ignore `go.work` files.

I guess a simpler way of doing this would have been running `go mod init golearn` in the top-level directory to have a single `go.mod` file, then adding subdirectories without their own `go.mod` files. I figured this out after learning more about Go modules. Read on.

### Go packages and modules

#### `Package is not a main package` errors

After moving `helloworld` to a subdirectory and running `go mod init` to establish the module, I first tried changing the `package` lines in `hello.go` and `hello_test.go` to say `package helloworld`, but that led to errors like `package command-line-arguments is not a main package` (in the top-level directory) and `package helloworld is not in GOROOT (/opt/homebrew/Cellar/go/1.20.4/libexec/src/helloworld)` (in the package directory). I also saw a notice in `hello.go` that `func main is unused`.

To resolve the errors, the module name in `go.mod` should say `module helloworld`, but the `package` lines in `hello.go` and `hello_test.go` should say `package main`. The Go CLI and VSCode extension could provide some clearer instructions here.

#### `Cannot import "main"` errors

Tutorials commonly include all source code in `package main` examples, which allows the examples to be run from the command-line, but this is not how a larger module would be built. The ["Learn Go with Tests" chapter on arrays and slices](https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/arrays-and-slices) explains:

> If you had initialized go mod with `go mod init main` you will be presented with an error `_testmain.go:13:2: cannot import "main"`. This is because according to common practice, package main will only contain integration of other packages and not unit-testable code and hence Go will not allow you to import a package with name `main`.
>
> To fix this, you can rename the main module in `go.mod` to any other name.

Did they mean "Go will not allow you to import a ~~package~~ _module_ with name `main`"? Tests seem to work on `package main` examples.

#### Definitions of package and module

_So what's the difference between a package and a module?_

The Go docs on "[How to write Go code](https://go.dev/doc/code)" explain,

> Go programs are organized into packages. A _package_ is a collection of source files in the same directory that are compiled together. Functions, types, variables, and constants defined in one source file are visible to all other source files within the same package.
>
> A repository contains one or more modules. A _module_ is a collection of related Go packages that are released together. A Go repository typically contains only one module, located at the root of the repository. A file named `go.mod` there declares the _module path_: the import path prefix for all packages within the module. The module contains the packages in the directory containing its go.mod file as well as subdirectories of that directory, up to the next subdirectory containing another `go.mod` file (if any).
>
> ...
>
> The first statement in a Go source file must be `package name`. Executable commands must always use `package main`.

The [Go modules docs](https://go.dev/ref/mod) provide more in-depth documentation.

This is effectively the inverse of Python - a [Python module](https://docs.python.org/3/tutorial/modules.html) is a `.py` file, whereas a [Python package](https://docs.python.org/3/tutorial/modules.html#packages) is a collection of modules in a directory with a `__init__.py` file in it. Packages can have subpackages, which are additional directories inside the package directory with `__init__.py` files in them.

Most other real-world definitions of "package" appear to follow the Python convention and not the Go convention. For example, GitHub calls their service "[GitHub Packages](https://docs.github.com/en/packages/learn-github-packages/introduction-to-github-packages)" not "GitHub Modules."

#### Setting up the Go module

At this point, I realized I didn't need Go workspaces. I deleted the `go.mod` files from each subdirectory, and added a single `go.mod` file to the top-level directory with `go mod init github.com/br3ndonland/golearn`, following the naming suggested in the [Go modules docs](https://go.dev/ref/mod).
