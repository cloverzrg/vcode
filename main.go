package vcode

import (
	"fmt"
	"image"
	"image/jpeg"
	"net/http"
)

func getImg(client *http.Client, url string) (img image.Image, err error) {
	resp, err := client.Get(url)
	if err != nil {
		return img, err
	}
	defer resp.Body.Close()

	image.RegisterFormat("jpeg", "jpeg", jpeg.Decode, jpeg.DecodeConfig)
	img, _, err = image.Decode(resp.Body)
	if err != nil {
		return img, err
	}
	return img, err
}

func getMap(pixels [][]Pixel) (m map[int]map[int]int) {
	m = make(map[int]map[int]int)
	for i := 0; i < 22; i++ {
		m[i] = make(map[int]int)
		for j := 0; j < 62; j++ {
			if pixels[i][j].B+pixels[i][j].G+pixels[i][j].R < 350 {
				m[i][j] = 1
			} else {
				m[i][j] = 0
			}
		}
	}
	return m
}

func printMap(m map[int]map[int]int) {
	for i := 3; i < 15; i++ {
		pstr := ""
		for j := 4; j < 45; j++ {
			if m[i][j] > 0 {
				pstr += "*"
			} else {
				pstr += " "
			}
		}
		fmt.Println(pstr)
	}
}

func GetVCodeCookie(client *http.Client, url string) (vcode string, err error) {
	img, err := getImg(client, url)
	if err != nil {
		return vcode, err
	}
	p, err := getPixels(img)
	if err != nil {
		return vcode, err
	}
	m := getMap(p)
	vcode = GetVCodeFromMap(m)
	printMap(m)
	fmt.Println(vcode)
	return vcode, nil
}

//func main() {
//	jar, err := cookiejar.New(&cookiejar.Options{})
//
//	client := http.Client{
//		Jar:     jar,
//		Timeout: 20 * time.Second,
//	}
//	_, err = GetVCodeCookie(&client, "http://kdjw.hnust.edu.cn:8080/kdjw/verifycode.servlet")
//	if err != nil {
//		log.Fatal(err)
//	}
//	//fmt.Println(cookie[0].String())
//	//fmt.Printf("%+v", client.Jar.Cookies(&url.URL{
//	//	Scheme: "http",
//	//	Host:   "kdjw.hnust.edu.cn:8080",
//	//	Path:   "/kdjw",
//	//}))
//
//}
