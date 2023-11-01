# Busca-cep-api
Sobe uma API na porta 8080 no qual seu path raiz busca os dados de um determinado cep, informado como query parameter na requisição http GET, no site da viacep. Caso o cep seja encontrado no site então é retornado o JSON com seus respectivos dados.

## Pré-requisitos
Antes de começar, certifique-se de ter as seguintes ferramentas instaladas em seu computador:

* Go (Golang 1.20.6)
* Disponibilidade da porta 8080

## Passo 1: Execute a API
No terminal, navegue até o diretório onde o projeto do aplicativo está localizado. Em seguida, execute o seguinte comando para executá-lo:

`go run main.go buscaCep.go`

## Passo 2: Requisição GET na API
No browser, entre na URL exemplificada abaixo e o retorno deve ser o JSON com os dados.

[http://localhost:8080/?cep=78075300](http://localhost:8080/?cep=78075300)

No exemplo acima o 78075300 é o cep escolhido para a busca.

```{
"cep": "78075-300",
"logradouro": "Rua Sabiá",
"complemento": "",
"bairro": "Recanto dos Pássaros",
"localidade": "Cuiabá",
"uf": "MT",
"ibge": "5103403",
"gia": "",
"ddd": "65",
"siafi": "9067"
}