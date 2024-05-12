# pos-golang-grpc
Repositório para as aulas de gRPC


## O que é o gRPC

É um framework desenvolvido pelo google com o objetivo de facilitar o processo de comunicação entre sistemas de uma forma extremamente rápida, leve, independente de linguagem. Autualmente ele é mantido plena CNCF.

Casos de uso:

- Ideal para microsserviços
- Mobile, Browsers (ainda não está maduro) e Backend
- Geração de bibliotecas de forma automática
- Streaming bidirecional utilizando HTTP/2

Linguagens com suporte oficial

- gRPC-GO
- gRPC-JAVA
- gRPC-C
    - C++
    - Python
    - Ruby
    - Objective C
    - PHP
    - C#
    - Node.js
    - Dart
    - Kotlin

## Protocol Buffers

"Protocol buffers are Google’s language-neutral, platform-neutral, extensible mechanism for serializing structured data – think XML, but smaller, faster, and simpler."


### Protocol Buffers vs JSON

- Arquivos binários < JSON
- Processo de serialização é mais leve (CPU) do que JSON
- Gasta menos recursos de rede
- Processo é mais veloz.

## HTTP/2

- Nome original criado pela Google era SPDY
- Lançado em 2015
- Dados trafegados em binário e não texto como no HTTP 1.1
- Utiliza a mesma conexão TCP para enviar e receber dados do cliente e do servidor (Multiplex)
- Server push
- Headers são comprimidos
- Gasta menos recursos de rede
- Processo é mais veloz

## Formas de comunicação

- gRPC - API "unary"
- gRPC - API "Server Streaming"
- gRPC - API "Client Streaming"
- gRPC - API "Bi directional Streaming"

## REST vs gRPC

Características do REST:

- Texto / JSON
- Unidirecional
- Alta Latência
- Sem contrato (maior chance de erros)
- Sem suporte a streaming (Request / Response)
- Design pré-definido
- Bibliotecas de terceiro

Características do gRPC:

- Protocol Buffers
- Bidirecional e Assíncrono
- Baixa latência
- Contrato deifido (.proto)
- Suporte a streaming
- Design é livre
- Geração de código

## Pré-requisitos para Go

Instalar o compilador do Protocol Buffer:


```bash

sudo apt update && sudo apt install -y protobuf-compiler
protoc --version  # Ensure compiler version is 3+

```

Instalar plugins para o compilador:

```bash

go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

```

Atualize seu PATH (se necessário) para que o compilador protoc possa encontrar os plugins:

```bash

export PATH="$PATH:$(go env GOPATH)/bin"

```

Feito isso, basta inicializar o projeto normalmente e começar a utilizar:

```bash

go mod init github.com/wandermaia/pos-golang-grpc

```



## Referências

gRPC IO

https://grpc.io

Protocol Buffers

https://protobuf.dev/


Protocol Buffer Compiler Installation

https://grpc.io/docs/protoc-installation/

