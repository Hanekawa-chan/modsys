package controllers

import (
	"awesomeProject/models/view"
	"github.com/google/uuid"
	"net/http"
)

func (t *TestHandler) teacherList(w http.ResponseWriter, r *http.Request) {
	studentId, err := t.GetAuthenticatedUserID(r)
	if err != nil {
		ReturnError(w, r, err)
		return
	}
	teachers := t.GetTeachers()
	teachersView := make([]view.Teacher, len(teachers))
	for i := range teachers {
		t, err := t.GetTeachersByStudentId(studentId)
		if err != nil {
			ReturnError(w, r, err)
			return
		}
		isAdded := false
		for j := range t {
			if t[j].Id == teachers[i].Id {
				isAdded = true
				break
			}
		}
		teachersView[i] = view.Teacher{
			Id:      teachers[i].Id,
			Name:    teachers[i].Name,
			Surname: teachers[i].Surname,
			Email:   teachers[i].Email,
			IsAdded: isAdded,
		}
	}
	data := map[string]interface{}{"title": "Преподаватели", "teachers": teachersView, "auth": true, "role": t.GetRole(r)}
	returnTemplateWithData(w, r, "teachers_list", data)
}

func (t *TestHandler) teacherGet(w http.ResponseWriter, r *http.Request) {
	studentId, err := t.GetAuthenticatedUserID(r)
	if err != nil {
		ReturnError(w, r, err)
		return
	}
	teachers, err := t.GetTeachersByStudentId(studentId)
	if err != nil {
		ReturnError(w, r, err)
		return
	}
	data := map[string]interface{}{"title": "Преподаватели", "teachers": teachers, "auth": true, "role": t.GetRole(r)}
	returnTemplateWithData(w, r, "teachers_get", data)
}

func (t *TestHandler) teacherAdd(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query().Get("id")
	teacherId, err := uuid.Parse(q)
	if err != nil {
		ReturnError(w, r, err)
		return
	}
	studentId, err := t.GetAuthenticatedUserID(r)
	if err != nil {
		ReturnError(w, r, err)
		return
	}
	err = t.AddTeacher(studentId, teacherId)
	if err != nil {
		ReturnError(w, r, err)
		return
	}
	http.Redirect(w, r, "http://localhost:8080/teacher/get", http.StatusFound)
}

func (t *TestHandler) teacherDelete(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query().Get("id")
	teacherId, err := uuid.Parse(q)
	if err != nil {
		ReturnError(w, r, err)
		return
	}
	studentId, err := t.GetAuthenticatedUserID(r)
	if err != nil {
		ReturnError(w, r, err)
		return
	}
	err = t.DeleteTeacher(studentId, teacherId)
	if err != nil {
		ReturnError(w, r, err)
		return
	}
	http.Redirect(w, r, "http://localhost:8080/teacher/get", http.StatusFound)
}

func (t *TestHandler) teacherTest(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query().Get("id")
	teacherId, err := uuid.Parse(q)
	if err != nil {
		ReturnError(w, r, err)
		return
	}
	teacher, err := t.GetUserByID(teacherId)
	if err != nil {
		ReturnError(w, r, err)
		return
	}
	tests := t.GetTestsByTeacherId(teacherId)
	data := map[string]interface{}{"title": "Тесты", "auth": true, "author": teacher.Name + "а " + teacher.Surname + "а", "tests": tests, "role": t.GetRole(r)}
	returnTemplateWithData(w, r, "teacher_tests", data)
}

func (t *TestHandler) studentsGet(w http.ResponseWriter, r *http.Request) {
	teacherId, err := t.GetAuthenticatedUserID(r)
	if err != nil {
		ReturnError(w, r, err)
		return
	}
	students, err := t.GetStudentsByTeacherId(teacherId)
	if err != nil {
		ReturnError(w, r, err)
		return
	}
	data := map[string]interface{}{"title": "Студенты", "students": students, "auth": true, "role": t.GetRole(r)}
	returnTemplateWithData(w, r, "students", data)
}
