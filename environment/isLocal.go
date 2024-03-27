package environment

import (
	"fmt"
	"os"
)

func PrintEnvVarsWithKey(env string,devfolder string){
	fmt.Printf("%s =%s\n",env, os.Getenv(env))
	if(IsLocalHTTPDevEnvironmentWithKey(env)) {
		fmt.Printf("%s =%s\n",devfolder, os.Getenv(devfolder))
	}
}
func IsLocalHTTPDevEnvironmentWithKey(env string) bool{
	KEY_ENV:= os.Getenv(env)
	return KEY_ENV=="LOCAL"
}
func GetDeveloperFolderWithKey(devfolder string) string{
	DEV_FOLDER:= os.Getenv(devfolder )
	return DEV_FOLDER
}
