package main

import (
	"bufio"
	"fmt"
	"os"
)

type Contact struct {
	Name               string
	ContactInformation string
}

func main() {
	var contacts []Contact = make([]Contact, 0)

	fmt.Println("==========")
	fmt.Println("INPUT CONTACT")
	fmt.Println("==========")

	for {
		var choice int
		fmt.Println("Menu:")
		fmt.Println("1. Lihat Kontak")
		fmt.Println("2. Tambah Kontak")
		fmt.Println("3. Hapus Kontak")
		fmt.Println("4. Quit")
		fmt.Println("Pilih menu: ")
		fmt.Scanln(&choice)

		if choice == 1 {
			if len(contacts) > 0 {
				fmt.Println("\nKontak: ")
				for i, Contact := range contacts {
					fmt.Printf("%d. %s - %s\n", i+1, Contact.Name, Contact.ContactInformation)
				}
			} else {
				fmt.Println("\nKontak kosong. Silahkan tambahkan kontak dulu")
			}
			fmt.Println("")
		} else if choice == 2 {
			fmt.Print("\nNama Kontak: ")
			scanner := bufio.NewScanner(os.Stdin)
			if scanner.Scan() {
				name := scanner.Text()

				fmt.Print("Informasi Kontak: ")
				contactScanner := bufio.NewScanner(os.Stdin)

				if contactScanner.Scan() {
					info := contactScanner.Text()

					contact := Contact{
						Name:               name,
						ContactInformation: info,
					}

					contacts = append(contacts, contact)
					fmt.Println("\nTambah kontak sukses\n")
				}
			}
			fmt.Println("")
		} else if choice == 3 {
			if len(contacts) > 0 {
				var deleteChoise int
				fmt.Println("\nKontak: ")
				for i, Contact := range contacts {
					fmt.Printf("%d. %s - %s\n", i+1, Contact.Name, Contact.ContactInformation)
				}
				fmt.Println("Input nomor yang hendak dihapus: ")
				fmt.Scanln(&deleteChoise)

				contacts = append(contacts[:deleteChoise-1], contacts[deleteChoise:]...)
				fmt.Println("\nHapus kontak sukses")
			} else {
				fmt.Println("\nKontak kosong. Silahkan tambahkan kontak dulu")
			}
			fmt.Println("")
		} else if choice == 4 {
			fmt.Println("Quitting.. Bye bye")
			os.Exit(0)
		} else {
			fmt.Println("Invalid input. Bye bye")
			os.Exit(1)
		}
	}
}
