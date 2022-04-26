package controllers

import (
	"Golang-API-tutoriel/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var Library []models.Book
var Counter int

func InitDatabase() {
	Counter = 1

	book1 := models.Book{
		Id:     1,
		Title:  "Le langage Go: Les fondamentaux du langage",
		Author: "Frédéric G. Marand",
	}
	Library = append(Library, book1)

}

func FindBooks(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": Library})
}

func CreateBook(c *gin.Context) {
	// Validate input
	var input models.Book
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	Counter++
	// Create book
	book := models.Book{Id: Counter, Title: input.Title, Author: input.Author}
	Library = append(Library, book)

	c.JSON(http.StatusOK, gin.H{"data": book})
}

func removeIt(ss models.Book, ssSlice []models.Book) []models.Book {
	for idx, v := range ssSlice {
		if v == ss {
			return append(ssSlice[0:idx], ssSlice[idx+1:]...)
		}
	}
	return ssSlice
}

func DeleteBook(c *gin.Context) {
	bookFound := false

	var bookFind models.Book

	for _, book := range Library {
		if c.Param("id") == strconv.Itoa(book.Id) {
			bookFound = true
			bookFind = book
		}
	}

	if !bookFound {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	Library = removeIt(bookFind, Library)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
