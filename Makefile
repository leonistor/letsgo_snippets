test:
	go test -v ./...

watch:
	watchexec -r go run ./cmd/web -e go,html,tmpl
