# Go Web I

## Aula 1
## Exercise 1 - Structuring a JSON

Depending on the chosen theme, generate a JSON that meets the following keys accordingly
with the theme.  

Products vary by id, name, color, price, stock, code (alphanumeric), publication
(yes-no), creation date.

Users vary by id, first name, last name, email, age, height, active (yes-no), date of
creation.

Transactions: id, transaction code (alphanumeric), currency, amount, sender (string), receiver
(string), transaction date.
>1. Inside the go-web folder create a theme.json file, the name has to be the theme
chosen, e.g. products.json.
>2. Inside it, I wrote a JSON that allows having an array of products, users or
transactions with all their variants.

## Exercise 2 - Hello {name}
>1. Create inside the go-web folder a file called main.go
>2. Create a web server with Gin that returns a JSON that has a key
“message” and say hello followed by your name.
>3. Access the endpoint to verify that the answer is correct.


## Exercise 3 - List Entity

Having already created and tested our API that receives us, we generate a route that returns a list
of the chosen theme.
>1. Inside "main.go", create a structure according to the theme with the fields
correspondents.
>2. Create an endpoint whose path is /thematic (plural). Example: “/products”
>3. Create a handler for the endpoint called "GetAll".
>4. Create a slice of the structure and return it through our endpoint.

## Aula 2
## Exercício 1 - Vamos filtrar nosso endpoint

Dependendo do tema escolhido, precisamos adicionar filtros ao nosso endpoint, ele deve ser
capaz de filtrar todos os campos.
1. Dentro do manipulador de endpoint, recebi os valores para filtrar do contexto.
2. Em seguida, ele gera a lógica do filtro para nossa matriz.
3. Retorne a matriz filtrada por meio do endpoint.

## Exercício 2 - Get one endpoint

Gere um novo endpoint que nos permita buscar um único resultado do array de temas.
Usando parâmetros de caminho o endpoint deve ser /theme/:id (lembre-se que o tema
sempre tem que ser plural). Uma vez que o id é recebido, ele retorna a posição
correspondente.
1. Gere uma nova rota.
2. Gera um manipulador para a rota criada.
3. Dentro do manipulador, procure o item que você precisa.
4. Retorna o item de acordo com o id.
Se você não encontrou nenhum elemento com esse id retorne como código de resposta 404.

# Go Web II

## Aula 1

## Exercício 1 - Criar Entidade
A funcionalidade para criar a entidade deve ser implementada. Se isso acontecer, os
seguintes passos devem ser seguidos:
1. Crie um endpoint por meio de POST que receba a entidade.
2. Você deve ter um array da entidade na memória (no nível global), no qual todas as
requisições que são feitas devem ser salvas.
3. No momento de fazer a solicitação, o ID deve ser gerado. Para gerar o ID, devemos
procurar o ID do último registro gerado, incrementá-lo em 1 e atribuí-lo ao nosso novo
registro (sem ter uma variável global do último ID).

## Exercício 2 - Validação de campo
As validações dos campos devem ser implementadas no momento do envio do pedido, para
isso devem ser seguidos os seguintes passos:
1. Todos os campos enviados na solicitação devem ser validados, todos os campos são
obrigatórios
2. Caso algum campo não esteja completo, um código de erro 400 deve ser retornado
com a mensagem “campo %s é obrigatório”.
(Em %s deve ir o nome do campo que não está completo).

## Exercício 3 - Validar Token
Para adicionar segurança à aplicação, o pedido deve ser enviado com um token, para isso
devem ser seguidos os seguintes passos:
1. No momento do envio da solicitação, deve ser validado que um token é enviado
2. Esse token deve ser validado em nosso código (o token pode ser codificado
permanentemente).
3. Caso o token enviado não esteja correto, devemos retornar um erro 401 e uma
mensagem que "você não tem permissão para fazer a solicitação solicitada".

## Aula 2

## Exercício 1 - Criar Entidade
A funcionalidade para criar a entidade deve ser implementada. Se isso acontecer, os
seguintes passos devem ser seguidos:
1. Crie um endpoint por meio de POST que receba a entidade.
2. Você deve ter um array da entidade na memória (no nível global), no qual todas as
requisições que são feitas devem ser salvas.
3. No momento de fazer a solicitação, o ID deve ser gerado. Para gerar o ID, devemos
procurar o ID do último registro gerado, incrementá-lo em 1 e atribuí-lo ao nosso novo
registro (sem ter uma variável global do último ID).

## Exercício 2 - Validação de campo
As validações dos campos devem ser implementadas no momento do envio do pedido, para
isso devem ser seguidos os seguintes passos:
1. Todos os campos enviados na solicitação devem ser validados, todos os campos são
obrigatórios
2. Caso algum campo não esteja completo, um código de erro 400 deve ser retornado
com a mensagem “campo %s é obrigatório”.
(Em %s deve ir o nome do campo que não está completo).

## Exercício 3 - Validar Token
Para adicionar segurança à aplicação, o pedido deve ser enviado com um token, para isso
devem ser seguidos os seguintes passos:
1. No momento do envio da solicitação, deve ser validado que um token é enviado
2. Esse token deve ser validado em nosso código (o token pode ser codificado
permanentemente).
3. Caso o token enviado não esteja correto, devemos retornar um erro 401 e uma
mensagem que "você não tem permissão para fazer a solicitação solicitada".