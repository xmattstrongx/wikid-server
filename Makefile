GOPATH=${PWD}

build: clean lint compile test

clean:
	rm -rf bin/**

lint:
	go vet `go list wikid-server/... | grep -v /vendor/`

compile:
	go install wikid-server

test:
	go test `go list wikid-server/... | grep -v /vendor/`

# ---------------------------------------------------------------------------- #

dev: build
	cd bin && ./wikid-server
