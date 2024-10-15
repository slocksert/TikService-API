package main

import (
	"encoding/json"
	"log"
	"net/http"
	"path/filepath"

	"github.com/streadway/amqp"
	"github.com/teris-io/shortid"
)

func handleDownload(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	var request VideoDto
	request.ID = shortid.MustGenerate() // Gera um ID único

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		log.Printf("Erro ao decodificar o JSON: %v", err)
		http.Error(w, "Erro ao decodificar o JSON", http.StatusBadRequest)
		return
	}

	// Adiciona (publish) na fila video-sent
		//
	if err != nil {
		log.Printf("Erro ao publicar na fila: %v", err)
		http.Error(w, "Erro ao publicar na fila", http.StatusInternalServerError)
		return
	}

	//log
	log.Printf("Pedido de download recebido: ID=%s, VideoURL=%s", request.ID, request.VideoURL)

	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(map[string]string{"message": "Download em progresso", "id": request.ID, "videoUrl": request.VideoURL})
}

func handleVideoDownload(w http.ResponseWriter, r *http.Request) {
	//
}

// Handler para baixar a thumbnail
func handleThumbnailDownload(w http.ResponseWriter, r *http.Request) {
	//
}
