package service

import (
	"randomgifsite/internal/utils"
)

func CatGift() (string, error) {
	links := utils.GetGifUrl

	gifUrl, err := utils.ParceUrl(links()[1])
	if err != nil {
		return "", err
	}
	gifname, err := utils.GetFileName(gifUrl)
	if err != nil {
		return "", err
	}
	utils.DownloadFile(gifUrl, gifname)

	return gifname, nil

}
