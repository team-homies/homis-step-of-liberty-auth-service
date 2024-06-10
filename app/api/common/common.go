package common

import (
	"context"
	"main/app/grpc/proto/dex"
	"main/config"
	"main/constant/common"
	"strconv"

	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

// 수집률로 시각적 성취도 단계 분류
func PercentCal(percentage uint64) (Code string) {

	switch {
	case percentage >= 80:
		Code = common.Ultimate
	case percentage >= 60:
		Code = common.Perfect
	case percentage >= 40:
		Code = common.Champion
	case percentage >= 20:
		Code = common.Rookie
	default:
		Code = common.Baby
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
