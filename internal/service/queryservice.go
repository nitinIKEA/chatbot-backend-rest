package service

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

// QueryOrderstatusManordrefGet - get data from database based on manual order reference
func (s *Service) QueryOrderstatusManordrefGet(w http.ResponseWriter, r *http.Request) {

	vals, err := getParams(r)
	if err != nil {
		log.Printf("error from getParams in QueryOrderstatusManordrefGet: %v", err)
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	dt, err := s.DBConns.QueryOrderstatusManordref(vals["ord_id_ref_req"].(string), vals["environment"].(string))
	if err != nil {
		log.Printf("error from QueryOrderstatusManordref in QueryOrderstatusManordrefGet: %v", err)
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(dt)
}

// QueryOrderstatusOrdidrefGet - get data from database based on original order id ref
func (s *Service) QueryOrderstatusOrdidrefGet(w http.ResponseWriter, r *http.Request) {

	vals, err := getParams(r)
	if err != nil {
		log.Printf("error from getParams in QueryOrderstatusOrdidrefGet: %v", err)
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	dt, err := s.DBConns.QueryOrderstatusOrdidref(vals["org_ord_ref"].(string), vals["environment"].(string))
	if err != nil {
		log.Printf("error from QueryOrderstatusOrdidref in QueryOrderstatusOrdidrefGet: %v", err)
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(dt)
}

// QueryOrderstatusOrdrefGet - get data from databasebased on sales order reference
func (s *Service) QueryOrderstatusOrdrefGet(w http.ResponseWriter, r *http.Request) {

	vals, err := getParams(r)
	if err != nil {
		log.Printf("error from getParams in QueryOrderstatusOrdrefGet: %v", err)
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	dt, err := s.DBConns.QueryOrderstatusOrdref(vals["ord_id_ref_sales"].(string), vals["environment"].(string))
	if err != nil {
		log.Printf("error from QueryOrderstatusOrdref in QueryOrderstatusOrdrefGet: %v", err)
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(dt)
}

// QueryOrderstatusWrkordrefGet - get data from database based on work order reference
func (s *Service) QueryOrderstatusWrkordrefGet(w http.ResponseWriter, r *http.Request) {

	vals, err := getParams(r)
	if err != nil {
		log.Printf("error from getParams in QueryOrderstatusWrkordrefGet: %v", err)
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	dt, err := s.DBConns.QueryOrderstatusWrkordref(vals["ord_id_ref_wrk"].(string), vals["environment"].(string))
	if err != nil {
		log.Printf("error from QueryOrderstatusWrkordref in QueryOrderstatusWrkordrefGet: %v", err)
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(dt)
}

// QueryPickingroupGet - get picking group store reports
func (s *Service) QueryPickingroupGet(w http.ResponseWriter, r *http.Request) {

	vals, err := getParams(r)
	if err != nil {
		log.Printf("error from getParams in QueryPickingroupGet: %v", err)
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	dt, err := s.DBConns.QueryPickingroup(vals["store_id"].(string), vals["environment"].(string))
	if err != nil {
		log.Printf("error from QueryPickingroup in QueryPickingroupGet: %v", err)
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(dt)
}

// QueryStockavailabilityAvailableGet - get available stock report based on bu_type, bu_code and item_no
func (s *Service) QueryStockavailabilityAvailableGet(w http.ResponseWriter, r *http.Request) {

	vals, err := getParams(r)
	if err != nil {
		log.Printf("error from getParams in QueryStockavailabilityAvailableGet: %v", err)
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	dt, err := s.DBConns.QueryStockavailabilityAvailable(vals["bu_code"].(string), vals["bu_type"].(string), vals["item_no"].(string), vals["environment"].(string))
	if err != nil {
		log.Printf("error from QueryStockavailabilityAvailable in QueryStockavailabilityAvailableGet: %v", err)
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(dt)
}

// QueryStockavailabilityBlockedGet - get blocked stock report based on bu_type, bu_code and item_no
func (s *Service) QueryStockavailabilityBlockedGet(w http.ResponseWriter, r *http.Request) {

	vals, err := getParams(r)
	if err != nil {
		log.Printf("error from getParams in QueryStockavailabilityBlockedGet: %v", err)
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	dt, err := s.DBConns.QueryStockavailabilityBlocked(vals["bu_code"].(string), vals["bu_type"].(string), vals["item_no"].(string), vals["environment"].(string))
	if err != nil {
		log.Printf("error from QueryStockavailabilityBlocked in QueryStockavailabilityBlockedGet: %v", err)
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(dt)
}

func getParams(r *http.Request) (map[string]interface{}, error) {
	vals := make(map[string]interface{})
	bt, err := io.ReadAll(r.Body)
	if err != nil {
		return vals, err
	}
	err = json.Unmarshal(bt, &vals)
	return vals, err
}
