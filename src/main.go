package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	// "os"
	tnd "tender_service/src/tender" // Путь к пакету с функцией работы с тендерами

	_ "github.com/lib/pq" // PostgreSQL driver
)

var db *sql.DB

func initDB() *sql.DB {
	// connStr := os.Getenv("POSTGRES_CONN")
	db, err := sql.Open("postgres", "host=db port=5432 user=avito password=1234 dbname=avito_db sslmode=disable")
	fmt.Println("conn not fail")
	if err != nil {
		fmt.Println("conn fail")
		log.Fatal("Failed to connect to database:", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Fail to connect to DB", err)
	} else {
		log.Println("Successfully connected to DB")
	}

	return db
}

func main() {
	db = initDB()

	http.HandleFunc("/api/ping", PingHandler)
	http.HandleFunc("/api/tenders", GetTendersHandler)

	fmt.Println("Server is running on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}

	fmt.Println("HTTP server started successfully and running on port 8080")
}

func PingHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
	fmt.Println("pinged")
}

func GetTendersHandler(w http.ResponseWriter, r *http.Request) {
	// serviceType := r.URL.Query().Get("serviceType") // Получаем параметр фильтрации по типу услуг

	tenders, err := GetTenders(db)
	if err != nil {
		http.Error(w, "Error fetching tenders", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tenders)
}

func GetTenders(db *sql.DB) ([]tnd.Tender, error) {
	rows, err := db.Query("SELECT id, name, description, service_type, status FROM tenders")
	if err != nil {
		fmt.Println("error query")
		return nil, err
	}
	defer rows.Close()

	var tenders []tnd.Tender
	for rows.Next() {
		var tender tnd.Tender
		if err := rows.Scan(&tender.ID, &tender.Name, &tender.Description, &tender.ServiceType, &tender.Status); err != nil {
			log.Println("Error scanning row:", err)
			continue
		}
		tenders = append(tenders, tender)
	}

	if err := rows.Err(); err != nil {
		fmt.Println("errorings", err)
		return nil, err
	}

	return tenders, nil
}
