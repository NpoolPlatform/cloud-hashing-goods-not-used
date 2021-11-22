package main

import (
	"fmt"
	"os"
	"path"
	"runtime"

	"github.com/NpoolPlatform/go-service-framework/pkg/app"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	mysqlconst "github.com/NpoolPlatform/go-service-framework/pkg/mysql/const"
	rabbitmqconst "github.com/NpoolPlatform/go-service-framework/pkg/rabbitmq/const"
	redisconst "github.com/NpoolPlatform/go-service-framework/pkg/redis/const"

	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/service-name" //nolint

	cli "github.com/urfave/cli/v2"
)

const serviceName = servicename.ServiceName

func main() {
	commands := cli.Commands{
		runCmd,
	}

	description := fmt.Sprintf("my %v service cli\nFor help on any individual command run <%v COMMAND -h>\n",
		serviceName, serviceName)

	_, mainPath, _, ok := runtime.Caller(0)
	if !ok {
		logger.Sugar().Errorf("cannot get path of main.go")
		return
	}
	err := app.Init(
		serviceName,
		description,
		"",
		"",
		path.Dir(mainPath),
		nil,
		commands,
		mysqlconst.MysqlServiceName,
		rabbitmqconst.RabbitMQServiceName,
		redisconst.RedisServiceName,
	)
	if err != nil {
		logger.Sugar().Errorf("fail to create %v: %v", serviceName, err)
		return
	}

	logger.Sugar().Infof("run main.go from: %v", mainPath)

	err = app.Run(os.Args)
	if err != nil {
		logger.Sugar().Errorf("fail to run %v: %v", serviceName, err)
	}
}
