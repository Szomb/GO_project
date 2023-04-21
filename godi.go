package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Bukkm2019 struct {
	Rajtszam  string
	Kategoria string
	Nev       string
	Egyesulet string
	Ido       string
	IdoT      time.Time
}

func main() {
	dfile, err := os.Open("bukkm2019.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer dfile.Close()

	reader := bufio.NewReader(dfile) //? READER NÉVEN CSINÁLOK EGY BEOLVASÁST

	csvReader := csv.NewReader(reader)
	csvReader.Comma = ';' //todo: MEGADOM MI ALAPJÁN VÁLASZTOM EL A FÁJLT
	csvReader.Read()
	var dataList []Bukkm2019
	for {
		line, err := csvReader.Read() //todo: BEOLVASOM ÉS ERRORHANDLINGET CSINÁLOK
		if err != nil {
			break
		} else {
			data := Bukkm2019{ //! A FELDARABOLT ADATOT BELERAKOM EGY STRUKTÚRÁLT LISTÁBA
				Rajtszam:  strings.TrimSpace(line[0]),
				Kategoria: strings.TrimSpace(line[1]),
				Nev:       strings.TrimSpace(line[2]),
				Egyesulet: strings.TrimSpace(line[3]),
				Ido:       strings.TrimSpace(line[4]),
				IdoT:      time.Time{},
			}
			t, err := time.Parse("15:04:05", strings.TrimSpace(line[4])) //? FELDARABOLOM AZ IDŐT
			if err != nil {                                              //?ERRORHANDLER
				fmt.Print(data.IdoT)
			}
			data.IdoT = t
			dataList = append(dataList, data) //!ELSŐ SOR BENNE VAN INDEXNÉL FIGYELNI

		}

	}
	osszIndulo := 691
	hossz := len(dataList)
	nemErtBe := 691 - hossz
	szazalek := float64(nemErtBe) / float64(osszIndulo) * 100
	noiVersenyzo := 0
	voltHat := false
	ffWindex := 660
	for i := 0; i < hossz; i++ {

		if strings.Contains(dataList[i].Rajtszam, "R") {
			if strings.Contains(dataList[i].Kategoria, "n") {
				noiVersenyzo++
			}
		}
		if string(dataList[i].Ido[:1]) == "6" {
			voltHat = true
		}

	}
	fmt.Printf("4. feladat: Versenytávot nem teljesítők: %.14f%%\n", szazalek)
	fmt.Println("5.feladat: A női versenyzők száma a rövidtávú versenyen", noiVersenyzo, " fő")
	if voltHat {
		fmt.Println("6.feladat: Volt ilyen versenyző!")
	} else {
		fmt.Println("Nem volt.")
	}
	fmt.Println("7.feladat: A felnőtt férfi (ff) kategória győztese rövid távon")
	for i := 0; i < hossz; i++ {
		alapT := strings.Split(dataList[i].Ido, ":")
		alapH, _ := strconv.Atoi(alapT[0])
		alapP, _ := strconv.Atoi(alapT[1])
		alapM, _ := strconv.Atoi(alapT[2])
		teljesM := alapH*60 + alapP + alapM/60
		legT := strings.Split(dataList[ffWindex].Ido, ":")
		legH, _ := strconv.Atoi(legT[0])
		legP, _ := strconv.Atoi(legT[1])
		legM, _ := strconv.Atoi(legT[2])
		maxP := legH*60 + legP + legM/60
		if strings.Contains(dataList[i].Kategoria, "ff") {
			if strings.Contains(dataList[i].Rajtszam, "R") && teljesM < maxP {
				ffWindex = i

			}

		}
	}

	fmt.Println("Rajtszám:", dataList[ffWindex].Rajtszam)
	fmt.Println("Név:", dataList[ffWindex].Nev)
	fmt.Println("Egysület:", dataList[ffWindex].Egyesulet)
	fmt.Println("Idő:", dataList[ffWindex].Ido)
}
