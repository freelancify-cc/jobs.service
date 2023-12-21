package jobshandler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/freelancify/jobs/internal/database"
	"github.com/freelancify/jobs/internal/models"
	"github.com/google/uuid"
)

func GetAllJobs(w http.ResponseWriter, r *http.Request) {
	job, err := database.SelectAllJobs()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	b, err := json.Marshal(&job)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(b)
	return
}

func CreateJob(w http.ResponseWriter, r *http.Request) {
	println("comes here")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id, err := uuid.Parse(r.Context().Value("id").(string))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var job models.JobModel
	err = json.Unmarshal(body, &job)
	job.PostedEmployer = id
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	database.InsertJob(&job)
	msg := map[string]interface{}{
		"id": job.Id,
	}

	j, err := json.Marshal(msg)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte(j))
	return
}

func GetJobDetails(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.Context().Value("id").(string))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	job, err := database.SelectJobById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	b, err := json.Marshal(&job)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(b)
	return
}
