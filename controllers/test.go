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

type Test struct {
}

type TestHandler struct {
	*services.Handler
	tests map[uuid.UUID]Test
}

func (t *TestHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/test/create":
		switch r.Method {
		case "GET":
			t.testCreateGet(w, r)
		case "POST":
			t.testCreatePost(w, r)
		}
	case "/test/update":
		switch r.Method {
		case "GET":
			t.testUpdateGet(w, r)
		case "POST":
			t.testUpdatePost(w, r)
		}
	case "/test/get":
		switch r.Method {
		case "GET":
			t.testGet(w, r)
		case "POST":
			t.testPost(w, r)
		}
	case "/test/delete":
		t.testDeleteGet(w, r)
	case "/teacher/list":
		t.teacherList(w, r)
	case "/teacher/add":
		t.teacherAdd(w, r)
	case "/teacher/delete":
		t.teacherDelete(w, r)
	case "/teacher/get":
		t.teacherGet(w, r)
	}
}

func (t *TestHandler) testCreateGet(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{"title": "Создание теста", "auth": true, "role": t.GetRole(r)}
	returnTemplateWithData(w, r, "test_create", data)
}

func (t *TestHandler) testCreatePost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		ReturnError(w, r, err)
		return
	}
	teacherId, err := t.GetAuthenticatedUserID(r)
	if err != nil {
		ReturnError(w, r, err)
		return
	}
	test := models.NewTest(teacherId, r.PostFormValue("name"))
	questions := make([]models.Question, len(r.PostForm["question"]))
	for i := range r.PostForm["question"] { // range over []string
		score, _ := strconv.Atoi(r.PostForm["score"][i])
		questions[i] = *models.NewQuestion(test.Id, r.PostForm["question"][i], r.PostForm["answer"][i], score)
	}
	test.Questions = questions
	err = t.SaveTest(test)
	if err != nil {
		ReturnError(w, r, err)
	}
	http.Redirect(w, r, "http://localhost:8080", http.StatusFound)
}

func (t *TestHandler) testGet(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query().Get("id")
	userId, err := t.GetAuthenticatedUserID(r)
	if err != nil {
		ReturnError(w, r, err)
		return
	}
	user, err := t.GetUserByID(userId)
	if user.Role.Id == services.Teacher {
		if q == "my" {
			tests := t.GetTestsByTeacherId(userId)
			data := map[string]interface{}{"title": "Тесты", "auth": true, "tests": tests, "role": t.GetRole(r)}
			returnTemplateWithData(w, r, "tests", data)
			return
		}
	} else {
		if q == "all" {
			tests := t.GetTests()
			testsAuthors := make([]view.TestAuthor, len(tests))
			for i, test := range tests {
				testsAuthors[i].Id = test.Id.String()
				testsAuthors[i].Name = test.Name
				teacher, err := t.GetUserByID(test.TeacherId)
				if err != nil {
					ReturnError(w, r, err)
					return
				}
				testsAuthors[i].Author = teacher.Name + " " + teacher.Surname
			}
			data := map[string]interface{}{"title": "Тесты", "auth": true, "tests": testsAuthors, "role": t.GetRole(r)}
			returnTemplateWithData(w, r, "tests", data)
			return
		}
		id, err := uuid.Parse(q)
		if err != nil {
			ReturnError(w, r, err)
			return
		}

		test, err := t.GetTestById(id)
		if err != nil {
			ReturnError(w, r, err)
			return
		}
		fmt.Println(test)
		data := map[string]interface{}{"test": test, "auth": true, "title": "Тест", "role": t.GetRole(r)}
		returnTemplateWithData(w, r, "test", data)
	}
}

func (t *TestHandler) testPost(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query().Get("id")
	testId, err := uuid.Parse(q)
	if err != nil {
		ReturnError(w, r, err)
		return
	}
	test, err := t.GetTestById(testId)
	if err != nil {
		ReturnError(w, r, err)
		return
	}
	err = r.ParseForm()
	if err != nil {
		ReturnError(w, r, err)
		return
	}
	userId, err := t.GetAuthenticatedUserID(r)
	if err != nil {
		ReturnError(w, r, err)
		return
	}
	var answers []models.Answer
	var result models.Result
	result.TestId = testId
	result.StudentId = userId
	result.Id = uuid.New()
	score := 0
	for i := range r.PostForm["answer"] { // range over []string
		answer := models.Answer{
			QuestionId: test.Questions[i].Id,
			ResultId:   result.Id,
			Answer:     r.PostForm["answer"][i],
			Question:   test.Questions[i],
		}
		answers = append(answers, answer)
		log.Info().Msg("Ответ ученика:" + answer.Answer + "; Правильный ответ:" + test.Questions[i].Answer)
		if answer.Answer == test.Questions[i].Answer {
			score += test.Questions[i].Score
		}
	}
	result.Score = int16(score)
	err = t.SaveResult(result)
	if err != nil {
		ReturnError(w, r, err)
		return
	}
	err = t.SaveAnswers(answers)
	if err != nil {
		ReturnError(w, r, err)
		return
	}
	http.Redirect(w, r, "http://localhost:8080", http.StatusFound)
}

func (t *TestHandler) testUpdateGet(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query().Get("id")
	testId, err := uuid.Parse(q)
	if err != nil {
		ReturnError(w, r, err)
		return
	}
	test, err := t.GetTestById(testId)
	if err != nil {
		ReturnError(w, r, err)
		return
	}
	data := map[string]interface{}{"title": "Редактирование теста", "auth": true,
		"role": t.GetRole(r), "test": test}
	returnTemplateWithData(w, r, "test_create", data)
}

func (t *TestHandler) testUpdatePost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		ReturnError(w, r, err)
		return
	}
	teacherId, err := t.GetAuthenticatedUserID(r)
	if err != nil {
		ReturnError(w, r, err)
		return
	}
	test := models.NewTest(teacherId, r.PostFormValue("name"))
	questions := make([]models.Question, len(r.PostForm["question"]))
	for i := range r.PostForm["question"] { // range over []string
		score, _ := strconv.Atoi(r.PostForm["score"][i])
		questions[i] = *models.NewQuestion(test.Id, r.PostForm["question"][i], r.PostForm["answer"][i], score)
	}
	test.Questions = questions
	err = t.SaveTest(test)
	if err != nil {
		ReturnError(w, r, err)
	}
	http.Redirect(w, r, "http://localhost:8080", http.StatusFound)
}

func (t *TestHandler) testDeleteGet(w http.ResponseWriter, r *http.Request) {
	//q := r.URL.Query().Get("id")
	//testId, err := uuid.Parse(q)
	//if err != nil {
	//	ReturnError(w, r, err)
	//	return
	//}
	//err = t.DeleteTestById(testId)
	//if err != nil {
	//	ReturnError(w, r, err)
	//	return
	//}
	http.Redirect(w, r, "http://localhost:8080/test/get?id=my", http.StatusFound)
}
