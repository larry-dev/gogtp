package main

import (
	"fmt"
	"github.com/larry-dev/gogtp"
	"time"
)

func main()  {
	contrl,err:=gogtp.NewControllerByStr("./run.sh")
	if err!=nil{
		panic(err)
	}

	engine,err:=gogtp.NewEngine(contrl)
	if err!=nil{
		panic(err)
	}
	fmt.Println(engine.ClearBoard())
	fmt.Println(engine.GenMove("B"))
	fmt.Println(engine.GenMove("W"))



	nameChan:=make(chan string,1)
	contrl.SendCommand(gogtp.BuildCommand(gogtp.CmdName("name")), func(response gogtp.Response) {
		fmt.Println(response)
		res,err:=response.GetResult()
		fmt.Println(res,err)
		nameChan<-res
	})
	name:=<-nameChan
	fmt.Println("name:",name)
	go func() {
		time.Sleep(10*time.Second)
		contrl.StopSendCommand(gogtp.BuildCommand(gogtp.CmdName("genmove B")), func(response gogtp.Response) {
			fmt.Println(response)
			fmt.Println(contrl.ResetStdErr())
		})
	}()
	contrl.SendCommand(gogtp.BuildCommand(gogtp.CmdName("kata-analyze"),gogtp.CmdArgs("60"),gogtp.CmdEnd(false)), func(response gogtp.Response) {
		fmt.Println(response)
		fmt.Println(contrl.ResetStdErr())
	})
}
