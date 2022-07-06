package models

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	Id int `gorm:"primaryKey;autoIncrement:true"`
	Isbn string
	Title string
	Author string
	Publisher string
	Category string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type BookModels struct {
	Db *gorm.DB
	books []Book
}


func (bm *BookModels) Init() {
	if !bm.Db.Migrator().HasTable(&Book{}) {
		bm.Migrate()
	}
}

func (bm *BookModels) GetAll() *[]Book {
	var b []Book
	bm.Db.Find(&b)
	return &b
}

func (bm *BookModels) Get(id int) *Book {
	var b Book
	bm.Db.First(&b, id)
	return &b
}

func (bm *BookModels) Create(book *Book) {
	bm.Db.Create(&book)
}

func (bm *BookModels) Update(book *Book) {
	bm.Db.Save(&book)
}

func (bm *BookModels) Delete(book *Book) {
	bm.Db.Delete(&book)
}

func (bm *BookModels) Query() {

}

func (bm *BookModels) Migrate() {
	bm.Db.AutoMigrate(&Book{})
}