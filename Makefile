
build: ./cmd/*.go 
	go build -o bin/trans_cmd.exe ./cmd/main.go
	go build -o bin/docx_cmd.exe ./cmd/docx.go

test:
	go test pkg/...

clean:
	rm -rf bin/*
