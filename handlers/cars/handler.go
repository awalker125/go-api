package cars

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"

	"github.com/awalker125/go-api/models/cars"
	"github.com/awalker125/go-api/store"
	"github.com/gosimple/slug"
)

type CarsHandler struct {
	Make  string
	store store.CarStore
}

var (
	CarRe       = regexp.MustCompile(`^/v1/cars/*$`)
	CarReWithID = regexp.MustCompile(`^/v1/cars/([a-z0-9]+(?:-[a-z0-9]+)+)$`)
)

func NewCarsHandler(s store.CarStore) *CarsHandler {
	return &CarsHandler{
		store: s,
	}
}

func (h *CarsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte(fmt.Sprintf("This is %s cars home page", h.Make)))

	log.Printf("In handler %s", r.Method)

	switch {
	case r.Method == http.MethodPost && CarRe.MatchString(r.URL.Path):
		log.Println("In handler")
		h.CreateCar(w, r)
		return
	case r.Method == http.MethodGet && CarRe.MatchString(r.URL.Path):
		h.ListCars(w, r)
		return
	case r.Method == http.MethodGet && CarReWithID.MatchString(r.URL.Path):
		h.GetCar(w, r)
		return
	case r.Method == http.MethodPut && CarReWithID.MatchString(r.URL.Path):
		h.UpdateCar(w, r)
		return
	case r.Method == http.MethodDelete && CarReWithID.MatchString(r.URL.Path):
		h.DeleteCar(w, r)
		return
	default:
		return
	}
}

// CreateCar example
//
//	@Summary		Add a new car
//	@Description	Create a car
//	@ID				post-car
//	@Accept			json
//	@Produce		json
//	@Param			car	body		cars.Cars		true	"Car"
//	@Success		200	{object}	cars.Cars		"this is a comment"
//	@Failure		400	{object}	web.APIError	"We need ID!!"
//	@Failure		404	{object}	web.APIError	"Can not find ID"
//	@Router			/cars/ [post]
func (h *CarsHandler) CreateCar(w http.ResponseWriter, r *http.Request) {

	log.Println("In handler")

	// Recipe object that will be populated from JSON payload
	var car cars.Cars
	if err := json.NewDecoder(r.Body).Decode(&car); err != nil {
		InternalServerErrorHandler(w, r)
		return
	}

	// Convert the name of the recipe into URL friendly string
	resourceID := slug.Make(car.Model)

	log.Println("resourceID %s", resourceID)
	// Call the store to add the recipe
	if err := h.store.Add(resourceID, car); err != nil {
		InternalServerErrorHandler(w, r)
		return
	}

	jsonBytes, err := json.Marshal(car)

	log.Print(fmt.Sprintf("%s", jsonBytes))

	if err != nil {
		InternalServerErrorHandler(w, r)
		return
	}
	// Set the status code to 200
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)

}

func (h *CarsHandler) ListCars(w http.ResponseWriter, r *http.Request)  {}
func (h *CarsHandler) GetCar(w http.ResponseWriter, r *http.Request)    {}
func (h *CarsHandler) UpdateCar(w http.ResponseWriter, r *http.Request) {}
func (h *CarsHandler) DeleteCar(w http.ResponseWriter, r *http.Request) {}

func InternalServerErrorHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("500 Internal Server Error"))
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("404 Not Found"))
}
