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
	./batterarch

graph: build
	echo "Generating graph"
	./batterarch graph

json: build
	echo "Generating json"
	./batterarch json

clean:
	rm -f batterarch
