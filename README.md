# API para Gerenciamento de Livros

## Visão Geral

Esta API permite gerenciar uma coleção de livros de usuários, oferecendo endpoints para criar, atualizar, listar e deletar livros. Os atributos de cada livro incluem:

- **ID**: Identificador único (gerado automaticamente).
- **Título**: Nome do livro.
- **Categoria**: Gênero ou classificação do livro.
- **Autor**: Nome do autor do livro.
- **Sinopse**: Breve descrição ou resumo do livro.

## Funcionalidades

- **Cadastrar um livro**: Permite criar um novo registro de livro.
- **Atualizar um livro**: Atualiza os dados de um livro existente.
- **Listar todos os livros**: Retorna todos os livros cadastrados.
- **Listar um livro**: Retorna os detalhes de um livro específico.
- **Deletar um livro**: Remove um livro pelo seu ID.

---

## Requisitos Técnicos

### Tecnologias Utilizadas

- **Linguagem**: GoLang
- **Banco de Dados**: PostgreSQL  
  > Escolhido por sua robustez e escalabilidade. Além disso, é amplamente adotado para aplicações web modernas.

- **Containerização**: Docker e Docker Compose  
  > Facilita o gerenciamento do ambiente e garante consistência entre os sistemas.

- **Bibliotecas e Frameworks**:
  > **Gin Gonic (v1.10.0)**: Framework web leve e eficiente, ideal para criar APIs RESTful.
    
  > **Validator (v10.23.0)**: Biblioteca para validação de estruturas e campos de entrada, garantindo que os dados sejam enviados no formato esperado.
  
  > **pq (v1.10.9)**: Driver para integração com o PostgreSQL, oferecendo suporte completo às funcionalidades do banco de dados.
  
  > **Bytedance Sonic (v1.12.6)**: Biblioteca de manipulação JSON de alto desempenho, utilizada para aprimorar o gerenciamento dos dados JSON.

---

## Configuração Inicial

1. Certifique-se de ter o Docker e o Docker Compose instalados.
   ```bash
   docker --version
   docker compose version    
   ```

2. Clone este repositório:
   ```bash
   git clone git@github.com:bruxo-dev/desafio_taghos_backend_jr.git
   cd desafio_taghos_backend_jr 
   ```
3. Suba os contêineres:
   ```bash
   docker-compose up -d
   ```
4. A API estará disponível em: `http://localhost:8080`.

---

## Endpoints

### 1. **Obter Todos os Livros**
**Endpoint**: `GET /books`  
**Descrição**: Recupera uma lista de todos os livros cadastrados.  

**Exemplo de Requisição BASH**:
```bash
curl -X GET http://localhost:8080/books
```
**OBS**: Se não houver livros no banco de dados, esta requisição não retorna nada no BASH.

**Resposta**:
- **200 OK**: Retorna uma lista de livros.
- **500 Internal Server Error**: Erro ao processar a solicitação.

---

### 2. **Obter Livro pelo ID**
**Endpoint**: `GET /books/:id`  
**Descrição**: Recupera os detalhes de um livro específico pelo ID.  

**Exemplo de Requisição BASH**:
```bash
curl -X GET http://localhost:8080/books/1
```

**Resposta**:
- **200 OK**: Retorna o livro.
- **404 Not Found**: Livro não encontrado.
- **400 Bad Request**: ID inválido.

---

### 3. **Cadastrar um Livro**
**Endpoint**: `POST /books`  
**Descrição**: Adiciona um novo livro ao banco de dados.  

**Exemplo de Requisição BASH**:
```bash
curl -X POST http://localhost:8080/books \
-H "Content-Type: application/json" \
-d '{
  "title": "O Livro",
  "category": "Ficção",
  "author": "Autor Exemplo",
  "synopsis": "Resumo do livro."
}'
```

**Resposta**:
- **201 Created**: Livro criado com sucesso.
- **400 Bad Request**: Dados inválidos.

---

### 4. **Atualizar um Livro**
**Endpoint**: `PUT /books/:id`  
**Descrição**: Atualiza as informações de um livro existente.  

**Exemplo de Requisição BASH**:
```bash
curl -X PUT http://localhost:8080/books/1 \
-H "Content-Type: application/json" \
-d '{
  "title": "Novo Título",
  "category": "Nova Categoria",
  "author": "Novo Autor",
  "synopsis": "Nova Sinopse."
}'
```

**Resposta**:
- **200 OK**: Livro atualizado com sucesso.
- **404 Not Found**: Livro não encontrado.
- **400 Bad Request**: Dados inválidos.

---

### 5. **Deletar um Livro**
**Endpoint**: `DELETE /books/:id`  
**Descrição**: Exclui um livro específico pelo seu ID.  

**Exemplo de Requisição BASH**:
```bash
curl -X DELETE http://localhost:8080/books/1
```

**Resposta**:
- **204 No Content**: Livro deletado com sucesso.
- **404 Not Found**: Livro não encontrado.
- **400 Bad Request**: ID inválido.

---

### Estrutura do Banco de Dados
A tabela de livros é definida da seguinte forma:
```sql
CREATE TABLE books (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    category VARCHAR(100) NOT NULL,
    author VARCHAR(100) NOT NULL,
    synopsis TEXT NOT NULL
);
```



---

## Licença
Este projeto é licenciado sob a [MIT License](LICENSE).
