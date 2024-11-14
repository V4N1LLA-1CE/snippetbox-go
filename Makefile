run:
	go run ./cmd/server/main.go

watch:
	air -- -addr=":8080"

.PHONY: watch run
