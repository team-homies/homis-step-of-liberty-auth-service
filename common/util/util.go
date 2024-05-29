package util

import "encoding/json"

func StructToMap(data any) (result map[string]any, err error) {
	result = make(map[string]any)

	jsonData, err := json.Marshal(data)
	if err != nil {
		return
	}

	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		return
	}

	return
}

// 수집률로 시각적 성취도 분류
func PercentCal(percentage uint) (Code string) {
	switch {
	case percentage >= 80:
		Code = "EM"
	case percentage >= 60:
		Code = "DM"
	case percentage >= 40:
		Code = "CM"
	case percentage >= 20:
		Code = "BM"
	default:
		Code = "AM"
	}
	return
}
