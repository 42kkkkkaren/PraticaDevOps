# Use a imagem base do Golang
FROM golang:1.19-alpine

# Crie um diretório de trabalho dentro do container
WORKDIR /app

# Copie os arquivos go para o diretório de trabalho
COPY *.go ./

# Baixe as dependências do projeto
RUN go mod download

# Construa o binário do projeto
RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-gs-ping

# Exponha a porta que a aplicação usa
EXPOSE 8080

# Comando para rodar a aplicação
CMD ["/docker-gs-ping"]