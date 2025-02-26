package randomname

import (
	"fmt"
	"github.com/disintegration/letteravatar"
	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"image/png"
	"os"
	"path"
	"unicode/utf8"
)

// GenerateAllNameAvatar 根据第一个字来生成头像，作为初始化函数调用
func GenerateAllNameAvatar() {
	dir := "uploads/chat_avatar"
	ttfpath := "uploads/file/喵字摄影体.ttf"
	err := initttf(ttfpath)

	if err != nil {
		fmt.Println("ttf初始化出错: ", err.Error())
		return
	}
	for _, s := range comicname {
		DrawImage(string([]rune(s)[0]), dir)
	}
	for _, s := range genshinname {
		DrawImage(string([]rune(s)[0]), dir)
	}
}

// 定义全局变量
var font *truetype.Font

// 全局初始化ttf
func initttf(ttfpath string) error {
	fontFile, err := os.ReadFile(ttfpath)
	font, err = freetype.ParseFont(fontFile)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

// DrawImage 生成单张图像,调用之前得先用initttf()初始化下字体

func DrawImage(name, dir string) {

	options := &letteravatar.Options{
		Font: font,
	}
	// 绘制文字
	firstLetter, _ := utf8.DecodeRuneInString(name)
	img, err := letteravatar.Draw(140, firstLetter, options)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 保存
	filePath := path.Join(dir, name+".png")
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = png.Encode(file, img)
	if err != nil {
		fmt.Println(err)
		return
	}
}
