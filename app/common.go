package app

import (
	"main/config"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func HistoryGetRateGrpc(c *fiber.Ctx) (err error) {
	// 0. grpc 연결 맺기
	var address string
	if viper.GetString(config.GRPC_HISTORY_HOST) == "localhost" {
		address = viper.GetString(config.GRPC_HISTORY_PORT)
	} else {
		address = viper.GetString(config.GRPC_HISTORY_HOST) + viper.GetString(config.GRPC_HISTORY_PORT)
	}
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return err
	}
	defer conn.Close()

	// dexClient := dex.NewDexEventServiceClient(conn)

	// 1. history 수집률 grpc호출
	// res, err := dexClient.GetRate(context.Background(), &dex.DexEventRequest{
	// 	Id: ,
	// })
	// if err != nil {
	// 	return err
	// }
	return

}
