package randomname

import (
	"fmt"
	"testing"
)

func TestGenerateAllNameAvatar(t *testing.T) {
	GenerateAllNameAvatar()
	fmt.Println("test")
}

func TestDrawImage(t *testing.T) {

	err := initttf("uploads/file/喵字摄影体.ttf")

	if err != nil {
		fmt.Println("ttf初始化出错: ", err.Error())
		return
	}
	DrawImage("温", "uploads/chat_avatar")
	fmt.Println("test")
}
