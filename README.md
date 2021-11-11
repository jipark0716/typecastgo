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
| 이름 | id | 미리듣기주소 |
| ---- | ---- | ---- |
| 박찬 기자 | 618b17e7ef7827cfea34e9c2 | https://light-wav-t2.s3-ap-northeast-2.amazonaws.com/data/actor/77dda966-26ba-453f-8b5a-74a26604a0ea.mp3| 
| 한준 | 618b1849ef7827cfea34ea1e | https://light-wav-t2.s3-ap-northeast-2.amazonaws.com/data/actor/c5ca2a27-553e-44be-809b-8bc0d84a81e0.mp3| 
| 인성 | 618b18b0d3fbf96c644dbc0d | https://light-wav-t2.s3-ap-northeast-2.amazonaws.com/data/actor/77e13bc2-7abd-4528-9049-262b92e415c3.mp3| 
| 지희 | 60a3c91b4bc87f0d62b09a50 | https://light-wav-t2.s3-ap-northeast-2.amazonaws.com/data/actor/5148aae4-f307-45ff-a65c-318362909477.mp3| 
| 영규 | 61130cd072f7b1a0c044b7f5 | https://light-wav-t2.s3-ap-northeast-2.amazonaws.com/data/actor/678a1a1f-28c0-4855-8555-fceea72ac800.mp3| 
| 유라 | 61130d6cf89dd58a4c13295d | https://light-wav-t2.s3-ap-northeast-2.amazonaws.com/data/actor/f667c0ec-40ac-4ce0-86b8-ab23aa55f2f5.mp3| 
| 덕구 | 618203826672d21ebf37748e | https://light-wav-t2.s3-ap-northeast-2.amazonaws.com/data/actor/d176310c-2a08-4558-b14a-0ecbd4048fe6.mp3| 
| 보라 | 618203f635ea62f8574c7d8a | https://light-wav-t2.s3-ap-northeast-2.amazonaws.com/data/actor/b8794de7-4523-4305-bfe6-b4ba7d5b475f.mp3| 
| 정원 | 5d01a529bb04140008a23f36 | https://light-wav-t2.s3-ap-northeast-2.amazonaws.com/data/actor/b996b398-faee-4a11-b736-e2827578cb5e.mp3| 
| 지우 | 5ebea14e433bd200073e0dce | https://light-wav-t2.s3-ap-northeast-2.amazonaws.com/data/actor/f336204c-79d4-43e1-9e63-a77a8845a03c.mp3| 
| 범수 | 5c3c52cb5827e00008dd7f3b | https://light-wav-t2.s3-ap-northeast-2.amazonaws.com/data/actor/0d608c8c-ec54-4c66-968b-40f5689b1dc1.mp3| 
| 현진 | 5d01a522bb04140008a23f34 | https://light-wav-t2.s3-ap-northeast-2.amazonaws.com/data/actor/2fe5fbdc-cb19-48cd-bed7-a239c5496348.mp3| 
| 준상 | 5ecbbc7399979700087711db | https://light-wav-t2.s3-ap-northeast-2.amazonaws.com/data/actor/705aeb0d-291d-4c77-9f5a-16449d5da40a.mp3| 
| 류은 | 61659c5818732016a95fe763 | https://light-wav-t2.s3-ap-northeast-2.amazonaws.com/data/actor/f28146bd-fea9-4c87-a8c0-9db29cd9f12d.mp3| 
| 하나 | 61659cc118732016a95fe7c6 | https://light-wav-t2.s3-ap-northeast-2.amazonaws.com/data/actor/72705182-4c2a-41dc-8c3e-bf409fb0eda7.mp3| 
| 지안 | 615c6a4e369566b08b8a7a71 | https://light-wav-t2.s3-ap-northeast-2.amazonaws.com/data/actor/0696e6ea-b557-4973-87c0-65b1ef0b3f5f.mp3| 
| 댄 | 615c6abd9d9ddf11c620ad8e | https://light-wav-t2.s3-ap-northeast-2.amazonaws.com/data/actor/0638e5a9-d5aa-48c2-bf06-437db73bedd0.mp3| 
| 비비안 | 615c6b331b008225e4d0fbe3 | https://light-wav-t2.s3-ap-northeast-2.amazonaws.com/data/actor/9bd1b8ce-1b1c-4625-8511-bd2a253c24c8.mp3| 
| 행크 | 602bd4fa8cfef478f1ceef0d | https://light-wav-t2.s3-ap-northeast-2.amazonaws.com/data/actor/a4d93900-c29b-47b5-a12b-174ff598d482.mp3| 
| 도넛 | 6075b407b8a5c6f36b877c88 | https://light-wav-t2.s3-ap-northeast-2.amazonaws.com/data/actor/fd5875d0-647e-48fb-af8f-977dff265f25.mp3| 
| 아봉 | 61532c5aed9bfa8b54d5dff6 | https://light-wav-t2.s3-ap-northeast-2.amazonaws.com/data/actor/80cf3203-91dd-4ab1-b0b9-1604322ce540.mp3| 
| 팡팡 | 61532cab9119555d352f5c69 | https://light-wav-t2.s3-ap-northeast-2.amazonaws.com/data/actor/842d0a48-cafa-49d5-94e5-7142b9714cac.mp3| 
| 바네사 | 61377f18a4d286537b5aba26 | https://light-wav-t2.s3-ap-northeast-2.amazonaws.com/data/actor/f9a8275f-d766-4b77-9012-ffc1bb8f2294.mp3| 
| 로미 | 61377f7ca4d286537b5aba6c | https://light-wav-t2.s3-ap-northeast-2.amazonaws.com/data/actor/ed8b960f-ef65-4095-bc86-032cd66b5ec5.mp3| 
| 로보 | 61377fdaa23e3c47ccf5e93e | https://light-wav-t2.s3-ap-northeast-2.amazonaws.com/data/actor/6e533368-d865-461a-afe9-c5508e3c3587.mp3| 
| 재이 | 612ed01c7eb720fddd3ddedf | https://light-wav-t2.s3-ap-northeast-2.amazonaws.com/data/actor/97574756-1328-4d68-9fa5-de810c5ec6db.mp3| 
| 준기 | 612ee95a222d88906b0e1b5e | https://light-wav-t2.s3-ap-northeast-2.amazonaws.com/data/actor/18eb0a99-edfe-4e94-8047-5a49f7330529.mp3| 
| 정진 | 61250adfc84cd70ce86622d0 | https://light-wav-t2.s3-ap-northeast-2.amazonaws.com/data/actor/a9ca42b8-e34b-435f-af0b-3bf44f79e9ad.mp3| 
| 한석필 | 61250b6622f4f8c50236dc71 | https://light-wav-t2.s3-ap-northeast-2.amazonaws.com/data/actor/6679adf6-ea0e-4328-b9f4-30d1387237cc.mp3| 
| 현우 | 611c3ef179a97a2c0d326a25 | https://light-wav-t2.s3-ap-northeast-2.amazonaws.com/data/actor/c0609e14-cab9-44e0-99e9-11b62237d407.mp3| 
| 세희 | 611c3f692fac944dff493a04 | https://light-wav-t2.s3-ap-northeast-2.amazonaws.com/data/actor/792b2079-bb37-4d29-93b6-79f7ec6c6ba0.mp3| 
| 예린 | 611c406d9540d10d3002e03a | https://light-wav-t2.s3-ap-northeast-2.amazonaws.com/data/actor/a5b04cfd-cbd1-4ca3-b168-30c64937ddc5.mp3| 
| 삼촌 | 61532d0f3ef1e9b007472bf0 | https://light-wav-t2.s3-ap-northeast-2.amazonaws.com/data/actor/f253b18a-325a-4b85-b6ea-061f1ceea92a.mp3| 
| 오빠 | 61532d6aec1b5e14b0fbc67a | https://light-wav-t2.s3-ap-northeast-2.amazonaws.com/data/actor/610af948-1554-4729-a7aa-ceb692f7e31d.mp3| 
| 언니 | 61532ddfd7042878b84aeeba | https://light-wav-t2.s3-ap-northeast-2.amazonaws.com/data/actor/f4c01e2c-9775-424d-a8a7-ebd540d2dbf7.mp3| 
| 이모 | 61532e3bcdbc843c98494d96 | https://light-wav-t2.s3-ap-northeast-2.amazonaws.com/data/actor/c14c1300-2204-45f2-9a44-b0c22f99dab1.mp3| 
| 카메론 | 610a664ce985c6893a292fc2 | https://light-wav-t2.s3-ap-northeast-2.amazonaws.com/data/actor/f9e42591-62d8-4d0c-8f1e-7c8f2769c3cb.mp3| 
| 터르난도 | 610028154ecbefd29fb421b3 | https://light-wav-t2.s3-ap-northeast-2.amazonaws.com/data/actor/580f61db-a3b2-4496-8a66-ddc077e6e312.mp3| 
| 키보 | 6100287f568d6198a78bac31 | https://light-wav-t2.s3-ap-northeast-2.amazonaws.com/data/actor/9e64c954-9b6d-4399-9949-18865e541bcb.mp3| 
| 신다린 | 61095ebc03c0e5b98f8ba69c | https://light-wav-t2.s3-ap-northeast-2.amazonaws.com/data/actor/b03a1f54-3867-4383-b2ad-01758bdc3ae0.mp3| 
| 케빈 | 61095f50876293488937d5db | https://light-wav-t2.s3-ap-northeast-2.amazonaws.com/data/actor/3b2654fb-8297-4627-b1f4-eed348c814a9.mp3| 
| 가을 | 60f7a217279ffba711ce2958 | https://light-wav-t2.s3-ap-northeast-2.amazonaws.com/data/actor/d45755d0-a243-4def-b13a-b06ea7f15869.mp3| 
| 노을 | 60f7a2946f86ce86906f8a4d | https://light-wav-t2.s3-ap-northeast-2.amazonaws.com/data/actor/a6094530-548c-4927-a201-ce461af61c68.mp3| 
| 헬런 | 60ee43c93a301a495e8e554e | https://light-wav-t2.s3-ap-northeast-2.amazonaws.com/data/actor/78641973-3390-444f-8921-1096a47874c6.mp3| 
| 데이비드 | 60e5421475f558ee6473a00a | https://light-wav-t2.s3-ap-northeast-2.amazonaws.com/data/actor/03f7d645-de88-4832-b335-ccc91686d21a.mp3| 
| 잭 | 60e5426de8b95f1d3000d7b5 | https://light-wav-t2.s3-ap-northeast-2.amazonaws.com/data/actor/fd955539-5f6d-4e85-a8e8-35355a23d939.mp3| 
| 애니 | 60db2fccdd71a40a1d3476f8 | https://light-wav-t2.s3-ap-northeast-2.amazonaws.com/data/actor/100c34d9-3094-4d50-9ac7-d58378218dcd.mp3| 
| 하준 | 60db308484130840f23e6ca0 | https://light-wav-t2.s3-ap-northeast-2.amazonaws.com/data/actor/410c83e8-a251-4794-bddb-bf4852324c0f.mp3| 
| 하영 | 60db35d96fdc060afd157d32 | https://light-wav-t2.s3-ap-northeast-2.amazonaws.com/data/actor/905eedd0-c5f1-4315-bbdd-85ea8968a05f.mp3| 
| 조지 | 60d1f8de6eff8301c397014a | https://light-wav-t2.s3-ap-northeast-2.amazonaws.com/data/actor/f910ed9b-6cd5-40ab-8985-41ccffe58ee9.mp3| 
| 경완 | 5d01a52163338e00072b8c9b | https://light-wav-t2.s3-ap-northeast-2.amazonaws.com/data/actor/36a149bd-fd5d-4467-a645-fe6792f4b7a5.mp3| 
| 글렌다 | 602292ef9089786f78c93ef1 | https://light-wav-t2.s3-ap-northeast-2.amazonaws.com/data/actor/0ded3892-7906-4ba3-bcc5-e99c46969d4d.mp3| 
| 더스트 | 60c8b9fcb1de2b377c6c2cb2 | https://light-wav-t2.s3-ap-northeast-2.amazonaws.com/data/actor/2f6e7fdf-c40d-4292-9cca-554d345a3743.mp3| 
| 데이지 | 60bf72699042ef1da40214c7 | https://light-wav-t2.s3-ap-northeast-2.amazonaws.com/data/actor/a051edd7-4dbd-41ce-82f6-61c191f3c422.mp3| 
| 리암 | 60b64250db6e52c09714b379 | https://light-wav-t2.s3-ap-northeast-2.amazonaws.com/data/actor/5cd9876a-2206-4a5b-a950-b937b728fe1c.mp3| 
| 리철용 | 60b642cc895d04d712086748 | https://light-wav-t2.s3-ap-northeast-2.amazonaws.com/data/actor/ab30ec9c-e530-4e51-9d76-332c4bfcfc25.mp3| 
| 순이 | 60ad0841061ee28740ec2e1c | https://light-wav-t2.s3-ap-northeast-2.amazonaws.com/data/actor/764681da-9dc0-4162-8257-6556cc6c1f36.mp3| 
| 창배 | 60ad08c829e878c3c7d73965 | https://light-wav-t2.s3-ap-northeast-2.amazonaws.com/data/actor/ca4b4755-5dc3-4b40-a01f-c68c0d35d203.mp3| 
| 성배 | 5d01a52bbb04140008a23f37 | https://light-wav-t2.s3-ap-northeast-2.amazonaws.com/data/actor/67acc052-cb07-4fa1-9b6e-770180665a5c.mp3| 
| 소영 | 5c789c317ad86500073a02cc | https://light-wav-t2.s3-ap-northeast-2.amazonaws.com/data/actor/c26a210e-de30-45dd-9e22-1c19043a92b6.mp3| 
| 채아 | 609a8f4362c4bdb3363bbfad | https://light-wav-t2.s3-ap-northeast-2.amazonaws.com/data/actor/633bb8ef-aa0a-4224-8c56-9f8a749ae4e4.mp3| 
| 마가렛 | 609a8fa4692037c4c168fe95 | https://light-wav-t2.s3-ap-northeast-2.amazonaws.com/data/actor/182a9989-3daa-4d95-be3d-6ea39709405e.mp3| 
| 루나 | 609b98c2e587f6dbd19414b9 | https://light-wav-t2.s3-ap-northeast-2.amazonaws.com/data/actor/01362776-eae8-49b2-905b-74f52616dcd4.mp3| 
| 우주 | 5f8e95eae146f10007b85f45 | https://light-wav-t2.s3-ap-northeast-2.amazonaws.com/data/actor/ee7ff9fc-2e40-478d-8e5f-94bb59dd3275.mp3| 
| 자비에르 | 60915a63276a6a7600d7acfa | https://light-wav-t2.s3-ap-northeast-2.amazonaws.com/data/actor/68b3b468-ae74-4154-a69e-fa0863f7112f.mp3| 
| 이보미 | 60915b5616d74069af8e8cab | https://light-wav-t2.s3-ap-northeast-2.amazonaws.com/data/actor/b68aecb1-8749-4b3a-82f1-2f45ddd15332.mp3| 
| 닥터두 | 6088202cb772c71a7b9e0af2 | https://light-wav-t2.s3-ap-northeast-2.amazonaws.com/data/actor/2a69061d-ecfa-410b-949f-788dddbd2238.mp3| 
| 토미 | 6088209815f2f9881e08e815 | https://light-wav-t2.s3-ap-northeast-2.amazonaws.com/data/actor/5629d970-a61e-4b7f-88f7-fb0f3d218088.mp3| 
| 정의로 | 608820eb05b71bc949df989d | https://light-wav-t2.s3-ap-northeast-2.amazonaws.com/data/actor/e6e92e35-03d7-4d0e-8a29-df38bd64868f.mp3| 
| 파캉 | 607ee367b245628c62b6ae98 | https://light-wav-t2.s3-ap-northeast-2.amazonaws.com/data/actor/5c0d391d-36ae-4640-9dc0-39043868b53e.mp3| 
| 토비 | 6080369d3211aa112ab131db | https://light-wav-t2.s3-ap-northeast-2.amazonaws.com/data/actor/c23f4d75-aaea-4993-9bd5-2cd71a992109.mp3| 
| 듀크 | 6075b391922523862ea7598d | https://light-wav-t2.s3-ap-northeast-2.amazonaws.com/data/actor/f58f2fc5-f9a3-4cd3-8bb2-4643074d4062.mp3| 
| 나나 | 6076e25ac80469168e3771cf | https://light-wav-t2.s3-ap-northeast-2.amazonaws.com/data/actor/9878bc5b-2405-4a29-b046-84cef8d374d3.mp3| 
| 개나리 | 606c6b127b9f53b4cd1743f5 | https://light-wav-t2.s3-ap-northeast-2.amazonaws.com/data/actor/986f92a9-c5f5-4890-aeaf-26238bf8e266.mp3| 
| 장태백 | 606c6c155e38f609c6789d2b | https://light-wav-t2.s3-ap-northeast-2.amazonaws.com/data/actor/96f0d0c5-3dff-4567-b5ab-af9534eed4f8.mp3| 
| 곽두필 | 606c6c684085209e5555abb0 | https://light-wav-t2.s3-ap-northeast-2.amazonaws.com/data/actor/2c2c3a1d-e531-4020-a602-874f9d4be213.mp3| 
| 하늘 캐스터 | 606324cd7b4da19a1f6100be | https://light-wav-t2.s3-ap-northeast-2.amazonaws.com/data/actor/2aca19d8-ac39-4042-9e02-ea2354689398.mp3| 
| 빠다가이 | 6063252471850cc8f04c7600 | https://light-wav-t2.s3-ap-northeast-2.amazonaws.com/data/actor/b44a6adf-f66d-4ae3-a31f-86db792c3eb0.mp3| 
| 박창수 | 6059dad0b83880769a50502f | https://light-wav-t2.s3-ap-northeast-2.amazonaws.com/data/actor/0106e5b3-3c7d-4280-89e9-821a0acd57b5.mp3| 
| 오미자 | 6059db64e2d12a32cce1b0c7 | https://light-wav-t2.s3-ap-northeast-2.amazonaws.com/data/actor/91e2e71b-8b9c-4957-b547-368194c49766.mp3| 
| 존 앵커 | 6059dbe12fbc0b7887584a05 | https://light-wav-t2.s3-ap-northeast-2.amazonaws.com/data/actor/9ddd82a1-a1f8-48fa-bd0e-df1b022a8b66.mp3| 
| 쇼린이 | 6050baed630c0d0906e65cc5 | https://light-wav-t2.s3-ap-northeast-2.amazonaws.com/data/actor/451a5d87-2bfd-4cc2-95e7-ac884fd7fef3.mp3| 
| TC 홈쇼핑 | 6050bb4a630c0d0906e65d5d | https://light-wav-t2.s3-ap-northeast-2.amazonaws.com/data/actor/ef749eec-9757-4225-aef7-9158f2074563.mp3| 
| 발키리 | 60478557f12456064b353409 | https://light-wav-t2.s3-ap-northeast-2.amazonaws.com/data/actor/f8cde5be-784d-48c8-a41f-d4ce6166b53d.mp3| 
| 카탈리나 기자 | 604785d9ecccaef0e390a635 | https://light-wav-t2.s3-ap-northeast-2.amazonaws.com/data/actor/54be1d6f-cdbb-4f42-9bd4-f8a966707992.mp3| 
| 아리 | 6047863af12456064b35354e | https://light-wav-t2.s3-ap-northeast-2.amazonaws.com/data/actor/301c5f49-3137-4d70-b165-30f056892862.mp3| 
| 지미 | 603fa0459441d42bc362eadf | https://light-wav-t2.s3-ap-northeast-2.amazonaws.com/data/actor/f382832c-b300-4184-9a95-3c85291ac802.mp3| 
| 미스타 변사 | 603fa172a669dfd23f450abd | https://light-wav-t2.s3-ap-northeast-2.amazonaws.com/data/actor/5f28b4ed-be19-4ef7-ab42-aa4a555ec0e7.mp3| 
| 민지 | 5eb55ce920d1b60016e91de6 | https://light-wav-t2.s3-ap-northeast-2.amazonaws.com/data/actor/dbf541ab-18ba-41a8-a3ea-91c7473dcccc.mp3| 
| 이카루스 | 603513d91860484c4dcb6a11 | https://light-wav-t2.s3-ap-northeast-2.amazonaws.com/data/actor/34ead40a-132b-46bd-a9da-f21de2501789.mp3| 
| MC 타캐 | 603514551860484c4dcb6acd | https://light-wav-t2.s3-ap-northeast-2.amazonaws.com/data/actor/83e57050-88d0-4119-b24b-cdbbe40fcbf0.mp3| 
| 단희빈 | 602bd483b1837361b08c1ab6 | https://light-wav-t2.s3-ap-northeast-2.amazonaws.com/data/actor/ffbc299c-700d-4ab5-a136-d0154e2d60f1.mp3| 
| 연우 | 5eb55cf1f0b0a700071f89c7 | https://light-wav-t2.s3-ap-northeast-2.amazonaws.com/data/actor/749c8333-13f1-4c8f-ad98-b873eeabe4ac.mp3| 
| 이영광 캐스터 | 60228e8775ecc9e7ab8ef334 | https://light-wav-t2.s3-ap-northeast-2.amazonaws.com/data/actor/e9354c34-6c21-4640-bfb7-b163fae9fa27.mp3| 
| 라미 | 60194446adfd1fbe5e0dd975 | https://light-wav-t2.s3-ap-northeast-2.amazonaws.com/data/actor/ed0bad18-da7a-4f9b-859a-15bf3785c4a6.mp3| 
| 제니퍼 | 601944fb9089786f78c285ef | https://light-wav-t2.s3-ap-northeast-2.amazonaws.com/data/actor/7c206e11-0654-4d28-bb02-fd22a3b756b8.mp3| 
| 레디오 | 6019457dadfd1fbe5e0ddaf1 | https://light-wav-t2.s3-ap-northeast-2.amazonaws.com/data/actor/7038091d-c16c-4850-9b96-3ec88e210254.mp3| 
| 연화 | 600697fd8a8ea9b977284703 | https://light-wav-t2.s3-ap-northeast-2.amazonaws.com/data/actor/22c58da6-6cf3-4bd4-8bb6-b724fc7051e8.mp3| 
| 삐뚤어진 찬구 | 6010088f885570093ad24d53 | https://light-wav-t2.s3-ap-northeast-2.amazonaws.com/data/actor/d6b2eeb3-119b-4340-83a0-5130e8b87a3d.mp3| 
| 하은 | 5de781b108fd6b000882c58f | https://light-wav-t2.s3-ap-northeast-2.amazonaws.com/data/actor/2c3d4a01-b881-4a78-86bf-832872f14c9e.mp3| 
| 스모크 | 60126fbf8e097503c73ed8d6 | https://light-wav-t2.s3-ap-northeast-2.amazonaws.com/data/actor/8a8b29ca-5547-4e80-8863-014ab97586b3.mp3| 
| 채린이 | 5ffda44bcba8f6d3d46fc41f | https://light-wav-t2.s3-ap-northeast-2.amazonaws.com/data/actor/607bd90c-739f-4374-82c6-398170d14aac.mp3| 
| 호빈이 | 5ffda49bcba8f6d3d46fc447 | https://light-wav-t2.s3-ap-northeast-2.amazonaws.com/data/actor/7c8b2421-b35f-4fba-bd2e-336962194ac1.mp3| 
| 용식이 | 5feb2085cca1a479e73bac37 | https://light-wav-t2.s3-ap-northeast-2.amazonaws.com/data/actor/eeeb9a5e-f588-4940-abc1-688a5415facc.mp3| 
| 김경화 앵커 | 5feb213228b7247f8c8eb6d9 | https://light-wav-t2.s3-ap-northeast-2.amazonaws.com/data/actor/cb21ac15-7c73-4400-8ff7-cab02a992b40.mp3| 
| VJ 싼타 | 5fcf92faf3a7f2000799cc0c | https://light-wav-t2.s3-ap-northeast-2.amazonaws.com/data/actor/365dec0d-7118-4f7e-9d3d-80f26692f7e8.mp3| 
| 다보나 기자 | 5fd89baaf8864c404f9097f4 | https://light-wav-t2.s3-ap-northeast-2.amazonaws.com/data/actor/34140a6b-a3ed-446a-af7b-0e6b80dfcc03.mp3| 
| 성호 | 5fe06471a9f79e8f959be96f | https://light-wav-t2.s3-ap-northeast-2.amazonaws.com/data/actor/f7010f60-0ffc-4cc8-87fa-82aec13aaddd.mp3| 
| 카밀라 | 5f8d7b0de146f10007b8042f | https://light-wav-t2.s3-ap-northeast-2.amazonaws.com/data/actor/99bd0221-cc39-4062-8427-5de97f86c37b.mp3| 
| 혜정 | 5f8e9644f6120100074475f0 | https://light-wav-t2.s3-ap-northeast-2.amazonaws.com/data/actor/1bce8da6-e4e7-45fd-a76c-6df3289a9c4d.mp3| 
| 한유격 교관 | 5faa3acfac283a00075d0d2e | https://light-wav-t2.s3-ap-northeast-2.amazonaws.com/data/actor/336b32e5-980d-484a-9b6c-a1f88f3f50da.mp3| 
| 성욱 | 5d01a52b63338e00072b8c9e | https://light-wav-t2.s3-ap-northeast-2.amazonaws.com/data/actor/68c8337a-977e-4e45-824f-bd15ec5242f0.mp3| 
| 찬구 | 5c547544fcfee90007fed455 | https://light-wav-t2.s3-ap-northeast-2.amazonaws.com/data/actor/b76de107-141c-40b9-8d87-6fe4f1b84380.mp3| 
| 줄리아 | 5e11db231c5eaa000b7d4c9c | https://light-wav-t2.s3-ap-northeast-2.amazonaws.com/data/actor/8813685d-f232-41db-ba0f-b310c403df21.mp3| 
| 영희 | 5c3c52caea9791000747155e | https://light-wav-t2.s3-ap-northeast-2.amazonaws.com/data/actor/05160586-bc96-4fb4-84f1-2871015152e0.mp3| 
| 애란 | 5d01a52c63338e00072b8c9f | https://light-wav-t2.s3-ap-northeast-2.amazonaws.com/data/actor/3247047b-76ae-4f34-8929-4967b25b1882.mp3| 
| 주하 | 5d01a52563338e00072b8c9c | https://light-wav-t2.s3-ap-northeast-2.amazonaws.com/data/actor/8e88b13e-338a-4ec4-8bf1-5342dfc0356f.mp3| 
| 신혜 | 5ecbbc4b0fbab10007bb3c7d | https://light-wav-t2.s3-ap-northeast-2.amazonaws.com/data/actor/c698f1b0-7b7a-480f-b3e4-e6d2d0c69699.mp3| 
| 윤성 | 5ecbbc696b7c780008be50c4 | https://light-wav-t2.s3-ap-northeast-2.amazonaws.com/data/actor/c4a1a949-ecf5-4666-a1b1-bfb37d367196.mp3| 
| 성규 | 5ecbbc6099979700087711d8 | https://light-wav-t2.s3-ap-northeast-2.amazonaws.com/data/actor/9b24ea99-783a-46e7-92ae-b3f74b1d242d.mp3| 
| 현경 | 5ecbbc5599979700087711d6 | https://light-wav-t2.s3-ap-northeast-2.amazonaws.com/data/actor/e854f048-d280-4c22-b1e8-0aa1996d567e.mp3| 
| 다희 | 5ebea251fcf5110007b77d0f | https://light-wav-t2.s3-ap-northeast-2.amazonaws.com/data/actor/0a6f7f5e-1e49-4bc9-a4d2-aba472fcab38.mp3| 
| 경숙 | 5ebea266728f5b00075e6215 | https://light-wav-t2.s3-ap-northeast-2.amazonaws.com/data/actor/e6a3c6cc-f319-48fb-a5c9-b5055b07d8e9.mp3| 
| 광용 | 5ebea235433bd200073e0e15 | https://light-wav-t2.s3-ap-northeast-2.amazonaws.com/data/actor/53ea3614-ac1d-4ffa-be71-8da4fcfd2073.mp3| 
| 정순 | 5ebea242728f5b00075e6211 | https://light-wav-t2.s3-ap-northeast-2.amazonaws.com/data/actor/6b38030d-b3b8-4c54-bb2c-e8f23c348daa.mp3| 
| 나진 | 5ebea185728f5b00075e61e3 | https://light-wav-t2.s3-ap-northeast-2.amazonaws.com/data/actor/08b3108a-198b-41e4-9a9c-578ba932d737.mp3| 
| 동우 | 5ebea0fefcf5110007b77cc4 | https://light-wav-t2.s3-ap-northeast-2.amazonaws.com/data/actor/8f3945a7-f9a8-44dd-a65e-37b67c4a5e0a.mp3| 
| 정희 | 5ebea1a364afaf00087fc2fb | https://light-wav-t2.s3-ap-northeast-2.amazonaws.com/data/actor/d7891ec6-be3a-43b1-80d2-dff07955946e.mp3| 
| 영길 | 5ebea13564afaf00087fc2e7 | https://light-wav-t2.s3-ap-northeast-2.amazonaws.com/data/actor/d9261acb-644e-4d80-98a7-91a9515aff58.mp3| 
| 쥬비 | 5ebea194728f5b00075e61e8 | https://light-wav-t2.s3-ap-northeast-2.amazonaws.com/data/actor/1775a2cd-c8af-4910-ad03-5f827a606018.mp3| 
| 용호 | 5ea7b4268727390007ba95fc | https://light-wav-t2.s3-ap-northeast-2.amazonaws.com/data/actor/ff6a5c07-8524-4f41-9361-5713b4a1a4fa.mp3| 
| 진혁 | 5c3c52ca5827e00008dd7f36 | https://light-wav-t2.s3-ap-northeast-2.amazonaws.com/data/actor/59708235-7547-44ab-bb85-2251fadcd812.mp3| 
| 민상 | 5c3c52ca5827e00008dd7f38 | https://light-wav-t2.s3-ap-northeast-2.amazonaws.com/data/actor/651301d9-0b7a-4914-b3f2-f742d26b5f3a.mp3| 
| 찬혁 | 5c547545fcfee90007fed458 | https://light-wav-t2.s3-ap-northeast-2.amazonaws.com/data/actor/b973b59a-1926-42aa-b369-d575804a526e.mp3| 
| 수진 | 5c3c52ca5827e00008dd7f3a | https://light-wav-t2.s3-ap-northeast-2.amazonaws.com/data/actor/5aa04b11-0e90-4b3e-a777-93a9a2343fea.mp3| 
| 미오 | 5c3c52caea9791000747155c | https://light-wav-t2.s3-ap-northeast-2.amazonaws.com/data/actor/37e75c48-889b-49cd-b1e4-64dd90191cb5.mp3| 
| 보노 | 5c547544fcfee90007fed454 | https://light-wav-t2.s3-ap-northeast-2.amazonaws.com/data/actor/32cbb891-0782-45a8-9f21-0d919fdd9f3d.mp3| 
| 덕춘 | 5c3c52c95827e00008dd7f34 | https://light-wav-t2.s3-ap-northeast-2.amazonaws.com/data/actor/046a9e59-fac2-41df-921e-67ef506b2d70.mp3| 
| 이승주 기자 | 5e4f7a5d716fd30009920721 | https://light-wav-t2.s3-ap-northeast-2.amazonaws.com/data/actor/0f83255a-08df-4d31-b7ce-f509cca63866.mp3| 
| 강수정 기자 | 5e4f7a5da82e1f000aca31af | https://light-wav-t2.s3-ap-northeast-2.amazonaws.com/data/actor/957ce44f-327c-4c3a-bcb9-a178c66495d5.mp3| 
| 주원 | 5c547545fcfee90007fed459 | https://light-wav-t2.s3-ap-northeast-2.amazonaws.com/data/actor/82d967d0-41f3-4ee6-8b11-58964d1c585d.mp3| 
| 의찬이 | 5c547545fcfee90007fed456 | https://light-wav-t2.s3-ap-northeast-2.amazonaws.com/data/actor/228c9093-cabe-4383-80ab-47e7cf84a206.mp3| 
| 마크 | 5e11db23b70e890009fb7ae7 | https://light-wav-t2.s3-ap-northeast-2.amazonaws.com/data/actor/2936d146-7d76-4af8-b764-c7823786d5a1.mp3| 
| 제임스 | 5c547545168d0a0007f5cf13 | https://light-wav-t2.s3-ap-northeast-2.amazonaws.com/data/actor/59cac92a-9602-43da-8518-edc94f9e0b01.mp3| 
| 신디 | 5c547545168d0a0007f5cf14 | https://light-wav-t2.s3-ap-northeast-2.amazonaws.com/data/actor/66f63ecf-4571-412b-9a4d-00b51120dcb6.mp3| 
| 재훈 | 5c789c32dabcfa0008b0a38d | https://light-wav-t2.s3-ap-northeast-2.amazonaws.com/data/actor/fced5a25-72a2-4c79-a4fe-d5636c194429.mp3| 
| 지영 | 5c789c34dabcfa0008b0a390 | https://light-wav-t2.s3-ap-northeast-2.amazonaws.com/data/actor/0e96f5e7-b180-4e9e-8468-6baa12f97654.mp3| 
| 선영 | 5c789c317ad86500073a02cb | https://light-wav-t2.s3-ap-northeast-2.amazonaws.com/data/actor/3998cbb1-efe8-419f-9b06-6cc56e35b3cf.mp3| 
| 국희 | 5c789c31dabcfa0008b0a38c | https://light-wav-t2.s3-ap-northeast-2.amazonaws.com/data/actor/9fa7ec3b-6a26-49be-a73d-f59922ea4363.mp3| 
| 정섭 | 5c789c32dabcfa0008b0a38e | https://light-wav-t2.s3-ap-northeast-2.amazonaws.com/data/actor/40dd3303-e293-445b-9f95-952e83ea7b06.mp3| 
| 상도 | 5c789c337ad86500073a02ce | https://light-wav-t2.s3-ap-northeast-2.amazonaws.com/data/actor/02e5bfd9-6129-4254-8705-fba1eb66bd08.mp3| 
| 지철 | 5c789c33dabcfa0008b0a38f | https://light-wav-t2.s3-ap-northeast-2.amazonaws.com/data/actor/66d68096-58a5-4779-af8a-4a2375c45f2c.mp3| 
| 금희 | 5c789c337ad86500073a02cd | https://light-wav-t2.s3-ap-northeast-2.amazonaws.com/data/actor/770e3a6c-c0c7-4dd7-875c-b2ae9776f38c.mp3| 
| 명희 | 5c789c347ad86500073a02cf | https://light-wav-t2.s3-ap-northeast-2.amazonaws.com/data/actor/9715e78f-bb7c-4dad-821a-086ab6904954.mp3| 
