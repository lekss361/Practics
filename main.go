package main

import (
	"time"
)

type Product struct {
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

type LogEntry struct {
	IP         string
	Method     string
	URL        string
	StatusCode int
	Size       int
}

type TaskStatus int

type Task struct {
	ID          int
	Title       string
	Description string
	Status      TaskStatus
	CreatedAt   time.Time
	CompletedAt *time.Time
}

type TaskManager struct {
	tasks  []Task
	nextID int
}

const (
	StatusNew TaskStatus = iota
	StatusInProgress
	StatusCompleted
)

// Задача 1: Система учёта товаров в магазине
// Цель: Закрепить переменные, условия, структурыОписание: Создать программу для учёта товаров в магазине. Нужно определить структуру товара, добавить несколько товаров и вывести информацию о дорогих товарах (цена > 1000).
//  // Создаём массив товаров
//     products := []Product{
//         {"Ноутбук", 45000, 5, "Электроника"},
//         {"Хлеб", 50, 20, "Продукты"},
//         {"Телефон", 25000, 3, "Электроника"},
//         {"Молоко", 80, 15, "Продукты"},
//         {"Планшет", 15000, 2, "Электроника"},
//     }

// Задача 2: Калькулятор зарплаты сотрудников
// Цель: Закрепить функции, циклы, условия, ссылкиОписание: Создать систему расчёта зарплаты. Функция должна принимать указатель на структуру сотрудника и рассчитывать итоговую зарплату с учётом премии (если отработано > 160 часов, премия 15%)
//  employees := []Employee{
//         {"Иван Петров", 500, 180, 0},
//         {"Мария Сидорова", 600, 150, 0},
//         {"Алексей Иванов", 450, 170, 0},
//     }

// Задача 3: Система бронирования отелей
// Цель: Комплексная задача на все темыОписание: Создать систему бронирования номеров в отеле. Нужно реализовать функции поиска свободных номеров, бронирования и отмены бронирования.
// hotel := Hotel{
//         Name: "Гранд Отель",
//         Rooms: []Room{
//             {101, "стандарт", 3000, false, ""},
//             {102, "стандарт", 3000, false, ""},
//             {201, "люкс", 8000, false, ""},
//             {202, "люкс", 8000, false, ""},
//             {301, "президентский", 15000, false, ""},
//         },
//     }

// Задача 4: Анализатор логов веб-сервера
// Цель: Работа с циклами, функциями и обработкой данныхОписание: Создать программу для анализа логов веб-сервера. Подсчитать количество запросов по статус-кодам и найти самые частые URL.
//     // Имитация логов
//     logs := []LogEntry{
//         {"192.168.1.1", "GET", "/", 200, 1024},
//         {"192.168.1.2", "GET", "/products", 200, 2048},
//         {"192.168.1.1", "POST", "/login", 401, 512},
//         {"192.168.1.3", "GET", "/", 200, 1024},
//         {"192.168.1.2", "GET", "/products", 404, 256},
//         {"192.168.1.1", "GET", "/about", 200, 1536},
//         {"192.168.1.4", "GET", "/", 200, 1024},
//         {"192.168.1.1", "POST", "/login", 200, 768},
//     }

// Задача 5: Система управления задачами (TODO)
// Цель: Комплексная задача с использованием всех концепцийОписание: Создать простую систему управления задачами с возможностью добавления, выполнения и удаления задач.

func main() {

}
