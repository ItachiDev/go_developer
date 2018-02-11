package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go_developer/edteam/09_proyecto/commons"
	"github.com/go_developer/edteam/09_proyecto/configuration"
	"github.com/go_developer/edteam/09_proyecto/models"
)

func CommentCreate(w http.ResponseWriter, r *http.Request) {
	comment := models.Comment{}
	m := models.Message{}

	err := json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = fmt.Sprintf("Error al leer el comentario: ", err)

		commons.DisplayMessage(w, m)
		return
	}

	db := configuration.GetConnection()
	defer db.Close()

	err = db.Create(&comment).Error
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = fmt.Sprintf("Hubo un error al registrar el comentario: ", err)

		commons.DisplayMessage(w, m)
		return
	}

	m.Code = http.StatusCreated
	m.Message = "Comentario creado con exito"
	commons.DisplayMessage(w, m)
}
