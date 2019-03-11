package main

import (
	"fmt"

	"github.com/codemk8/cloud_translator/pkg/docx"
	"github.com/codemk8/cloud_translator/pkg/translate"
)

func main() {
	str1, strMap, err := docx.DocxToText("./002.docx")
	if err != nil {
		panic(err)
	}
	fmt.Printf("str %v ", str1)
	fmt.Printf("strMap %v", strMap)
	eng, err := translate.Translate(&str1)
	if err != nil {
		panic(err)
	}
	fmt.Printf("translated %s.", eng)
	/*
		r, err := docx.ReadDocxFile("./002.docx")
		if err != nil {
			panic(err)
		}
		docx1 := r.Editable()
		err = docx1.Translate()
		if err != nil {
			panic(err)
		}
		docx1.WriteToFile("./002_en.docx")
	*/
}
