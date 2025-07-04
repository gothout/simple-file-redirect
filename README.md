# Simple File Redirect

Este projeto é uma API escrita em Go que permite:

- Enviar arquivos para o servidor;
- Baixar arquivos enviados;
- Converter arquivos (atualmente de MP3 para OGG).

O objetivo é disponibilizar um serviço simples de armazenamento e conversão que possa evoluir para lidar com diferentes formatos. A aplicação é open source e está disponível para qualquer pessoa usar ou contribuir.

## Rotas da API

Todas as rotas abaixo estão protegidas por token Bearer (`TOKEN_APPLICATION`). Elas podem ser encontradas no código dentro de `internal/app/handler`.

| Método | Caminho | Descrição |
| ------ | ------ | -------- |
| `POST` | `/manager/v1/upload` | Faz upload de um arquivo via `multipart/form` |
| `GET` | `/manager/v1/download?path=<caminho>` | Baixa um arquivo salvo. O arquivo é removido após o envio |
| `POST` | `/manager/v1/convert` | Converte um arquivo MP3 para OGG |

A documentação completa gerada pelo Swagger pode ser acessada após subir o servidor em `http://<HOST>:<HTTP_PORT>/swagger/index.html`.

## Arquitetura

O projeto utiliza uma arquitetura em camadas:

- **Handlers** (`internal/app/handler`): definem as rotas HTTP.
- **Controllers** (`internal/app/controller`): contém a lógica para lidar com as requisições.
- **Services** (`internal/app/service`): regras de negócio e interface com o armazenamento.
- **Storage** (`internal/storage`): implementação de salvamento e conversão de arquivos.
- **Configuration** (`internal/configuration`): leitura das variáveis de ambiente.

Essa organização facilita a manutenção e possibilita novas extensões, como suporte a outros formatos de conversão.

## Dependências e Deploy

1. Copie `.env.example` para `.env` e ajuste as variáveis:
   - `HOST` – endereço que o servidor irá escutar.
   - `HTTP_PORT` – porta de acesso.
   - `DNS` – DNS utilizado para expor os arquivos.
   - `TOKEN_APPLICATION` – token necessário para acessar as rotas.
2. Instale o [`ffmpeg`](https://ffmpeg.org/), pois ele é utilizado na conversão de MP3 para OGG.
3. Compile e execute a aplicação:
   ```bash
   go build -o simple-file-redirect
   ./simple-file-redirect
   ```

Com o servidor em execução você poderá acessar a documentação via Swagger e testar cada rota.

## Licença

Distribuído sob a licença MIT. Sinta‑se à vontade para abrir issues ou pull requests!
