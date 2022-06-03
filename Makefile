build_all: build_gcom build_gpom

build_gcom:
	go build -o bin/gcom ./cmd/gcom

build_gpom:
	go build -o bin/gpom ./cmd/gpom

install:
	go install ./cmd/gpom
	go install ./cmd/gcom
