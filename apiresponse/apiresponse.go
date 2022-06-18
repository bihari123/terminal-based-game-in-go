package apiresponse

type APIResponse struct{
  Code int `json:"status"`
  Status bool 
  Data interface{} `json:"data"`
  Message string `json:"message"`
  Error APIError `json:"error"`
}

type APIError struct {
  Code int `json:"code"`
  Message string `json:"message"`
}

func Error(errCode int , message string)APIResponse{
  return APIResponse{
    Code: errCode,
    Status: false,
    Message: "",
    Data:nil,
    Error: APIError{
      Code: errCode, 
      Message: message,
    },
  }
}

func Success(successCode int, message string,data interface{})APIResponse{
  return APIResponse{
    Code: successCode,
    Status:true,
    Message:message,
    Data:data,
    Error:APIError{},
    
  }
}
