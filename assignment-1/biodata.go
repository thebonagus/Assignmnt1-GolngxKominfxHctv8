package main

import (
	"fmt"
	"os"
	"strconv"
)

type Classmate struct {
	nama      string
	alamat    string
	Pekerjaan string
	alasan    string
}

var classmate = []Classmate{
	{
		nama:      "Adam",
		alamat:    "Merkurius",
		Pekerjaan: "President",
		alasan:    "Want to clean code",
	},
	{
		nama:      "Budi",
		alamat:    "Venus",
		Pekerjaan: "Commander",
		alasan:    "Want to refactors",
	},
	{
		nama:      "Candra",
		alamat:    "Bumi",
		Pekerjaan: "Blacksmits",
		alasan:    "Want to debuging",
	},
	{
		nama:      "Danu",
		alamat:    "Mars",
		Pekerjaan: "Farmer",
		alasan:    "Want to code",
	},
	{
		nama:      "Erik",
		alamat:    "Jupiter",
		Pekerjaan: "Trader",
		alasan:    "Want to rich",
	},
	{
		nama:      "Febri",
		alamat:    "Saturnus",
		Pekerjaan: "Athlete",
		alasan:    "Want to health",
	},
	{
		nama:      "Guntur",
		alamat:    "Uranus",
		Pekerjaan: "God of Olympus",
		alasan:    "Want to Endorse like Zeus",
	},
	{
		nama:      "Hasan",
		alamat:    "Neptunus",
		Pekerjaan: "Sailor",
		alasan:    "Want to pirate king",
	},
	{
		nama:      "Ibrahim",
		alamat:    "Pluto",
		Pekerjaan: "Vetteranian",
		alasan:    "Want to go heaven's cat's",
	},
	{
		nama:      "Jefri",
		alamat:    "Nibiru",
		Pekerjaan: "Alien",
		alasan:    "Want to colonize",
	},
}

func main() {
	printClassmate(os.Args[1])
}

func error() {
	fmt.Println("Nomor yang diinput salah, input kembali dalam rentang 1 s/d 10")
}

func printClassmate(numb string) {
	number, _ := strconv.Atoi(numb)
	if number > 0 && number <= len(classmate) {
		fmt.Printf("Data : %+v\n", classmate[number-1])
	} else {
		error()
	}
}
