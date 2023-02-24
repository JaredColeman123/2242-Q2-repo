package main

import "fmt"

type building interface{
	doors () int
	windows () int
	roof () string

}

type warehouse struct {
	location string
	colour string
}

func ( w warehouse) doors () int {
	return 4
}

func ( w warehouse) windows () int {
	return 8
}

func ( w warehouse) roof () string {
	return "Zinc"
}

type house struct {
	houseNum int
	location string
	colour string
}

func (h house) doors () int{
	return 2
}

func (h house) windows () int {
	return 15
}

func (h house ) roof () string {
	return "Titanuim"
}


func (h house) bathrooms () int {
	return 3
}

func main (){
	wh := warehouse{
		location: "Belmopan City",
		colour: "Peach",
	}

	hs := house{
		houseNum: 14,
		location: "Belize City",
		colour: "White" ,
	}

	printInfo (wh)
	printInfo (hs)
}

func printInfo ( b building){
	fmt.Println("This building has", b.doors(),"doors along with", b.windows(), "windows, meanwhile the roof is made out of", b.roof() )
}
