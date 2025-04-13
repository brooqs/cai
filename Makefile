
.PHONY: build docker-build clean docker-run

DIST_DIR=dist

# Build locally (for testing)
build:
	go build -o $(DIST_DIR)/cai ./cmd/cai

# Clean output
clean:
	rm -rf $(DIST_DIR)

# Multi-platform Docker build
docker-build:
	docker build -t cai-multibuild -f Dockerfile .

docker-run:
	docker run --rm -v $(PWD)/$(DIST_DIR):/dist cai-multibuild
