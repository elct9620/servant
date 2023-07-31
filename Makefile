DESTDIR=./bin/build

.PHONY: bin-servant
bin-servant:
	CGO_ENABLED=0 go build -trimpath -o "${DESTDIR}/docker-servant" ./cmd/servant

.PHONY: install
install: bin-servant
	mkdir -p ~/.docker/cli-plugins
	install bin/build/docker-servant ~/.docker/cli-plugins/docker-servant

