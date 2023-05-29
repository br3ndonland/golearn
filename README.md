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
- [Go Time podcast](https://changelog.com/gotime)
- [Domain-Driven Design with Go](https://www.packtpub.com/product/domain-driven-design-with-golang/9781804613450) (found via [Go Time 273](https://changelog.com/gotime/273))

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

The first time I tried to run tests with `go test`, I got the error `go: cannot find main module; see 'go help modules'`. The "Hello, World" chapter of Learn Go with Tests instructed me to run `go mod init helloworld`.

After running `go mod init` in a subfolder and then hovering over the module or package name in a Go file with the [VSCode Go extension](https://open-vsx.org/extension/golang/Go) enabled (extension ID `golang.go`), I see an error:

> gopls was not able to find modules in your workspace.
> When outside of GOPATH, gopls needs to know which modules you are working on.
> You can fix this by opening your workspace to a folder inside a Go module, or
> by using a go.work file to specify multiple modules.
> See the documentation for more information on setting up your workspace:
> https://github.com/golang/tools/blob/master/gopls/doc/workspace.md. `go list`

<img src="" alt="Screenshot of VSCode showing gopls error" width="50%" />

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
