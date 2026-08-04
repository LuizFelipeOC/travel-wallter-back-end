package database 

func Connect() {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")

	cnnStr := fmt.Sprintf("host=%s port=%s user=%s password=%s sslmode=disable", dbHost, dbPort, dbUser, dbPassword)


	db, err := sql.Open("godror", cnnStr)

	if err != nil {
		return nil, NewDatabaseError("DB_CONNECTION", "Erro ao conectar no banco", err)
	}

	err = db.Ping()

	if err != nil {
		return nil, NewDatabaseError("DB_CONNECTION", "Erro ao conectar no banco", err)
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(2)

	return db, nil
}