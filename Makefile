# Variables
CERT_FILE=cert.pem
KEY_FILE=key.pem
UNENCRYPTED_KEY_FILE=key_unencrypted.pem
HOST=myserver.local
PORT=5000

# Comandos
run: check-certs
	@echo "Iniciando servidor en https://$(HOST):$(PORT)..."
	@go run main.go

check-certs: $(CERT_FILE) $(UNENCRYPTED_KEY_FILE)

$(CERT_FILE) $(KEY_FILE):
	@echo "Generando certificados..."
	@openssl req -x509 -newkey rsa:4096 -nodes -keyout $(KEY_FILE) -out $(CERT_FILE) -days 365 -subj "/C=ES/ST=Madrid/L=Madrid/O=MyProject/CN=$(HOST)"

$(UNENCRYPTED_KEY_FILE): $(KEY_FILE)
	@echo "Eliminando contraseña de la clave privada..."
	@openssl rsa -in $(KEY_FILE) -out $(UNENCRYPTED_KEY_FILE)

clean:
	@echo "Limpiando archivos generados..."
	@if [ -f $(CERT_FILE) ]; then rm $(CERT_FILE); fi
	@if [ -f $(KEY_FILE) ]; then rm $(KEY_FILE); fi
	@if [ -f $(UNENCRYPTED_KEY_FILE) ]; then rm $(UNENCRYPTED_KEY_FILE); fi


help:
	@echo "Comandos disponibles:"
	@echo "  make run        - Genera certificados (si no existen) y lanza el servidor"
	@echo "  make clean      - Limpia los archivos generados (certificados y claves)"
	@echo "  make help       - Muestra esta ayuda"

