package webRequesthandlers

import "net/http"

type TeacherRequestHandler struct{}

func (teacherRequestHandler *TeacherRequestHandler) serveHTTP(w http.ResponseWriter, r *http.Request) {
}