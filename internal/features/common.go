package features

import (
	"encoding/json"
	"fmt"
	"go_api/internal/models"
)

func MapToStruct(data map[string]interface{}, result *models.Memo) {
	jsonbody, err := json.Marshal(data)
	if err != nil {
		// do error check
		fmt.Println(err)
	}
	fmt.Println(jsonbody)
	if err := json.Unmarshal(jsonbody, &result); err != nil {
		// do error check
		fmt.Println(err)
	}
	fmt.Println(result)
}
