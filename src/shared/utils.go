package shared

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetPathParam(req *http.Request, name string) string {
	return mux.Vars(req)[name]
}

func StringToObjectId(hex string) primitive.ObjectID {
	objID, err := primitive.ObjectIDFromHex(hex)
	if err != nil {
		panic(err)
	}

	return objID
}

func DecodeBody(req *http.Request, v interface{}) error {
	err := json.NewDecoder(req.Body).Decode(v)
	if err != nil {
		return err
	}
	return nil
}
