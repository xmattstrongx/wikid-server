GOPATH=${PWD}

build: clean lint compile test package

clean:
	rm -rf bin/**

lint:
	go vet `go list wikid-server/... | grep -v /vendor/`

compile:
	go install wikid-server

test:
	go test `go list wikid-server/... | grep -v /vendor/`

package:
	cp -r src/swagger-ui bin

# ---------------------------------------------------------------------------- #

dev: build
	cd bin && ./wikid-server
