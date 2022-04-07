# GoLang API 용 서버 프로젝트 입니다.

0. GOROOT 및 GOPATH 지정
   - GOROOT : go sdk 가 존재하는 위치
   - GOPATH : 적용될 프로젝트(혹은 모듈)가 있는 위치


1. 개발 서버 실행 Auto Build 적용(https://github.com/gravityblast/fresh)

```re
//main.go가 존재하는 디렉토리
cd /go_api/api  
// auto build 적용
./fresh -c runner.conf
```

2. golang으로 api를 학습, 구현, 테스트 하는 프로젝트
   - 프로젝트 구조가 언제든 변경될 수 있음
   - 학습용 예제를 삭제하지 말고 유지