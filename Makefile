test:
	go test -v ./...

watch:
	watchexec -r go run ./cmd/web -e go,html,tmpl

gen_tls_cert:
	go run $GOROOT/src/crypto/tls/generate_cert.go --rsa-bits=2048 --host=localhost

genkey:
	openssl rand -base64 32
