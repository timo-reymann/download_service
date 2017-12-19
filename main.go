package main

import(
	"net/http"
	"encoding/json"
	"io/ioutil"
	"os"
	"fmt"
)

// Application embed server port to listen
const Port = "8086"

// Download storage for current downloads
var downloads = make(map[string][]Download)

// Entry point
func main() {
	fmt.Println("Listening on port " + Port)
	http.HandleFunc("/",ServeIndex)
	http.HandleFunc("/request", Request)
	http.HandleFunc("/check", Status)
	http.HandleFunc("/download", DownloadBundle)
	http.ListenAndServe(":" + Port, nil)
}

// Serve form for sending requests useing ajax
func ServeIndex(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "form.html")
}

// Download bundle requested
func DownloadBundle(w http.ResponseWriter, r *http.Request) {
	identifier := r.FormValue("identifier")
	f := identifier + ".tar.gz";

	Log(identifier,"Download file")

	w.Header().Set("Content-Disposition", "attachement;filename=" + identifier + ".tar.gz")
	http.ServeFile(w, r, f)

	defer os.Remove(f)
}

// Request file download
func Request(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	var req RequestDownloads
	err = json.Unmarshal(body, &req)
	if err != nil {
		panic(err)
	}

	fmt.Println(req)

	uuid := GenerateUUID()
	json, err := json.Marshal(Response{uuid})
	downloads[uuid] = req.Downloads
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)

	Log(uuid, "Start downloads")

	if len(downloads[uuid]) > 0 {
		os.MkdirAll(uuid,os.ModePerm)
	}

	for _, dl := range downloads[uuid] {
		go DownloadFile(uuid, dl)
	}
}

// Get current status for download
func Status(w http.ResponseWriter, r *http.Request) {
	identifier := r.FormValue("identifier")

	// Some downloads are missing to make tar archive
	if len(downloads[identifier]) > 0 {
		w.Write([]byte("DOWNLOADING"))
		return
	}

	// Tar file is not created now or download is complete
	if _,err := os.Stat(identifier + ".tar.gz"); os.IsNotExist(err) {
		w.Write([]byte("PACKAGING"))
	} else {
		Log(identifier, "Notified download is complete")
		w.Write([]byte("COMPLETE"))
	}
}
