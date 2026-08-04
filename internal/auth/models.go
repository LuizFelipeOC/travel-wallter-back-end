package auth

import "time"

type User struct {
	ID        int       `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	Email     string    `json:"email" db:"email"`
	Password  string    `json:"-" db:"password"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type RegisterRequest struct {
	Name     string `json:"name" binding:"required,min=3,max=100"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type AuthResponse struct {
	Token string `json:"token"`
	User  User   `json:"user"`
}

type Travel struct {
	ID        int       `json:"id" db:"id"`
	UserID    int       `json:"user_id" db:"usuario_id"`
	Destino   string    `json:"destino" db:"destino"`
	DataInicio time.Time `json:"data_inicio" db:"data_inicio"`
	DataFim   time.Time `json:"data_fim" db:"data_fim"`
	Descricao string    `json:"descricao" db:"descricao"`
	Ativo     int       `json:"ativo" db:"ativo"`
	CreatedAt time.Time `json:"created_at" db:"criado_em"`
	UpdatedAt time.Time `json:"updated_at" db:"atualizado_em"`
}

type CreateTravelRequest struct {
	Destino    string    `json:"destino" binding:"required,min=3,max=100"`
	DataInicio string    `json:"data_inicio" binding:"required"`
	DataFim    string    `json:"data_fim" binding:"required"`
	Descricao  string    `json:"descricao" binding:"max=500"`
}

type Gasto struct {
	ID        int       `json:"id" db:"id"`
	ViagemID  int       `json:"viagem_id" db:"viagem_id"`
	Descricao string    `json:"descricao" db:"descricao"`
	Valor     float64   `json:"valor" db:"valor"`
	Categoria string    `json:"categoria" db:"categoria"`
	DataGasto time.Time `json:"data_gasto" db:"data_gasto"`
	CreatedAt time.Time `json:"created_at" db:"criado_em"`
	UpdatedAt time.Time `json:"updated_at" db:"atualizado_em"`
}

type CreateGastoRequest struct {
	Descricao string  `json:"descricao" binding:"required,min=3,max=255"`
	Valor     float64 `json:"valor" binding:"required,gt=0"`
	Categoria string  `json:"categoria" binding:"max=50"`
	DataGasto string  `json:"data_gasto" binding:"required"`
}
