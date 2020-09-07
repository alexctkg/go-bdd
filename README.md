
## Getting started
1. Clone o repositório  no `$GOPATH/src` 
2. Entre na pasta `$ cd Gogo-bdd`
3. Execute os testes `$ godog`
4. Executar o projeto `$ go run main.go` 


# Observações
O teste contêm apenas a aplicação dos endpoints que coletam a máxima velocidade e a última velocidade pelo id.
Os enpoints podem ser acessados pelas rotas
- http://localhost:8080/max-speed-allowed
- http://localhost:8080/last-speed?id=

Para acessar o endpoint  com o id inserido no meio da url (/{id}/last-speed) seria necessãrio o uso de um framework web como o gin e o mux, porém como o godog usa os protocolos net/http simplifiquei e fiz dessa maneira com um parâmetro. Além do mais, o teste desse endpoint falha na response. Como a linguagem de Go usa muito o tratamento de erros resolvi
tratar os erros, isso entrou em conflito com a feature descrita.