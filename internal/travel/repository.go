package travel

import (
	"database/sql"
	"time"

	"github.com/luizfelipeeoliveiraac/travel-wallter-back-end/internal/auth"
)

type TravelRepository struct {
	db *sql.DB
}

func NewTravelRepository(db *sql.DB) *TravelRepository {
	return &TravelRepository{db: db}
}

func (r *TravelRepository) CreateTravel(userID int, req auth.CreateTravelRequest) (*auth.Travel, error) {
	dataInicio, _ := time.Parse("2006-01-02", req.DataInicio)
	dataFim, _ := time.Parse("2006-01-02", req.DataFim)

	query := `INSERT INTO viagens (usuario_id, destino, data_inicio, data_fim, descricao, ativo, criado_em, atualizado_em)
	          VALUES (:1, :2, :3, :4, :5, 1, SYSTIMESTAMP, SYSTIMESTAMP)
	          RETURNING id`

	var travelID int
	err := r.db.QueryRow(query, userID, req.Destino, dataInicio, dataFim, req.Descricao).Scan(&travelID)
	if err != nil {
		return nil, err
	}

	travel := &auth.Travel{
		ID:        travelID,
		UserID:    userID,
		Destino:   req.Destino,
		DataInicio: dataInicio,
		DataFim:   dataFim,
		Descricao: req.Descricao,
		Ativo:     1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return travel, nil
}

func (r *TravelRepository) GetTravelByID(travelID, userID int) (*auth.Travel, error) {
	query := `SELECT id, usuario_id, destino, data_inicio, data_fim, descricao, ativo, criado_em, atualizado_em
	          FROM viagens WHERE id = :1 AND usuario_id = :2`

	var travel auth.Travel
	err := r.db.QueryRow(query, travelID, userID).Scan(
		&travel.ID, &travel.UserID, &travel.Destino, &travel.DataInicio, &travel.DataFim,
		&travel.Descricao, &travel.Ativo, &travel.CreatedAt, &travel.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &travel, nil
}

func (r *TravelRepository) ListTravels(userID int) ([]auth.Travel, error) {
	query := `SELECT id, usuario_id, destino, data_inicio, data_fim, descricao, ativo, criado_em, atualizado_em
	          FROM viagens WHERE usuario_id = :1 AND ativo = 1 ORDER BY criado_em DESC`

	rows, err := r.db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var travels []auth.Travel
	for rows.Next() {
		var travel auth.Travel
		err := rows.Scan(
			&travel.ID, &travel.UserID, &travel.Destino, &travel.DataInicio, &travel.DataFim,
			&travel.Descricao, &travel.Ativo, &travel.CreatedAt, &travel.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		travels = append(travels, travel)
	}

	return travels, nil
}

func (r *TravelRepository) UpdateTravel(travelID, userID int, req auth.CreateTravelRequest) (*auth.Travel, error) {
	dataInicio, _ := time.Parse("2006-01-02", req.DataInicio)
	dataFim, _ := time.Parse("2006-01-02", req.DataFim)

	query := `UPDATE viagens SET destino = :1, data_inicio = :2, data_fim = :3, descricao = :4, atualizado_em = SYSTIMESTAMP
	          WHERE id = :5 AND usuario_id = :6`

	_, err := r.db.Exec(query, req.Destino, dataInicio, dataFim, req.Descricao, travelID, userID)
	if err != nil {
		return nil, err
	}

	return r.GetTravelByID(travelID, userID)
}

func (r *TravelRepository) DeleteTravel(travelID, userID int) error {
	query := `UPDATE viagens SET ativo = 0 WHERE id = :1 AND usuario_id = :2`
	_, err := r.db.Exec(query, travelID, userID)
	return err
}
