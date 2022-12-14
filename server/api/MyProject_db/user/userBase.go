package user

/**
 * 
 * 
  _____                      _              _ _ _     _   _     _        __ _ _      
 |  __ \                    | |            | (_) |   | | | |   (_)      / _(_) |     
 | |  | | ___    _ __   ___ | |_    ___  __| |_| |_  | |_| |__  _ ___  | |_ _| | ___ 
 | |  | |/ _ \  | '_ \ / _ \| __|  / _ \/ _` | | __| | __| '_ \| / __| |  _| | |/ _ \
 | |__| | (_) | | | | | (_) | |_  |  __/ (_| | | |_  | |_| | | | \__ \ | | | | |  __/
 |_____/ \___/  |_| |_|\___/ \__|  \___|\__,_|_|\__|  \__|_| |_|_|___/ |_| |_|_|\___|
 
 * DO NOT EDIT THIS FILE!! 
 * 
 *  TO CUSTOMIZE UserBase.go PLEASE EDIT ./User.go
 * 
 *  -- THIS FILE WILL BE OVERWRITTEN ON THE NEXT SKAFFOLDER'S CODE GENERATION --
 * 
 */

import (
	"encoding/json"
	"net/http"

	"github.com/globalsign/mgo/bson"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

// Data model
type User struct {
	ID		bson.ObjectId	`bson:"_id,omitempty" json:"_id"`
	Fullname	string		`json:"fullname"`
	Mail	string		`json:"mail"`
	Password	string		`json:"password"`
	Roles	string		`json:"roles"`
	Username	string		`json:"username"`


}

// Functions

// CRUD METHODS
	
// CRUD - CREATE
func (config *Config) create(writer http.ResponseWriter, req *http.Request) {

	var body User
	json.NewDecoder(req.Body).Decode(&body)
	config.Database.C("User").Insert(body)
	render.JSON(writer, req, &body)
}

// CRUD - DELETE
func (config *Config) delete(writer http.ResponseWriter, req *http.Request) {

	id := chi.URLParam(req, "id")
	config.Database.C("User").RemoveId(bson.ObjectIdHex(id))
	render.JSON(writer, req, id)
}

// CRUD - GET
func (config *Config) get(writer http.ResponseWriter, req *http.Request) {

	var result User
	id := chi.URLParam(req, "id")
	err := config.Database.C("User").FindId(bson.ObjectIdHex(id)).One(&result)
	if err != nil {
		panic(err)
	}

	render.JSON(writer, req, result)
}
	
// CRUD - GET LIST
func (config *Config) list(writer http.ResponseWriter, req *http.Request) {

	var result []User
	err := config.Database.C("User").Find(nil).All(&result)
	if err != nil {
		panic(err)
	}

	render.JSON(writer, req, result)
}

// CRUD - UPDATE
func (config *Config) update(writer http.ResponseWriter, req *http.Request) {

	id := chi.URLParam(req, "id")
	var body User
	json.NewDecoder(req.Body).Decode(&body)
	config.Database.C("User").UpdateId(bson.ObjectIdHex(id), body)
	render.JSON(writer, req, body)
}

/*
 * CUSTOM SERVICES
 * 
 *	These services will be overwritten and implemented in  Ext.go
 */


/*
Name: changePassword
Description: Change password of user from admin
Params: 
*/
func (config *Config) changePassword(writer http.ResponseWriter, request *http.Request) {
	response := make(map[string]string)
	render.JSON(writer, request, response)
}		
