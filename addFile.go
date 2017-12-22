package main

import(
	"archive/tar"
	"io"
	"os"
)

// Add file to tar writer useing the full path
func addFile(tw * tar.Writer, path string) error {
	// Open file to add to tar
	file, err := os.Open(path)

	// Error opening file
	if err != nil {
		return err
	}

	// Close file when not needed anymore
	defer file.Close()

	// Check file exists
	if stat, err := file.Stat(); err == nil {

		// Create header
		header := new(tar.Header)

		// Set tar entry meta
		header.Name = path
		header.Size = stat.Size()
		header.Mode = int64(stat.Mode())
		header.ModTime = stat.ModTime()

		// write the header to the tarball archive
		if err := tw.WriteHeader(header); err != nil {
			return err
		}

		// copy the file data to the tarball 
		if _, err := io.Copy(tw, file); err != nil {
			return err
		}
	}

	// File doesnt exist return nil
	return nil
}
