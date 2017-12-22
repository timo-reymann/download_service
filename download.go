package main

import (
	"io"
	"net/http"
	"os"
	"strings"
	"io/ioutil"
	"archive/tar"
	"compress/gzip"
)

// Download file and return file name as string
func DownloadFile(uuid string, download Download, w http.ResponseWriter) string {
	r, err := http.Get(download.Url)

	// Error getting file from remote server
	if err != nil {
		Log(uuid, "Aborting download, error download file "+download.Url)
		removeDownloadByUrl(uuid)

		// War nur ein Link
		if len(downloads[uuid]) == 0 {
			delete(downloads, uuid)
			Log(uuid, "Build empty tar file and clear uuid from list");
			buildTar(uuid)
		}
		return ""
	}

	// Get parts of url
	tokens := strings.Split(download.Url, "/")

	// Get last part (hopefully filename)
	filename := uuid + "/" + tokens[len(tokens)-1]

	// Create downloaded file
	out, err := os.Create(filename)

	// Error creating downloaded file
	if err != nil {
		return ""
	}

	defer out.Close()
	defer r.Body.Close()

	_, err = io.Copy(out, r.Body)

	if err != nil {
		return ""
	}

	removeDownloadByUrl(uuid)

	if len(downloads[uuid]) == 0 {
		delete(downloads, uuid)
		Log(uuid, "Clear UUID from list");
		buildTar(uuid)
	}

	return filename
}

// Remove download from download map indexed by uuid
func removeDownloadByUrl(uuid string) {

	for index, value := range downloads[uuid] {
		if value.Url == value.Url {
			downloads[uuid] = remove(downloads[uuid], index)
			break
		}
	}
}

// Remove download from downloads list by index
func remove(s []Download, i int) []Download {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func buildTar(directory string) {
	files, _ := ioutil.ReadDir("./" + directory)

	// Create file
	file, _ := os.Create(BuildFileName(directory) + ".tmp")
	defer file.Close()

	// Create the new gz writer
	gw := gzip.NewWriter(file)
	defer gw.Close()

	// Create new tar writer
	tw := tar.NewWriter(gw)
	defer tw.Close()

	// Add all files of uuid directory
	for _, f := range files {
		addFile(tw, "./"+directory+"/"+f.Name());
		Log(directory, "Adding file "+f.Name())
	}

	Log(directory, "Tar created!")

	// Delete folde rnamed as identifier and it contents
	defer os.RemoveAll("./" + directory)
	Log(directory, "Delete tmp files")

	// Wait for file to be closed an move it to final name without .tmp at end
	Log(directory, "Move file to final")
	defer os.Rename(BuildFileName(directory)+".tmp", BuildFileName(directory))
}
