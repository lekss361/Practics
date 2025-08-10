package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type Room struct {
	Number    int
	Type      string
	Price     float64
	IsBooked  bool
	GuestName string
}
type Hotel struct {
	Name  string
	Rooms []Room
}

func main() {
	//Система бронирования отелей
	//Цель: Комплексная задача на все темы
	//Описание: Создать систему бронирования номеров в отеле. Нужно реализовать функции поиска свободных номеров, бронирования и отмены бронирования.
	/*type Product struct {
	  	Name     string
	  	Price    float64
	  	Quantity int
	  	Category string
	  }
	  type Employee struct {
	  	Name        string
	  	HourlyRate  float64
	  	HoursWorked int
	  	Salary      float64
	  }

	  func main() {
	  	products := []Product{
	  		{"Ноутбук", 45000, 5, "Электроника"},
	  		{"Хлеб", 50, 20, "Продукты"},
	  		{"Телефон", 25000, 3, "Электроника"},
	  		{"Молоко", 80, 15, "Продукты"},
	  		{"Планшет", 15000, 2, "Электроника"},
	  	}
	  	for _, product := range products {
	  		if product.Quantity < 5 {
	  			fmt.Println(product)
	  		}
	  	}
	  	//Создать систему расчёта зарплаты. Функция должна принимать
	  	//указатель на структуру сотрудника и рассчитывать итоговую зарплату с
	  	//учётом премии (если отработано > 160 часов, премия 15%).
	  	employees := []Employee{
	  		{"Иван Петров", 500, 180, 0},
	  		{"Мария Сидорова", 600, 150, 0},
	  		{"Алексей Иванов", 450, 170, 0},
	  	}
	  	for _, employee := range employees {
	  		salary(&employee)
	  	}
	  }

	  func salary(emp *Employee) {
	  	Zarplata := (emp.HourlyRate) * float64(emp.HoursWorked)

	  	if emp.HoursWorked > 160 {
	  		emp.Salary = Zarplata + Zarplata*0.15
	  		fmt.Println("премию получили :", emp)
	  	} else {
	  		emp.Salary = Zarplata
	  		fmt.Println("премию не получили :", emp)
	  	}
	*/
	hotel := Hotel{
		Name: "Гранд Отель",
		Rooms: []Room{
			{101, "стандарт", 3000, true, ""},
			{102, "стандарт", 3000, false, ""},
			{201, "люкс", 8000, false, ""},
			{202, "люкс", 8000, false, ""},
			{301, "президентский", 15000, false, ""},
		},
	}

	for {
		fmt.Println("ВЫхотите забронировать номер или отменить бронь?")
		cmd, ok := scan()
		if !ok {
			return
		}
		if cmd == "забронировать" {
			fmt.Println("Свободные номера:")
			booking(hotel)
			fmt.Println("Какой номер Вы хотите забронировать?")
			cmd2, ok := scan()
			if !ok {
				return
			}
			bron(&hotel, cmd2)
			fmt.Println(hotel)
		} else if cmd == "отменить" {
			otmena(&hotel)
			fmt.Println("Ваша бронь отменена!")
			fmt.Println(hotel)
		} else if cmd == "выйти" {
			fmt.Println("Всего доброго")
			return
		} else {
			fmt.Println("Неопозноная команда(")

		}
	}
}
func booking(hotel Hotel) {
	fmt.Println("В отеле ", hotel.Name)
	for _, room := range hotel.Rooms {
		if !room.IsBooked {
			fmt.Println(room)
		}

	}

}
func bron(hotel *Hotel, nomerHotel string) {
	if isInt(nomerHotel) {
		x, err := strconv.Atoi(nomerHotel)
		if err != nil {
		}
		fmt.Println("Укажите Ваше имя")
		cmd, ok := scan()
		if !ok {
			return
		}
		for q, i := range hotel.Rooms {
			if x == i.Number {
				hotel.Rooms[q].IsBooked = true
				hotel.Rooms[q].GuestName = cmd
			}
		}
	}
}
func otmena(hotel *Hotel) {
	fmt.Println("Укажите Ваше имя")

	cmd, ok := scan()
	if !ok {
		return
	}
	for q, i := range hotel.Rooms {
		if cmd == i.GuestName {
			hotel.Rooms[q].IsBooked = false
			hotel.Rooms[q].GuestName = ""
		}
	}
}
func isInt(s string) bool {
	for _, c := range s {
		if !unicode.IsDigit(c) {
			return false
		}
	}
	return true
}

func scan() (string, bool) {
	scanner := bufio.NewScanner(os.Stdin)
	if ok := scanner.Scan(); !ok {
		fmt.Println("Ошибка ввода!")
		return "", false
	}
	text := scanner.Text()
	fields := strings.Fields(text)
	if len(fields) == 0 {
		fmt.Println("Вы ничего не написали")
		return "", false
	}
	return fields[0], true
}
