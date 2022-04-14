package gin

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//Todo
// 1단계 : 메모 crud 기능 구현하기
// 2단계 : jwt 인증을 통한 로그인관리 - 메모만 사용한다고하면 그냥 로그인 관련 시간만 추가 해놓을것
// 3단계 : 사용자별로 메모 관리 기능

func ReadMemo(c *gin.Context) {
	//Todo
	// 메모 불러오기 기능

	c.String(http.StatusOK, "Read")
}

func DeleteMemo(c *gin.Context) {
	//Todo
	// 메모 삭제 기능

	c.String(http.StatusOK, "Delete")
}

func UpdateMemo(c *gin.Context) {
	//Todo
	// 메모 업데이트 기능

	c.String(http.StatusOK, "Update")
}

func CreateMemo(c *gin.Context) {
	//Todo
	// 메모 생성 기능

	c.String(http.StatusOK, "Create")
}
