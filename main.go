package main

import (
	"fmt"
	"os"

	"github.com/slack-go/slack"
)

func main(){
	//Token:The Token Access Authentication scheme provides a method for making authenticated HTTP requests using a token - an identifier used to denote an access grant with specific scope, duration, cryptographic properties, and other attributes

	//To set a key/value pair, use os.Setenv.
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-3764556107300-3747577450327-hIrsOvr73iDYogknfGIB8qtb")
	os.Setenv("CHANNEL_ID", "C03P3SDRW0Y")
	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))          //using slack package and creating a new connection
	//To get a value for a key, use os.Getenv. This will return an empty string if the key isn’t present in the environment.

	channelArr:= []string{os.Getenv("CHANNEL_ID")}  //slice getting channel id
	fileArr:= []string{"Infosys Online Test_Samples.pdf"}  //slice 

	for i:=0; i<len(fileArr); i++ {
		parameters := slack.FileUploadParameters{                //fileuploadpar. is a function of slack package. It will have 2 parameters
			Channels : channelArr,
			File : fileArr[i],
		}  
		file, err := api.UploadFile(parameters)    //will call this inbuilt func in slack api with parameters as specified
		if err!=nil{
			fmt.Printf("%s\n",err)
			return
		}
		fmt.Printf("Name: %s, URL: %s",file.Name,file.URL)
	}

}