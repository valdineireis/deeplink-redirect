# deeplink-redirect

🔗 [Acesse aqui o Deeplink Redirect](https://deeplink-redirect-1059480473666.us-central1.run.app/)

### Execute os testes:

```bash
go test -v
```

### Suba o container com Docker Compose:

```bash
docker compose up --build
```

#### ou se preferir

```bash
# Reconstruir a imagem
docker build -t deeplink-redirect .

# Executar o contêiner
docker run -p 8080:8080 deeplink-redirect
```

### Teste as rotas em outro terminal:

- Sucesso:

```bash
curl http://localhost:8080/
# Resposta esperada: Página HML com um formulário
```

- Sucesso:

```bash
curl http://localhost:8080/redirect?deeplink=myapp://home
# Resposta esperada: Redirecionamento para o APP
```

- Erro:

```bash
curl http://localhost:8080/redirect
# Resposta esperada: "Missing deeplink parameter"
```
