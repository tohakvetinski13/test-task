package main

import (
	"reflect"
	"strings"
	"testing"
)

var file = `<?xml version="1.0" encoding="utf-8"?>
			<root>
			<item city="Барнаул" street="Дальняя улица" house="56" floor="2" />
			<item city="Братск" street="Большая Октябрьская улица" house="65" floor="5" />
			<item city="Братск" street="Большая Октябрьская улица" house="65" floor="5" />
			<item city="Балаково" street="Барыши, местечко" house="67" floor="2" />
			</root>`

var data = strings.NewReader(file)

var addresses = map[Address]int{
	{City: "Балаково", Street: "Барыши, местечко", House: "67", Floor: 2}:        1,
	{City: "Барнаул", Street: "Дальняя улица", House: "56", Floor: 2}:            1,
	{City: "Братск", Street: "Большая Октябрьская улица", House: "65", Floor: 5}: 2,
}

func TestReadFromXML(t *testing.T) {
	got, _ := ReadFromXML(data)
	want := addresses
	if !reflect.DeepEqual(got, want) {
		t.Errorf("incorrect, got\n %v,\n want\n %v\n", got, want)
	}
}

func TestFindDuplicate(t *testing.T) {

	got := FindDuplicate(addresses)
	want := map[Address]int{
		{City: "Братск", Street: "Большая Октябрьская улица", House: "65", Floor: 5}: 2,
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("incorrect, got\n %v,\n want\n %v\n", got, want)
	}
}

func TestCountHousesInCity(t *testing.T) {

	got := CountHousesInCity(addresses)
	want := map[string]House{
		"Балаково": {Floor2: 1},
		"Барнаул":  {Floor2: 1},
		"Братск":   {Floor5: 1},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("incorrect, got\n %v,\n want\n %v\n", got, want)
	}
}

func BenchmarkReadFromXML(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ReadFromXML(data)
	}
}

func BenchmarkFindDuplicate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FindDuplicate(addresses)
	}
}

func BenchmarkCountHousesInCity(b *testing.B) {
	m1 := FindDuplicate(addresses)
	for i := 0; i < b.N; i++ {
		CountHousesInCity(m1)
	}
}
