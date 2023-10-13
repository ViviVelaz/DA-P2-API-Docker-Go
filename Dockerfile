# Usar la imagen oficial de golang
FROM golang:1.21
# Asignar un directorio de trabajo
WORKDIR /app
# Copear lo que este en raíz al directorio raíz de trabajo
COPY . .
# Descargar e instalar las dependencias
RUN go get -d -v ./...
# Construir la aplicación Go
RUN go build -o api .
# Exponer el puerto
EXPOSE 8080
# Correr el ejecutable
CMD ["./api"]
