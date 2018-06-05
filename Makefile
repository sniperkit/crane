test:
	@(go list ./... | grep -v "vendor/" | xargs -n1 go test -v -cover)

fmt:
	@(gofmt -w crane)

build: build-linux build-darwin build-darwin-pro build-windows build-windows-pro

build-linux:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o crane_linux_amd64 -v github.com/sniperkit/crane

build-darwin:
	GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 go build -o crane_darwin_amd64 -v github.com/sniperkit/crane

build-darwin-pro:
	GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 go build -tags pro -o crane_darwin_amd64_pro -v github.com/sniperkit/crane

build-windows:
	GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -o crane_windows_amd64.exe -v github.com/sniperkit/crane

build-windows-pro:
	GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -tags pro -o crane_windows_amd64_pro.exe -v github.com/sniperkit/crane
