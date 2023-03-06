package bookhandler

import (
	"encoding/json"
	"example/server/db"
	book_migration "example/server/db/migration/book"
	globalhandler "example/server/handler/global_handler"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"gopkg.in/validator.v2"
)

type BooksConfiguration struct {
	Name        string `validate:"nonzero"`
	Description string `validate:"nonzero"`
	Price       string `validate:"nonzero"`
	IsActive    string `validate:"nonzero"`
}

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	var books []book_migration.Books
	db.DataBase.Find(&books)

	res := globalhandler.BasicOutputResponse{
		Response: globalhandler.BasicResponse{
			Status:    true,
			Message:   "Success Get data",
			StatusMSG: "success",
		},
		Results: books,
		Param:   make(map[string]string),
	}
	res.Param["book_id"] = ""
	results, _ := json.Marshal(res)
	w.WriteHeader(http.StatusOK)
	w.Write(results)
}

func Store(w http.ResponseWriter, r *http.Request) {
	post, _ := ioutil.ReadAll(r.Body)
	var book book_migration.Books
	json.Unmarshal(post, &book)
	// validation
	validation := BooksConfiguration{
		Name:        book.Name,
		Description: book.Description,
		Price:       book.Price,
		IsActive:    book.IsActive,
	}
	if errs := validator.Validate(validation); errs != nil {
		resErr := globalhandler.BasicOutputResponse{
			Response: globalhandler.BasicResponse{
				Status:    false,
				Message:   errs.Error(),
				StatusMSG: "Validation failed",
			},
		}
		resultErr, _ := json.Marshal(resErr)
		w.Header().Set("Content-type", "aplication/json")
		w.WriteHeader(http.StatusOK)
		w.Write(resultErr)
		return
	}
	db.DataBase.Create(&book)

	res := globalhandler.BasicOutputResponse{
		Response: globalhandler.BasicResponse{
			Status:    true,
			Message:   "Success add dataa",
			StatusMSG: "success",
		},
		Results: post,
		Param:   make(map[string]string),
	}
	res.Param["book_id"] = ""
	results, _ := json.Marshal(res)
	w.WriteHeader(http.StatusOK)
	w.Write(results)
}

func Detail(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	bookId := param["book_id"]

	var book book_migration.Books
	if err := db.DataBase.First(&book, bookId).Error; err != nil {
		res := globalhandler.BasicOutputResponse{
			Response: globalhandler.BasicResponse{
				Status:    false,
				Message:   "Failed",
				StatusMSG: "Record not found",
			},
			Param: make(map[string]string),
		}
		res.Param["book_id"] = bookId
		results, _ := json.Marshal(res)
		w.WriteHeader(http.StatusOK)
		w.Write(results)
		return
	}
	res := globalhandler.BasicOutputResponse{
		Response: globalhandler.BasicResponse{
			Status:    true,
			Message:   "Success",
			StatusMSG: "Success",
		},
		Param:   make(map[string]string),
		Results: book,
	}
	res.Param["book_id"] = bookId
	results, _ := json.Marshal(res)
	w.WriteHeader(http.StatusOK)
	w.Write(results)
}

func Update(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	bookId := param["book_id"]

	post, _ := ioutil.ReadAll(r.Body)
	var book book_migration.Books
	json.Unmarshal(post, &book)
	// validation
	validation := BooksConfiguration{
		Name:        book.Name,
		Description: book.Description,
		Price:       book.Price,
		IsActive:    book.IsActive,
	}
	if errs := validator.Validate(validation); errs != nil {
		resErr := globalhandler.BasicOutputResponse{
			Response: globalhandler.BasicResponse{
				Status:    false,
				Message:   errs.Error(),
				StatusMSG: "Validation failed",
			},
		}
		resultErr, _ := json.Marshal(resErr)
		w.Header().Set("Content-type", "aplication/json")
		w.WriteHeader(http.StatusOK)
		w.Write(resultErr)
		return
	}
	var getBook book_migration.Books
	if err := db.DataBase.First(&getBook, bookId).Error; err != nil {
		resErr := globalhandler.BasicOutputResponse{
			Response: globalhandler.BasicResponse{
				Status:    false,
				Message:   err.Error(),
				StatusMSG: "Record Not found",
			},
		}
		resultErr, _ := json.Marshal(resErr)
		w.Header().Set("Content-type", "aplication/json")
		w.WriteHeader(http.StatusOK)
		w.Write(resultErr)
		return
	}

	db.DataBase.Model(&getBook).Updates(book)
	res := globalhandler.BasicOutputResponse{
		Response: globalhandler.BasicResponse{
			Status:    true,
			Message:   "Success updated data",
			StatusMSG: "Success",
		},
		Results: getBook,
		Param:   make(map[string]string),
	}
	res.Param["book_id"] = bookId

	results, _ := json.Marshal(res)
	w.WriteHeader(http.StatusOK)
	w.Write(results)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	bookId := param["book_id"]

	var book book_migration.Books
	if err := db.DataBase.First(&book, bookId).Error; err != nil {
		res := globalhandler.BasicOutputResponse{
			Response: globalhandler.BasicResponse{
				Status:    false,
				Message:   "Record Not Found",
				StatusMSG: "failed",
			},
			Param: make(map[string]string),
		}
		res.Param["book_id"] = bookId
		result, _ := json.Marshal(res)
		w.WriteHeader(http.StatusOK)
		w.Write(result)
		return
	}
	db.DataBase.Delete(&book)
	res := globalhandler.BasicOutputResponse{
		Response: globalhandler.BasicResponse{
			Status:    true,
			Message:   "Success Deleted Data",
			StatusMSG: "Success",
		},
		Param: make(map[string]string),
	}
	res.Param["book_id"] = bookId
	result, _ := json.Marshal(res)
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}
