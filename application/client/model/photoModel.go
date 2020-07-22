package model

import (
	"fmt"
	// "fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path"
)

type PhotoModel struct {
}

func (this *PhotoModel) DownPhoto(address string) (localAddress string) {
	// imgUrl := "http://simisimisimiru.gitee.io/imagebed/3.png"
	localAddress = "../data/image/" + path.Base(address)
	fmt.Println(localAddress, address)
	_, err := os.Stat(localAddress)
	if err != nil {
		res, _ := http.Get(address)
		defer res.Body.Close()
		data, _ := ioutil.ReadAll(res.Body)
		ioutil.WriteFile(localAddress, data, os.ModePerm)
		return
	} else {
		return
	}

}
