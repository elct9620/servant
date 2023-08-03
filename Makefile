DESTDIR=./bin/build

bin-servant:
	CGO_ENABLED=0 go build -trimpath -o "${DESTDIR}/docker-servant" ./cmd/servant

install: bin-servant
	@mkdir -p ~/.docker/cli-plugins
	install bin/build/docker-servant ~/.docker/cli-plugins/docker-servant

test:
	@go test ./...

serve:
	@go run ./cmd/servantd
