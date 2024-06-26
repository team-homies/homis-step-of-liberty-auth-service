package common

import (
	"context"
	"main/app/api/auth/resource"
	"main/app/grpc/proto/dex"
	"main/config"
	"main/database/repository"
	"strconv"

	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

// 수집률로 시각적 성취도 단계 분류
func PercentCal(percentage uint64) (codePercent uint64) {

	switch {
	case percentage >= 80:
		codePercent = 80
	case percentage >= 60:
		codePercent = 60
	case percentage >= 40:
		codePercent = 40
	case percentage >= 20:
		codePercent = 20
	default:
		codePercent = 0
	}
	return
}

func GetSinglePercent(percentage uint64) (single uint64, err error) {
	singleRepotitory := repository.NewRepository()
	// 1. 레포지토리 호출하여 테이블 카운트 가져오기
	count, err := singleRepotitory.GetSinglePercent()
	if err != nil {
		return
	}
	Num := uint64(count)

	// 2. 가져온 코드정보에서 원하는 정보 가져오기
	visualInfo, err := singleRepotitory.FindVisualCode(percentage)
	if err != nil {
		return
	}
	// 3. 가져온 정보에서 level가져오기
	level := &resource.FindVisualCodeResponse{
		DisplayLevel: visualInfo.DisplayLevel,
	}
	// 4. 싱글 퍼센트 계산
	single = (percentage - (100/Num)*uint64(level.DisplayLevel)) * uint64(level.DisplayLevel)

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
