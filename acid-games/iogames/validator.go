package iogames

import (
	"log"
	"net/url"

	"gopkg.in/go-playground/validator.v9"
)

var validate *validator.Validate

//StartValidator .
func StartValidator() {
	validate = validator.New()
	validate.RegisterValidation("max-length-20", ValidateMaxLength20)
	validate.RegisterValidation("max-length-50", ValidateMaxLength50)
	validate.RegisterValidation("valid-category", ValidateCategory)
	validate.RegisterValidation("valid-URL", ValidateURL)
	validate.RegisterValidation("unique-dev-username", ValidateUniqueDevUsername)
	validate.RegisterValidation("unique-game-title", ValidateUniqueGameTitle)
	validate.RegisterValidation("unique-category-name", ValidateUniqueCategoryName)

}

//ValidateUniqueCategoryName .
func ValidateUniqueCategoryName(fl validator.FieldLevel) bool {
	category, err := GetCategoryByName(fl.Field().String())

	//is empty
	if (Category{}) == category {
		return true
	}

	if err != nil {
		log.Printf("[ERROR]: %v\n", err)
		//error is "not found"
		return true
	}

	return false
}

//ValidateUniqueGameTitle .
func ValidateUniqueGameTitle(fl validator.FieldLevel) bool {
	game, err := GetGameByTitle(fl.Field().String())

	//is empty
	if (Game{}) == game {
		return true
	}

	if err != nil {
		log.Printf("[ERROR]: %v\n", err)
		//error is "not found"
		return true
	}

	return false
}

//ValidateUniqueDevUsername .
func ValidateUniqueDevUsername(fl validator.FieldLevel) bool {
	dev, err := GetDeveloperByUserName(fl.Field().String())

	//is empty
	if (DBDeveloper{}) == dev {
		return true
	}

	if err != nil {
		log.Printf("[ERROR]: %v\n", err)
		//error is "not found"
		return true
	}

	return false
}

//ValidateMaxLength20 .
func ValidateMaxLength20(fl validator.FieldLevel) bool {
	return len(fl.Field().String()) <= 20
}

//ValidateMaxLength50 .
func ValidateMaxLength50(fl validator.FieldLevel) bool {
	return len(fl.Field().String()) <= 50
}

//ValidateCategory .
func ValidateCategory(fl validator.FieldLevel) bool {
	categories, err := GetAllCategories()

	if err != nil {
		log.Printf("%v\n", err)
		return false
	}

	for i := 0; i < len(categories); i++ {
		if categories[i].Name == fl.Field().String() {
			return true
		}
	}

	return false
}

//ValidateURL .
func ValidateURL(fl validator.FieldLevel) bool {
	_, err := url.ParseRequestURI(fl.Field().String())
	return err == nil
}
