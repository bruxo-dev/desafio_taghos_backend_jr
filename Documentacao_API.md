
# Documentação da API para Gerenciamento de Livros

## Visão Geral
Esta API fornece endpoints para gerenciar uma coleção de livros. Cada livro possui os seguintes atributos:
- **ID**: Identificador único (inteiro).
- **Title**: Título do livro (string).
- **Category**: Categoria ou gênero do livro (string).
- **Author**: Autor do livro (string).
- **Synopsis**: Uma breve sinopse do livro (string).

A API suporta as seguintes operações:
- Recuperar todos os livros
- Recuperar um livro pelo ID
- Adicionar um novo livro
- Atualizar um livro pelo ID
- Excluir um livro pelo ID

## Endpoints

### 1. **Obter Todos os Livros**
**Endpoint**: `GET /books`

**Descrição**: Recupera uma lista de todos os livros na coleção.

**Resposta**:
- **200 OK**: Retorna um array de objetos de livro.
  ```json
  [
    {
      "id": 1,
      "title": "Título do Livro",
      "category": "Ficção",
      "author": "Nome do Autor",
      "synopsis": "Descrição breve."
    }
  ]
  ```
- **500 Internal Server Error**: Se ocorrer um erro ao recuperar os livros.

---

### 2. **Obter Livro pelo ID**
**Endpoint**: `GET /books/:id`

**Descrição**: Recupera os detalhes de um livro específico pelo seu ID.

**Parâmetros de Caminho**:
- `id` (inteiro): O ID do livro a ser recuperado.

**Resposta**:
- **200 OK**: Retorna o objeto do livro.
  ```json
  {
    "id": 1,
    "title": "Título do Livro",
    "category": "Ficção",
    "author": "Nome do Autor",
    "synopsis": "Descrição breve."
  }
  ```
- **400 Bad Request**: Se o `id` não for um inteiro válido.
- **404 Not Found**: Se nenhum livro for encontrado com o ID fornecido.
- **500 Internal Server Error**: Se ocorrer um erro ao recuperar o livro.

---

### 3. **Criar um Livro**
**Endpoint**: `POST /books`

**Descrição**: Adiciona um novo livro à coleção.

**Corpo da Requisição**:
- Um objeto JSON com os seguintes campos:
  ```json
  {
    "title": "Título do Livro",
    "category": "Ficção",
    "author": "Nome do Autor",
    "synopsis": "Descrição breve."
  }
  ```

**Resposta**:
- **201 Created**: Retorna o objeto do livro criado.
- **400 Bad Request**: Se o corpo da requisição for inválido ou algum campo obrigatório estiver ausente.
- **500 Internal Server Error**: Se ocorrer um erro ao adicionar o livro.

---

### 4. **Atualizar um Livro pelo ID**
**Endpoint**: `PUT /books/:id`

**Descrição**: Atualiza os detalhes de um livro específico.

**Parâmetros de Caminho**:
- `id` (inteiro): O ID do livro a ser atualizado.

**Corpo da Requisição**:
- Um objeto JSON com os campos atualizados:
  ```json
  {
    "title": "Título Atualizado",
    "category": "Categoria Atualizada",
    "author": "Autor Atualizado",
    "synopsis": "Sinopse atualizada."
  }
  ```

**Resposta**:
- **200 OK**: Retorna o objeto do livro atualizado.
- **400 Bad Request**: Se o `id` não for válido ou o corpo da requisição for inválido.
- **404 Not Found**: Se nenhum livro for encontrado com o ID fornecido.
- **500 Internal Server Error**: Se ocorrer um erro ao atualizar o livro.

---

### 5. **Excluir um Livro pelo ID**
**Endpoint**: `DELETE /books/:id`

**Descrição**: Exclui um livro específico pelo seu ID.

**Parâmetros de Caminho**:
- `id` (inteiro): O ID do livro a ser excluído.

**Resposta**:
- **204 No Content**: Se o livro for excluído com sucesso.
- **400 Bad Request**: Se o `id` não for válido.
- **404 Not Found**: Se nenhum livro for encontrado com o ID fornecido.
- **500 Internal Server Error**: Se ocorrer um erro ao excluir o livro.

---

## Notas
- Todos os corpos de requisição devem ser enviados no formato JSON.
- Erros de validação fornecem mensagens detalhadas sobre os campos inválidos.
- Todas as entradas são sanitizadas para evitar injeção de dados maliciosos.

Esta documentação está sujeita a atualizações conforme a evolução da API.
