package utils

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"compress/zlib"
	"fmt"
	"io"
	"os"
	"path"
)

func UntarTo(r []byte, dst string) error {
	var tarReader *tar.Reader
	uncompressedStream, err := gzip.NewReader(bytes.NewReader(r))
	if err != nil {
		// if it fails we try zlib
		gRead, err2 := zlib.NewReader(bytes.NewReader(r))
		if err2 != nil {
			return fmt.Errorf("ExtractTarGz: NewReader failed: %s", err2.Error())
		}
		tarReader = tar.NewReader(gRead)
	} else {
		tarReader = tar.NewReader(uncompressedStream)
	}

	for {
		header, err := tarReader.Next()

		if err == io.EOF {
			break
		}

		if err != nil {
			return fmt.Errorf("ExtractTarGz: Next() failed: %s", err.Error())
		}

		switch header.Typeflag {
		case tar.TypeDir:
			if err := os.MkdirAll(path.Join(dst, header.Name), 0755); err != nil {
				return fmt.Errorf("ExtractTarGz: Mkdir() failed: %s", err.Error())
			}
		case tar.TypeReg:
			p := path.Join(dst, header.Name)
			if err := os.MkdirAll(path.Dir(p), 0755); err != nil {
				return fmt.Errorf("ExtractTarGz: MkdirAll() failed: %s", err.Error())
			}
			outFile, err := os.Create(p)
			defer outFile.Close()
			if err != nil {
				return fmt.Errorf("ExtractTarGz: Create() failed: %s", err.Error())
			}
			if _, err := io.Copy(outFile, tarReader); err != nil {
				return fmt.Errorf("ExtractTarGz: Copy() failed: %s", err.Error())
			}
			if err := os.Chmod(p, os.FileMode(header.Mode)); err != nil {
				return fmt.Errorf("ExtractTarGz: Chmod() failed: %s", err.Error())
			}

		default:
			return fmt.Errorf("ExtractTarGz: uknown type in %s", header.Name)
		}
	}
	return nil
}
