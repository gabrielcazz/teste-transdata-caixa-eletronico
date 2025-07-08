# Desafio Caixa Eletrônico em Golang

Gabriel Feliciano da Silva 
Teste Transdata Desenvolvedor Pleno

Este projeto é uma solução para o desafio de desenvolvimento que simula o funcionamento de um caixa eletrônico. A aplicação, desenvolvida em Golang, recebe um valor monetário e calcula a quantidade mínima de notas de R$200, R$100, R$50, R$20, R$10 e R$5 necessárias para compor esse valor.

## Requisitos

* Go (Golang) instalado em sua máquina.

## Como Executar a Aplicação

1.  **Clone o Repositório**
    ```bash
    git clone [https://github.com/gabrielcazz/teste-transdata-caixa-eletronico.git](https://github.com/gabrielcazz/teste-transdata-caixa-eletronico.git)
    cd teste-transdata-caixa-eletronico
    ```

2.  **Execute o Programa**
    Você pode executar o arquivo `main.go` diretamente através do terminal. O programa solicitará que você insira um valor para o saque.

    ```bash
    go run main.go
    ```

3.  **Exemplo de Uso**
    Ao executar o comando acima, você verá a seguinte mensagem:
    ```
    Digite o valor monetário para o saque:
    ```
    Se você digitar `770` e pressionar Enter, a saída será:
    ```
    Para sacar R$770, você receberá:
    3 nota(s) de R$200
    1 nota(s) de R$100
    1 nota(s) de R$50
    1 nota(s) de R$20
    ```
