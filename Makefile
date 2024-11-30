buildall:
	go build -o bin ./days/...
.PHONY: buildall

testall:
	go test ./
.PHONY: testall