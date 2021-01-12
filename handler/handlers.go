package handler

import (
	"log"

	. "github.com/cleslley/api-rest-go/model"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//MoviesDB armazena os dados do BD
type MoviesDB struct {
	Server   string
	Database string
}

var db *mgo.Database

//Collection armazena o nome da coleção
const Collection = "movies"

//Connect inicia a conexão com o mongoDB
func (m *MoviesDB) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)

}

//GetAll retorna todos os dados do banco
func (m *MoviesDB) GetAll() ([]Movie, error) {
	var movies []Movie
	err := db.C(Collection).Find(bson.M{}).All(&movies)
	return movies, err
}

//GetByID faz uma consulta por ID
func (m *MoviesDB) GetByID(id string) (Movie, error) {
	var movie Movie
	err := db.C(Collection).FindId(bson.ObjectIdHex(id)).One(&movie)
	return movie, err
}

//Create insere um novo dado pelo metodo POST
func (m *MoviesDB) Create(movie Movie) error {
	err := db.C(Collection).Insert(&movie)
	return err
}

//Delete remove um dado por ID
func (m *MoviesDB) Delete(id string) error {
	err := db.C(Collection).RemoveId(bson.ObjectIdHex(id))
	return err
}
