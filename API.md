Описание эндпоинтов:

Метод / Путь
GET /api/cabinet/orders - список заказов
Source: mock
Response body: (ответ в файле storage.go)
Успешный ответ: 
	{	
		ID: 			   40364,
	 	Number: 		   "ЗК-2024-0901", 
		Title: 			   "Насосная станция НС-150",
		StatusCode: 	   "C1:NEW",
		StatusGroup: 	   "NEW",
		StatusLabel: 	   "Новый",
		Date:		       "2026-05-21",
		Equipment:		   "Насосная станция НС-150",
		Quantity:  		    1,
		QuotationValidTo:  "2026-06-20",
		InvoiceNumber:	   "СЧ-123",
		InvoiceValidTo:	   "2026-06-20",
	},

	{	
		ID: 			   40365,
	 	Number: 		   "ЗК-2024-0902", 
		Title: 			   "Насосная станция НС-151",
		StatusCode: 	   "C1:NEW",
		StatusGroup: 	   "NEW",
		StatusLabel: 	   "Новый",
		Date:		       "2026-05-21",
		Equipment:		   "Насосная станция НС-151",
		Quantity:  		    1,
		QuotationValidTo:  "2026-06-20",
		InvoiceNumber:	   "СЧ-124",
		InvoiceValidTo:	   "2026-06-20",

	},

	{
		ID: 			   40366,
	 	Number: 		   "ЗК-2024-0903", 
		Title: 			   "Насосная станция НС-152",
		StatusCode: 	   "C1:NEW",
		StatusGroup: 	   "NEW",
		StatusLabel: 	   "Новый",
		Date:		       "2026-05-21",
		Equipment:		   "Насосная станция НС-152",
		Quantity:  		    1,
		QuotationValidTo:  "2026-06-20",
		InvoiceNumber:	   "СЧ-125",
		InvoiceValidTo:	   "2026-06-20",
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
		ID:				40364,
		Number:	 		"ЗК-2024-0901",
		Title: 			"Насосная станция НС-150",
		StatusCode: 	"C1:NEW",
		StatusGroup: 	"NEW",
		StatusLabel: 	"Новый",
		Equipment: 		"Насосная станция НС-150",
		Quantity: 		1,
		PaymentPercent: 50,
		Category: 		"Оборудование",
		DeliveryType: 	"ТК",
		DeliveryAddress: "Москва, ул. Ленина",
		Consignee: 		"Покупатель",
		Payer: 			"Покупатель",
		TransportCompany: "СДЭК",
		TransportWaybill: "ТН-123456",
		PublicComment: 	  "Нет комментариев",
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
{"code": "ERROR",   "message": "Комментарий не должен быть пустым"}



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

TODO / planned
Метод / Путь
POST /api/cabinet/auth/logout
Source: mock
Response body:
Успешный ответ - {"message": "logged out"}

Ошибки:
Клиент имеет возможность выйти из платформы
{"code": "METHOD_NOT_ALLOWED", "message": "Метод не поддерживается"}
{"code": "UNAUTHORIZED", "message": "Не авторизован"}



//TODO / planned
Метод / Путь
GET /api/cabinet/auth/me
Source: mock
Response body: {"name": "Ivan", "email": "i_ivanov@test.ru", "area": "project-manager"}
Ответ - {"name": "Ivan", "email": "i_ivanov@test.ru", "area": "project-manager"}
Ошибки:
{"code": "METHOD_NOT_ALLOWED", "message": "Метод не поддерживается"}
{"code": "ACCESS_DENIED", "message": "Доступ запрещен"}
{"code": "UNAUTHORIZED", "message": "Не авторизован"}

Клиент имеет доступ к платформе при успешной авторизации

Метод / Путь
GET /api/cabinet/orders/{id}/documents
Auth required: yes
Source: mock
Успешный ответ - [{"ID": 1, "order_id": 40366, "title": "Спецификация", "type": "specification", "VisibleForClient": true}],

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

1. GET /api/cabinet/orders читает данные из Timeweb DB.
2. GET /api/cabinet/orders/{id} читает данные из Timeweb DB.
3. GET /api/cabinet/orders/{id}/documents читает метаданные документов из Timeweb DB.
4. raw_data не отдается frontend-у напрямую.
5. Bitrix24 используется для записи комментариев, файлов, запросов КП/счета и задач менеджеру.
6. Timeweb DB обновляется каждые 30 минут.


