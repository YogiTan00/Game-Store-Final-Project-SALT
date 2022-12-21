package test_data

import "game-store-final-project/project/domain/entity/customer"

func GetTestDataCustomer() *customer.Customer {
	customer, _ := customer.NewCustomer(customer.DTOCustomer{
		Nik:          "3204223423442582",
		Nama:         "Taupik Pirdian",
		Alamat:       "Bandung",
		NoTlp:        "085846342122",
		JenisKelamin: "LK",
	})

	return customer
}

func GetTestDataCountCustomer(count int) []*customer.Customer {
	listCustomer := make([]*customer.Customer, 0)
	for i := 1; i <= count; i++ {
		customer, _ := customer.NewCustomer(customer.DTOCustomer{
			Nik:          "3204223423442582",
			Nama:         "Taupik Pirdian",
			Alamat:       "Bandung",
			NoTlp:        "085846342122",
			JenisKelamin: "LK",
		})
		listCustomer = append(listCustomer, customer)
	}
	return listCustomer
}
