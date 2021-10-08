package account

import (
	"gorm.io/gorm"
)

func AuthorName(db *gorm.DB, authorId float64) (string, error) {
	var authorName string
	result := db.Raw("SELECT name FROM s_user WHERE id = ?", authorId).Scan(&authorName)

	err := result.Error
	if err != nil {
		return "", err
	}

	println(authorName + " Selected")
	return authorName, nil
}
