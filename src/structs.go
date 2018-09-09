package main

// Requested download for downloads list sent via form
type Download struct {
	Url string
}

// Requested downloads sent via form
type RequestDownloads struct {
	Downloads []Download
}

// Response for download request if successful
type Response struct {
	Identifier string `json:"identifier"`
}

// Api structure for uuid
type ApiUUIDResponse struct {
	Status   int      `json:"status"`
	Messages []string `json:"messages"`
	Data     UUID     `json:"data"`
}

// UUID wrapped in api structure for uuid
type UUID struct {
	Uuid    string `json:"uuid"`
	Version int    `json:"version"`
}
