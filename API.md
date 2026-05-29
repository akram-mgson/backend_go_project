Описание эндпоинтов:
Метод / Путь
GET /api/cabinet/orders - список заказов

Ответ - [{"id":40364,"title":"Заказ №1","stageId":"C1:NEW"},
		{"id":40365,"title":"Заказ №2","stageId":"C1:WON"}]


Пользовательский интерфейс:

Клиент видит номер заказа, статус и название


Метод / Путь
GET /api/cabinet/orders/{id} - карточка заказов

Ответ - {"id":40364,"title":"Заказ №1","stageId":"C1:NEW"}

Ошибки:

INVALID_REQUEST (400) - браузер отправил запрос на сайт с ошибкой
ORDER_NOT_FOUND (404) - страница или файл не найдены

Пользовательский интерфейс:

Клиент видит конкретные разделы, такие как, номер, статус, оборудование
Конкретная информация 


Метод / Путь
POST /api/cabinet/orders/{id}/comment
Ответ - {"message": "comment sent"}
Ошибки - INVALID_REQUEST (400) - Неправильный запрос


Метод / Путь
POST /api/cabinet/auth/login
Request:
{
"login": "[client@example.com](mailto:client@example.com)",
"password": "password"
}
Response:
{
	"token": "fake-token"
}
Errors:

	{
		"code":"INVALID_CREDENTIALS",
		"message": "Неверный логин и пароль"
	}


Ответ - {"token": "fake-token"}
Ошибки - METHOD_NOT_ALLOWED (405) - если не устанавливаем POST


Метод / Путь
GET /api/cabinet/profile

Ответ - {"phone_number":"11111111","email":"i_ivanov@example.ru","manager":"Ivan Ivanov"}

Ошибки - METHOD_NOT_ALLOWED


Клиент видит свои данные


Метод / Путь
POST /api/cabinet/auth/logout
Ответ - {"message": "logged out"}
Ошибки:
METHOD_NOT_ALLOWED (405) - если метод не POST
UNAUTHORIZED (401) - нет токена
Клиент имеет возможность выйти из платформы




Метод / Путь
GET/ api/cabinet/auth/me
Ответ - {"name": "Ivan", "email": "i_ivanov@test.ru", "area": "project-manager"}
Ошибки:
UNAUTHORIZED (401) - Запрос отклонен из-за отсутствия, недействительности, просрочки утечки данных
ACCESS_DENIED (403) - нет прав
METHOD_NOT_ALLOWED (405) - если метод не GET

Клиент имеет доступ к платформе при успешной авторизации

Метод / Путь
GET/api/cabinet/orders/{id}/documents
Ответ - {"message": "document is not available"}
Ошибки - VALIDATION_ERROR (ошибка валидации)
Клиент видит список документов, привязанных к определенному заказу



Формат ошибок:


	{
	"code": "ORDER_NOT_FOUND",
	"message": "Заказ не найден"
	}



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




