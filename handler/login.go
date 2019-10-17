package handler

import "net/http"

func PingHandler(s service.Services) {
	return func(w http.ResponseWriter, r *http.Request) {
		resp := s.PaymentMethodService.GetAllPaymentMethods()
		s.ResponseService.SendJSONResponse(w, r, http.StatusOK, resp)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("{\"success\": \"pong\"}"))
}

