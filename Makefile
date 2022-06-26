build_all: build_gcom build_gpom build_gcb

build_gcom:
	go build -o bin/gcom ./cmd/gcom

build_gpom:
	go build -o bin/gpom ./cmd/gpom

build_gcb:
	go build -o bin/gcb ./cmd/gcb

install:
	go install ./cmd/gpom
	go install ./cmd/gcom
	go install ./cmd/gcb
