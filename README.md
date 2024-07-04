# API de Gerenciamento de Produtos

# Implementação
Este projeto foi implementado com simplicidade para focar no básico do gerenciamento de produtos, utilizando Go. Optei por não adicionar uma tabela separada para preços, pois o maior desafio seria programar em go.

## Visão Geral

Este projeto implementa uma API básica de CRUD para gerenciar produtos. Ele fornece endpoints para criar, excluir,
recuperar por ID, listar todos os produtos e atualizar detalhes do produto.

## Endpoints

### Criar um Produto

- **URL**: `/product`
- **Método**: `POST`
```json
  Request: {
              "name": "name",
              "description": "description",
              "price": 20.00,
              "stock": 2
  }
```

- **Resposta**: Objeto JSON com os detalhes do produto criado

### Excluir um Produto

- **URL**: `/product/{id}`
- **Método**: `DELETE`
- **Resposta**: Mensagem dizendo que foi excluído com sucesso

### Obter Produto por ID

- **URL**: `/product/{id}`
- **Método**: `GET`
- **Resposta**: Objeto JSON com os detalhes do produto correspondente ao ID

### Obter Todos os Produtos

- **URL**: `/product`
- **Método**: `GET`
- **Resposta**: Array JSON com todos os produtos no banco de dados

### Atualizar um Produto

- **URL**: `/product`
- **Método**: `PUT`
```json
  Request: {
              "name": "name",
              "description": "description",
              "price": 20.00,
              "stock": 2
  }
```
- **Resposta**: Objeto JSON com os detalhes atualizados do produto

# Stack Utilizada

- Optei por utilizar GO, Echo, Mysql.

# Arquitetura

- Utilizei a arquitetura Onion, mas fiz apenas a estrutura de pastas para passar a ideia. Para um projeto maior, poderia
  separar em módulos.

# Postman

- Deixei um arquivo na pasta raiz do projeto com as requisições da api.

# Como executar

#### Deixei um docker-compose na pasta raiz do projeto, e também o .env, tendo em vista que não é um projeto que vai para produção. Esse compose é pra você executar junto com alguma ide. Após rodar o comando abaixo é só rodar a aplicação.

- docker-compose up


