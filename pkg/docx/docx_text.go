package docx

import (
	"os"

	"code.sajari.com/docconv"
)

func DocxToText(fileName string) (string, map[string]string, error) {
	r, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	return docconv.ConvertDocx(r)
}
