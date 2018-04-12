# Govm [ WIP ]
Govm is a fast and flexible version manager for Go built with Go.

Govm lets you work with multiple versions of Go, actively.


# installation

## Step 1
#### from binary
download binary from [here](https://github.com/golang-vm/govm/releases)
and copy the binary to a place under `$PATH`.

[or]
#### with Go
```
go get github.com/golang-vm/govm
```

## step 2
#### configure govm
```
govm configure
```

# List versions
List all versions available for download
```
govm ls-remote
```

List versions installed locally.
```
govm ls
```

# Download / Install a Go version
```
govm install 1.10
```

# Use a version
```
govm use 1.10
```

# uninstall a version
```
govm uninstall 1.10
```

# Execute a command with specific version
```
govm exec 1.10 go env
```

# Contributing

Contributions to GoVM are welcome and encouraged.  please read CONTRIBUTING.md

# Bugs ?
Please report them under [issues](https://github.com/golang-vm/govm/issues)