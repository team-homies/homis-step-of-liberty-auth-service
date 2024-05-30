package common

// middleware
const LOCALS_USER_ID string = "userId"

// 시각적 성취도 단계
const (
	Baby     string = "AM"
	Rookie   string = "BM"
	Champion string = "CM"
	Perfect  string = "AM"
	Ultimate string = "AM"
)

// 수집률로 시각적 성취도 단계 분류
func PercentCal(percentage uint) (Code string) {
	switch {
	case percentage >= 80:
		Code = Ultimate
	case percentage >= 60:
		Code = Perfect
	case percentage >= 40:
		Code = Champion
	case percentage >= 20:
		Code = Rookie
	default:
		Code = Baby
	}
	return
}
