# 🔐 API de Autenticação JWT com Go

API segura para gerenciamento de usuários com autenticação JWT, criptografia bcrypt e proteção de dados sensíveis.

## 🚀 Começando

### Pré-requisitos
- Go 1.18+
- Docker 20+ e docker-compose

### Passos:

1. Clone o repositório:
   ```bash
   git clone https://github.com/ericsanto/api_authentication
   cd api_authentication
   ```

2. Suba a aplicação com Docker:
   ```bash
   docker-compose up -d
   ```

3. Acesse a API em `http://localhost:8080`.

## 📡 Endpoints da API

## User

| Método | Rota                     | Descrição                          |
| ------ | ------------------------ | ---------------------------------- |
| POST   | `/v1/register`           | Cria um novo usuário               |
| POST   | `/v1/login             ` | Faz login no sistema               |

### ✅ Exemplo de Request: `POST /v1/register`

```json
{
    "name": "josé felisberto cruz de formiga",
    "email": "felisjoseformiga1.com",
    "password": "Fe1isJoseA2345"
}
```
### ✅ Respostas:
```
    201 Created: Sucesso

    400 Bad Request: Body da requisição inválido

    422 UnprocessableEntity: campos da requisição é inválido ou está faltando
```

---

### ✅ Exemplo de Request: `POST /v1/login`
```json
{
    "email": "felisjoseformiga1.com",
    "password": "Fe1isJoseA2345"
}
```

### ✅ Exemplo de Response: `POST /v1/login`
```json
{
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6IiQyYSQxMCRNTjVvalFTeTZVQlJaa1VwODNOSGYuTEZrRUphbGZKQ1I5dFZtTFozalVmWldrZUwzTzVTSyIsImV4cCI6MTc0NjE0NTg1OCwiaWQiOjJ9.Ysog61L6CrDS2x6pvl55exs42UuX9eaysuceLD6OOVE"
}
```

```
  200 OK: Sucesso

    400 Bad Request: Body da requisição inválido

    422 UnprocessableEntity: Campos do json é inválido ou está faltando

    404 NotFound: Não existe email cadastrado

    401 Unauthorized: Senha incorreta
```