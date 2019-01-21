MOD=gooey
TESTS=somemod someothermod .


all:
	make test
	make build

clean:
	rm ${MOD}

test:
	for TEST in $(TESTS); do \
		go test -v ${MOD}/$$TEST; \
	done

build:
	echo $$GOPATH
	go build -o ${MOD}
