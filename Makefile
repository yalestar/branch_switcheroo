LDFLAGS="-X 'main.version=`date -u +"%Y%m%d_%H%M%S"`'"

build_all: build_gcom build_gpom build_gcb

build_gcom:
	go build -ldflags=${LDFLAGS} -o bin/gcom ./cmd/gcom

build_gpom:
	go build -ldflags=${LDFLAGS} -o bin/gpom ./cmd/gpom

build_gcb:
	go build -ldflags=${LDFLAGS} -o bin/gcb ./cmd/gcb

install:
	go install -ldflags=${LDFLAGS} ./cmd/gpom
	go install -ldflags=${LDFLAGS} ./cmd/gcom
	go install -ldflags=${LDFLAGS} ./cmd/gcb
