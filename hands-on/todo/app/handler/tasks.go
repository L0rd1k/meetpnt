package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/L0rd1k/meetpnt/app/model"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func GetTasks(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	projectTitle := vars["title"]
	project := findProject(db, projectTitle, w, r)
	if project == nil {
		return
	}
	tasks := []model.Task{}
	if err := db.Model(&project).Related(&tasks).Error; err != nil {
		respondErr(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, tasks)
}

func CreateTask(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	projectTitle := vars["title"]
	project := findProject(db, projectTitle, w, r)
	if project == nil {
		return
	}
	task := model.Task{ProjectID: project.ID}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&task); err != nil {
		respondErr(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()
	if err := db.Save(&task).Error; err != nil {
		respondErr(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusCreated, task)
}

func GetTask(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	projectTitle := vars["title"]
	project := findProject(db, projectTitle, w, r)
	if project == nil {
		return
	}
	id, _ := strconv.Atoi(vars["id"])
	task := findTask(db, id, w, r)
	if task == nil {
		return
	}
	respondJSON(w, http.StatusOK, task)
}

func DeleteTask(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	projectTitle := vars["title"]
	project := findProject(db, projectTitle, w, r)
	if project == nil {
		return
	}
	id, _ := strconv.Atoi(vars["id"])
	task := findTask(db, id, w, r)
	if task == nil {
		return
	}
	if err := db.Delete(&project).Error; err != nil {
		respondErr(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusNoContent, nil)
}

func UpdateTask(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	projectTitle := vars["title"]
	project := findProject(db, projectTitle, w, r)
	if project == nil {
		return
	}
	id, _ := strconv.Atoi(vars["id"])
	task := findTask(db, id, w, r)
	if task == nil {
		return
	}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&task); err != nil {
		respondErr(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()
	if err := db.Save(&task).Error; err != nil {
		respondErr(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, task)
}

func findTask(db *gorm.DB, id int, w http.ResponseWriter, r *http.Request) *model.Task {
	task := model.Task{}
	// SELECT * FROM task ORDER BY id LIMIT 1;
	if err := db.First(&task, id).Error; err != nil {
		respondErr(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &task
}
