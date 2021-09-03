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
  * Mostra todas as transações 

- Rotas: (Post) localhost:3000/transfer
  * Realiza uma transação se:
    * Os IDS das contas de origem e de destino existirem.
    * O saldo da conta de origem seja maior que o valor transferido. 

- Rotas: (Post) localhost:3000/login
  * Valida a senha e o cpf do usuário. 

### Para rodar aplicação
Baixe as depêndencias citadas acima.

Foi ultizado o aplicativo Postman para acessos as funcionalidades das rotas.

A primeira rota a ser utilizada deve ser
  * /accounts  como o método Post 

Crie duas contas para testar outras rotas.

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
A rota para ver as contas existentes.
  * /accounts com o método Get.

A Rota para ver o valor do saldo existente em uma conta especifica passando ID como paramêtro.
  * /accounts/id/balance

A rota para realizar uma transferência.
  * /transfer com o metodo Post.
  * Status code: 200
  * Content-Type: application/json
  * Body
```
{
    "Amount" : float,
    "Account_origin_id":int,
    "Account_destination_id":int
}
```
Essa Rota tem os seguintes retornos.

Caso o saldo da conta seja menor do que o valor
transferido:
```
"Insufficient balance of source account"
```
Caso o id inserido no corpo da requisição não exista:
```
"id does not exist"
```
Caso todas as informações sejam válidas : 
```
"Transaction performed successfully"
```
A rota mostrar todas transações realizadas com sucesso.
  * /transfer com o metodo Get.

A rota para validação de login deve se inserir as seguintes informações :
  * /login
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
Caso as informações sejam identicas as que foram salvas no "banco".
```
"Access allowed"
```
O comando utilizado para rodar o servidor e utilizar as rotas:
no terminal. ( *atalho ctrl+j* )
  * Navegue até o diretório da aplicação utilizando o seguinte comando:
```
cd src
```
Ainda no terminal execute:
```
go run
```

Execute o comando:
```
go test
```
Resultados dos testes feitos até o momento:
```
PASS
ok      exemple.com/DesafioStone/src    0.157s
```










