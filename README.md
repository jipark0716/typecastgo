# typecastGo

### get start
```
# go get github.com/jipark0716/typecastGo
```

### sample
```
import (
    "github.com/jipark0716/typecastGo"
    "io/ioutil"
)

func main() {
    typeCast, err := typecastGo.NewTypeCast("{id}", "{password}")
    typeCast.ActorName = "찬구"
    blob, _ := typeCast.Exec("ㅋㅋ루ㅋㅋ")

    ioutil.WriteFile("ㅋㅋㄹㅋㅋ.wav", blob, 0707)
}
```

### 목소리 지원
찬구, 주하, 카밀라, 우주, 혜정, 한유격, 영희, 소영, 지영
경완, 성욱, 애란, 줄리아, 민지, 신혜, 준상, 현경, 성규
윤성, 관용, 정순, 다희, 경숙, 동우, 연길, 지우, 나진, 쥬비
정희, 연우, 용호, 재훈, 정섭, 상도, 진혁, 민상, 찬혁, 지철
수진, 국희, 제임스, 선영, 금희, 명희, 미오, 보노, 덕춘
강수정, 이승주, 주원, 의찬, 마크, 트럼프, 김정은
신디, 현진, 정원, 성배, 리춘희, 하은, 범수
