# Travel Walter Backend

REST API para gerenciar viagens e despesas, desenvolvida em Go com Oracle Database.

## Requisitos

- Docker 20.10+
- Docker Compose 2.0+
- Go 1.22+ (opcional, para desenvolvimento local)

## Quick Start

### 1. Clonar e configurar
```bash
git clone <repository>
cd travel-wallter-back-end
cp .env.example .env
```

### 2. Compilar e rodar
```bash
docker-compose build
docker-compose up -d
```

### 3. Verificar
```bash
docker-compose ps
docker-compose logs -f
```

A API estará disponível em `http://localhost:8080`

## Estrutura do Projeto

```
.
├── Dockerfile              # Build da aplicação Go
├── docker-compose.yml      # Orquestração (App + Oracle)
├── .env.example           # Variáveis de ambiente
├── .dockerignore          # Arquivos ignorados no build
├── Makefile               # Comandos úteis
├── go.mod                 # Dependências Go
├── go.sum                 # Checksums das dependências
└── main.go                # Aplicação principal
```

## Variáveis de Ambiente

As seguintes variáveis estão definidas no `docker-compose.yml`:

- `DB_HOST` - Host do Oracle (padrão: oracle)
- `DB_PORT` - Porta do Oracle (padrão: 1521)
- `DB_USER` - Usuário do banco (padrão: admin)
- `DB_PASSWORD` - Senha do banco (padrão: oracle123)

## Comandos Úteis

```bash
# Build
make build

# Iniciar containers
make up

# Parar containers
make down

# Ver logs em tempo real
make logs

# Reiniciar containers
make restart

# Limpar volumes
make clean
```

## Desenvolvimento

### Dependência do Oracle

```bash
go get github.com/godror/godror
```

### Exemplo de conexão

```go
import (
    "database/sql"
    _ "github.com/godror/godror"
)

db, err := sql.Open("godror", 
    "admin/oracle123@oracle:1521/XEPDB1")
```

## Banco de Dados

O Oracle estará acessível em `localhost:1521` após o container estar pronto (~60s).

**Credenciais:**
- Usuário: `admin`
- Senha: `oracle123`
- Database: `XEPDB1`

## Troubleshooting

**Oracle não conecta**
- Aguarde 60+ segundos para o Oracle iniciar
- Verifique: `docker-compose logs oracle`
- Confirme que a porta 1521 está disponível

**Go não compila**
- Certifique-se que `go.mod` e `main.go` existem
- Delete volumes e rebuild: `make clean && make build`

## Arquitetura

- **App**: Go API rodando em container
- **Database**: Oracle Express Edition

Ambos conectados na mesma rede Docker (`app-network`).
