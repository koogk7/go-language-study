package main

import "fmt"

type Phone struct {
	model string
}

func (p Phone) Call(num string){
	fmt.Println("Ring Ring...", num)
}

func (p Phone) GetModel() string{
	return p.model
}

type Camera struct {
	model string
}

func (c Camera) TakePicture() {
	fmt.Println("Cheese~~~")
}

func (c Camera) GetModel() string{
	return c.model
}

type SmartPhone struct {
	Phone
	Camera
	owner string
}


func main()  {
	myPhone := SmartPhone{
		Phone:  Phone{model: "iphone-x"},
		Camera: Camera{model: "sony-500"},
		owner:  "pnu",
	}

	myPhone.Call("010-1111-4444")
	myPhone.TakePicture()

	// 모호해서 에러가 남
	//fmt.Println(myPhone.GetModel())
	fmt.Println(myPhone.Phone.GetModel())
	fmt.Println(myPhone.Camera.GetModel())
}

