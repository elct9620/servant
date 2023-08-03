DESTDIR=./bin/build

bin-servant:
	CGO_ENABLED=0 go build -trimpath -o "${DESTDIR}/docker-servant" ./cmd/servant

servantd:
	@docker build -t ghcr.io/elct9620/servantd:dev .

install: bin-servant
	@mkdir -p ~/.docker/cli-plugins
	install bin/build/docker-servant ~/.docker/cli-plugins/docker-servant

test:
	@go test -v ./...

serve:
	@go run ./cmd/servantd

ping:
	@go run ./cmd/servantd ping
