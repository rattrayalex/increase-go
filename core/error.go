package core

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type RequestError struct {
	Cause    error
	Request  *http.Request
	Response *http.Response
}

func (e RequestError) StatusCode() int {
	if res := e.Response; res != nil {
		return res.StatusCode
	}
	return -1
}

func (e RequestError) URL() string {
	if req := e.Request; req == nil {
		return "<nil>"
	} else if url := req.URL; url == nil {
		return "<nil>"
	} else {
		return url.String()
	}
}

func (e RequestError) Error() string {
	return fmt.Errorf("%s: %d: %w", e.URL(), e.StatusCode(), e.Cause).Error()
}

type APIError interface {
	Status() int
	ErrorBody() interface{}
	Message() string
	Headers() http.Header
	Error() string
}

type GenericAPIError struct {
	request       *CoreRequest
	response      *http.Response
	status        int
	errorBody     interface{}
	errorBodyJSON *string
	message       string
	headers       http.Header
}

func (e GenericAPIError) Status() int {
	return e.status
}

func (e GenericAPIError) ErrorBody() interface{} {
	return e.errorBody
}

func (e GenericAPIError) Message() string {
	return e.message
}

func (e GenericAPIError) Headers() http.Header {
	return e.headers
}

func (e GenericAPIError) errorjSON() string {
	if json := e.errorBodyJSON; json != nil {
		return *json
	}
	if body := e.errorBody; body == nil {
		e.errorBodyJSON = &e.message
		return e.message
	} else if bytes, err := json.MarshalIndent(body, "", "  "); err != nil {
		json := fmt.Sprintf("{\"error\":\"go runtime error while dumping error body: %s\"}", err.Error())
		e.errorBodyJSON = &json
		return json
	} else {
		json := string(bytes)
		e.errorBodyJSON = &json
		return json
	}
}

func (e GenericAPIError) Method() string {
	if coreRequest := e.request; coreRequest == nil {
		return "<nil>"
	} else if request := coreRequest.Request; request == nil {
		return "<nil>"
	} else {
		return request.Method
	}
}

func (e GenericAPIError) URL() string {
	if coreReq := e.request; coreReq == nil {
		return "<nil>"
	} else if req := coreReq.Request; req == nil {
		return "<nil>"
	} else if url := req.URL; url == nil {
		return "<nil>"
	} else {
		return url.String()
	}
}

func (e GenericAPIError) Error() string {
	return fmt.Sprintf("%s %s: %d -- %s", e.Method(), e.URL(), e.status, e.errorjSON())
}

type BadRequestError struct{ GenericAPIError }
type AuthenticationError struct{ GenericAPIError }
type PermissionDeniedError struct{ GenericAPIError }
type NotFoundError struct{ GenericAPIError }
type ConflictError struct{ GenericAPIError }
type UnprocessableEntityError struct{ GenericAPIError }
type RateLimitError struct{ GenericAPIError }
type InternalServerError struct{ GenericAPIError }

func NewAPIError(req *CoreRequest, res *http.Response, status int, error interface{}, message string, headers http.Header) APIError {
	err := GenericAPIError{req, res, status, error, nil, message, headers}

	if status == 400 {
		return BadRequestError{err}
	}
	if status == 401 {
		return AuthenticationError{err}
	}
	if status == 403 {
		return PermissionDeniedError{err}
	}
	if status == 404 {
		return NotFoundError{err}
	}
	if status == 409 {
		return ConflictError{err}
	}
	if status == 422 {
		return UnprocessableEntityError{err}
	}
	if status == 429 {
		return RateLimitError{err}
	}
	if status > 500 {
		return InternalServerError{err}
	}

	return err
}

func extractResponseText(res *http.Response) (string, error) {
	buf := new(strings.Builder)
	if _, err := io.Copy(buf, res.Body); err != nil {
		return "", err
	}
	return buf.String(), nil
}

func NewAPIErrorFromResponse(req *CoreRequest, res *http.Response) APIError {
	var errorJSON interface{}
	var message string

	if text, err := extractResponseText(res); err == nil {
		jsonDecoder := json.NewDecoder(strings.NewReader(text))
		_ = jsonDecoder.Decode(&errorJSON)
		if errorJSON != nil {
			message = ""
		} else {
			message = text
		}
	}

	return NewAPIError(req, res, res.StatusCode, errorJSON, message, res.Header)
}
