package requestHandler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"

	"github.com/labstack/echo"
)

// TokenType type of token
type TokenType int8

// const of token type
const (
	Basic = iota
	Bearer
	Other
)

func (s TokenType) String() string {
	switch s {
	case Basic:
		return "Basic"
	case Bearer:
		return "Bearer"
	case Other:
		return "Other"
	default:
		return "Other"
	}
}

// RequestModel model
type RequestModel struct {
	URL         string
	TokenType   TokenType
	Token       string
	Username    string
	Password    string
	Body        string
	Type        string
	ContentType string
}

// SimpleResponse response
type SimpleResponse struct {
	Status  int
	Message string
	Body    interface{}
}

// Base request
func Base(requestModel RequestModel, result *interface{}) error {

	body := bytes.NewBufferString(requestModel.Body)

	req, err := http.NewRequest(requestModel.Type, requestModel.URL, body)
	switch requestModel.TokenType {
	case Basic:
		req.SetBasicAuth(requestModel.Username, requestModel.Password)
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	case Bearer:
		header := fmt.Sprintf("Bearer %s", requestModel.Token)
		req.Header.Set("Authorization", header)
		req.Header.Add("Content-Type", "application/json")
	default:
		if requestModel.ContentType == "" {
			requestModel.ContentType = "application/json"
		}
		req.Header.Set("Authorization", requestModel.Token)
		req.Header.Add("Content-Type", requestModel.ContentType)
		break
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Accept-Encoding", "identity")
	req.Header.Add("x-xss-protection", "1; mode=block")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if err = json.NewDecoder(resp.Body).Decode(result); err != nil {
		// fmt.Printf("%s-URL:%s-ERROR:%v\n", requestModel.Type, requestModel.URL, err)
		return nil
	}
	if resp.StatusCode >= http.StatusBadRequest {
		message := fmt.Sprintf("%s-URL:%s-ERROR CODE:%v\n", requestModel.Type, requestModel.URL, resp.StatusCode)
		if result != nil && !reflect.ValueOf(result).IsNil() {
			switch (*result).(type) {
			case string:
				message = fmt.Sprintf("%s %v", message, *result)
			default:
				temp, ok := (*result).(map[string]interface{})
				fmt.Printf("res    %+v\n", temp)
				if ok && temp != nil && temp["data"] != nil {
					data := (temp["data"]).(map[string]interface{})
					if data != nil {
						msg := data["message"]
						if msg != nil {
							msgStr := msg.(string)

							if msgStr != "" {
								message = msgStr
							}
						}
					}
				}
			}
		}
		return echo.NewHTTPError(resp.StatusCode, message)
	}

	return nil
}

// Get request
func Get(requestModel RequestModel, result *interface{}) error {
	requestModel.Type = "GET"
	return Base(requestModel, result)
}

// Post request
func Post(requestModel RequestModel, result *interface{}) error {
	requestModel.Type = "POST"
	return Base(requestModel, result)
}

// Put request
func Put(requestModel RequestModel, result *interface{}) error {
	requestModel.Type = "PUT"
	return Base(requestModel, result)
}

// Patch request
func Patch(requestModel RequestModel, result *interface{}) error {
	requestModel.Type = "PATCH"
	return Base(requestModel, result)
}

// Delete request
func Delete(requestModel RequestModel, result *interface{}) error {
	requestModel.Type = "DELETE"
	return Base(requestModel, result)
}

// GetV2 request
func GetV2(requestModel RequestModel, result interface{}) error {
	requestModel.Type = "GET"
	return Base(requestModel, &result)
}

// PostV2 request
func PostV2(requestModel RequestModel, result interface{}) error {
	requestModel.Type = "POST"
	return Base(requestModel, &result)
}

// PutV2 request
func PutV2(requestModel RequestModel, result interface{}) error {
	requestModel.Type = "PUT"
	return Base(requestModel, &result)
}

// PatchV2 request
func PatchV2(requestModel RequestModel, result interface{}) error {
	requestModel.Type = "PATCH"
	return Base(requestModel, &result)
}

// DeleteV2 request
func DeleteV2(requestModel RequestModel, result interface{}) error {
	requestModel.Type = "DELETE"
	return Base(requestModel, &result)
}
