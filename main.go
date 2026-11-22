package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/BigDuck10/PS-naloga5/vreme"
	"github.com/urfave/cli/v3"
)

func main() {

	cmd := &cli.Command{
		Name:  "vremenska-postaja",
		Usage: "Program za upravljanje vremenskih postaj",
		Flags: []cli.Flag{
			&cli.Float64Flag{
				Name:  "minTemp",
				Value: -50,
				Usage: "Najnižja dovoljena temperatura",
			},
			&cli.Float64Flag{
				Name:  "maxTemp",
				Value: 50,
				Usage: "Najvišja dovoljena temperatura",
			},
			&cli.IntFlag{
				Name:  "stMeritev",
				Value: 3,
				Usage: "Minimalno število meritev za statistiko",
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {

			minTemp := cmd.Float64("minTemp")
			maxTemp := cmd.Float64("maxTemp")
			stMeritev := cmd.Int("stMeritev")

			return run(minTemp, maxTemp, stMeritev)
		},
	}
	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}

}

func run(minTemp float64, maxTemp float64, stMeritev int) error {

	if minTemp > maxTemp {
		return fmt.Errorf("minTemp ne sme biti večji od maxTemp")
	}

	if stMeritev < 1 {
		return fmt.Errorf("stMeritev mora biti vsaj 1")
	}
	postajemap := make(map[string]vreme.Postaja)
	vreme.DodajPostajo(postajemap, "LJ01", "Bavarski_dvor", "LJ")
	vreme.DodajPostajo(postajemap, "LJ02", "Drama", "LJ")
	vreme.DodajPostajo(postajemap, "LJ03", "Ajdovscina", "LJ")
	vreme.DodajPostajo(postajemap, "LJ04", "Stadion", "LJ")
	vreme.DodajPostajo(postajemap, "AR01", "Arktika", "Severni pol")
	vreme.DodajPostajo(postajemap, "SA01", "Sahara", "Timbuktu")
	vreme.DodajPostajo(postajemap, "DC01", "Pentagon", "DC")
	dodajMeritve(postajemap, "LJ01", []float64{3.7, 2.173, 5}, minTemp, maxTemp)
	dodajMeritve(postajemap, "LJ02", []float64{1, 2, 3.98}, minTemp, maxTemp)
	dodajMeritve(postajemap, "LJ03", []float64{69, 6.7, 4.20}, minTemp, maxTemp)
	dodajMeritve(postajemap, "LJ04", []float64{1, 2, 3, 35}, minTemp, maxTemp)
	dodajMeritve(postajemap, "AR01", []float64{-50, -49, -273.15, -35, -44}, minTemp, maxTemp)
	dodajMeritve(postajemap, "SA01", []float64{45, 47.3, 33, 29.99, 50}, minTemp, maxTemp)
	dodajMeritve(postajemap, "DC01", []float64{15, 2}, minTemp, maxTemp)
	vreme.DodajMeritev(postajemap, "SA01", 50, minTemp, maxTemp)
	vreme.DodajMeritev(postajemap, "AR01", -40, minTemp, maxTemp)
	vreme.DodajMeritev(postajemap, "ne_obstaja", -40, minTemp, maxTemp)
	vreme.DodajMeritev(postajemap, "AR01", -51, minTemp, maxTemp)

	vreme.IzpisMeritev(postajemap)
	fmt.Println()
	vreme.IzpisPorocila(postajemap, stMeritev)

	return nil
}

func dodajMeritve(postajemap map[string]vreme.Postaja, oznaka string, meritve []float64, minTemp float64, maxTemp float64) {
	for _, temperatura := range meritve {
		vreme.DodajMeritev(
			postajemap, oznaka, temperatura, minTemp, maxTemp,
		)
	}
}
