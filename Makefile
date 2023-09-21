build:
	@go build  -o bin/main cmd/main.go
run: build
	@ ./bin/main
build-metriko:
	@go build  -o bin/main2 metriko-agent/main.go
run-metriko: build-metriko
	@ ./bin/main2