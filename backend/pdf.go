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

func mergePDFs(elements MergeElements) ([]byte, error) {
	if len(elements) == 0 {
		return nil, fmt.Errorf("empty merge list")
	}

	errGroup, _ := errgroup.WithContext(context.Background())

	parts := make([]io.ReadSeeker, len(elements))

	for i, el := range elements {
		errGroup.Go(func() error {
			part := bytes.Buffer{}

			if len(el.Data) == 0 {
				return fmt.Errorf("empty data for pdf %s", el.Name)
			}

			err := pdfcpuAPI.Collect(bytes.NewReader(el.Data), &part, el.Pages, nil)
			if err != nil {
				return fmt.Errorf("collecting pages from %s: %w", el.Name, err)
			}

			parts[i] = bytes.NewReader(part.Bytes())

			return nil
		})
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
