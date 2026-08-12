Описание эндпоинтов:

Метод / Путь
GET /api/cabinet/orders - список заказов
Source: mock
Response body: (ответ в файле storage.go)
Успешный ответ: 
	{	
		"id": 			   40364,
	 	"number": 		   "ЗК-2024-0901", 
		"title": 			   "Насосная станция НС-150",
		"status_code": 	   "C1:NEW",
		"status_group": 	   "NEW",
		"status_label": 	   "Новый",
		"date":		       "2026-05-21",
		"equipment":		   "Насосная станция НС-150",
		"quantity":  		    1,
		"quotation_valid_to":  "2026-06-20",
		"invoice_number":	   "СЧ-123",
		"invoice_valid_to":	   "2026-06-20",
	},

	{	
		"id": 			   40365,
	 	"number": 		   "ЗК-2024-0902", 
		"title": 			   "Насосная станция НС-151",
		"status_code": 	   "C1:NEW",
		"status_group": 	   "NEW",
		"status_label": 	   "Новый",
		"date":		       "2026-05-21",
		"equipment":		   "Насосная станция НС-151",
		"quantity":  		    1,
		"quotation_valid_to":  "2026-06-20",
		"invoice_number":	   "СЧ-124",
		"invoice_valid_to":	   "2026-06-20",

	},

	{
		"id": 			   40366,
	 	"number": 		   "ЗК-2024-0903", 
		"title": 			   "Насосная станция НС-152",
		"status_code": 	   "C1:NEW",
		"status_group": 	   "NEW",
		"status_label": 	   "Новый",
		"date":		       "2026-05-21",
		"equipment":		   "Насосная станция НС-152",
		"quantity":  		    1,
		"quotation_valid_to":  "2026-06-20",
		"invoice_number":	   "СЧ-125",
		"invoice_valid_to":	   "2026-06-20",
	},

Ошибка: {"code": "METHOD_NOT_ALLOWED", "message": "Метод не поддерживается"}


Пользовательский интерфейс:

Клиент видит номер заказа, статус и название


Метод / Путь
GET /api/cabinet/orders/{id} - карточка заказов
Auth required: yes
Source: mock
Responce body: 

Успешный ответ:
	{
		"id":				 40364,
		"number":	 		"ЗК-2024-0901",
		"title": 			"Насосная станция НС-150",
		"status_code": 	    "C1:NEW",
		"status_group": 	"NEW",
		"status_label": 	"Новый",
		"equipment": 		"Насосная станция НС-150",
		"quantity": 		 1,
		"payment_percent":   50,
		"category": 		"Оборудование",
		"delivery_type": 	"ТК",
		"delivery_address": "Москва, ул. Ленина",
		"consignee": 		"Покупатель",
		"payer": 			"Покупатель",
		"transport_company": "СДЭК",
		"transport_waybill": "ТН-123456",
		"public_comment": 	  "Нет комментариев",
},


Ошибки:

{"code": "METHOD_NOT_ALLOWED", "message": "Метод не поддерживается"}
{"code": "INVALID_REQUEST", "message": "Неверный запрос"}
{"code": "ORDER_NOT_FOUND", "message": "Заказ не найден"}


Пользовательский интерфейс:

Клиент видит конкретные разделы, такие как, номер, статус, оборудование
Конкретная информация 


Метод / Путь
POST /api/cabinet/orders/{id}/comment
Auth required: yes 
Source: mock
Request body: {"text": "текст комментария"}
Успешный ответ - {"message": "Комментарий отправлен. Данные обновятся после синхронизации."}

Ошибки:
{"code": "METHOD_NOT_ALLOWED", "message": "Метод не поддерживается"}
{"code": "INVALID_REQUEST",    "message": "Некорректный запрос"}
{"code": "ORDER_NOT_FOUND",    "message": "Заказ не найден"}
{"code": "VALIDATION_ERROR",   "message": "Некорректный JSON"}
{"code": "VALIDATION_ERROR",   "message": "Комментарий не должен быть пустым"}



Метод / Путь
POST /api/cabinet/auth/login
Auth required: no
Source: mock
Request body: {"login": "client@example.com", "password": "password"}
Response body: {"token": "fake-token"}

Успешный ответ: 
{"token": "fake-token"}

Ошибки:

{"code": "METHOD_NOT_ALLOWED", "message": "Метод не поддерживается"}
{"code": "VALIDATION_ERROR",  "message": "Неправильный запрос"}
{"code": "INVALID_CREDENTIALS", "message": "Неправильный логин или пароль"}


Метод / Путь
GET /api/cabinet/profile
Auth requiered: yes
Source: mock
Responce body: {"phone_number":"11111111","email":"i_ivanov@example.ru","manager":"Ivan Ivanov"}
Успешный ответ - {"phone_number":"11111111","email":"i_ivanov@example.ru","manager":"Ivan Ivanov"}
Ошибки - METHOD_NOT_ALLOWED (405)
{"code": "METHOD_NOT_ALLOWED", "message": "Метод не поддерживается"}


Клиент видит свои данные


Метод / Путь
POST /api/cabinet/auth/logout
Auth required: да
Source: mock
Response body:
Успешный ответ - {"text": "logged out"}

Ошибки:
Клиент имеет возможность выйти из платформы
{"code": "METHOD_NOT_ALLOWED", "message": "Метод не поддерживается"}
{"code": "UNAUTHORIZED", "message": "Неверный токен"}




Метод / Путь
GET /api/cabinet/auth/me
Auth required: да
Source: mock
Response body: {"name": "Ivan", "email": "i_ivanov@test.com"}
Ответ - {"name": "Ivan", "email": "i_ivanov@test.com"}
Ошибки:
{"code": "METHOD_NOT_ALLOWED", "message": "Метод не поддерживается"}
{"code": "UNAUTHORIZED", "message": "Неверный токен"}

Клиент имеет доступ к платформе при успешной авторизации

Метод / Путь
GET /api/cabinet/orders/{id}/documents
Auth required: yes
Source: mock
Успешный ответ - [{"id": 1, "order_id": 40366, "title": "Спецификация", "type": "specification", "visible_for_client": true}],

Ошибки:
{"code": "METHOD_NOT_ALLOWED", "message": "Метод не поддерживается"}
{"code": "INVALID_REQUEST", "message": "Неверный запрос"}
{"code": "ORDER_NOT_FOUND", "message": "Заказ не найден"}


Формат ошибок возвращается в следующем формате:


{"code": "ORDER_NOT_FOUND", "message": "Заказ не найден"}




Описание модели доступа:

Клиент взаимодействует с contact_id. В данном поле содержатся заказы.
Возвращаются только те заказы, где уникальный номер клиента есть в CONTACT_IDS.
Исходя из этого, клиент видит только свои заказы.


Список методов Bitrix24:

crm.deal.list - список заказов
crm.deal.get - карточки заказов
disk.file.list - файлы с заказами
tasks.task.add - создание задачи менеджеру
user.get - данные менеджера

Источники данных:

Эндпоинты использую mock - данные

- `ordersDTO` — список заказов
- `details` — детали заказов
- `documents` — документы
- `profile` — профиль клиента
- `req` — учётные данные для логина
