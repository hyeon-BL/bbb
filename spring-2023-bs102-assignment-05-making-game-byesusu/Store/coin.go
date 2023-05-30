package main

import (
	"fmt"
)

var myCoin = 400 //보유한 코인 수

func main() {
	cropLimit := 13 // 작물 제한 수
	var cropNumber int
	var remainingCoins int

	fmt.Println("n번 작물까지 살 수 있습니다. 몇 번 작물을 사시겠습니까?")
	fmt.Scanln(&cropNumber)

	if cropNumber <= cropLimit {
		var cropQuantity int
		fmt.Printf("%d번 작물을 몇 개 사시겠습니까?", cropNumber)
		fmt.Scanln(&cropQuantity)

		cropPrice := calculateCropPrice(cropNumber)
		totalPrice := cropPrice * cropQuantity

		if cropPrice <= myCoin {
			remainingCoins = myCoin - totalPrice
			fmt.Printf("%d번 작물을 %d개 심었습니다.", cropNumber, cropQuantity)

			//작물을 심고 코인이 빠져나가는 프로그램 작성해야함
		} else {
			fmt.Println("코인이 부족하여 구매할 수 없습니다.\n", "남은 코인은 ", remainingCoins, "냥 입니다.")
		}

	} else {
		fmt.Printf("죄송합니다. %d번 작물은 구매할 수 없습니다.\n", cropNumber)
	}
}

func calculateCropPrice(cropNumber int) int {
	var cropPrice int
	switch cropNumber {
	case 1:
		cropPrice = 10
	case 2:
		cropPrice = 15
	case 3:
		cropPrice = 25
	case 4:
		cropPrice = 40
	case 5:
		cropPrice = 50
	case 6:
		cropPrice = 60
	case 7:
		cropPrice = 80
	case 8:
		cropPrice = 110
	case 9:
		cropPrice = 200
	case 10:
		cropPrice = 350
	case 11:
		cropPrice = 500
	case 12:
		cropPrice = 750
	case 13:
		cropPrice = 2000

	default:
		cropPrice = 0
	}

	if cropPrice > 0 {
		myCoin -= cropPrice
	}
	return cropPrice
}
