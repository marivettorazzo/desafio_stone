# API Desafio Stone
API com rotas GET/POST em Rest CRUD

### Dependências
github.com/gorilla/mux v1.8.0

github.com/mitchellh/hashstructure v1.1.0 

## Versão - go1.16.6

### Funcionalidades da Aplicação

- Rotas: (get) localhost:3000/accounts
  * mostra todas as contas

- Rotas: (post) localhost:3000/accounts 
  * Criar uma conta 

- Rotas: (get) localhost:3000/transfer
  * Para mostrar todas as transações 

- Rotas: (Post) localhost:3000/transfer
  * Realiza uma transação se:
    * Os IDS das contas de origem e de destino existirem & forem diferentes.
    * O saldo da conta de origem seja maior que o valor transferido e o valor transferido seja um valor válido. 

- Rotas: (Post) localhost:3000/login
  * Valida a senha e o cpf do usuário. 

### Para rodar aplicação
Baixe as depêndencias citadas acima.

Foi ultizado o aplicativo Postman para acessos as funcionalidades das rotas.

A primeira rota a ser utilizada deve ser
  * /accounts  como o método Post 
  * Content-Type: application/json


Inserindo os dados da conta a ser criada ex:
* Status code: 200
* Content-Type: application/json
* Body 
```
{
    "Name": "string",
    "Cpf": int,
    "Balance" : float,
    "Secret" : "string"
}
```
Essa Rota tem os seguintes retornos.

Caso o CPF inserido já exista :

```
"invalid cpf"
```
Caso  o CPF contenha menos de 11 dígitos:
```
"invalid cpf"
```
Caso  o nome esteja com zero caracteres:
```
"invalid name"
```

A rota para ver as contas existentes.
  * Status code: 200
  * Content-Type: application/json
  * /accounts com o método Get.

A Rota para ver o valor do saldo existente em uma conta especifica passando ID como paramêtro.
  * Status code: 200
  * Content-Type: application/json
  * /accounts/id/balance

A rota para realizar uma transferência.
  * /transfer com o metodo Post.
  * Status code: 200
  * Content-Type: application/json
  * Body : 
```
{
    "Amount" : float,
    "Account_origin_id":int,
    "Account_destination_id":int
}
```
Essa Rota tem os seguintes retornos.

Caso o saldo da conta seja menor do que o valor transferido:
```
"Not enough balance for transaction"
```
Caso o valor a ser transferido seja menor ou igual 0:

```
"Invalid transfer value"
```
Caso o id inserido no corpo da requisição não exista:
```
"invalid id"
```
Caso os ids de origem e destino sejam iguais:
```
"the ids must be different"
```
Caso todas as informações sejam válidas : 
```
"Success"
```
A rota mostra todas transações realizadas com sucesso.
  * Status code: 200
  * Content-Type: application/json
  * /transfer com o metodo Get.

A rota para validação de login deve se inserir as seguintes informações :
  * /login
  * Status code: 200
  * Content-Type: application/json
  * Body : 
```
{
    "Cpf" : int,
    "Secret" : "string"
}
```
Essa Rota tem os seguintes retornos.
Caso as informações não sejam identicas as que foram salvas no "banco".
```
"Access denied"
```
Caso as informações sejam identicas as que foram salvas no "banco" ele traz as informações da conta inserida. 
```
{
    "ID": 1680,
    "Name": "Mariana",
    "Cpf": 25288855251,
    "Balance": 1000,
    "Secret": "f065e062a06bf53438210057971560e4a499826a08b083d0ab9a47089dcd08a6",
    "create_at": "2021-10-26T11:32:28.6445063-03:00"
}
```
O comando utilizado para rodar o servidor e utilizar as rotas:
no terminal. ( *atalho ctrl+j* )
  * Navegue até o diretório da aplicação utilizando o seguinte comando:
```
cd main
```
Ainda no terminal execute:
```
go run main.go
```

Executando comando para testes unitários:
```
go test
```
Resultados dos testes feitos até o momento:
```
$ go test
PASS
ok      exemple.com/DesafioStone/desafio_stone/domain/usecases  0.068s
```










