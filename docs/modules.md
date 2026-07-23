# Modules

A `module` is a collection of packages stored in a file tree under `$GOPATH/pkg` folder with a `go.mod` file at its root. This file defines the module's path which is also the import path used for the root directory and its dependency requirements.

The Go command automatically checks and adds dependencies required for imports provided the current directory or the parent directory has a `go.mod` file.

A Go module will have several Go files or packages in addition to two important files in the root, the `go.mod` file and `go.sum` file. These files are maintained by the Go tool, and it is used to track the module's configuration.

Before creating a module, you need to identify a directory where the module will reside. This directory can be anywhere on the computer and need not be in any specific Go directory. You can use an existing directory or create a new one.

```bash
go mod init MyModule
```

This will create the go.mod file under the MyModule folder.

The newly created go.mod file will have the module name and the go version which the module is targeting. This file will expand as more information is added to the module.

Now, you can start adding files to the newly created module. First, create the `main.go` file to run the go module. The `main.go` file is the starting point of a go program. The name of the file is not important (can have any name) but the `main()` function within this file is the entry point for the program. So having the file name as main.go makes it easier to find the starting point.

## Module Versioning

A specific version of the Go module can be used as Go modules are distributed from a version control repository and they can use version control features like tags, branches, and commits. You can specify the version of the module that will be used in the dependency with `@` symbol at the end of the module path as shown below.

```bash
go get sample.com/sales@latest
```

**The function names inside any file in the go module should be starting with a capital letter.**

For more info: [Using Go Modules](https://go.dev/blog/using-go-modules)
