package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/L0rd1k/meetpnt/app/model"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func GetProjects(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	projects := []model.Project{}
	db.Find(&projects)

	fmt.Println(projects)

	respondJSON(w, http.StatusOK, projects)
}

func CreateProject(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create project")
	project := model.Project{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&project); err != nil {
		respondErr(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()
	if err := db.Save(&project).Error; err != nil {
		respondErr(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusCreated, project)
}

func GetProject(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title := vars["title"]
	project := findProject(db, title, w, r)
	if project == nil {
		return
	}
	respondJSON(w, http.StatusOK, project)
}

func findProject(db *gorm.DB, title string, w http.ResponseWriter, r *http.Request) *model.Project {
	project := model.Project{}
	if err := db.First(&project, model.Project{Title: title}).Error; err != nil {
		respondErr(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &project
}

func UpdateProject(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title := vars["title"]
	project := findProject(db, title, w, r)
	if project == nil {
		return
	}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&project); err != nil {
		respondErr(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()
	if err := db.Save(&project).Error; err != nil {
		respondErr(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, project)
}

func DeleteProject(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title := vars["title"]
	project := findProject(db, title, w, r)
	if project == nil {
		return
	}
	if err := db.Unscoped().Delete(&project).Error; err != nil {
		respondErr(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusNoContent, project)
}
