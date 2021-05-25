package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

type Address struct {
	City   string `xml:"city,attr"`
	Street string `xml:"street,attr"`
	House  string `xml:"house,attr"`
	Floor  int    `xml:"floor,attr"`
}

type House struct {
	Floor1 int
	Floor2 int
	Floor3 int
	Floor4 int
	Floor5 int
}

// Add() метод добавляет +1 к количеству n-этажных домов в структуре House
func (h House) Add(n int) House {
	switch n {
	case 5:
		h.Floor5 += 1
	case 4:
		h.Floor4 += 1
	case 3:
		h.Floor3 += 1
	case 2:
		h.Floor2 += 1
	case 1:
		h.Floor1 += 1

	}
	return h
}

func main() {
	var Now time.Time = time.Now()
	printLogo()
	arguments := os.Args
	if len(arguments) == 1 {
		log.Fatal("Пожайлуста проверьте имя файла!")
		return
	}

	file, err := os.Open(arguments[1])
	if err != nil {
		panic(err)
	}

	defer file.Close()

	v, err := ReadFromXML(file)
	if err != nil {
		log.Fatal(err)
	}

	//поиск дубликатов
	dup := FindDuplicate(v)
	printDuplicate(dup) //вывод на печать

	//Кол-во н этажных домов в каждом городе
	count := CountHousesInCity(v)
	printHousesInCity(count) //вывод на печать

	//Общее кол-во уникальных записей(без дублей)
	entries := len(v)
	printEntries(entries)
	fmt.Println("Time work:", time.Since(Now))
}

// printHousesCity печатает количество домов в городе этажностью  1,2,3,4,5
func printHousesInCity(m1 map[string]House) {
	hr := "========================================================================================================\n"
	print(hr)
	header := "Город:"
	header2 := "Кол-во зданий этажностью 1,2,3,4,5:"
	s := fmt.Sprintf("| %20v | %77v |\n", header, header2)
	print(s)
	print(hr)
	for k, v := range m1 {
		fmt.Printf("| %20v | 1-Этаж: %4v | 2-Этажа: %4v | 3-Этажа: %4v | 4-Этажа: %4v | 5-Этажей: %4v |\n", k, v.Floor1, v.Floor2, v.Floor3, v.Floor4, v.Floor5)
	}
	print(hr)
	print("\n \n")
}

// печать
func print(s string) {
	fmt.Print(s)
}
func printLogo() {
	stat := []string{
		`╔═══╗╔════╗╔═══╗╔════╗`,
		`║╔═╗║║╔╗╔╗║║╔═╗║║╔╗╔╗║`,
		`║╚══╗╚╝║║╚╝║║─║║╚╝║║╚╝`,
		`╚══╗║──║║──║╚═╝║──║║──`,
		`║╚═╝║──║║──║╔═╗║──║║──`,
		`╚═══╝──╚╝──╚╝─╚╝──╚╝──`}
	for _, v := range stat {
		fmt.Println(v)
	}
}
func printEntries(e int) {
	hr := "===================================================================\n"
	print(hr)
	header := "Количество уникальных записей:"
	s := fmt.Sprintf("| %40v | %20v |\n", header, e)
	print(s)
	print(hr)
	print("\n \n")
}

// printDoublicate печатает дублирущиеся записи и их кол-во в таблицу
func printDuplicate(m map[Address]int) {
	hr := "===================================================================\n"
	print(hr)
	header := "Дублирующиеся записи:"
	header2 := "Кол-во дублей:"
	s := fmt.Sprintf("| %40v | %20v |\n", header, header2)
	print(s)
	print(hr)
	for k, v := range m {
		s := fmt.Sprintf(" %v %v %v %v", k.City, k.Street, k.House, k.Floor)
		s1 := fmt.Sprintf("| %40s | %20v |\n", s, v)
		print(s1)
	}
	print(hr)
	print("\n \n")
}

// countHousesInCity считает кол-во домов этажностью 1,2,3,4,5 в каждом городе
func CountHousesInCity(v map[Address]int) map[string]House {
	m := make(map[string]House, len(v))
	for key, _ := range v {
		m[key.City] = m[key.City].Add(key.Floor)
	}
	return m
}

// Поиск дубликатов
func FindDuplicate(m map[Address]int) map[Address]int {
	k := make(map[Address]int)
	for key, value := range m {
		if value > 1 {
			k[key] = value
		}
	}
	return k
}

// Чтение файла XML
func ReadFromXML(file io.Reader) (map[Address]int, error) {

	decoder := xml.NewDecoder(file)

	// Чтение item по частям
	m := make(map[Address]int)

	for {

		tok, err := decoder.Token()
		if err != nil && err != io.EOF {
			return nil, err
		} else if err == io.EOF {
			break
		}

		switch tp := tok.(type) {
		case xml.StartElement:
			if tp.Name.Local == "item" {

				// Декодирование элемента в структуру
				var b Address
				err = decoder.DecodeElement(&b, &tp)
				if err != nil {
					return nil, err
				}

				m[b] += 1

			}
		}
	}
	return m, nil
}
