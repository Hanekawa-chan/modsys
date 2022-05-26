package controllers

import (
	"awesomeProject/models"
	"awesomeProject/services"
	"net/http"
)

type ResultHandler struct {
	*services.Handler
}

func (h *ResultHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/result/get":
		h.resultGet(w, r)
	}
}

func (h *ResultHandler) resultGet(w http.ResponseWriter, r *http.Request) {
	userId, err := h.GetAuthenticatedUserID(r)
	if err != nil {
		ReturnError(w, r, err)
		return
	}
	results := h.GetResults(userId)
	resultsView := make([]models.ResultView, len(results))
	for i := range results {
		resultsView[i].Id = results[i].Id
		test, err := h.GetTestByID(results[i].TestId)
		if err != nil {
			ReturnError(w, r, err)
			return
		}
		resultsView[i].Name = test.Name
		teacherId := test.TeacherId
		teacher, err := h.GetUserByID(teacherId)
		if err != nil {
			ReturnError(w, r, err)
			return
		}
		resultsView[i].Author = teacher.Name + " " + teacher.Surname
		resultsView[i].Score = results[i].Score
	}
	data := map[string]interface{}{"title": "Результаты", "results": resultsView}
	returnTemplateWithData(w, r, "result", data)
}
