Installation
------------

    $ brew install go glfw

Setup your `$GOPATH`.

    $ go get \
        github.com/go-gl/mathgl/mgl32 \
        github.com/go-gl/gl/v4.1-core/gl \
        github.com/go-gl/glfw/v3.2/glfw \
        github.com/slimsag/binpack

Clone the repo into `$GOPATH/src`.

Running
-------

Prepare texture pack:

    $ go run cmd/texpack.go \
        -in resources/ \
        -out resources/assets.tp

Run the example code:

    $ go run cmd/ex01.go
