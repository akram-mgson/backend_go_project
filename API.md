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

{"code": "INVALID_REQUEST", "message": "Неверный запрос"}
{"code": "ORDER_NOT_FOUND", "message": "Заказ не найден"}

Пользовательский интерфейс:

Клиент видит конкретные разделы, такие как, номер, статус, оборудование
Конкретная информация 


Метод / Путь
POST /api/cabinet/orders/{id}/comment
Ответ - {"message": "comment sent"}
Ошибки - VALIDATION_ERROR (400) - Неправильный запрос
{"code": "VALIDATION_ERROR", "message": "Неправильный запрос"}


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

Ошибки:
METHOD_NOT_ALLOWED (405) - если не устанавливаем POST
INVALID_CREDENTIALS (401) - неправильный логин или пароль


{"code": "METHOD_NOT_ALLOWED", "message": "Метод не поддерживается"}
{"code": "INVALID_CREDENTIALS", "message": "Неправильный логин и пароль"}


Метод / Путь
GET /api/cabinet/profile
Ответ - {"phone_number":"11111111","email":"i_ivanov@example.ru","manager":"Ivan Ivanov"}
Ошибки - METHOD_NOT_ALLOWED (405)
{"code": "METHOD_NOT_ALLOWED", "message": "Метод не поддерживается"}


Клиент видит свои данные

TODO / planned
Метод / Путь
POST /api/cabinet/auth/logout
Ответ - {"message": "logged out"}
Ошибки:
METHOD_NOT_ALLOWED (405) - если метод не POST
UNAUTHORIZED (401) - нет токена
Клиент имеет возможность выйти из платформы
{"code": "METHOD_NOT_ALLOWED", "message": "Метод не поддерживается"}
{"code": "UNAUTHORIZED", "message": "Не авторизован"}



//TODO / planned
Метод / Путь
GET /api/cabinet/auth/me
Ответ - {"name": "Ivan", "email": "i_ivanov@test.ru", "area": "project-manager"}
Ошибки:
UNAUTHORIZED (401) - Запрос отклонен из-за отсутствия, недействительности, просрочки утечки данных
ACCESS_DENIED (403) - нет прав
METHOD_NOT_ALLOWED (405) - если метод не GET

{"code": "UNAUTHORIZED", "message": "Не авторизован"}
{"code": "ACCESS_DENIED", "message": "Доступ запрещен"}
{"code": "METHOD_NOT_ALLOWED", "message": "Метод не доступен"}


Клиент имеет доступ к платформе при успешной авторизации

Метод / Путь
GET /api/cabinet/orders/{id}/documents
Ответ - [{ID: 1, Title: "Спецификация", Body: "В процессе", OrderID: 40366}]
Ошибки:
INVALID_REQUEST - 400 (если id - не число)
ORDER_NOT_FOUND (отсутствие заказа)
Клиент видит список документов, привязанных к определенному заказу

{"code": "INVALID_REQUEST", "message": "Ошибка валидации"}



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




