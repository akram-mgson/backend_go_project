package main

type OrderListItemDTO struct {
    ID                   int    `json:"id"`
    Number               string `json:"number"`
    Title                string `json:"title"`
    StatusCode           string `json:"status_code"`
    StatusGroup          string `json:"status_group"`
    StatusLabel          string `json:"status_label"`
    Date                 string `json:"date"`
    Equipment            string `json:"equipment"`
    Quantity             int    `json:"quantity"`
    QuotationValidTo     string `json:"quotation_valid_to"`
    InvoiceNumber        string `json:"invoice_number"`
    InvoiceValidTo       string `json:"invoice_valid_to"`
}

type OrderDetailsDTO struct{
	ID				 int	 `json:"id"`
	Number			 string	 `json:"number"`
	Title			 string  `json:"title"`
	StatusCode		 string  `json:"status_code"`
	StatusGroup		 string	 `json:"status_group"`
	StatusLabel		 string	 `json:"status_label"`
	Equipment		 string	 `json:"equipment"`
	Quantity		 int	 `json:"quantity"` 
	PaymentPercent   int	 `json:"payment_percent"`
	Category		 string	 `json:"category"`
	DeliveryType	 string  `json:"Delivery_type"`
	DeliveryAddress  string	 `json:"Delivery_address"`
	Consignee		 string	 `json:"consigne"`
	Payer 			 string  `json:"payer"`
	TransportCompany string	 `json:"transport_company"`
	TransportWaybill string  `json:"transport_waybill"`
	PublicComment    string  `json:"public_comment"`
}

type Info struct {
	PhoneNumber 	string `json:"phone_number"`
	Email       	string `json:"email"`
	Manager     	string `json:"manager"`
}

type ErrorResp struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type SuccessResp struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type Comment struct {
	Text string `json:"text"`
}

type DocumentDTO struct {
	ID      int    `json:"id"`
	OrderID int    `json:"order_id"`
	Title   string `json:"title"`
	Type 	string `json:"type"`
	VisibleForClient bool `json:"visible_for_client"`
}

type Auth struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type LoginRequest struct {
	Login   	string		 `json:"login"`
	Password 	string       `json:"password"`
}

type LogResp struct {
	Token string `json:"token"`
}
