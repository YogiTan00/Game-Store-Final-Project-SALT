package model

type TransactionModel struct {
	Id              int    `dbq:"id"`
	CustomerId      int    `dbq:"customer_id"`
	CodeTransaction string `dbq:"code_transaction"`
}

func GetTableTransaction() string {
	return "transactions"
}

func TableTransaction() []string {
	return []string{
		"id",
		"customer_id",
		"code_transaction",
	}
}
