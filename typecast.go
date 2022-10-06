package typecast

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Session struct {
	Token  *GoogleOauth2Response
	Client *http.Client
}

type LoginRequest struct {
	Email             string `json:"email"`
	Password          string `json:"password"`
	ReturnSecureToken bool   `json:"returnSecureToken"`
}

type TypecastOauth2Request struct {
	Token string `json:"token"`
}

type GoogleOauth2Request struct {
	Token             string `json:"token"`
	ReturnSecureToken bool   `json:"returnSecureToken"`
}

type GoogleLoginResponse struct {
	Kind           string `json:"kind"`
	LocalId        string `json:"localId"`
	Email          string `json:"email"`
	DisplayName    string `json:"displayName"`
	IdToken        string `json:"idToken"`
	Registered     bool   `json:"registered"`
	ProfilePicture string `json:"profilePicture"`
	RefreshToken   string `json:"refreshToken"`
	ExpiresIn      int    `json:"expiresIn"`
}

type TypecastOauth2Response struct {
	Result *TypecastOauth2Result `json:"result"`
}

type TypecastOauth2Result struct {
	Status        string  `json:"status"`
	AccessToken   string  `json:"access_token"`
	GuardianEmail *string `json:"guardian_email"`
}

type GoogleOauth2Response struct {
	Kind         string `json:"kind"`
	IdToken      string `json:"idToken"`
	RefreshToken string `json:"refreshToken"`
	ExpiresIn    int    `json:"expiresIn"`
	IsNewUser    bool   `json:"isNewUser"`
}

type TypecastExecuteRequest struct {
	ActorId     string  `json:"actor_id"`
	Text        string  `json:"text"`
	Lang        string  `json:"lang"`
	MaxSeconds  int     `json:"max_seconds"`
	Naturalness float64 `json:"naturalness"`
	SpeedX      int     `json:"speed_x"`
	Gid         string  `json:"gid"`
	StyleIdx    int     `json:"style_idx"`
}

func NewSession() (s Session) {
	s.Client = &http.Client{}
	return
}

func NewRequest() (request *TypecastExecuteRequest) {
	request.Gid = TypecastGid
	request.StyleIdx = 0
	request.SpeedX = 1
	request.Naturalness = 0.8
	request.MaxSeconds = 30
	return
}

func (s *Session) Request(method string, endpoint string, data interface{}, headers map[string]string) (responseBody []byte, err error) {
	requestBody, err := json.Marshal(data)
	if err != nil {
		return
	}

	request, err := http.NewRequest(
		method,
		endpoint,
		bytes.NewBuffer(requestBody),
	)
	if err != nil {
		return
	}

	for headerKey, headerContent := range headers {
		request.Header.Add(headerKey, headerContent)
	}

	response, err := s.Client.Do(request)
	if err != nil {
		return
	}
	defer response.Body.Close()

	responseBody, err = ioutil.ReadAll(response.Body)
	return
}

func (s *Session) RequestWithToken(method string, endpoint string, data interface{}) ([]byte, error) {
	return s.Request(
		method,
		endpoint,
		data,
		map[string]string{
			"Content-Type":  "application/json",
			"Authorization": fmt.Sprintf("Bearer %s", s.Token.IdToken),
		},
	)
}

func (s *Session) googleLogin(loginRequest *LoginRequest) (oauth2Token *GoogleLoginResponse, err error) {
	responseBody, err := s.Request(
		"POST",
		EndpointGoogleSecurePassword,
		loginRequest,
		map[string]string{},
	)
	if err != nil {
		return
	}

	json.Unmarshal(responseBody, &oauth2Token)
	if err != nil {
		err = fmt.Errorf(string(responseBody))
		return
	}

	return
}

func (s *Session) typecastOauth2(oauth2Token *GoogleLoginResponse) (typecastOauth2Response *TypecastOauth2Response, err error) {
	responseBody, err := s.Request(
		"POST",
		EndpointTypecastOauth,
		&TypecastOauth2Request{
			Token: oauth2Token.IdToken,
		},
		map[string]string{},
	)
	if err != nil {
		return
	}

	json.Unmarshal(responseBody, &typecastOauth2Response)
	if err != nil {
		err = fmt.Errorf("%#v", responseBody)
		return
	}

	return
}

func (s *Session) googleOauth2(typecastOauth2Response *TypecastOauth2Response) (googleOauth2Response *GoogleOauth2Response, err error) {
	responseBody, err := s.Request(
		"POST",
		EndpointGoogleSecureToken,
		&GoogleOauth2Request{
			Token:             typecastOauth2Response.Result.AccessToken,
			ReturnSecureToken: true,
		},
		map[string]string{},
	)
	if err != nil {
		return
	}

	json.Unmarshal(responseBody, &googleOauth2Response)
	if err != nil {
		err = fmt.Errorf("%#v", responseBody)
		return
	}

	return
}

func (s *Session) Connect(loginRequest *LoginRequest) (err error) {
	googleOauth2Token, err := s.googleLogin(loginRequest)
	if err != nil {
		return
	}

	typecastOauth2Response, err := s.typecastOauth2(googleOauth2Token)
	if err != nil {
		return
	}

	s.Token, err = s.googleOauth2(typecastOauth2Response)

	return
}

func (s *Session) Do(request *TypecastExecuteRequest) (err error) {

}
