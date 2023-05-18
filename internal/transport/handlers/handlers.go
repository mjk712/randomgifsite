package handlers

import (
	"net/http"
	"randomgifsite/internal/service"
	"randomgifsite/internal/utils"
)

func GetCatGif(w http.ResponseWriter, r *http.Request) {

	gifName, _ := service.CatGift()

	gifPath := "/Users/grigorijmatukov/projects/randomgifsite/cmd/app/" + gifName

	http.ServeFile(w, r, gifPath)

	utils.DeleteFile(gifPath)
}
