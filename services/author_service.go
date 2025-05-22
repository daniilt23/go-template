package services

import (
	"fmt"
	"go-test/database"
	"go-test/models"
)

func GetAllAuthors(authors *[]models.Author) error {
	var err error = database.DB.Preload("Book").Find(authors).Error
	if err != nil {
		return err
	}

	return nil
}

func GetAuthorByID(author *models.Author, id string) error {
	var err error = database.DB.Where("id = ?", id).First(author).Error
	if err != nil {
		return err
	}

	return nil
}

func CreateAuthor(author *models.Author) error {
	var err error = database.DB.Create(author).Error
	if err != nil {
		return err
	}

	return nil
}

func UpdateAuthor(author *models.Author, id string) error {
	var existingAuthor models.Author
    if err := database.DB.First(&existingAuthor, id).Error; err != nil {
        return fmt.Errorf("author with id %s not found", id)
    }
	result := database.DB.Model(author).Where("id = ?", id).Updates(author)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("we havent key")
	}

	return nil
}

func DeleteAuthor(author *models.Author, id string) error {
	result := database.DB.Where("id = ?", id).Delete(author)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("author with id %s dont found", id)
	}
	
	return nil
}

