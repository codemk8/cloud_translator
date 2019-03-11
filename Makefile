
build: ./cmd/*.go 
	go build -o bin/trans_cmd.exe ./cmd/main.go
	go build -o bin/docx_cmd.exe ./cmd/docx.go

test: ./pkg/*/*.go
	go test -v github.com/codemk8/cloud_translator/pkg/...

clean:
	rm -rf bin/*
