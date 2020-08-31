package backend

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/ledongthuc/pdf"
)

func mustRead(fileName string) []byte {
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	return content
}

func TestMergePDFsParallel(t *testing.T) {
	var mergeElements = MergeElements{
		{Name: "a.pdf", Data: mustRead("testdata/a.pdf"),
			Pages: []string{"-2", "5"}},
		{Name: "b2.pdf", Data: mustRead("testdata/b.pdf"),
			Pages: []string{"1", "4"}},
		{Name: "a.pdf", Data: mustRead("testdata/a.pdf"),
			Pages: []string{"4-8"}},
	}

	expected := []string{"a1", "a2", "a5", "b1", "b4", "a4", "a5", "a6", "a7", "a8"}

	merged, err := MergePDFs(mergeElements)
	if err != nil {
		t.Fatalf("merging PDFs: %v", err)
	}

	contents, err := readPDFPages(merged)
	if err != nil {
		t.Fatalf("reading merged PDF: %v", err)
	}

	if len(contents) != len(expected) {
		t.Fatalf("unexpected number of pages: %d instead of %d",
			len(contents), len(expected))
	}

	for i, expectedContent := range expected {
		actualContent := contents[i]

		if expectedContent != actualContent {
			t.Errorf("unecpected content on output page %d: %q (expected %q)",
				i, actualContent, expectedContent)
		}
	}
}

func readPDFPages(pdfData []byte) ([]string, error) {
	pdfReader, err := pdf.NewReader(bytes.NewReader(pdfData), int64(len(pdfData)))
	if err != nil {
		return nil, fmt.Errorf("creating pdf reader: %w", err)
	}

	numPages := pdfReader.NumPage()
	contents := make([]string, 0, numPages)

	for i := 1; i <= numPages; i++ {
		p := pdfReader.Page(i)
		if p.V.IsNull() {
			continue
		}

		pageContent, err := p.GetPlainText(nil)
		if err != nil {
			return nil, fmt.Errorf("getting content of page %d: %w", i, err)
		}

		// this pdf library duplicates the last character
		if len(pageContent) > 1 {
			pageContent = pageContent[:len(pageContent)-1]
		}

		contents = append(contents, pageContent)
	}

	return contents, nil
}
