package backend

import (
	"encoding/base64"
	"fmt"
	"strings"
)

// MergePDFs merges selected pages from multiple PDFs
func MergePDFs(data map[string]interface{}) (string, error) {
	pdfsValue, ok := data["mergeElements"]
	if !ok {
		return "", fmt.Errorf("data key 'mergeElements' missing")
	}

	mergeElements, ok := pdfsValue.([]interface{})
	if !ok {
		return "", fmt.Errorf("'mergeElements' value has unexpected type %T", pdfsValue)
	}

	mergeConfig := make([]*PDFSelection, 0, len(mergeElements))

	for i, element := range mergeElements {
		pdfConfig, ok := element.(map[string]interface{})
		if !ok {
			return "", fmt.Errorf("pdf config %d has unexpected type %T", i, element)
		}

		nameInterface, ok := pdfConfig["name"]
		if !ok {
			return "", fmt.Errorf("missing 'name' in pdf config %d", i)
		}

		name, ok := nameInterface.(string)
		if !ok {
			return "", fmt.Errorf("'name' in pdf config %d is not a string but a %T", i, nameInterface)
		}

		dataInterface, ok := pdfConfig["data"]
		if !ok {
			return "", fmt.Errorf("missing 'name' in pdf config %d", i)
		}

		dataStr, ok := dataInterface.(string)
		if !ok {
			return "", fmt.Errorf("'data' in pdf config %d is not a string but a %T", i, nameInterface)
		}

		data, err := base64.StdEncoding.DecodeString(dataStr)
		if err != nil {
			return "", fmt.Errorf("base64 decode data: %w", err)
		}

		pagesInterface, ok := pdfConfig["pages"]
		if !ok {
			return "", fmt.Errorf("missing 'name' in pdf config %d", i)
		}

		pages, ok := pagesInterface.(string)
		if !ok {
			return "", fmt.Errorf("'pages' in pdf config %d is not a string but a %T", i, nameInterface)
		}

		mergeConfig = append(mergeConfig, &PDFSelection{
			Name:  name,
			Data:  data,
			Pages: strings.Split(pages, ","),
		})
	}

	merged, err := mergePDFs(mergeConfig)
	if err != nil {
		return "", fmt.Errorf("merging pdfs: %w", err)
	}

	return base64.StdEncoding.EncodeToString(merged), nil
}
