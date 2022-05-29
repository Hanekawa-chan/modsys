package controllers

import (
	"awesomeProject/models"
	"awesomeProject/models/view"
	"awesomeProject/services"
	"fmt"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"net/http"
	"strconv"
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
	q := r.URL.Query().Get("id")
	userId, err := h.GetAuthenticatedUserID(r)
	if err != nil {
		ReturnError(w, r, err)
		return
	}
	if q == "all" {
		results := h.GetResults(userId)
		resultsView := make([]view.ResultView, len(results))
		for i := range results {
			resultsView[i].Id = results[i].Id
			test, err := h.GetTestById(results[i].TestId)
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
		data := map[string]interface{}{"title": "Результаты", "auth": true,
			"results": resultsView, "role": h.GetRole(r)}
		returnTemplateWithData(w, r, "results", data)
	} else {
		resultId, err := uuid.Parse(q)
		if err != nil {
			ReturnError(w, r, err)
			return
		}
		result := h.GetResultById(resultId)
		log.Info().Msg(result.Id.String())
		var answers []models.Answer
		answers = h.GetAnswersByResultId(result.Id)
		for i := range answers {
			answers[i].Question = h.GetQuestionById(answers[i].QuestionId)
		}
		fmt.Println(result)
		test, err := h.GetTestById(result.TestId)
		if err != nil {
			ReturnError(w, r, err)
			return
		}
		fmt.Println(test.Name)
		teacher, err := h.GetUserByID(test.TeacherId)
		if err != nil {
			ReturnError(w, r, err)
			return
		}
		log.Info().Msg(strconv.Itoa(len(answers)))
		resultFull := view.ResultFull{Name: test.Name, Author: teacher.Name + " " +
			test.Teacher.Surname, Score: result.Score, Answers: answers}
		data := map[string]interface{}{"title": "Результаты Теста", "auth": true,
			"result": resultFull, "role": h.GetRole(r)}
		returnTemplateWithData(w, r, "result", data)
	}
}
