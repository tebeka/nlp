all:
	$(error please pick a target)

build-docker:
	docker build -f ./cmd/nlpd/Dockerfile -t tebeka/nlpd .

test:
	go test -v ./...

benchmark: go test -bench . -run '^$'
