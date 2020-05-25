package main

import (
	"fmt"
	"path"

	"github.com/manga-community/mangadex"
)

func main() {

	md := mangadex.Initilize()

	links, err := md.Latest("2ZevhabKgkstB6DPzQpMcdSRnxwf78uC")

	// r, err := md.RetrieveImageLinks(810915)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(len(links))
	for _, element := range links {
		fmt.Println(path.Base(element))
	}

	// res, err := md.GetInfo("1")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// for _, element := range res.Chapters {
	// 	r, err := md.RetrieveImageLinks(element.ID)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		return
	// 	}
	// 	fmt.Println(r.Links)
	// }

}
