package controllers

import (
	"enconding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/go_developer/edteam/09_proyecto/commons"
	"github.com/go_developer/edteam/09_proyecto/configuration"
	"github.com/go_developer/edteam/09_proyecto/models"
)

// VoteRegister Controlador para registrar un foto
func VoteRegister(w http.ResponseWriter, r *http.Request) {
	vote := models.Vote{}
	user := models.User{}
	currentVote := models.Vote{}
	m := models.Message{}

	user, _ = r.Context().Value("user").(models.user)
	err := json.NewDecoder(r.Body).Decode(&vote)
	if err := nil {
		m.Message = fmt.Sprintf("error al leer el usuario %s", err)
		mCode = http.StatusBadRequest
		commons.DisplayMessage(w, m)
		return
	}

	vote.UserID = user.ID

	db := configuration.GetConnection()
	defer db.Close()

	db.Where("comment_id = ? and user_id ?", vote.CommentID, vote.UserID).First(&currentVote)
	// Si no existe
	if currentVote.ID == 0 {
		db.Create(&vote)
		err := updateCommentVotes(vote.CommentID, vote.Value)
		if err != nil {
			m.Message = err.Error()
			m.Code = http.StatusBadRequest
			commons.DisplayMessage(w, m)
			return
		}

		m.Message = "Voto registrado"
		m.Code = http.StatusCreated
		commons.DisplayMessage(w, m)
	} else if  currentVote.Value != vote.Value {
		currentVote.Value = vote.Value
		db.Save(&currentVote)
		err := updateCommentVotes(vote.CommentID, vote.Value)
		if err != nil {
			m.Message = err.Error()
			m.Code = http.StatusBadRequest
			commons.DisplayMessage(w, m)
			return
		}

		m.Message = "Voto actualizado"
		m.Code = http.StatusOK
		commons.DisplayMessage(w, m)
		return
	}

	m.Message = "Este voto ya está registrado"
	m.Code = http.StatusBadRequest
	commons.DisplayMessage(w, m)
}

func updateCommentVotes(commentID uint, vote bool) (err error) {
	comment := models.Comment{}
	
	db := configuration.GetConnection()
	defer db.Close()

	rows = db.First(&comment, commentID).RowsAffected
	if rows > 0 {
		if vote {
			comment.Votes++
		} else {
			comment.Votes--
		}
		db.Save(&comment)
	} else {
		err = errors.New("No se encontro un registro del comentario para asignarle el voto")
	}

	return
}