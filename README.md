# deeplink-redirect

🔗 [Copie e informe o deeplink](https://deeplink-redirect-1059480473666.us-central1.run.app/redirect?deeplink=)

### Execute os testes:

```bash
go test -v
```

### Suba o container com Docker Compose:

```bash
docker compose up --build
```

### Teste as rotas em outro terminal:

- Sucesso:

```bash
curl http://localhost:8080/redirect?deeplink=myapp://home
# Resposta esperada: Redirecionamento para o APP
```

- Erro:

```bash
curl http://localhost:8080/redirect
# Resposta esperada: Erro com código de status 400
```
