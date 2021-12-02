package shared

import (
	"crypto/md5"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func Setup() {
	SaveConfig(Config{
		StoragePath: "/srv/papyrus",
		CLIConfig: CLIConfig{
			JenkinsURL: "https://jenkins.example.com",
			JenkinsFilePath: "{url}/job/{project}/{build}/artifact/{file}",
			Webhook: false,
			WebhookID: "",
			WebhookToken: "",
			SuccessEmbed: EmbedConfig{
				Title: "Build Successful",
				Description: "**Project:** {project} {version}\n" +
					"**Build:** {build}\n" +
					"**Status:** {result}\n" +
					"\n" +
					"**Changes:**\n" +
					"{changes}",
				Changes: "- `{short_hash}` *{title} - {author}*\n",
				Color: 3066993,
			},
			FailureEmbed: EmbedConfig{
				Title: "Build Failed",
				Description: "**Project:** {project} {version}\n" +
					"**Build:** {build}\n" +
					"**Status:** {result}\n" +
					"\n" +
					"**Changes:**\n" +
					"{changes}",
				Changes: "- `{short_hash}` *{title} - {author}*\n",
				Color: 10038562,
			},
		},
		WebConfig: WebConfig{
			IP: "localhost:3000",
			Dev: true,
		},
	})

	err := os.MkdirAll(GetConfig().StoragePath, os.ModePerm)
	if err != nil {
		panic(err)
	}

	SaveData(Data{
		Projects: nil,
	})
}

func Reset() {
	SaveData(Data{
		Projects: nil,
	})
}

func PrintDebug() {
	fmt.Printf("%+v", GetConfig())
	println("\n\n")
	fmt.Printf("%+v", GetData())
}

func GetMD5(path string) string {
	file, err := os.Open(path)

	hash := md5.New()
	_, err = io.Copy(hash, file)

	if err != nil {
		panic(err)
	}

	return fmt.Sprintf("%x", hash.Sum(nil))
}

func DownloadFile(url string, path string) {
	resp, err := http.Get(url)
	file, err := os.Create(path)
	_, err = io.Copy(file, resp.Body)
	err = resp.Body.Close()

	if err != nil {
		panic(err)
	}
}

func Before(string string, sep string) string {
	i := strings.Index(string, sep)
	if i == -1 {
		return string
	}
	return string[:i]
}

func After(string string, sep string) string {
	i := strings.Index(string, sep)
	if i == -1 {
		return string
	}
	return string[i+len(sep):]
}

func First(string string, x int) string {
	if len(string) < x {
		return string
	}
	return string[:x]
}
