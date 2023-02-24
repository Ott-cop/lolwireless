package main

import (
	"fmt"
	"image"
	_ "image/png"
	"os"
	"sync"
	"time"

	"github.com/go-vgo/robotgo"
	"github.com/vcaesar/gcv"
)

func main() {
	wg := sync.WaitGroup{}

	steps := getSteps()
	images := loadImages()
	for {
		validationStruct := make(chan struct {
			x, y     int
			detected bool
			imgType  string
		}, 3)
		baseImg := robotgo.ToImage(robotgo.CaptureScreen())

		wg.Add(len(steps))

		for s := range steps {
			go func(s string) {
				defer wg.Done()
				x, y, d := verifyImagePresence(images[s], baseImg)
				validationStruct <- struct {
					x, y     int
					detected bool
					imgType  string
				}{
					x:        x,
					y:        y,
					detected: d,
					imgType:  s,
				}
			}(s)
		}

		// aceitarX, aceitarY, isMatchFound := verifyImagePresence(images["notBannedYet"], baseImg)
		// banirX, banirY, isBanning := verifyImagePresence(images["notBannedYet"], baseImg)
		// confirmarX, confirmarY, isConfirmChamp := verifyImagePresence(images["confirmChamp"], baseImg)

		wg.Wait()
		close(validationStruct)

		for v := range validationStruct {
			switch v.imgType {
			case "matchFound":
				if v.detected {
					fmt.Printf("Partida Encontrada!...")
					robotgo.Move(v.x, v.y)
					robotgo.Click()
					time.Sleep(time.Millisecond * 800)
					steps[v.imgType] = true
					if steps["notBannedYet"] || steps["confirmChamp"] {
						steps = getSteps()
					}
				}

			case "notBannedYet":
				if v.detected {
					if steps[v.imgType] {
						continue
					}
					fmt.Printf("Banimento de Campeões!...")
					banning(images["lolIcon"], baseImg)
					robotgo.Move(v.x, v.y)
					robotgo.Click()
					time.Sleep(time.Millisecond * 800)
					robotgo.Move(1096, 258)
					time.Sleep(time.Millisecond * 800)
					steps[v.imgType] = true
				}

			case "confirmChamp":
				if v.detected {
					if steps[v.imgType] {
						continue
					}
					fmt.Printf("Confirmando Campeão!...")
					selectChampion(images["lolIcon"], images["searchChamp"], baseImg)
					robotgo.Move(v.x, v.y)
					robotgo.Click()
					time.Sleep(time.Millisecond * 800)
					steps[v.imgType] = true
				}

			default:
				fmt.Println("Carregando...")
				fmt.Println(v)
			}
		}
		fmt.Println("===================================================")
		time.Sleep(1 * time.Second)
	}
	//robotgo.Move(811, 953)
}

func getSteps() map[string]bool {
	return map[string]bool{
		"matchFound":   false,
		"notBannedYet": false,
		"confirmCham":  false,
	}
}

func loadImages() map[string]image.Image {
	return map[string]image.Image{
		"lolIcon":      loadImage("img/icone-lol.png"),
		"searchChamp":  loadImage("img/pesquisar-champ.png"),
		"matchFound":   loadImage("img/partida-encontrada.png"),
		"notBannedYet": loadImage("img/banir-nao-selecionado.png"),
		"confirmChamp": loadImage("img/confirmar-campeao.png"),
	}
}

func loadImage(imgPath string) image.Image {
	imageFile, _ := os.Open(imgPath)
	defer imageFile.Close()
	img, _, _ := image.Decode(imageFile)
	return img
}

func verifyImagePresence(searchedImg, baseImg image.Image) (int, int, bool) {
	var x = 0
	var y = 0
	var verificacao bool

	ress := gcv.FindAllImg(searchedImg, baseImg)

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

func banning(lolIcon, baseImg image.Image) {
	var campeao string = ""

	fmt.Println("Digite o nome do campeão que deseja banir: ")
	fmt.Scanf("%s", &campeao)
	time.Sleep(time.Millisecond * 800)
	lolIconX, lolIconY, _ := verifyImagePresence(lolIcon, baseImg)
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

func selectChampion(lolIcon, searchChamp, baseImg image.Image) {
	var campeao string = ""
	fmt.Println("Digite o nome do seu campeão: ")
	fmt.Scanf("%s", &campeao)
	time.Sleep(time.Millisecond * 800)
	lolIconX, lolIconY, _ := verifyImagePresence(lolIcon, baseImg)
	robotgo.Move(lolIconX, lolIconY)
	robotgo.Click()
	time.Sleep(time.Millisecond * 800)
	pesqSelectX, pesqSelectY, _ := verifyImagePresence(searchChamp, baseImg)
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
