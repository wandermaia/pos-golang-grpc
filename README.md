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

## Utilização do protoc

Exemplo:


```bash

protoc --go_out=. --go-grpc_out=. proto/course_category.proto

go mod tidy

```

Serão gerados dois arquivos na pasta `internal/pb`.

## Dicas

Instale a extensão `vscode-proto3` para auxiliar na criaçaõ dos protofiles.

## Evans


Para interagir com o servidor gRPC, vamos utilizar o `Evans`. Ele necessita de reflection.


Criar a tabela no banco de dados

```bash

wander@bsnote283:~/repos/CURSO-GO/12gRPC/pos-golang-grpc$ sqlite3 db.sqlite
SQLite version 3.37.2 2022-01-06 13:25:41
Enter ".help" for usage hints.
sqlite> create table categories (id string, name string, description string);
sqlite> 

```

Abaixo segue um exemplo de utilização do Evans:

```bash

wander@bsnote283:~/repos/CURSO-GO/12gRPC$ evans -r repl

  ______
 |  ____|
 | |__    __   __   __ _   _ __    ___
 |  __|   \ \ / /  / _. | | '_ \  / __|
 | |____   \ V /  | (_| | | | | | \__ \
 |______|   \_/    \__,_| |_| |_| |___/

 more expressive universal gRPC client


127.0.0.1:50051> package pb

pb@127.0.0.1:50051> service CategoryService

pb.CategoryService@127.0.0.1:50051> call CreateCategory
name (TYPE_STRING) => Cat 2
description (TYPE_STRING) => Cat 2 Desc
{
  "description": "Cat 2 Desc",
  "id": "8dd8b38d-7e2c-46ae-b35a-d64d83802ebc",
  "name": "Cat 2"
}

pb.CategoryService@127.0.0.1:50051> 

```


## Referências

gRPC IO

https://grpc.io


Protocol Buffers

https://protobuf.dev/


Protocol Buffer Compiler Installation

https://grpc.io/docs/protoc-installation/


Evans

https://github.com/ktr0731/evans