package typecastGo

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"bytes"
	"time"
	"fmt"
	"os"
)

const (
	GOOGLE_SECURE_PASSWORD = "https://www.googleapis.com/identitytoolkit/v3/relyingparty/verifyPassword?key=AIzaSyA7rq_sg-scoKIMWPbq7MJ-3kGu7W1Uouw"
	GOOGLE_SECURE_TOKEN = "https://www.googleapis.com/identitytoolkit/v3/relyingparty/verifyCustomToken?key=AIzaSyA7rq_sg-scoKIMWPbq7MJ-3kGu7W1Uouw"
	TYPECAST_OAUTH = "https://typecast.ai/api/auth-fb/custom-token"
	TYPECAST_SPEAK_POST = "https://typecast.ai/api/speak/batch/post"
	TYPECAST_SPEAK_GET = "https://typecast.ai/api/speak/batch/get"

	TYPECAST_GID = "f8ON1ZpeF0mDFjZTQlr9G"
)

var Actor map[string]string

func init() {
	Actor = map[string]string{
		"찬구": "5c547544fcfee90007fed455",
		"주하": "5d01a52563338e00072b8c9c",
		"카밀라": "5f8d7b0de146f10007b8042f",
		"우주": "5f8e95eae146f10007b85f45",
		"혜정": "5f8e9644f6120100074475f0",
		"한유격": "5faa3acfac283a00075d0d2e",
		"영희": "5c3c52caea9791000747155e",
		"소영": "5c789c317ad86500073a02cc",
		"지영": "5c789c34dabcfa0008b0a390",
		"경완": "5d01a52163338e00072b8c9b",
		"성욱": "5d01a52b63338e00072b8c9e",
		"애란": "5d01a52c63338e00072b8c9f",
		"줄리아": "5e11db231c5eaa000b7d4c9c",
		"민지": "5eb55ce920d1b60016e91de6",
		"신혜": "5ecbbc4b0fbab10007bb3c7d",
		"준상": "5ecbbc7399979700087711db",
		"현경": "5ecbbc5599979700087711d6",
		"성규": "5ecbbc6099979700087711d8",
		"윤성": "5ecbbc696b7c780008be50c4",
		"관용": "5ebea235433bd200073e0e15",
		"정순": "5ebea242728f5b00075e6211",
		"다희": "5ebea251fcf5110007b77d0f",
		"경숙": "5ebea266728f5b00075e6215",
		"동우": "5ebea0fefcf5110007b77cc4",
		"연길": "5ebea13564afaf00087fc2e7",
		"지우": "5ebea14e433bd200073e0dce",
		"나진": "5ebea185728f5b00075e61e3",
		"쥬비": "5ebea194728f5b00075e61e8",
		"정희": "5ebea1a364afaf00087fc2fb",
		"연우": "5eb55cf1f0b0a700071f89c7",
		"용호": "5ea7b4268727390007ba95fc",
		"재훈": "5c789c32dabcfa0008b0a38d",
		"정섭": "5c789c32dabcfa0008b0a38e",
		"상도": "5c789c337ad86500073a02ce",
		"진혁": "5c3c52ca5827e00008dd7f36",
		"민상": "5c3c52ca5827e00008dd7f38",
		"찬혁": "5c547545fcfee90007fed458",
		"지철": "5c789c33dabcfa0008b0a38f",
		"수진": "5c3c52ca5827e00008dd7f3a",
		"국희": "5c789c31dabcfa0008b0a38c",
		"제임스":"5c547545168d0a0007f5cf13",
		"선영": "5c789c317ad86500073a02cb",
		"금희": "5c789c337ad86500073a02cd",
		"명희": "5c789c347ad86500073a02cf",
		"미오": "5c3c52caea9791000747155c",
		"보노": "5c547544fcfee90007fed454",
		"덕춘": "5c3c52c95827e00008dd7f34",
		"강수정": "5e4f7a5da82e1f000aca31af",
		"이승주": "5e4f7a5d716fd30009920721",
		"주원": "5c547545fcfee90007fed459",
		"의찬": "5c547545fcfee90007fed456",
		"마크": "5e11db23b70e890009fb7ae7",
		"트럼프": "5c3c52c9ea9791000747155b",
		"김정은": "5c3c52ca5827e00008dd7f35",
		"문재인":"5cf9ecb62cae6a0008b446af",
		"신디": "5c547545168d0a0007f5cf14",
		"현진": "5d01a522bb04140008a23f34",
		"정원": "5d01a529bb04140008a23f36",
		"성배": "5d01a52bbb04140008a23f37",
		"리춘희": "5d357d34ba9d2d0007195f9e",
		"하은": "5de781b108fd6b000882c58f",
		"범수": "5c3c52cb5827e00008dd7f3b",
	}
}

type TypeCast struct {
	Email string
	Password string
	ActorName string
	Speed int
	Naturalness float64
	MaxSeconds int
	Token OauthToken
}

type OauthToken struct {
	AccessToken string
	RefreshToken string
	ExpiresIn time.Time
}

/**
 * email string
 * password string
 */
func NewTypeCast(args ...string) (this TypeCast, err error) {
	this = TypeCast{
		Speed: 1,
		Naturalness: 0.8,
		MaxSeconds: 30,
	}
	if len(args) > 2 {
		err = fmt.Errorf("too many arguments")
		return
	}

	if len(args) > 0 && args[0] != "" {
		this.Email = args[0]
	} else if os.Getenv("TYPECAST_EMAIL") != "" {
		this.Email = os.Getenv("TYPECAST_EMAIL")
	} else {
		err = fmt.Errorf("email이 설정 되어 있지않습니다. 첫번째 파라미터를 입력하거나 환경변수 TYPECAST_EMAIL 를 설정하세요.")
		return
	}

	if len(args) > 1 && args[1] != "" {
		this.Password = args[1]
	} else if os.Getenv("TYPECAST_PASSWORD") != "" {
		this.Password = os.Getenv("TYPECAST_PASSWORD")
	} else {
		err = fmt.Errorf("password가 설정 되어 있지않습니다. 첫번째 파라미터를 입력하거나 환경변수 TYPECAST_PASSWORD 를 설정하세요.")
		return
	}
	err = this.TokenInit()

	if err != nil {
		return
	}

	return
}

/**
 * url string
 * setTokenHeader bool
 */
func (this *TypeCast) Get(url string, args ...bool) (responseJson map[string]interface{}, err error) {
	request, _ := http.NewRequest("GET", url, nil)
	if len(args) == 0 || args[0] {
		if this.TokenExpired() {
			this.TokenRefresh()
		}
		this.setTokenHeader(request)
	}
	client := &http.Client{}
	response, err := client.Do(request)
	jsonString, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	if err != nil {
		return
	}
	json.Unmarshal(jsonString, &responseJson)

	return
}

/**
 * url string
 * requestJson map[string]interface{}
 * setTokenHeader bool
 */
func (this *TypeCast) Post(url string, requestJson interface{}, args ...bool) (responseJson map[string]interface{}, err error) {
	jsonByte, err := json.Marshal(requestJson)
	jsonBuffer := bytes.NewBuffer(jsonByte)
	if err != nil {
		return
	}
	request, _ := http.NewRequest("POST", url, jsonBuffer)

	if len(args) == 0 || args[0] {
		if this.TokenExpired() {
			this.TokenRefresh()
		}
		this.setTokenHeader(request)
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return
	}
	jsonString, err := ioutil.ReadAll(response.Body)
	// fmt.Printf("%s \n\n", string(jsonString))
	defer response.Body.Close()
	if err != nil {
		return
	}
	json.Unmarshal(jsonString, &responseJson)

	return
}

/**
 * text string
 * actorName string
 */
func (this *TypeCast) Exec(text string, args ...string) (blob []byte, err error) {
	var ok bool
	var ActorId string
	ActorName := this.ActorName
	if len(args) > 0 {
		ActorName = args[0]
	}
	if ActorId, ok = Actor[ActorName]; !ok {
		err = fmt.Errorf("undefind Actor")
		return
	}
	response, err := this.Post(TYPECAST_SPEAK_POST, []interface{}{
		map[string]interface{}{
			"actor_id": ActorId,
			"text": text,
			"lang": "auto",
			"max_seconds": this.MaxSeconds,
			"naturalness": this.Naturalness,
			"speed_x": this.Speed,
			"gid": TYPECAST_GID,
			"style_idx": 0,
		},
	})
	if err != nil {
		return
	}
	var speakUrls interface{}
	if speakUrls, ok = response["result"].(map[string]interface{})["speak_urls"]; !ok {
		err = fmt.Errorf("undefind speark_urls")
		return
	}

	var audioUrl string
	getTryCnt := 0
	for {
		getTryCnt += 1
		time.Sleep(time.Millisecond * 100)

		response, err = this.Post(TYPECAST_SPEAK_GET, speakUrls)

		var status string
		if status, ok = response["result"].([]interface{})[0].(map[string]interface{})["status"].(string); !ok {
			fmt.Errorf("undefind status")
			return
		}
		if status == "done" {
			if audioUrl, ok = response["result"].([]interface{})[0].(map[string]interface{})["audio"].(map[string]interface{})["url"].(string); !ok {
				err = fmt.Errorf("undefind audio url")
				return
			}
			break
		}
		if getTryCnt > 10 {
			err = fmt.Errorf("timeout")
			return
		}
	}

	response, err = this.Get(audioUrl + "/no-redirect")
	if err != nil {
		return
	}

	if audioUrl, ok = response["result"].(string); !ok {
		err = fmt.Errorf("undefind result")
		return
	}

	client := &http.Client{}
	audioResponse, err := client.Get(audioUrl)
	if err != nil {
		return
	}
	defer audioResponse.Body.Close()
	blob, err = ioutil.ReadAll(audioResponse.Body)
	if err != nil {
		return
	}
	return
}

func (this *TypeCast) TokenExpired() bool {
	return time.Now().After(this.Token.ExpiresIn)
}

func (this *TypeCast) TokenRefresh() error {
	return this.TokenInit()
}

func (this *TypeCast) setTokenHeader(req *http.Request) {
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", this.Token.AccessToken))
}

func (this *TypeCast) TokenInit() (err error) {
	//
	response, err := this.Post(GOOGLE_SECURE_PASSWORD, map[string]interface{}{
		"email": this.Email,
		"password": this.Password,
		"returnSecureToken": true,
	}, false)

	if err != nil {
		return
	}
	var idToken string
	var ok bool
	if idToken, ok = response["idToken"].(string); !ok {
		err = fmt.Errorf("token encode 실패")
		return
	}

	//
	response, err = this.Post(TYPECAST_OAUTH, map[string]interface{}{
		"token": idToken,
	}, false)

	if err != nil {
		return
	}

	var accessToken string
	if accessToken, ok = response["result"].(map[string]interface{})["access_token"].(string); !ok {
		err = fmt.Errorf("typecast custom token 실패")
		return
	}

	//
	response, err = this.Post(GOOGLE_SECURE_TOKEN, map[string]interface{}{
		"token": accessToken,
		"returnSecureToken": true,
	}, false)

	if err != nil {
		return
	}

	if accessToken, ok = response["idToken"].(string); !ok {
		err = fmt.Errorf("fail typecast accessToken 1")
		return
	}
	var refreshToken string
	if refreshToken, ok = response["refreshToken"].(string); !ok {
		err = fmt.Errorf("fail typecast accessToken 2")
		return
	}
	var expiresInString string
	if expiresInString, ok = response["expiresIn"].(string); !ok {
		err = fmt.Errorf("fail typecast accessToken 3")
		return
	}
	var expiresIn int
	if expiresIn, err = strconv.Atoi(expiresInString); err != nil {
		err = fmt.Errorf("fail typecast accessToken 4")
		return
	}

	this.Token.AccessToken = accessToken
	this.Token.RefreshToken = refreshToken
	this.Token.ExpiresIn = time.Now().Add(time.Second * time.Duration(expiresIn))

	return
}
