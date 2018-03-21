package cli

import (
	"superBrain/model"
	"superBrain/db"
	"github.com/spf13/cobra"
	"os"
	"net/http"
	//"superBrain/kafka"
	"github.com/ethereum/go-ethereum/log"
	"superBrain/config"
	"superBrain/handle"
)

func TestDataCmd(name string) *cobra.Command {
	cjqc := &cobra.Command{
		Use:   name,
		Short: "superBrain services",
	}
	cjqc.AddCommand(startCmd())
	return cjqc
}

func startCmd() *cobra.Command {
	var addrPort string
	var addr[] string
	var user, passwd, ip, port, dbname string
	start := &cobra.Command{
		Use:   "start",
		Short: "Start superBrain system service",
		PreRun: func(cmd *cobra.Command, args []string) {
			if err := db.Init(user, passwd, ip, port, dbname); err != nil {
				log.Error("reporter init db err %v\n", err)
				os.Exit(1)
			}
			if err := model.Init(db.GetDB()); err != nil {
				log.Error("reporter init model err %v\n", err)
				os.Exit(1)
			}
		},
		Run: func(cmd *cobra.Command, args []string) {
			router := handle.CreateHttpRouter()
			if err := http.ListenAndServe(addrPort, router); err != nil {
				log.Error("reporter start err %v\n", err)
				os.Exit(1)
			}
		},
	}
	flags := start.Flags()
	flags.StringVar(&addrPort, "listen", ":8088", " http request listen port")
	flags.StringArrayVar(&addr, "kafuka", config.Get().Kafka.KafkaBrokers,"Kafuka IpAddr And Port")
	flags.StringVar(&user, "db_user", "root", "App Database User")
	flags.StringVar(&passwd, "db_passwd", "123456", "App Database Passwd")
	flags.StringVar(&ip, "db_ip", "192.168.90.241", "App Database IP")
	flags.StringVar(&port, "db_port", "3306", "App Database Port")
	flags.StringVar(&dbname, "db_name", "ruidong", "App Database Name")
	return start
}

