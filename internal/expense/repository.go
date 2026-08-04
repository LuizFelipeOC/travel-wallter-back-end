package expense

import (
	"database/sql"
	"time"

	"github.com/luizfelipeeoliveiraac/travel-wallter-back-end/internal/auth"
)

type ExpenseRepository struct {
	db *sql.DB
}

func NewExpenseRepository(db *sql.DB) *ExpenseRepository {
	return &ExpenseRepository{db: db}
}

func (r *ExpenseRepository) CreateExpense(travelID int, userID int, req auth.CreateGastoRequest) (*auth.Gasto, error) {
	dataGasto, _ := time.Parse("2006-01-02", req.DataGasto)

	query := `INSERT INTO gastos (viagem_id, descricao, valor, categoria, data_gasto, criado_em, atualizado_em)
	          SELECT :1, :2, :3, :4, :5, SYSTIMESTAMP, SYSTIMESTAMP FROM DUAL
	          WHERE EXISTS (SELECT 1 FROM viagens WHERE id = :1 AND usuario_id = :6)
	          RETURNING id`

	var gastoID int
	err := r.db.QueryRow(query, travelID, req.Descricao, req.Valor, req.Categoria, dataGasto, userID).Scan(&gastoID)
	if err != nil {
		return nil, err
	}

	gasto := &auth.Gasto{
		ID:        gastoID,
		ViagemID:  travelID,
		Descricao: req.Descricao,
		Valor:     req.Valor,
		Categoria: req.Categoria,
		DataGasto: dataGasto,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return gasto, nil
}

func (r *ExpenseRepository) GetExpenseByID(gastoID, travelID, userID int) (*auth.Gasto, error) {
	query := `SELECT g.id, g.viagem_id, g.descricao, g.valor, g.categoria, g.data_gasto, g.criado_em, g.atualizado_em
	          FROM gastos g
	          JOIN viagens v ON g.viagem_id = v.id
	          WHERE g.id = :1 AND g.viagem_id = :2 AND v.usuario_id = :3`

	var gasto auth.Gasto
	err := r.db.QueryRow(query, gastoID, travelID, userID).Scan(
		&gasto.ID, &gasto.ViagemID, &gasto.Descricao, &gasto.Valor, &gasto.Categoria, &gasto.DataGasto, &gasto.CreatedAt, &gasto.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &gasto, nil
}

func (r *ExpenseRepository) ListExpenses(travelID, userID int) ([]auth.Gasto, error) {
	query := `SELECT g.id, g.viagem_id, g.descricao, g.valor, g.categoria, g.data_gasto, g.criado_em, g.atualizado_em
	          FROM gastos g
	          JOIN viagens v ON g.viagem_id = v.id
	          WHERE g.viagem_id = :1 AND v.usuario_id = :2
	          ORDER BY g.data_gasto DESC`

	rows, err := r.db.Query(query, travelID, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var gastos []auth.Gasto
	for rows.Next() {
		var gasto auth.Gasto
		err := rows.Scan(
			&gasto.ID, &gasto.ViagemID, &gasto.Descricao, &gasto.Valor, &gasto.Categoria, &gasto.DataGasto, &gasto.CreatedAt, &gasto.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		gastos = append(gastos, gasto)
	}

	return gastos, nil
}

func (r *ExpenseRepository) UpdateExpense(gastoID, travelID, userID int, req auth.CreateGastoRequest) (*auth.Gasto, error) {
	dataGasto, _ := time.Parse("2006-01-02", req.DataGasto)

	query := `UPDATE gastos SET descricao = :1, valor = :2, categoria = :3, data_gasto = :4, atualizado_em = SYSTIMESTAMP
	          WHERE id = :5 AND viagem_id = :6 AND EXISTS (SELECT 1 FROM viagens WHERE id = :6 AND usuario_id = :7)`

	_, err := r.db.Exec(query, req.Descricao, req.Valor, req.Categoria, dataGasto, gastoID, travelID, userID)
	if err != nil {
		return nil, err
	}

	return r.GetExpenseByID(gastoID, travelID, userID)
}

func (r *ExpenseRepository) DeleteExpense(gastoID, travelID, userID int) error {
	query := `DELETE FROM gastos WHERE id = :1 AND viagem_id = :2 AND EXISTS (SELECT 1 FROM viagens WHERE id = :2 AND usuario_id = :3)`
	_, err := r.db.Exec(query, gastoID, travelID, userID)
	return err
}

func (r *ExpenseRepository) GetTravelTotal(travelID, userID int) (float64, error) {
	query := `SELECT COALESCE(SUM(valor), 0) FROM gastos g
	          JOIN viagens v ON g.viagem_id = v.id
	          WHERE g.viagem_id = :1 AND v.usuario_id = :2`

	var total float64
	err := r.db.QueryRow(query, travelID, userID).Scan(&total)
	return total, err
}
