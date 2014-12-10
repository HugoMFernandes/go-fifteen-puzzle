# Go Fifteen Puzzle

A simple (term-based) Go implementation of 15-puzzle. It relies on termbox for rendering, so it should work on any platform. 

Note: I wrote this while learning Go's syntax and package structure, so the code could probably use some improvements (this was literally the first program I wrote after completing the tour).

## Installing
Make sure you have both Git and Go installed and that your `$GOPATH` is defined.

### Dependencies
You need to get [termbox] (used for rendering everything on the screen regardless of the platform).

### Install command
Get the package
```
$ go get github.com/HugoMFernandes/go-fifteen-puzzle
```
Move to the package directory and install it to `$GOPATH/bin`
```
$ go install
```

## Running
Run a simple 15-puzzle (4x4, with one empty space)
```
$ go-fifteen-puzzle
```
Or specify custom width/heights for the puzzle (as long as they fit on the terminal!)
```
$ go-fifteen-puzzle [-w width] [-h height]
```

[termbox]:https://github.com/nsf/termbox-go
