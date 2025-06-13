# Etapa de compilación
FROM golang:1.24-alpine AS builder

# Instalar dependencias necesarias
RUN apk add --no-cache git gcc musl-dev

# Establecer directorio de trabajo
WORKDIR /app

# Copiar archivos go.mod y go.sum
COPY go.mod go.sum ./

# Descargar dependencias
RUN go mod download

# Copiar el código fuente
COPY . .

# Compilar la aplicación
RUN go build -o portfolio ./cmd/app

# Etapa de ejecución
FROM alpine:latest

# Instalar dependencias necesarias para la aplicación en producción
RUN apk add --no-cache ca-certificates tzdata

# Crear un usuario no root para ejecutar la aplicación
RUN adduser -D -g '' appuser

# Establecer directorio de trabajo
WORKDIR /app

# Copiar el binario compilado desde la etapa de compilación
COPY --from=builder /app/portfolio /app/

# Copiar archivos estáticos y plantillas
COPY --from=builder /app/src /app/src
COPY --from=builder /app/.env /app/.env

# Asegurar que los directorios tienen los permisos correctos
RUN chown -R appuser:appuser /app

# Cambiar al usuario no root
USER appuser

# Exponer el puerto que usa la aplicación
EXPOSE 9999

# Comando para ejecutar la aplicación
CMD ["/app/portfolio"]
