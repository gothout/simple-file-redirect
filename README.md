# Simple File Redirect

Este projeto fornece uma API escrita em Go para fazer upload, download e conversão de arquivos de `mp3` para `ogg`. A aplicação utiliza o framework [Gin](https://github.com/gin-gonic/gin) e o `ffmpeg` para realizar as conversões de áudio.

## Como executar

1. Copie o arquivo `.env.example` para `.env` e ajuste as variáveis de ambiente:
   - `HOST`: Endereço em que o servidor escutará.
   - `HTTP_PORT`: Porta utilizada pela aplicação.
   - `DNS`: DNS utilizado para exposição dos arquivos.
   - `TOKEN_APPLICATION`: Token utilizado para autenticação das rotas.
2. Certifique-se de ter o `ffmpeg` instalado e disponível no `PATH` do sistema.
3. Execute a aplicação:
   ```bash
   go run main.go
   ```

## Documentação da API

A documentação gerada via Swagger fica no diretório [`docs/`](docs/). Ao rodar o servidor, é possível acessar a interface diretamente em `http://<HOST>:<HTTP_PORT>/swagger/index.html`. Também é possível abrir o arquivo `swagger.yaml` com alguma outra ferramenta de visualização (por exemplo, [Swagger UI](https://swagger.io/tools/swagger-ui/)).

## Modo de conversão

A API expõe a rota `POST /manager/v1/convert` para converter arquivos `mp3` em `ogg`. O processo consiste em:
1. Realizar o upload do arquivo `mp3` via multipart/form.
2. O serviço salva o arquivo temporariamente e executa o `ffmpeg` para gerar a versão `ogg`.
3. Após a conversão, o arquivo original `mp3` é removido.

## Deleção após download

Quando você realiza o download por meio da rota `GET /manager/v1/download`, o arquivo é enviado ao cliente e, em seguida, o serviço o remove do armazenamento local.

## Dependência ffmpeg

O `ffmpeg` é essencial para a funcionalidade de conversão. Garanta que o binário esteja instalado em sua máquina e acessível no `PATH`. Sem ele, a rota de conversão não funcionará.


## Arquitetura do Projeto

O projeto segue uma estrutura modular:

- `cmd/server`: inicializa o servidor HTTP.
- `internal/app`: contém handlers, controllers e serviços da aplicação.
- `internal/storage`: implementa a lógica de armazenamento e conversão de arquivos.
- `internal/configuration`: gerencia variáveis de ambiente e configurações.
- `docs/`: arquivos gerados pelo Swagger.

Essa separação facilita a manutenção e a evolução do código.
