Go Started Snippets
===============

### Install Go & GVM
#### install GVM
```sh
zsh < <(curl -s -S -L https://raw.githubusercontent.com/moovweb/gvm/master/binscripts/gvm-installer)
```

#### Go 可用版本
```sh
gvm listall
```

#### Go 已安装版本
```sh
gvm list
```

#### install Go
```sh
gvm install go1.7.5
gvm use go1.7.5 --default
```

### GOPATH env variable
```sh
    $ mkdir $HOME/stack/goenv
    $ export GOPATH=$HOME/stack/goenv
    $ export PATH=$PATH:$GOPATH/bin
```

### Uninstall Go
```sh
    sudo vi /etc/profile
    rm -rf /usr/local/go
    rm -rf /etc/paths.d/go
```

### Import Paths
```sh
    $ mkdir -p $GOPATH/src/github.com/user
```

### Package
```sh
    $ mkdir $GOPATH/src/github.com/user/hello
```

### Build and Install
```sh
    $ go install github.com/user/hello

    $ cd $GOPATH/src/github.com/user/hello
    $ go install
```

### Run
```sh
    $ $GOPATH/bin/hello
    $ hello
```

### Tutorials
-	[A Tour of Go](http://tour.golang.org) - Interactive tour of Go.
-	[Go By Example](https://gobyexample.com) - A hands-on introduction to Go using annotated example programs.
-	[Working with Go](https://github.com/mkaz/working-with-go) - An intro to go for experienced programmers.
-	[Go database/sql tutorial](http://go-database-sql.org) - Introduction to database/sql.
-	[Go database/sql jmoiron](http://jmoiron.net/blog/gos-database-sql) - Introduction to Go's database/sql.

### Resource
-	[Download the Go distribution & Install the Go tools](https://golang.org/doc/install)
-	[How to Write Go Code](https://golang.org/doc/code.html)
-	[Effective Go](https://golang.org/doc/effective_go.html)

