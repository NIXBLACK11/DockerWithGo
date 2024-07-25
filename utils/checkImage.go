package utils

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"os"
	"path/filepath"
)

// ExtractAndRenameTarGz extracts a .tar.gz file into a specified directory
func ExtractAndRenameTarGz(tarGzPath, newFolderName string) error {
	// Open the .tar.gz file
	file, err := os.Open(tarGzPath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Create a gzip reader
	gzr, err := gzip.NewReader(file)
	if err != nil {
		return err
	}
	defer gzr.Close()

	// Create a tar reader
	tr := tar.NewReader(gzr)

	// Ensure the newFolderName directory exists
	if err := os.MkdirAll(newFolderName, os.ModePerm); err != nil {
		return err
	}

	for {
		// Get the next file from the archive
		header, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		// Construct the file path inside the new folder
		fpath := filepath.Join(newFolderName, header.Name)

		switch header.Typeflag {
		case tar.TypeDir:
			// Create directories
			if err := os.MkdirAll(fpath, os.ModePerm); err != nil {
				return err
			}
		case tar.TypeReg:
			// Create the necessary directories
			if err := os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
				return err
			}

			// Create the file and write the contents
			outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.FileMode(header.Mode))
			if err != nil {
				return err
			}
			if _, err := io.Copy(outFile, tr); err != nil {
				outFile.Close()
				return err
			}
			outFile.Close()
		}
	}

	return nil
}

// FolderExists checks if a folder exists
func FolderExists(folderPath string) bool {
	info, err := os.Stat(folderPath)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}

// CheckImage checks if the rootfs folder exists, and extracts the tar.gz file if it does not
func CheckImage() error {
	if !FolderExists("./rootfs") {
		err := ExtractAndRenameTarGz("ubuntu-base-14.04-core-i386.tar.gz", "rootfs")
		if err != nil {
			return err
		}
	}
	return nil
}
