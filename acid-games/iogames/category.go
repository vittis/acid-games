package iogames

//Category .
type Category struct {
	Name string `json:"name" validate:"required,max-length-20,unique-category-name"`
}

//RegisterCategory .
/* func RegisterCategory(w http.ResponseWriter, r *http.Request) {

	var c Category

	err := DecodeJSON(r.Body, &c)
	if err != nil {
		log.Printf("%v\n", err)
		w.Write([]byte(responses["errorJSON"]))
		return
	}

	log.Printf("Category populado: %v", c)

	err = validate.Struct(c)
	if err != nil {
		log.Printf("%v\n", err)
		w.Write([]byte(responses["errorValidation"]))
		return
	}

	err = InsertCategory(c)
	if err != nil {
		log.Printf("%v\n", err)
		w.Write([]byte(responses["errorInsertDB"]))
		return
	}

	w.Write([]byte(responses["sucess"]))

	return
} */

//ListAllCategories .
/* func ListAllCategories(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	categories, err := GetAllCategories()

	if err != nil {
		log.Printf("%v\n", err)
		w.Write([]byte(responses["get-from-db-error"]))
		return
	}

	categoriesToJSON, err := EncodeJSON(categories)

	if err != nil {
		log.Printf("%v\n", err)
		w.Write([]byte(responses["bad-encode-json"]))
		return
	}

	w.Write([]byte(categoriesToJSON))
}
*/
