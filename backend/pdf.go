package backend

import (
	"bytes"
	"context"
	"fmt"
	"io"

	pdfcpuAPI "github.com/pdfcpu/pdfcpu/pkg/api"
	"golang.org/x/sync/errgroup"
)

type MergeElements []*PDFSelection

type PDFSelection struct {
	Name  string
	Data  []byte
	Pages []string
}

func MergePDFs(elements MergeElements) ([]byte, error) {
	errGroup, _ := errgroup.WithContext(context.Background())

	parts := make([]io.ReadSeeker, len(elements))

	for i, el := range elements {
		func(i int, el *PDFSelection) {
			errGroup.Go(func() error {
				part := bytes.Buffer{}

				err := pdfcpuAPI.Collect(bytes.NewReader(el.Data), &part, el.Pages, nil)
				if err != nil {
					return fmt.Errorf("collecting pages from %s: %w", el.Name, err)
				}

				parts[i] = bytes.NewReader(part.Bytes())

				return nil
			})
		}(i, el)
	}

	err := errGroup.Wait()
	if err != nil {
		return nil, err
	}

	var output bytes.Buffer

	err = pdfcpuAPI.Merge(parts, &output, nil)
	if err != nil {
		return nil, fmt.Errorf("merging parts: %w", err)
	}

	return output.Bytes(), nil
}
