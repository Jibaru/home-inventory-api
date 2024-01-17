package responses

type DataResponse struct {
	Data interface{} `json:"data"`
}

func NewDataResponse(data interface{}) *DataResponse {
	return &DataResponse{
		data,
	}
}
