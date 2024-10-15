package main

import (
    "log"
    "net/http"
)

func main() {

	if err := RabbitMQ(); err != nil {
		log.Fatal(err)
	}
	defer CloseRabbitMQ()

    go startConsumer()

    http.HandleFunc("/download", handleDownload)
	http.HandleFunc("/get_video/{id}", handleVideoDownload)
	http.HandleFunc("/get_thumbnail/{id}", handleThumbnailDownload)

    log.Println("API rodando na porta 3000...")
    if err := http.ListenAndServe(":3000", nil); err != nil {
        log.Fatalf("Erro ao iniciar o servidor: %v", err)
    }
}

