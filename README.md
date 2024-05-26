Package Discovery: https://pkg.go.dev/

# Go CLI

- go run <filename.go>: Compiles and executes one or two files

- go build <filename.go>: Compiles a bunch of go source code files

- go fmt: Automatically format all the code of each file of current directory

- go install: Compiles and installs a package

- go get: Downloads the raw source code of someone else's package

-go test: Runs any tests associated with the current project

### Types of packages

- Executable: Generates a file that we can run
- Reusable: Code used as 'helpers'. Good place to put resuable logic

- `package main` is sacred, specialized package name, i.e; anything other than `main` is used to create reusable package or dependency.

- Executable file should have a function named `main`

- "fmt" library is used to format or used for debugging, -- golang.org/pkg

