package docx

import (
	"archive/zip"
	"encoding/xml"
	"io"
	"strings"
)

type Body struct {
	Paragraph []string `xml:"p>r>t"`
}

type Document struct {
	XMLName xml.Name `xml:"document"`
	Body    Body     `xml:"body"`
}

// Extract converts the xml elements into their respective struct fields
func (d *Document) Extract(xmlContent string) error {

	return xml.Unmarshal([]byte(xmlContent), d)
}

// UnpackDocx unzips the doc file
func UnpackDocx(filePath string) (*zip.ReadCloser, []*zip.File) {
	reader, err := zip.OpenReader(filePath)
	if err != nil {
		panic(err)
	}
	return reader, reader.File
}

// WordDocToString converts the file interface object into a raw string
func WordDocToString(reader io.Reader) (content string) {
	_content := make([]string, 100)
	data := make([]byte, 100)

	for {
		n, err := reader.Read(data)
		_content = append(_content, string(data))
		if err == io.EOF && n == 0 {
			break
		}
	}
	content = strings.Join(_content, "")
	return
}

// RetrieveWordDoc simply loops over the files looking for the file with name "word/document"
func RetrieveWordDoc(files []*zip.File) (file *zip.File) {

	for _, f := range files {
		if f.Name == "word/document.xml" {
			file = f
		}
	}
	return
}
