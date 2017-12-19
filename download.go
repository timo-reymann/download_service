package main

import(
	"io"
	"net/http"
	"os"
	"strings"
	"fmt"
	"io/ioutil"
	"archive/tar"
	"compress/gzip"
)

// Download file and return file name as string

func DownloadFile(uuid string, download Download) string {
	r,err := http.Get(download.Url)
	tokens := strings.Split(download.Url, "/")
	filename := uuid + "/" + tokens[len(tokens)-1]
	out, err := os.Create(filename)

	if err != nil {
		return ""
	}

	defer out.Close()
	defer r.Body.Close()

	_,err = io.Copy(out, r.Body)

	if err != nil {
		return ""
	}

	for index, value := range downloads[uuid] {
		if value.Url == value.Url {
			downloads[uuid] = remove(downloads[uuid], index)
			break;
		}
	}

	if len(downloads[uuid]) == 0 {
		delete(downloads, uuid)
		fmt.Println("Clear UUID from list");
		buildTar(uuid)
	}

	return filename
}

func remove(s []Download, i int) []Download {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func buildTar(directory string) {
	files,_ := ioutil.ReadDir("./" + directory)

	file, _ := os.Create(directory + ".tar.gz.tmp")
	defer file.Close()

	gw := gzip.NewWriter(file)
	defer gw.Close()

	tw := tar.NewWriter(gw)
	defer tw.Close()

	for _,f := range files {
		addFile(tw,"./" + directory + "/" + f.Name());
	}

	fmt.Println("Tar created!");
	os.RemoveAll("./" + directory)
	fmt.Println("Delete tmp files")

	fmt.Println("Move file to final")
	os.Rename("./" +  directory + ".tar.gz.tmp", "./" + directory + ".tar.gz")
}
