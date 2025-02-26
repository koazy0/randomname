package randomname

import (
	"math/rand"
)

const (
	Genshin = iota
	Comic
)

// GenerateNameAndReturnFirst 生成随机昵称并返回第一个字
func GenerateNameAndReturnFirst() (nickName string, nickNameFirst string) {

	selectedType := rand.Intn(2)
	var name string
	switch selectedType {
	case Genshin:
		name = genshinname[rand.Intn(genshinnameCount)]

	case Comic:
		name = comicname[rand.Intn(comicnameCount)]

	}

	nickNameFirst = string([]rune(name)[0])
	nickName = adjwords[rand.Intn(adjCount)] + name
	return
}
