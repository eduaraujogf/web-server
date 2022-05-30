# Go Web I

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