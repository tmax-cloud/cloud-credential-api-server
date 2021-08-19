package util

import (
	"encoding/json"
	"net/http"
)

const (
	QUERY_PARAMETER_STARTTIME   = "startTime"
	QUERY_PARAMETER_ENDTIME     = "endTime"
	QUERY_PARAMETER_SORT        = "sort"
	QUERY_PARAMETER_GRANULARITY = "granularity"
	QUERY_PARAMETER_METRICS     = "metrics"
	QUERY_PARAMETER_DIMENSION   = "dimension"
)

func SetResponse(res http.ResponseWriter, outString string, outJson interface{}, status int) http.ResponseWriter {

	//set Cors
	// res.Header().Set("Access-Control-Allow-Origin", "*")
	res.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	res.Header().Set("Access-Control-Max-Age", "3628800")
	res.Header().Set("Access-Control-Expose-Headers", "Content-Type, X-Requested-With, Accept, Authorization, Referer, User-Agent")

	//set Out
	if outJson != nil {
		res.Header().Set("Content-Type", "application/json")
		js, err := json.Marshal(outJson)
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
		}
		//set StatusCode
		res.WriteHeader(status)
		res.Write(js)
		return res

	} else {
		//set StatusCode
		res.WriteHeader(status)
		res.Write([]byte(outString))
		return res

	}
}
