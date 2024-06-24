package cmd


import(
	"flag"
	"fmt"
	"os"
)

func Echo(){
	echoCmd:=flag.NewFlagSet("head",flag.ExitOnError)
	trailing:=echoCmd.Bool("n",false,"newLine")
	echoCmd.Parse(os.Args[2:])
	args:=echoCmd.Args()

	for i:=0;i < len(args);i++{
		if(*trailing){
			fmt.Print(args[i], " ")
		}else{
			fmt.Println(args[i])
		}
	}
}