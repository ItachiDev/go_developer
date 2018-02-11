package migration

import (
	"github.com/go_developer/edteam/09_proyecto/configuration"
	"github.com/go_developer/edteam/09_proyecto/models"
)

func Migrate() {
	db := configuration.GetConnection()
	defer db.Close()

	db.CreateTable(&models.User{})
	db.CreateTable(&models.Comment{})
	db.CreateTable(&models.Vote{})
	db.Model(&models.Vote{}).AddUniqueIndex("comment_id_user_id_unique", "comment_id", "user_id")
}
