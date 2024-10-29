package main

import (
	. "fmt"
	. "os"
)

type Coffee struct {
	mWater     int
	mMilk      int
	mBeanGrams int
	mCups      int
	mMoney     int
	mIceCubes  int
}

const (
	waterEsprReq    int = 250
	iceCubesReq     int = 7
	waterFreddoReq  int = 35
	beanGramEsprReq int = 16
	costEspresso    int = 4
	costFreddo      int = 5

	latteWaterReq int = 350
	latteMilkReq  int = 75
	icedMilkReq   int = 200
	latteBeanReq  int = 20
	costLatte     int = 7
	costIcedLatte int = 8

	cappuWaterReq      int = 200
	cappuMilkReq       int = 100
	cappuBeanReq       int = 12
	icedFroth          int = 25
	costCappuccino     int = 6
	costIcedCappuccino int = 8
)

func main() {
	// write your code here
	coffeeMachine := Coffee{
		mWater:     400,
		mMilk:      540,
		mBeanGrams: 120,
		mCups:      9,
		mMoney:     550,
		mIceCubes:  21,
	}
	coffeeMachine.promptChoice()
	Println()
}

func (cm *Coffee) showStatus() {
	Println("\nThe coffee machine has:")
	Printf("%d ml of water\n", cm.mWater)
	Printf("%d ml of milk\n", cm.mMilk)
	Printf("%d g of coffee beans\n", cm.mBeanGrams)
	Printf("%d disposable cups\n", cm.mCups)
	Printf("%d ice cubes\n", cm.mIceCubes)
	Printf("$%d of money\n", cm.mMoney)
	Println()
	cm.promptChoice()
}

func (cm *Coffee) promptChoice() {
	var option string
	Println("Write action (buy, fill, take, remaining, exit):")
	Scan(&option)

	switch option {
	case "buy":
		cm.promptBuy()
	case "fill":
		cm.refill()
	case "take":
		cm.takeMoney()
	case "remaining":
		cm.showStatus()
	case "exit":
		Exit(0)
	default:
		cm.promptChoice()
	}
}

func (cm *Coffee) promptBuy() {
	var coffeeType string
	Println("\nWhat do you want to buy? 1 - espresso, 2 - latte, 3 - cappuccino, back - to main menu:")
	Scan(&coffeeType)

	switch coffeeType {
	case "1":
		cm.buyEspresso()
	case "2":
		cm.buyLatte()
	case "3":
		cm.buyCappuccino()
	case "back":
		cm.promptChoice()
	default:
		cm.promptBuy()
	}

	cm.promptChoice()
}

func (cm *Coffee) buyEspresso() {
	switch {
	case cm.mWater < waterEsprReq:
		Println("Sorry, not enough water!")
	case cm.mBeanGrams < beanGramEsprReq:
		Println("Sorry, not enough beans!")
	case cm.mCups < 1:
		Println("Sorry, not enough cups!")
	default:
		var hotOrIce string
		Println("Would you like your espresso hot or cold?")
		Scan(&hotOrIce)
		switch hotOrIce {
		case "hot":
			Println("I have enough resources, making you a coffee!")
			cm.mWater -= waterEsprReq
			cm.mBeanGrams -= beanGramEsprReq
			cm.mCups--
			cm.mMoney += costEspresso
		case "cold":
			if cm.mIceCubes < iceCubesReq {
				Println("Sorry, not enough ice cubes!")
			} else {
				Println("I have enough resources, making you a coffee!")
				cm.mWater -= waterFreddoReq
				cm.mBeanGrams -= beanGramEsprReq
				cm.mIceCubes -= iceCubesReq
				cm.mCups--
				cm.mMoney += costFreddo
			}
		}
	}

	Println()
	cm.promptChoice()
}

func (cm *Coffee) buyLatte() {
	switch {
	case cm.mWater < latteWaterReq:
		Println("Sorry, not enough water!")
	case cm.mMilk < latteMilkReq:
		Println("Sorry, not enough milk!")
	case cm.mBeanGrams < latteBeanReq:
		Println("Sorry, not enough beans!")
	case cm.mCups < 1:
		Println("Sorry, not enough cups!")
	case cm.mIceCubes < iceCubesReq:
		Println("Sorry, not enough ice cubes!")
	default:
		var hotOrIce string
		Println("Would you like your cappuccino hot or cold?")
		Scan(&hotOrIce)
		switch hotOrIce {
		case "hot":
			Println("I have enough resources, making you a coffee!")
			cm.mWater -= latteWaterReq
			cm.mBeanGrams -= latteBeanReq
			cm.mMilk -= latteMilkReq
			cm.mCups--
			cm.mMoney += costLatte
		case "cold":
			if cm.mIceCubes < iceCubesReq {
				Println("Sorry, not enough ice cubes!")
			} else {
				Println("I have enough resources, making you a coffee!")
				cm.mBeanGrams -= beanGramEsprReq
				cm.mIceCubes -= iceCubesReq
				cm.mMilk -= icedMilkReq + 25 + icedFroth
				cm.mCups--
				cm.mMoney += costIcedLatte
			}
		}
	}

	Println()
	cm.promptChoice()
}

func (cm *Coffee) buyCappuccino() {

	switch {
	case cm.mWater < cappuWaterReq:
		Println("Sorry, not enough water!")
	case cm.mMilk < cappuMilkReq:
		Println("Sorry, not enough milk!")
	case cm.mBeanGrams < cappuBeanReq:
		Println("Sorry, not enough beans!")
	case cm.mCups < 1:
		Println("Sorry, not enough cups!")
	default:
		var hotOrIce string
		Println("Would you like your cappuccino hot or cold?")
		Scan(&hotOrIce)
		switch hotOrIce {
		case "hot":
			Println("I have enough resources, making you a coffee!")
			cm.mWater -= cappuWaterReq
			cm.mBeanGrams -= cappuBeanReq
			cm.mMilk -= cappuMilkReq
			cm.mCups--
			cm.mMoney += costLatte
		case "cold":
			if cm.mIceCubes < iceCubesReq {
				Println("Sorry, not enough ice cubes!")
			} else {
				Println("I have enough resources, making you a coffee!")
				cm.mBeanGrams -= beanGramEsprReq
				cm.mIceCubes -= iceCubesReq
				cm.mMilk -= icedMilkReq
				cm.mCups--
				cm.mMoney += costIcedLatte
			}
		}
	}

	Println()
	cm.promptChoice()
}

func (cm *Coffee) refill() {
	var cups, water, milk, beanGrams, iceCubes int

	Println("\nWrite how many ml of water you want to add:")
	Scan(&water)
	Println("Write how many ml of milk you want to add:")
	Scan(&milk)
	Println("Write how many grams of coffee beans you want to add:")
	Scan(&beanGrams)
	Println("Write how many ice cubes you want to add:")
	Scan(&iceCubes)
	Println("Write how many disposable cups you want to add:")
	Scan(&cups)

	cm.mWater += water
	cm.mMilk += milk
	cm.mBeanGrams += beanGrams
	cm.mIceCubes += iceCubes
	cm.mCups += cups
	Println()
	cm.promptChoice()
}

func (cm *Coffee) takeMoney() {
	if cm.mMoney == 0 {
		Println("Nothing to give you...")
	}
	Printf("\nI gave you $%d\n", cm.mMoney)
	cm.mMoney -= cm.mMoney
	Println()
	cm.promptChoice()
}
