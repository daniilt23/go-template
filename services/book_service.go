package services

import (
	"fmt"
	"go-test/database"
	"go-test/models"
)

func GetAllBooks(books *[]models.Book) error {
	var err error = database.DB.Find(books).Error
	if err != nil {
		return err
	}

	return nil
}

func GetBookByID(book *models.Book, id string) error {
	var err error = database.DB.Where("id = ?", id).First(book).Error
	if err != nil {
		return err
	}

	return nil
}

func CreateBook(book *models.Book) error {
	var err error = database.DB.Create(book).Error
	if err != nil {
		return err
	}

	return nil
}

func UpdateBook(book *models.Book, id string) error {
	result := database.DB.Model(book).Where("id = ?", id).Updates(book)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("book with id %s not found", id)
	}

	return nil
}

func DeleteBook(book *models.Book, id string) error {
	result := database.DB.Where("id = ?", id).Delete(book)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("book with id %s not found", id)
	} 
	
	return nil
}

