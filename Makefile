run:
	CompileDaemon -build="go build -o ./bin/go-exec" -command="./bin/go-exec"

build:
	go build -o ./bin/go-exec