# i3refgen
Reads your i3 configuration file and generates a HTML file with the keybindings to print out and put near your screen.

# Build
You need to install [dep](https://golang.github.io/dep/) to manage dependencies.
```
    go get github.com/ronthecookie/i3refgen
    cd $GOPATH/src/github.com/ronthecookie/i3refgen
    dep ensure
    go build
```
# Usage
Just run `i3refgen` and it'll write a out.html in the current directory.
