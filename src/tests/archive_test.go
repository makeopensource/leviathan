package tests

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"github.com/makeopensource/leviathan/utils"
	"io"
	"testing"
)

// TestArchiveJobData tests the ArchiveJobData function.
func TestArchiveJobData(t *testing.T) {
	// Sample files to archive
	files := map[string][]byte{
		"file1.txt": []byte("Hello, World!"),
		"file2.txt": []byte("Go is awesome!"),
	}

	// Create archive
	archive, err := utils.ArchiveJobData(files)
	if err != nil {
		t.Fatalf("Failed to create archive: %v", err)
	}
	defer archive.Close()

	// Read back the archive and extract its contents
	extractedFiles := make(map[string][]byte)
	err = UnarchiveJobData(archive, extractedFiles)
	if err != nil {
		t.Fatalf("Failed to extract archive: %v", err)
	}

	// Verify extracted content
	for name, originalContent := range files {
		extractedContent, exists := extractedFiles[name]
		if !exists {
			t.Errorf("Missing file in extracted output: %s", name)
			continue
		}
		if !bytes.Equal(originalContent, extractedContent) {
			t.Errorf("Content mismatch for file %s", name)
		}
	}
}

// UnarchiveJobData extracts files from a tar.gz archive into a map.
func UnarchiveJobData(archive io.Reader, extractedFiles map[string][]byte) error {
	gzr, err := gzip.NewReader(archive)
	if err != nil {
		return err
	}
	defer gzr.Close()

	tr := tar.NewReader(gzr)
	for {
		header, err := tr.Next()
		if err == io.EOF {
			break // End of archive
		}
		if err != nil {
			return err
		}

		// Read file content
		var buf bytes.Buffer
		if _, err := io.Copy(&buf, tr); err != nil {
			return err
		}

		// Store in map
		extractedFiles[header.Name] = buf.Bytes()
	}
	return nil
}
