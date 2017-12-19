package main

type Download struct {
	Url string
}

type RequestDownloads struct {
	Downloads []Download
}

type Response struct {
	Identifier string `json:"identifier"`
}

