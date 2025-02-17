package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"time"

	"github.com/go-vgo/robotgo"
	pdfcpu "github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
	hook "github.com/robotn/gohook"
	"github.com/vcaesar/imgo"
)

/*
pos: 269, 164
pos: 802, 996
pos: 911, 149
pos: 1468, 997
*/

func main() {

	// 기본 경로
	path := "./book"

	// 책 설정 값
	start := 1
	end := 660
	name := "시작하세요_도커_쿠버네티스"

	// 캡처 설정 값. TestCapture에서 값 확인
	spots := [][2]int{{270, 160}, {910, 160}}
	w := 550 //530
	h := 840 //820

	// 코드 시작
	ok := hook.AddEvent(robotgo.Enter)
	if ok {
		fmt.Println("Event Start")
	}

	for i := start; i <= end; i += 2 {
		if i%50 == 1 && i != start {
			refresh()
		}

		capture(spots[0], w, h, fmt.Sprintf("%s/img%03d.png", path, i))
		capture(spots[1], w, h, fmt.Sprintf("%s/img%03d.png", path, i+1))
		move()
	}

	mergeToPDH(path, name)

}

func capture(spot [2]int, w, h int, name string) {

	// time.Sleep(time.Duration((5*rand.Float32() + 1)) * time.Second)
	bit := robotgo.CaptureScreen(spot[0], spot[1], w, h)
	defer robotgo.FreeBitmap(bit)

	img := robotgo.ToImage(bit)
	imgo.Save(name, img)
}

func refresh() {
	robotgo.KeyTap("r", "cmd")
	time.Sleep(10 * time.Second)
	robotgo.MoveClick(876, 504)
	time.Sleep(1 * time.Second)
	robotgo.MoveClick(1235, 159)
	time.Sleep(1 * time.Second)
}

func move() {
	robotgo.MoveClick(1613, 573)
	time.Sleep(1 * time.Second)
}

func mergeToPDH(path string, name string) {

	root := os.DirFS(path)

	pngFiles, err := fs.Glob(root, "*.png")
	if err != nil {
		log.Fatal(err)
	}

	files := make([]string, len(pngFiles))
	for i, f := range pngFiles {
		files[i] = fmt.Sprintf("%s/%s", path, f)
	}

	imp, _ := pdfcpu.Import("form:A3, pos:c, s:1.0", types.POINTS)
	pdfcpu.ImportImagesFile(files, name+".pdf", imp, nil)
}
