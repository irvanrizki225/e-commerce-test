package objects

type Params struct {
	Search string `json:"search"`
}

type Product struct {
	ID       	uint           `json:"id"`
	Name     	string         `json:"name"`
	Price    	int            `json:"price"`
	Stock    	int            `json:"stock"`
}

type Cart struct {
	ProductID	uint           `json:"product_id"`
	UserID   	uint           `json:"user_id"`
	Quantity 	int            `json:"quantity"`
}

type CheckOut struct {
	PaymentMethod string `json:"payment_method"`
}

type OrderDetail struct {
	OrderID  	uint           `json:"order_id"`
	ProductName	string         `json:"product_name"`
}

type Order struct {
	ID       		uint           `json:"id"`
	PaymentMethod 	string         `json:"payment_method"`
}