all: spacedex

spacedex:
	go build

test: test-models


test-models:
	go test github.com/spacedock-io/index/models -c
	./models.test
