package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

//Consumer que irá consumir a fila data-sent
func startConsumer() {
    msgs, err := channel.Consume("data-sent", "", true, false, false, false, nil)
    if err != nil {
        log.Fatalf("Erro ao consumir a fila data-sent: %v", err)
    }

    for msg := range msgs {
        var data DataDto
        if err := json.Unmarshal(msg.Body, &data); err != nil {
            log.Printf("Erro ao decodificar mensagem: %v", err)
            continue
        }

        log.Printf("Dados recebidos da fila data-sent: %+v\n", data)

        log.Printf("Baixando vídeo: %s\n", data.VideoURL)

        // Cria uma nova requisição HTTP
            //implementar requisição HTTP

        // Adiciona headers
        req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/105.0.0.0 Safari/537.36")
        req.Header.Set("Referer", "https://www.tiktok.com/")
        req.Header.Set("Accept", "application/json, text/plain, */*")
        req.Header.Set("Accept-Language", "en-US,en;q=0.9")
        req.Header.Set("Connection", "keep-alive")
        req.Header.Set("Cookie", data.Cookies)

        // Request para baixar o video
            //

        // Cria o diretório se não existir
            //

		// Adicionando .mp4 ao arquivo
            //

        // Copia o conteúdo do corpo da resposta para o arquivo
            //

        //log
        log.Printf("Download concluído: %s\n", outFileName)

        // Gera a thumbnail
        thumbnailPath := filepath.Join("downloads", data.ID+".png")
        // chamar função para gerar thumbnail

        log.Printf("Thumbnail gerada: %s\n", thumbnailPath)
    }
}