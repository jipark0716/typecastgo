package typecast

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

const STAUTS_DONE = "done"

type Session struct {
	Token  *GoogleOauth2Response
	Client *http.Client
}

type LoginRequest struct {
	Email             string `json:"email"`
	Password          string `json:"password"`
	ReturnSecureToken bool   `json:"returnSecureToken"`
}

type TypecastActorResponse struct {
	Result []TypecastActor `jsno:"result"`
}

type Dist struct {
	CardImageUrl   string `json:"card_image_url"`
	OriginImageUrl string `json:"origin_image_url"`
	ResourceUrl    string `json:"resource_url"`
}

type Language struct {
	Ko string `json:"ko"`
	En string `json:"en"`
}

type TypecastActor struct {
	ActorId             string             `json:"actor_id"`
	ActorUrl            string             `json:"actor_url"`
	AfxName             *string            `json:"afx_name"`
	Age                 string             `json:"age"`
	AudioLengthFactor   interface{}        `json:"audio_length_factor"`
	AudioQuality        string             `json:"audio_quality"`
	AudioUrl            string             `json:"audio_url"`
	Bookmark            bool               `json:"bookmark"`
	Dist                Dist               `json:"dist"`
	EnhanceSpeaker      string             `json:"enhance_speaker"`
	Flags               []string           `json:"flags"`
	Hidden              bool               `json:"hidden"`
	Ice                 int                `json:"ice"`
	ImgUrl              string             `json:"img_url"`
	Language            string             `json:"language"`
	Name                Language           `json:"name"`
	NumSpeeds           int                `json:"num_speeds"`
	NumStyles           int                `json:"num_styles"`
	Order               int                `json:"order"`
	Price               string             `json:"price"`
	Recent              bool               `json:"recent"`
	RefActorId          *string            `json:"ref_actor_id"`
	Rid                 string             `json:"rid"`
	Roles               []string           `json:"roles"`
	Score               int                `json:"score"`
	Sex                 []string           `json:"sex"`
	SpecPow             *string            `json:"spec_pow"`
	SpeedParams         map[string]float32 `json:"speed_params"`
	StyleLabel          interface{}        `json:"style_label"`
	StyleLabelInfo      interface{}        `json:"style_label_info"`
	StyleLabelV2        interface{}        `json:"style_label_v2"`
	StyleLabelV2Info    interface{}        `json:"style_label_v2_info"`
	StyleRecommendation interface{}        `json:"style_recommendation"`
	StyleSource         string             `json:"style_source"`
	Tag                 []string           `json:"tag"`
	TagPrimary          interface{}        `json:"tag_primary"`
	TagV2               interface{}        `json:"tag_v2"`
	Tuning              []string           `json:"tuning"`
	UniqueId            string             `json:"unique_id"`
	Ssage               []string           `json:"usage"`
	Version             string             `json:"version"`
	VideoUrl            *string            `json:"video_url"`
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
	ExpiresIn      string `json:"expiresIn"`
}

type TypecastOauth2Response struct {
	Result *TypecastOauth2Result `json:"result"`
}

type TypecastOauth2Result struct {
	Status        string  `json:"status"`
	AccessToken   string  `json:"access_token"`
	GuardianEmail *string `json:"guardian_email"`
}

type TypecastSpeakResponse struct {
	Result string `json:"result`
}

type GoogleOauth2Response struct {
	Kind         string `json:"kind"`
	IdToken      string `json:"idToken"`
	RefreshToken string `json:"refreshToken"`
	ExpiresIn    string `json:"expiresIn"`
	ExpiresAt    time.Time
	IsNewUser    bool `json:"isNewUser"`
}

type TypecastExecuteRequest struct {
	ActorId           string  `json:"actor_id"`
	Text              string  `json:"text"`
	Lang              string  `json:"lang"`
	MaxSeconds        int     `json:"max_seconds"`
	Naturalness       float64 `json:"naturalness"`
	SpeedX            int     `json:"speed_x"`
	Gid               string  `json:"gid"`
	StyleIdx          int     `json:"style_idx"`
	LastPitch         *string `json:"last_pitch"`
	Mode              string  `json:"mode"`
	Pitch             int     `json:"pitch"`
	StyleLabel        string  `json:"style_label"`
	StyleLabelVersion string  `json:"style_label_version"`
	Tempo             int     `json:"tempo"`
}

type TypecastExecuteQuery struct {
	Ttext             string  `json:"text"`
	ActorId           string  `json:"actor_id"`
	Lang              string  `json:"lang"`
	MaxSeconds        int     `json:"max_seconds"`
	Actor             *string `json:"actor"`
	Namespace         string  `json:"namespace"`
	SpeakerDb         string  `json:"speaker_db"`
	SpeakerId         string  `json:"speaker_id"`
	TacoSvc           *string `json:"taco_svc"`
	VocoSvc           *string `json:"voco_svc"`
	Version           string  `json:"version"`
	AudioQuality      string  `json:"audio_quality"`
	Platform          *string `json:"platform"`
	EnableLoudness    bool    `json:"enable_loudness"`
	EnableAfterEffect bool    `json:"enable_after_effect"`
	SynthesisType     string  `json:"synthesis_type"`
	Title             bool    `json:"title"`
	PlayId            bool    `json:"play_id"`
	UseWavernn        bool    `json:"use_wavernn"`
	UseWghd           bool    `json:"use_wghd"`
	StyleIdx          int     `json:"style_idx"`
	StyleLabel        string  `json:"style_label"`
	StyleLabelVersion string  `json:"style_label_version"`
}

type TypecastVoiceResponse struct {
	Result []TypecastVoiceResult `json:"result"`
}

type TypecastVoiceResult struct {
	Id              *TypecastId            `json:"_id"`
	ActorId         *TypecastId            `json:"actor_id"`
	Uid             string                 `json:"uid"`
	Query           *TypecastExecuteQuery  `json:"query"`
	Status          string                 `json:"status"`
	TaskId          string                 `json:"task_id"`
	SpeakUrl        string                 `json:"speak_url"`
	Audio           *TypecastAudio         `json:"audio"`
	AudioPath       *string                `json:"audio_path"`
	Quality         *string                `json:"quality"`
	SentenceTaskIds *[]string              `json:"sentence_task_ids"`
	Callback        *TypecastVoiceCallback `json:"callback"`
	Download        *string                `json:"download"`
	TextCount       int                    `json:"text_count"`
	Duration        *float32               `json:"duration"`
}

type TypecastExtension struct {
	Url       *string `json:"url"`
	Extension *string `json:"extension"`
}

type TypecastAudio struct {
	TypecastExtension
	High *TypecastExtension `json:"high"`
	Hd1  *TypecastExtension `json:"hd1"`
	Low  *TypecastExtension `json:"low"`
}

type TypecastVoiceCallback struct {
	Status   string      `json:"status"`
	PlayId   *TypecastId `json:"play_id"`
	ErrorMsg *string     `json:"error_msg"`
}

type TypecastId struct {
	Oid         *string `json:"$oid"`
	CreatedDate *int    `json:"created_date"`
}

type TypecastExecuteResponse struct {
	Result *TypecastExecuteResult `json:"result"`
}

type TypecastExecuteResult struct {
	SpeakUrls []string `json:"speak_urls"`
}

func NewSession() (s Session) {
	s.Client = &http.Client{}
	return
}

func NewRequest() (request *TypecastExecuteRequest) {
	request = new(TypecastExecuteRequest)
	request.Gid = TypecastGid
	request.Lang = "auto"
	request.StyleIdx = 0
	request.SpeedX = 1
	request.Naturalness = 0.8
	request.MaxSeconds = 30
	request.Pitch = 0
	request.StyleLabel = "0"
	request.StyleLabelVersion = "v2"
	request.Tempo = 1
	return
}

func (s *Session) Request(method string, endpoint string, data interface{}, headers map[string]string) (responseBody []byte, err error) {
	var body io.Reader

	if data != nil {
		var requestBody []byte
		requestBody, err = json.Marshal(data)
		if err != nil {
			return
		}
		body = bytes.NewBuffer(requestBody)
	}

	request, err := http.NewRequest(
		method,
		endpoint,
		body,
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

func (s *Session) IsTokenExpired() bool {
	return time.Now().After(s.Token.ExpiresAt)
}

func (s *Session) TokenRefresh() {
	// @todo init 처음부터 하지 않고 Refresh Token 사용해서 토큰 갱신하기
}

func (s *Session) TokenHeader() map[string]string {
	if !s.IsTokenExpired() {
		s.TokenRefresh()
	}

	return map[string]string{
		"Content-Type":  "application/json",
		"Authorization": fmt.Sprintf("Bearer %s", s.Token.IdToken),
	}
}

func (s *Session) RequestWithToken(method string, endpoint string, data interface{}) ([]byte, error) {
	return s.Request(
		method,
		endpoint,
		data,
		s.TokenHeader(),
	)
}

func (s *Session) RequestJson(method string, endpoint string, data interface{}, headers map[string]string, response interface{}) (err error) {
	responseBody, err := s.Request(
		method,
		endpoint,
		data,
		headers,
	)
	if err != nil {
		return
	}

	err = json.Unmarshal(responseBody, &response)
	if err != nil {
		responseString := string(responseBody)
		if len(responseString) > 200 {
			responseString = responseString[:200]
		}
		err = fmt.Errorf("%+v\n%#v", err, responseString)
		return err
	}

	return err
}

func (s *Session) RequestJsonWithToken(method string, endpoint string, data interface{}, response interface{}) error {
	return s.RequestJson(
		method,
		endpoint,
		data,
		s.TokenHeader(),
		response,
	)
}

func (s *Session) googleLogin(loginRequest *LoginRequest) (oauth2Token *GoogleLoginResponse, err error) {
	err = s.RequestJson(
		"POST",
		EndpointGoogleSecurePassword,
		loginRequest,
		nil,
		&oauth2Token,
	)
	return
}

func (s *Session) typecastOauth2(oauth2Token *GoogleLoginResponse) (typecastOauth2Response *TypecastOauth2Response, err error) {
	err = s.RequestJson(
		"POST",
		EndpointTypecastOauth,
		&TypecastOauth2Request{
			Token: oauth2Token.IdToken,
		},
		nil,
		&typecastOauth2Response,
	)
	return
}

func (s *Session) googleOauth2(typecastOauth2Response *TypecastOauth2Response) (googleOauth2Response *GoogleOauth2Response, err error) {
	err = s.RequestJson(
		"POST",
		EndpointGoogleSecureToken,
		&GoogleOauth2Request{
			Token:             typecastOauth2Response.Result.AccessToken,
			ReturnSecureToken: true,
		},
		nil,
		&googleOauth2Response,
	)

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
	if err != nil {
		return
	}

	expiresIn, err := strconv.Atoi(s.Token.ExpiresIn)
	if err != nil {
		return
	}

	s.Token.ExpiresAt = time.Now().Add(time.Second * time.Duration(expiresIn))

	return
}

func (s *Session) getAudioUrl(speakUrls []string) (audioUrl string, err error) {
	for tryCount := 0; tryCount < 10; tryCount++ {
		time.Sleep(time.Microsecond * 100)

		var typecastVoiceResponse *TypecastVoiceResponse
		err = s.RequestJsonWithToken(
			"POST",
			EndpointTypecastSpeakGet,
			speakUrls,
			&typecastVoiceResponse,
		)
		if err != nil {
			return
		}

		result := typecastVoiceResponse.Result[0]

		if result.Status == STAUTS_DONE {
			audioUrl = EndpointTypecastNoredirect(*(*result.Audio).Url)
			return
		}
	}

	err = fmt.Errorf("wait timeout")
	return
}

func (s *Session) Do(request []*TypecastExecuteRequest) (audio []byte, err error) {
	var typecastExecuteResponse *TypecastExecuteResponse
	err = s.RequestJsonWithToken(
		"POST",
		EndpointTypecastSpeakPost,
		request,
		&typecastExecuteResponse,
	)
	if err != nil {
		return
	}

	audioUrl, err := s.getAudioUrl(typecastExecuteResponse.Result.SpeakUrls)
	if err != nil {
		return
	}

	var typecastSpeakResponse *TypecastSpeakResponse
	err = s.RequestJsonWithToken(
		"GET",
		audioUrl,
		nil,
		&typecastSpeakResponse,
	)
	if err != nil {
		return
	}

	audio, err = s.Request(
		"GET",
		typecastSpeakResponse.Result,
		nil,
		nil,
	)

	return
}

func (s *Session) GetActors() (actors []TypecastActor, err error) {
	var typecastActorResponse *TypecastActorResponse
	err = s.RequestJsonWithToken(
		"GET",
		EndpointTypecast,
		nil,
		&typecastActorResponse,
	)

	actors = typecastActorResponse.Result
	return
}
