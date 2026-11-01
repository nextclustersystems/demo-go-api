# Demo Go (API com chi) na Nextcluster

Exemplo mínimo para a "Aplicação com build (Docker)": a plataforma lê o `Dockerfile`, baixa os módulos do `go.mod` e publica um binário estático.

- Publicar por ZIP: compacte esta pasta e envie na aba **Publicar**.
- Publicar pelo GitHub: aponte o painel Git para o repositório e a branch.
- A API escuta na porta da variável `PORT`. Em **Configurações**, use o health check `/health`.
- Rotas: `/`, `/health`, `/ola/{nome}`.

Local: `docker build -t demo-go . && docker run -p 8080:8080 demo-go` e abra http://localhost:8080.
