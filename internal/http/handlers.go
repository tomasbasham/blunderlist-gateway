package http

import (
	"net/http"

	"github.com/google/jsonapi"

	"github.com/tomasbasham/blunderlist-gateway/internal/comment"
	"github.com/tomasbasham/blunderlist-gateway/internal/todo"
)

func (s *Server) listTasks(w http.ResponseWriter, r *http.Request) {
	jsonapiRuntime := jsonapi.NewRuntime().Instrument("tasks.list")

	tasks, err := s.store.ListTasks(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	if err := jsonapiRuntime.MarshalPayload(w, tasks); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (s *Server) getTask(w http.ResponseWriter, r *http.Request) {
	jsonapiRuntime := jsonapi.NewRuntime().Instrument("tasks.get")

	var id uint
	if err := requestVar(r, "id", &id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	task, err := s.store.GetTask(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	if err := jsonapiRuntime.MarshalPayload(w, task); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (s *Server) createTask(w http.ResponseWriter, r *http.Request) {
	jsonapiRuntime := jsonapi.NewRuntime().Instrument("tasks.create")

	var in todo.Task
	if err := jsonapiRuntime.UnmarshalPayload(r.Body, &in); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	task, err := s.store.CreateTask(r.Context(), &in)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	w.WriteHeader(http.StatusCreated)

	if err := jsonapiRuntime.MarshalPayload(w, task); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (s *Server) updateTask(w http.ResponseWriter, r *http.Request) {
	jsonapiRuntime := jsonapi.NewRuntime().Instrument("tasks.update")

	var in todo.Task
	if err := jsonapiRuntime.UnmarshalPayload(r.Body, &in); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	task, err := s.store.UpdateTask(r.Context(), &in)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	if err := jsonapiRuntime.MarshalPayload(w, task); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (s *Server) deleteTask(w http.ResponseWriter, r *http.Request) {
	jsonapi.NewRuntime().Instrument("tasks.delete")

	var id uint
	if err := requestVar(r, "id", &id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := s.store.DeleteTask(r.Context(), id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (s *Server) listComments(w http.ResponseWriter, r *http.Request) {
	jsonapiRuntime := jsonapi.NewRuntime().Instrument("comments.list")

	var id uint
	if err := requestVar(r, "id", &id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	comments, err := s.store.ListComments(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	if err := jsonapiRuntime.MarshalPayload(w, comments); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (s *Server) getComment(w http.ResponseWriter, r *http.Request) {
	jsonapiRuntime := jsonapi.NewRuntime().Instrument("comments.get")

	var id uint
	if err := requestVar(r, "id", &id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	comment, err := s.store.GetComment(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	if err := jsonapiRuntime.MarshalPayload(w, comment); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (s *Server) createComment(w http.ResponseWriter, r *http.Request) {
	jsonapiRuntime := jsonapi.NewRuntime().Instrument("comments.create")

	var in comment.Comment
	if err := jsonapiRuntime.UnmarshalPayload(r.Body, &in); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	comment, err := s.store.CreateComment(r.Context(), &in)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	w.WriteHeader(http.StatusCreated)

	if err := jsonapiRuntime.MarshalPayload(w, comment); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (s *Server) updateComment(w http.ResponseWriter, r *http.Request) {
	jsonapiRuntime := jsonapi.NewRuntime().Instrument("comments.update")

	var in comment.Comment
	if err := jsonapiRuntime.UnmarshalPayload(r.Body, &in); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	comment, err := s.store.UpdateComment(r.Context(), &in)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	if err := jsonapiRuntime.MarshalPayload(w, comment); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (s *Server) deleteComment(w http.ResponseWriter, r *http.Request) {
	jsonapi.NewRuntime().Instrument("comments.delete")

	var id uint
	if err := requestVar(r, "id", &id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := s.store.DeleteComment(r.Context(), id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (s *Server) healthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
