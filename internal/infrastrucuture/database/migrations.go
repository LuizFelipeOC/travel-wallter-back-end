package database

import (
	"database/sql"
	"fmt"
)

func RunMigrations(db *sql.DB) error {
	migrations := []string{
		`CREATE TABLE usuarios (
			id NUMBER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
			nome VARCHAR2(100) NOT NULL,
			email VARCHAR2(100) NOT NULL UNIQUE,
			senha VARCHAR2(255) NOT NULL,
			ativo NUMBER(1) DEFAULT 1,
			criado_em TIMESTAMP DEFAULT SYSTIMESTAMP,
			atualizado_em TIMESTAMP DEFAULT SYSTIMESTAMP
		)`,
		`CREATE TABLE viagens (
			id NUMBER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
			usuario_id NUMBER NOT NULL,
			destino VARCHAR2(100) NOT NULL,
			data_inicio DATE NOT NULL,
			data_fim DATE NOT NULL,
			descricao CLOB,
			ativo NUMBER(1) DEFAULT 1,
			criado_em TIMESTAMP DEFAULT SYSTIMESTAMP,
			atualizado_em TIMESTAMP DEFAULT SYSTIMESTAMP,
			CONSTRAINT fk_viagens_usuarios
				FOREIGN KEY (usuario_id) REFERENCES usuarios(id) ON DELETE CASCADE
		)`,
		`CREATE TABLE gastos (
			id NUMBER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
			viagem_id NUMBER NOT NULL,
			descricao VARCHAR2(255) NOT NULL,
			valor NUMBER(10,2) NOT NULL,
			categoria VARCHAR2(50),
			data_gasto DATE NOT NULL,
			criado_em TIMESTAMP DEFAULT SYSTIMESTAMP,
			atualizado_em TIMESTAMP DEFAULT SYSTIMESTAMP,
			CONSTRAINT fk_gastos_viagens
				FOREIGN KEY (viagem_id) REFERENCES viagens(id) ON DELETE CASCADE
		)`,
		`CREATE INDEX idx_usuarios_email ON usuarios(email)`,
		`CREATE INDEX idx_viagens_usuario_id ON viagens(usuario_id)`,
		`CREATE INDEX idx_gastos_viagem_id ON gastos(viagem_id)`,
	}

	for _, migration := range migrations {
		_, err := db.Exec(migration)
		if err != nil {
			return NewDatabaseError("DB_MIGRATION", "Erro ao executar migration", err)
		}
	}

	fmt.Println("✓ Migrations executadas com sucesso")
	return nil
}