repackr
==========
`repackr` combines multiple different Go projects with `main` packages into a single Go program or library.

`repackr` is useful in situations where one may want to vendor and use the functionality of multiple Go `main`
packages that don't provide libraries for accessing their functionality. Because `repackr` programmatically rewrites
the packages and provides an invocation mechanism for them, it provides a maintainable solution for using functionality
provided by such `main` packages without having to manually fork or rewrite the libraries.

`repackr` takes a configuration file that specifies the `main` packages that should be combined and an output
directory as arguments. It then does the following:

* Creates a directory named `internal` in the output directory
  * This directory acts as a de facto vendoring directory with rewritten imports
* Copies the projects for the specified inputs into the `repacked` directory
  * By default, it is assumed that the `main` package is the project root
  * If this is not the case, the configuration can be used to specify the degrees of separation between the `main`
    package and the root of the project package
* Rewrites all of the imports of the copied projects to point to the copied version in `repacked`
* All files that have a package value of `main` are renamed to `repacked`
  * Only the package name in the Go file is changed (the name of the directory containing the file will not be changed)
  * The `main` function is renamed to `RepackedMain`
* Writes a new Go file `{{package_name}}.go` in the output directory
  * If the specified package name is `main`, the Go file that is written contains a `main` function that provides a way to invoke the repacked commands by name
  * If the specified package name is not `main`, a library Go file is written. The library file contains a `Run` method
    that allows the wrapped program to be invoked by name and a `Commands` method that returns the valid commands

Usage
-----
Install the package:

```
go get github.com/sniperkit/snk.fork.michaelsauter-crane/cmd/repackr
```

Run the command:

```
repackr --config repackr.yml --output-dir outpkg --pkg main
```

The above command runs `repackr` on the files specified in `repackr.yml` and writes the output source files into a
new directory called `outpkg`. `outpkg` will contain an `repacked` directory that contains all of the repacked
projects and a `main.go` file that contains a `main` method for invoking the repacked libraries.

Configuration
-------------
`repackr` uses a configuration file to determine the packages that should be used as input and the name of the 
command that should be used for that package. The configuration is a `yml` file that contains an entry for each program
that should be repackaged:

```yml
packages:
  sample:
    main: github.com/sniperkit/go-sample
  inner:
    main: github.com/sniperkit/go-project/main
    distance-to-project-pkg: 1
```

Each package must have a unique name (this will be the value that the generated Go wrapper will use to reference the
program). The package must specify a `main` package. The package will be resolved in the same way it would if it were in
a Go source file contained in the output directory (including vendoring behavior). If the program being wrapped is in a
subdirectory of a main project, then the `distance-to-project-pkg` parameter can be used to specify the distance between
the `main` package and the project root package. When a program is being wrapped, the project package is copied into the
vendor directory of the output directory, so this parameter can be used in cases where the `main` package is in a
subdirectory of a project but more files need to be copied in order for the import to function correctly.
