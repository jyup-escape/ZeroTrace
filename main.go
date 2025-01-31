package main

import (
	"fmt"
	"os"
	"time"
)


func secureDelete(filePath string) error {
	var err error
	for i := 0; i < 5; i++ {
		err = os.Remove(filePath)
		if err == nil {
			return nil 
		}
		fmt.Printf("ファイル削除に失敗しました: %v\n", err)
	}
	return fmt.Errorf("ファイル削除に失敗しました: %v", err)
}

func main() {
	filePath := "C:/Users/haruh/Downloads/test.txt" // ここで削除するファイルのパスを指定


	err := secureDelete(filePath)
	if err != nil {
		fmt.Println("エラー:", err)
	} else {
		fmt.Println("ファイルが安全に削除されました。")
	}
}
