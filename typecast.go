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
		"박찬_기자": "618b17e7ef7827cfea34e9c2",
		"한준": "618b1849ef7827cfea34ea1e",
		"인성": "618b18b0d3fbf96c644dbc0d",
		"지희": "60a3c91b4bc87f0d62b09a50",
		"영규": "61130cd072f7b1a0c044b7f5",
		"유라": "61130d6cf89dd58a4c13295d",
		"덕구": "618203826672d21ebf37748e",
		"보라": "618203f635ea62f8574c7d8a",
		"정원": "5d01a529bb04140008a23f36",
		"지우": "5ebea14e433bd200073e0dce",
		"범수": "5c3c52cb5827e00008dd7f3b",
		"현진": "5d01a522bb04140008a23f34",
		"준상": "5ecbbc7399979700087711db",
		"류은": "61659c5818732016a95fe763",
		"지안": "615c6a4e369566b08b8a7a71",
		"댄": "615c6abd9d9ddf11c620ad8e",
		"비비안": "615c6b331b008225e4d0fbe3",
		"행크": "602bd4fa8cfef478f1ceef0d",
		"도넛": "6075b407b8a5c6f36b877c88",
		"아봉": "61532c5aed9bfa8b54d5dff6",
		"팡팡": "61532cab9119555d352f5c69",
		"바네사": "61377f18a4d286537b5aba26",
		"로미": "61377f7ca4d286537b5aba6c",
		"로보": "61377fdaa23e3c47ccf5e93e",
		"재이": "612ed01c7eb720fddd3ddedf",
		"준기": "612ee95a222d88906b0e1b5e",
		"정진": "61250adfc84cd70ce86622d0",
		"한석필": "61250b6622f4f8c50236dc71",
		"현우": "611c3ef179a97a2c0d326a25",
		"세희": "611c3f692fac944dff493a04",
		"예린": "611c406d9540d10d3002e03a",
		"삼촌": "61532d0f3ef1e9b007472bf0",
		"오빠": "61532d6aec1b5e14b0fbc67a",
		"언니": "61532ddfd7042878b84aeeba",
		"이모": "61532e3bcdbc843c98494d96",
		"카메론": "610a664ce985c6893a292fc2",
		"터르난도": "610028154ecbefd29fb421b3",
		"키보": "6100287f568d6198a78bac31",
		"신다린": "61095ebc03c0e5b98f8ba69c",
		"케빈": "61095f50876293488937d5db",
		"가을": "60f7a217279ffba711ce2958",
		"노을": "60f7a2946f86ce86906f8a4d",
		"헬런": "60ee43c93a301a495e8e554e",
		"데이비드": "60e5421475f558ee6473a00a",
		"잭": "60e5426de8b95f1d3000d7b5",
		"애니": "60db2fccdd71a40a1d3476f8",
		"하준": "60db308484130840f23e6ca0",
		"하영": "60db35d96fdc060afd157d32",
		"조지": "60d1f8de6eff8301c397014a",
		"경완": "5d01a52163338e00072b8c9b",
		"글렌다": "602292ef9089786f78c93ef1",
		"더스트": "60c8b9fcb1de2b377c6c2cb2",
		"데이지": "60bf72699042ef1da40214c7",
		"리암": "60b64250db6e52c09714b379",
		"리철용": "60b642cc895d04d712086748",
		"순이": "60ad0841061ee28740ec2e1c",
		"창배": "60ad08c829e878c3c7d73965",
		"성배": "5d01a52bbb04140008a23f37",
		"소영": "5c789c317ad86500073a02cc",
		"채아": "609a8f4362c4bdb3363bbfad",
		"마가렛": "609a8fa4692037c4c168fe95",
		"루나": "609b98c2e587f6dbd19414b9",
		"우주": "5f8e95eae146f10007b85f45",
		"자비에르": "60915a63276a6a7600d7acfa",
		"이보미": "60915b5616d74069af8e8cab",
		"닥터두": "6088202cb772c71a7b9e0af2",
		"토미": "6088209815f2f9881e08e815",
		"정의로": "608820eb05b71bc949df989d",
		"파캉": "607ee367b245628c62b6ae98",
		"토비": "6080369d3211aa112ab131db",
		"듀크": "6075b391922523862ea7598d",
		"나나": "6076e25ac80469168e3771cf",
		"개나리": "606c6b127b9f53b4cd1743f5",
		"장태백": "606c6c155e38f609c6789d2b",
		"곽두필": "606c6c684085209e5555abb0",
		"하늘_캐스터": "606324cd7b4da19a1f6100be",
		"빠다가이": "6063252471850cc8f04c7600",
		"박창수": "6059dad0b83880769a50502f",
		"오미자": "6059db64e2d12a32cce1b0c7",
		"존_앵커": "6059dbe12fbc0b7887584a05",
		"쇼린이": "6050baed630c0d0906e65cc5",
		"TC_홈쇼핑": "6050bb4a630c0d0906e65d5d",
		"발키리": "60478557f12456064b353409",
		"카탈리나_기자": "604785d9ecccaef0e390a635",
		"아리": "6047863af12456064b35354e",
		"지미": "603fa0459441d42bc362eadf",
		"미스타_변사": "603fa172a669dfd23f450abd",
		"민지": "5eb55ce920d1b60016e91de6",
		"이카루스": "603513d91860484c4dcb6a11",
		"MC_타캐": "603514551860484c4dcb6acd",
		"단희빈": "602bd483b1837361b08c1ab6",
		"연우": "5eb55cf1f0b0a700071f89c7",
		"이영광_캐스터": "60228e8775ecc9e7ab8ef334",
		"라미": "60194446adfd1fbe5e0dd975",
		"제니퍼": "601944fb9089786f78c285ef",
		"레디오": "6019457dadfd1fbe5e0ddaf1",
		"연화": "600697fd8a8ea9b977284703",
		"삐뚤어진_찬구": "6010088f885570093ad24d53",
		"하은": "5de781b108fd6b000882c58f",
		"스모크": "60126fbf8e097503c73ed8d6",
		"채린이": "5ffda44bcba8f6d3d46fc41f",
		"호빈이": "5ffda49bcba8f6d3d46fc447",
		"용식이": "5feb2085cca1a479e73bac37",
		"김경화_앵커": "5feb213228b7247f8c8eb6d9",
		"VJ_싼타": "5fcf92faf3a7f2000799cc0c",
		"다보나_기자": "5fd89baaf8864c404f9097f4",
		"성호": "5fe06471a9f79e8f959be96f",
		"카밀라": "5f8d7b0de146f10007b8042f",
		"혜정": "5f8e9644f6120100074475f0",
		"한유격_교관": "5faa3acfac283a00075d0d2e",
		"성욱": "5d01a52b63338e00072b8c9e",
		"찬구": "5c547544fcfee90007fed455",
		"줄리아": "5e11db231c5eaa000b7d4c9c",
		"영희": "5c3c52caea9791000747155e",
		"애란": "5d01a52c63338e00072b8c9f",
		"주하": "5d01a52563338e00072b8c9c",
		"신혜": "5ecbbc4b0fbab10007bb3c7d",
		"윤성": "5ecbbc696b7c780008be50c4",
		"성규": "5ecbbc6099979700087711d8",
		"현경": "5ecbbc5599979700087711d6",
		"다희": "5ebea251fcf5110007b77d0f",
		"경숙": "5ebea266728f5b00075e6215",
		"광용": "5ebea235433bd200073e0e15",
		"정순": "5ebea242728f5b00075e6211",
		"나진": "5ebea185728f5b00075e61e3",
		"동우": "5ebea0fefcf5110007b77cc4",
		"정희": "5ebea1a364afaf00087fc2fb",
		"영길": "5ebea13564afaf00087fc2e7",
		"쥬비": "5ebea194728f5b00075e61e8",
		"용호": "5ea7b4268727390007ba95fc",
		"진혁": "5c3c52ca5827e00008dd7f36",
		"민상": "5c3c52ca5827e00008dd7f38",
		"수진": "5c3c52ca5827e00008dd7f3a",
		"미오": "5c3c52caea9791000747155c",
		"보노": "5c547544fcfee90007fed454",
		"덕춘": "5c3c52c95827e00008dd7f34",
		"이승주_기자": "5e4f7a5d716fd30009920721",
		"강수정_기자": "5e4f7a5da82e1f000aca31af",
		"주원": "5c547545fcfee90007fed459",
		"마크": "5e11db23b70e890009fb7ae7",
		"제임스": "5c547545168d0a0007f5cf13",
		"신디": "5c547545168d0a0007f5cf14",
		"재훈": "5c789c32dabcfa0008b0a38d",
		"지영": "5c789c34dabcfa0008b0a390",
		"선영": "5c789c317ad86500073a02cb",
		"국희": "5c789c31dabcfa0008b0a38c",
		"정섭": "5c789c32dabcfa0008b0a38e",
		"상도": "5c789c337ad86500073a02ce",
		"지철": "5c789c33dabcfa0008b0a38f",
		"금희": "5c789c337ad86500073a02cd",
		"명희": "5c789c347ad86500073a02cf",
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

	if speakUrls, ok = response["result"]; !ok {
		err = fmt.Errorf("undefind speark_urls")
		return
	}

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
