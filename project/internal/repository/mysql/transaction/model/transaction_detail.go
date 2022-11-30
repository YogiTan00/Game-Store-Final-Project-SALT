package model

//type TransactionDetailModel struct {
//	Id              int    `dbq:"id"`
//	TransactionId      int    `dbq:"transaction_id"`
//	ItemId string `dbq:"item_id"`
//}

func GetTableTransactionDetail() string {
	return "transaction_detail"
}

//func TableTransaction() []string {
//		return []string{
//			"id",
//			"transaction_id",
//			"item_id",
//		}
//}
