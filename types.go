package main

type VideoDto struct {
	ID       string `json:"id"`
	VideoURL string `json:"videoUrl"`
}

type DataDto struct {
	ID          string `json:"id"`
	VideoName   string `json:"videoName"`
	VideoURL    string `json:"videoUrl"`
	ChannelName string `json:"channelName"`
	Cookies     string `json:"cookies"`
}
