all: spacedex

spacedex:
	go build

test: test-models


test-models:
	go test github.com/yawnt/index.spacedock/models -c
	./models.test
