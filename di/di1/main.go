package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/ludwig125/ludwig125_gosample/di/di1/model"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	config := NewConfig()

	db, err := ConnectDatabase(config)

	if err != nil {
		panic(err)
	}

	personRepository := NewPersonRepository(db)

	personService := NewPersonService(config, personRepository)

	server := NewServer(config, personService)

	server.Run()
}

type Config struct {
	Enabled      bool
	DatabasePath string
	Port         string
}

func NewConfig() *Config {
	return &Config{
		Enabled:      true,
		DatabasePath: "./example.db",
		Port:         "8080",
	}
}

func ConnectDatabase(config *Config) (*sql.DB, error) {
	return sql.Open("sqlite3", config.DatabasePath)
}

type PersonRepository interface {
	FindAll() []*model.Person
}

type personRepository struct {
	database *sql.DB
}

func (repository *personRepository) FindAll() []*model.Person {
	rows, _ := repository.database.Query(`SELECT id, name, age FROM people;`)
	defer rows.Close()

	people := []*model.Person{}

	for rows.Next() {
		var (
			id   int
			name string
			age  int
		)

		rows.Scan(&id, &name, &age)

		people = append(people, &model.Person{
			Id:   id,
			Name: name,
			Age:  age,
		})
	}

	return people
}

func NewPersonRepository(database *sql.DB) *personRepository {
	return &personRepository{database: database}
}

type PersonService interface {
	FindAll() []*model.Person
}

type personService struct {
	config     *Config
	repository PersonRepository
	// repository *PersonRepository
}

func (service *personService) FindAll() []*model.Person {
	if service.config.Enabled {
		return service.repository.FindAll()
	}

	return []*model.Person{}
}

func NewPersonService(config *Config, repository PersonRepository) *personService {
	return &personService{config: config, repository: repository}
}

type Server struct {
	config        *Config
	PersonService PersonService
}

func (server *Server) Handler() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/people", server.findPeople)

	return mux
}

func (server *Server) Run() {
	httpServer := &http.Server{
		Addr:    "localhost:" + server.config.Port,
		Handler: server.Handler(),
	}

	log.Println("please check `http://localhost:8080/people`")
	httpServer.ListenAndServe()
}

func (server *Server) findPeople(writer http.ResponseWriter, request *http.Request) {
	people := server.PersonService.FindAll()
	bytes, _ := json.Marshal(people)

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(bytes)
}

func NewServer(config *Config, PersonService PersonService) *Server {
	return &Server{
		config:        config,
		PersonService: PersonService,
	}
}
