package main

import (
	"fmt"
	"image"
	_ "image/png"
	"os"
	"time"

	"github.com/go-vgo/robotgo"
	"github.com/vcaesar/gcv"
)

func loadImage(imgPath string) image.Image {
	imageFile, _ := os.Open(imgPath)
	defer imageFile.Close()
	img, _, _ := image.Decode(imageFile)
	return img
}
func opencv(imgPath string) (int, int, bool) {
	var x = 0
	var y = 0
	var verificacao bool

	img := loadImage(imgPath)

	img2 := robotgo.CaptureScreen()
	img_full := robotgo.ToImage(img2)
	ress := gcv.FindAllImg(img, img_full)

	for i := range ress {
		x = ress[i].Middle.X
		y = ress[i].Middle.Y

	}
	if x == 0 {
		verificacao = false
	} else {
		verificacao = true
	}

	return x, y, verificacao
}
func banimento() {
	var campeao string = ""

	fmt.Println("Digite o nome do campeão que deseja banir: ")
	fmt.Scanf("%s", &campeao)
	time.Sleep(time.Millisecond * 800)
	lolIconX, lolIconY, _ := opencv("img/icone-lol.png")
	robotgo.Move(lolIconX, lolIconY)
	robotgo.Click()
	time.Sleep(time.Millisecond * 800)
	// pesqBanX, pesqBanY, _ := opencv("img/banir-pesquisar.png")
	robotgo.Move(1096, 258)
	time.Sleep(time.Millisecond * 800)
	robotgo.Click()
	robotgo.Click()
	robotgo.KeyPress("delete")
	time.Sleep(time.Millisecond * 800)

	robotgo.TypeStr(campeao)
	time.Sleep(time.Millisecond * 800)
	robotgo.Move(722, 313)
	time.Sleep(time.Millisecond * 800)
	robotgo.Click()
	time.Sleep(time.Millisecond * 800)
}

func confirmar_campeao() {
	var campeao string = ""
	fmt.Println("Digite o nome do seu campeão: ")
	fmt.Scanf("%s", &campeao)
	time.Sleep(time.Millisecond * 800)
	lolIconX, lolIconY, _ := opencv("img/icone-lol.png")
	robotgo.Move(lolIconX, lolIconY)
	robotgo.Click()
	time.Sleep(time.Millisecond * 800)
	pesqSelectX, pesqSelectY, _ := opencv("img/pesquisar-champ.png")
	robotgo.Move(pesqSelectX, pesqSelectY)
	robotgo.Click()
	robotgo.Click()
	robotgo.KeyPress("delete")
	time.Sleep(time.Millisecond * 800)
	robotgo.TypeStr(campeao)

	time.Sleep(time.Millisecond * 800)
	robotgo.Move(722, 313)
	robotgo.Click()
	time.Sleep(time.Millisecond * 800)

}

func main() {

	for {
		aceitarX, aceitarY, aceitarVerif := opencv("img/partida-encontrada.png")
		banirX, banirY, banirVerif := opencv("img/banir-nao-selecionado.png")
		confirmarX, confirmarY, confirmarVerf := opencv("img/confirmar-campeao.png")
		time.Sleep(time.Millisecond * 800)
		if aceitarVerif {
			fmt.Printf("Partida Encontrada!...")
			robotgo.Move(aceitarX, aceitarY)

			robotgo.Click()
			time.Sleep(time.Millisecond * 800)

		}
		if banirVerif {
			fmt.Printf("Banimento de Campeões!...")
			banimento()
			robotgo.Move(banirX, banirY)

			robotgo.Click()
			time.Sleep(time.Millisecond * 800)
			robotgo.Move(1096, 258)

		}
		if confirmarVerf {
			fmt.Printf("Confirmando Campeão!...")
			confirmar_campeao()
			robotgo.Move(confirmarX, confirmarY)

			robotgo.Click()
			time.Sleep(time.Millisecond * 800)

		} else {
			fmt.Println("Carregando...")
		}

	}

	//robotgo.Move(811, 953)
}
