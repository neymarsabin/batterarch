##
# batterarch
#
# @file
# @version 0.1
build:
	echo "Building batterarch"
	go build -o batterarch

server: build
	echo "Running batterarch"
	./batterarch server

graph: build
	echo "Generating graph"
	./batterarch graph

json: build
	echo "Generating json"
	./batterarch json

test: build
	echo "Running tests"
	./batterarch json
	./batterarch graph
	./batterarch server

clean:
	rm -f batterarch
