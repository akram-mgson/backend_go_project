package main

var profile = Info{PhoneNumber: "11111111", Email: "i_ivanov@example.ru", Manager: "Ivan Ivanov"}

var ordersDTO = []OrderListItemDTO{
	{
		ID:               40364,
		Number:           "ЗК-2024-0901",
		Title:            "Насосная станция НС-150",
		StatusCode:       "C1:NEW",
		StatusGroup:      "NEW",
		StatusLabel:      "Новый",
		Date:             "2026-05-21",
		Equipment:        "Насосная станция НС-150",
		Quantity:         1,
		QuotationValidTo: "2026-06-20",
		InvoiceNumber:    "СЧ-123",
		InvoiceValidTo:   "2026-06-20",
	},

	{
		ID:               40365,
		Number:           "ЗК-2024-0902",
		Title:            "Насосная станция НС-151",
		StatusCode:       "C1:NEW",
		StatusGroup:      "NEW",
		StatusLabel:      "Новый",
		Date:             "2026-05-21",
		Equipment:        "Насосная станция НС-151",
		Quantity:         1,
		QuotationValidTo: "2026-06-20",
		InvoiceNumber:    "СЧ-124",
		InvoiceValidTo:   "2026-06-20",
	},

	{
		ID:               40366,
		Number:           "ЗК-2024-0903",
		Title:            "Насосная станция НС-152",
		StatusCode:       "C1:NEW",
		StatusGroup:      "NEW",
		StatusLabel:      "Новый",
		Date:             "2026-05-21",
		Equipment:        "Насосная станция НС-152",
		Quantity:         1,
		QuotationValidTo: "2026-06-20",
		InvoiceNumber:    "СЧ-125",
		InvoiceValidTo:   "2026-06-20",
	},
}

var details = []OrderDetailsDTO{
	{
		ID:               40364,
		Number:           "ЗК-2024-0901",
		Title:            "Насосная станция НС-150",
		StatusCode:       "C1:NEW",
		StatusGroup:      "NEW",
		StatusLabel:      "Новый",
		Equipment:        "Насосная станция НС-150",
		Quantity:         1,
		PaymentPercent:   50,
		Category:         "Оборудование",
		DeliveryType:     "ТК",
		DeliveryAddress:  "Москва, ул. Ленина",
		Consignee:        "Покупатель",
		Payer:            "Покупатель",
		TransportCompany: "СДЭК",
		TransportWaybill: "ТН-123456",
		PublicComment:    "Нет комментариев",
	},

	{
		ID:               40365,
		Number:           "ЗК-2024-0902",
		Title:            "Насосная станция НС-151",
		StatusCode:       "C1:NEW",
		StatusGroup:      "NEW",
		StatusLabel:      "Новый",
		Equipment:        "Насосная станция НС-151",
		Quantity:         1,
		PaymentPercent:   50,
		Category:         "Оборудование",
		DeliveryType:     "ТК",
		DeliveryAddress:  "Москва, ул. Ленина",
		Consignee:        "Покупатель",
		Payer:            "Покупатель",
		TransportCompany: "СДЭК",
		TransportWaybill: "ТН-123456",
		PublicComment:    "Нет комментариев",
	},

	{
		ID:               40366,
		Number:           "ЗК-2024-0902",
		Title:            "Насосная станция НС-151",
		StatusCode:       "C1:NEW",
		StatusGroup:      "NEW",
		StatusLabel:      "Новый",
		Equipment:        "Насосная станция НС-151",
		Quantity:         1,
		PaymentPercent:   50,
		Category:         "Оборудование",
		DeliveryType:     "ТК",
		DeliveryAddress:  "Москва, ул. Ленина",
		Consignee:        "Покупатель",
		Payer:            "Покупатель",
		TransportCompany: "СДЭК",
		TransportWaybill: "ТН-123456",
		PublicComment:    "Нет комментариев",
	},
}

var documents = []DocumentDTO{

	{ID: 1, OrderID: 40366, Title: "Спецификация", Type: "Specification", VisibleForClient: true},
	{ID: 2, OrderID: 40365, Title: "Спецификация", Type: "Specification", VisibleForClient: false},
}

var req = []LoginRequest{
	{Login: "client@example.com", Password: "password"},
}

var clientOrders = map[int][]int{
	1: {40364, 40365, 40366},
}
