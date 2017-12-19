package main

import(
	"net/http"
	"encoding/json"
	"io/ioutil"
	"os"
	"fmt"
)

const Port = "8086"

var downloads = make(map[string][]Download)

func main() {
	fmt.Println("Listening on port " + Port)
	http.HandleFunc("/",ServeIndex)
	http.HandleFunc("/request", Request)
	http.HandleFunc("/check", Status)
	http.HandleFunc("/download", DownloadBundle)
	http.ListenAndServe(":" + Port, nil)
}

func ServeIndex(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "form.html")
}

func DownloadBundle(w http.ResponseWriter, r *http.Request) {
	identifier := r.FormValue("identifier")
	f := identifier + ".tar.gz";

	Log(identifier,"Download file")

	w.Header().Set("Content-Disposition", "attachement;filename=" + identifier + ".tar.gz")
	http.ServeFile(w, r, f)

	defer os.Remove(f)
}

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

func Status(w http.ResponseWriter, r *http.Request) {
	identifier := r.FormValue("identifier")
	if len(downloads[identifier]) > 0 {
		w.Write([]byte("DOWNLOADING"))
		return
	}

	if _,err := os.Stat(identifier + ".tar.gz"); os.IsNotExist(err) {
		w.Write([]byte("PACKAGING"))
	} else {
		Log(identifier, "Notified download is complete")
		w.Write([]byte("COMPLETE"))
	}
}
