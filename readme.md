# <h1 align="center">Projeto Devbook</h1>

## 📗 Sobre o Projeto
Este projeto foi criado para aprofundar meus estodos em golang, Docker e Nest.js

## 🔨 Ferramentas e Bibliotecas
- [Golang](https://go.dev)
- [MySQL](https://www.mysql.com/)
- [Next.js](https://nextjs.org)
- [Docker](https://www.docker.com)
- [JWT](https://jwt.io/)

## 🦾 Iniciando o Projeto
Para executar o projeto, siga os passos abaixo:

1. Clone o repositório usando o seguinte comando no terminal:
    ```bash
    git clone https://github.com/caio1459/devbook
    ```

2. Acesse o diretório do projeto:
    ```bash
    cd devbook
    ```

3. Para rodar a API entre no diretório `api` e certifique-se de ter o docker instalado na sua máquina
    ```bash
    cd api
    ```
    Em seguinte execute os seguintes comando para iniciar a API
    ```bash
    docker compose up
    ```
    ```bash
    ./devbook
    ```

4. Abra o projeto no VSCode e execute o script SQL contido no arquivo `api/sql/devbook.sql` no banco de dados MySQL para criar a estrutura necessária.

    Agora, a API estará disponível para uso localmente. Certifique-se de ajustar as configurações e váriaveis de ambiente conforme necessário.

5. Para o front entre no diretório `client` e certifique-se de ter o docker instalado na sua máquina
    ```bash
    cd client
    ```

6. Instale as dependências do projeto
    ```bash
    npm install
    ```

6. Inicie o servidor, certificando-se de ter a API rodando
    ```bash
    npm run dev
    ```