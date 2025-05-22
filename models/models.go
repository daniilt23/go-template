package models

type Book struct {
    ID          uint   `json:"id" gorm:"primaryKey"`
    Title       string `json:"title"`
    Description string `json:"description"`
    AuthorID    uint   `json:"author_id"`
}

type Author struct {
    ID uint `json:"id" gorm:"primaryKey"`
    Name string `json:"name"`
    Book []Book `json:"books" gorm:"foreignKey:AuthorID"`  
}