package Model

import "awesomeProject/Model/User"

var FakeDB = []User.SUser{
	{
		Id:   1,
		Age:  20,
		Name: "Pedro",
		Wallet: User.SWallet{
			Money: 0,
			Cards: []User.SCard{
				{
					Bank:   "Inter",
					Secret: 123,
				},
				{
					Bank:   "C6",
					Secret: 321,
				},
			},
		},
	},
}
