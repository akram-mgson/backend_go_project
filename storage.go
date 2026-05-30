package main

var profile = Info{PhoneNumber: "11111111", Email: "i_ivanov@example.ru", Manager: "Ivan Ivanov"}

var orders = []Order{
	{ID: 40364, Title: "Заказ №1", StageID: "C1:NEW"},
	{ID: 40365, Title: "Заказ №2", StageID: "C1:WON"},
}

// создаем список(срез)
var documents = []Document{
	{ID: 1, Title: "первый документ", Body: "Содержит 1"},
	{ID: 2, Title: "второй документ", Body: "Содержит 2"},
	{ID: 3, Title: "третий документ", Body: "Содержит 3"},
}
