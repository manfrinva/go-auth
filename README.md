Instalar dependências:
- go mod init go-auth
- go get github.com/gorilla/mux github.com/golang-jwt/jwt/v4

Rodar o servidor:
- go run main.go

Testar a API:
Enviar um POST para /login com um corpo JSON:

{
  "email": "user@example.com",
  "password": "password123"
}

Usar o token recebido para acessar a rota protegida /protected com o cabeçalho Authorization.
