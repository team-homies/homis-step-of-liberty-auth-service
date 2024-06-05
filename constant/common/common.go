package common

import (
	"context"
	"main/app/grpc/proto/dex"
	"main/config"
	"strconv"

	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

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

// 수집률 grpc
func GetRateGrpc(userId uint) (result uint64, err error) {

	// 0. grpc 연동
	var address string
	if viper.GetString(config.GRPC_HISTORY_HOST) == "localhost" {
		address = viper.GetString(config.GRPC_HISTORY_PORT)
	} else {
		address = viper.GetString(config.GRPC_HISTORY_HOST) + viper.GetString(config.GRPC_HISTORY_PORT)
	}
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return
	}
	defer conn.Close()

	dexClient := dex.NewDexEventServiceClient(conn)

	// 1. userId를 이용해서 user의 수집률을 구하고 변수에 담는다
	rate, err := dexClient.GetRate(context.Background(), &dex.RateRequest{
		UserId: uint64(userId),
	})
	if err != nil {
		return
	}
	// 2. 형변환
	result, err = strconv.ParseUint(rate.Rate, 10, 64)
	if err != nil {
		return
	}
	return
}
