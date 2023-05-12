package main

import (
	"bufio"
	"fmt"
	"net/url"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func main(){
	// 必要な情報を取得する
	project := inputProject()
	title := inputTitle()
	articleUrl := inputArticleUrl()

	// ScrapboxのURLを生成する
	pageUrl := generatePageUrl(project,title,articleUrl)
	askDecision(pageUrl)
}

func generatePageUrl(project,title,articleUrl string)string{
	query := url.Values{}
	query.Set("url",articleUrl)
	encodedQuery := query.Encode()
	pageUrl := fmt.Sprintf("https://scrapbox.io/%s/%s?body=%s",project,title,encodedQuery)
	return pageUrl
}

// プロジェクト名を取得する
func inputProject()string{
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Project:")
	scanner.Scan()
	project:= strings.TrimSpace(scanner.Text())
	return project
}

// タイトルを取得する
func inputTitle()string{
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Title:")
	scanner.Scan()
	title:= strings.TrimSpace(scanner.Text())
	return title
}

// 記事のURLを取得する
func inputArticleUrl()string{
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("URL:")
	scanner.Scan()
	url:= strings.TrimSpace(scanner.Text())
	return url
}

// 入力した内容でページを作成するか確認する
func askDecision(url string){
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("\nOK? [y/n]:")
	scanner.Scan()
	decision:= strings.TrimSpace(scanner.Text())
	if decision == "y"{
		// ブラウザで新規ページを開く
		openUrl(url)
	}else if decision == "n"{
		// キャンセルしてプログラムを終了する
		fmt.Println("キャンセルしました")
		os.Exit(0)
	}else{
		fmt.Println("y/nで入力してください")
		askDecision(url)
	}
}

func openUrl(url string)error{
	var cmd string
	var args []string

	switch runtime.GOOS{
		case "linux":
			if os.Getenv("WSL_DISTRO_NAME") != ""{
				cmd = "cmd.exe"
				args = []string{"/c","start",url}
				fmt.Print("wsl")
			}else{
				cmd = "xdg-open"
				args = []string{url}
				fmt.Print("linux")
			}
		case "windows":
			cmd = "cmd.exe"
			args = []string{"/c","start"}
			fmt.Print("windows")
		case "darwin":
			cmd = "open"
			args = []string{url}
			fmt.Print("darwin")
	}

	command := exec.Command(cmd,args...)
	err := command.Start()
	if err != nil{
		return err
	}
	return nil
}
