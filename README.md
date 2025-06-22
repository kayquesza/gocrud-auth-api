# gocrud-auth-api

API RESTful de autenticação e CRUD de usuários utilizando Go, Gin e MongoDB.

## ✨ Tecnologias Utilizadas
- Go (Golang)
- Gin (framework web)
- MongoDB (banco de dados)
- JWT (autenticação)
- Zap (logger)

## ⚙️ Variáveis de Ambiente
Crie um arquivo `.env` na raiz do projeto com base no `.env.example`:

```
MONGODB_URL=mongodb://localhost:27017
MONGODB_USER_DB=crudInit
MONGODB_USER_COLLECTION=users
JWT_SECRET_KEY=your-secret-key-here
LOG_OUTPUT=stdout
LOG_LEVEL=info
```

> Se `MONGODB_USER_COLLECTION` não for definida, será usado o valor padrão `users`.

## 🚀 Como rodar localmente
1. **Clone o repositório:**
   ```sh
   git clone https://github.com/seu-usuario/gocrud-auth-api.git
   cd gocrud-auth-api
   ```
2. **Configure o arquivo `.env`:**
   ```sh
   cp .env.example .env
   # Edite o arquivo conforme necessário
   ```
3. **Suba o MongoDB localmente** (ou use Docker):
   ```sh
   docker run -d -p 27017:27017 --name mongo mongo:latest
   ```
4. **Instale as dependências:**
   ```sh
   go mod tidy
   ```
5. **Rode a aplicação:**
   ```sh
   go run main.go init_dependencies.go
   # ou
   go build -o gocrud-auth-api.exe
   ./gocrud-auth-api.exe
   ```

## 🛠️ Makefile
O projeto possui um `Makefile` para facilitar o desenvolvimento. Os principais comandos são:

| Comando                | Descrição                                 |
|------------------------|--------------------------------------------|
| `make help`            | Lista todos os comandos disponíveis        |
| `make install`         | Instala as dependências do projeto         |
| `make build`           | Compila o projeto                         |
| `make run`             | Roda a aplicação localmente                |
| `make run-build`       | Compila e roda a aplicação                |
| `make test`            | Executa os testes                         |
| `make clean`           | Remove arquivos de build                   |
| `make docker-build`    | Build da imagem Docker                     |
| `make docker-run`      | Sobe tudo com Docker Compose               |
| `make docker-stop`     | Para os containers Docker                  |
| `make mongo-start`     | Sobe o MongoDB local com Docker            |
| `make mongo-stop`      | Para o MongoDB local                       |

Exemplo:
```sh
make run           # Roda a aplicação localmente
make docker-run    # Sobe app + MongoDB com Docker Compose
make test          # Executa os testes
```

Veja todos os comandos disponíveis com:
```sh
make help
```

## 🐳 Rodando com Docker
### Opção 1: Docker Compose (Recomendado)
O projeto inclui um `docker-compose.yml` que configura automaticamente o MongoDB e a aplicação:

```sh
# Build e start dos containers
docker-compose up --build

# Para rodar em background
docker-compose up -d --build

# Para parar os containers
docker-compose down

# Para parar e remover volumes
docker-compose down -v
```

### Opção 2: Dockerfile apenas
Se quiser rodar apenas a aplicação com Docker:

```sh
# Build da imagem
docker build -t gocrud-auth-api .

# Run do container
docker run -p 8080:8080 --env-file .env gocrud-auth-api
```

### Variáveis de Ambiente no Docker
O `docker-compose.yml` já está configurado com as variáveis necessárias:
- `MONGODB_URL=mongodb://mongo:27017` (aponta para o container do MongoDB)
- `MONGODB_USER_DB=crudInit`
- `MONGODB_USER_COLLECTION=users`
- `JWT_SECRET_KEY=your-secret-key-here`
- `LOG_OUTPUT=stdout`
- `LOG_LEVEL=info`

## 📚 Rotas Disponíveis

| Método | Rota                        | Descrição                        |
|--------|-----------------------------|-----------------------------------|
| POST   | `/createUser`               | Cria um novo usuário              |
| GET    | `/getUserById/:userId`      | Busca usuário por ID              |
| GET    | `/getUserByEmail/:userEmail`| Busca usuário por e-mail          |
| PUT    | `/updateUser/:userId`       | Atualiza dados do usuário         |
| DELETE | `/deleteUser/:userId`       | Remove usuário                    |
| POST   | `/login`                    | Realiza login e retorna JWT       |

### Exemplo de payload para criação de usuário
```json
{
  "email": "user@email.com",
  "password": "senha123",
  "name": "Nome do Usuário",
  "age": 25
}
```

## 🤝 Contribuição
Sinta-se à vontade para abrir issues, pull requests ou sugerir melhorias!

---

> Projeto desenvolvido por Kayque Souza. 