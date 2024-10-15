# TikService API

Esta é uma API desenvolvida em Go para gerenciar o envio e download de vídeos e thumbnails. A API utiliza RabbitMQ para comunicação entre serviços e implementa funcionalidades de download automático.

## Funcionalidades

- **Upload de Vídeos**: Recebe URLs de vídeos para processamento.
- **Download Automático**: Permite o download de vídeos e thumbnails via URLs.
- **Integração com RabbitMQ**: Publica e consome mensagens entre serviços para comunicação eficiente.
- **Múltiplas Instâncias de Selenium**: Suporta múltiplas instâncias do serviço de Selenium para processamento paralelo.

## Requisitos

- Go 1.23.2
- RabbitMQ
- Docker (opcional, para ambiente de desenvolvimento)

## Estrutura do Projeto

- `main.go`: Arquivo principal que inicia a aplicação.
- `rabbitmq.go`: Lógica de conexão e comunicação com RabbitMQ.
- `handlers.go`: Funções de manipuladores de requisições para download de vídeos e thumbnails.
- `types.go`: Definições de tipos e estruturas utilizadas na aplicação.
- `consumer.go`: Lógica para consumo de mensagens da fila RabbitMQ.
- `utils.go`: Funções utilitárias que auxiliam em várias operações da aplicação.
- `downloads/`: Diretório para armazenar vídeos e thumbnails baixados.

## Executando a Aplicação

Para rodar a aplicação, use o seguinte comando:

```bash
make run
```