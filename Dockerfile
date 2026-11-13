# ==========================================
# ETAPA 1: Construcción de la aplicación Go
# ==========================================
FROM golang:1.22-alpine AS builder

# Directorio de trabajo dentro del contenedor
WORKDIR /app

# Copiar los archivos de dependencias
COPY go.mod ./
RUN go mod download

# Copiar todo el código fuente del proyecto
COPY . .

# Compilar la aplicación estáticamente
# (Desactiva CGO para que el binario sea completamente autónomo)
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/mi-servidor .

# ==========================================
# ETAPA 2: Imagen final ligera para producción
# ==========================================
FROM alpine:latest

# Instalar certificados raíz por si tu app hace peticiones HTTPS externas
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copiar únicamente el archivo ejecutable compilado desde la etapa anterior
COPY --from=builder /app/mi-servidor .

# Exponer el puerto en el que corre tu servidor (ej. 8080)
EXPOSE 8080

# Comando por defecto para ejecutar la aplicación al iniciar el contenedor
CMD ["./mi-servidor"]