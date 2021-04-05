package stats

import (
	"github.com/Alexanders95/bank/pkg/bank/types"
	"fmt"
)

func ExampleAvg(){
	payments := []types.Payment{
		{
			ID:2,
			Amount:53_00,
			Category: "Cat",
		},
		{
			ID:1,
			Amount:51_00,
			Category: "Cat",
		},
		{
			ID:3,
			Amount:52_00,
			Category: "Cat",
		},
	}

	fmt.Println(Avg(payments))  
	//Output: 5200
}

func ExampleTotalInCategory() {
	payments := []types.Payment{
		{
			ID: 1,
			Amount: 55_00,
			Category: "Books",
		},{
			ID: 2,
			Amount: 57_00,
			Category: "Shop",
		},{
			ID: 3,
			Amount: 57_00,
			Category: "Shop",
		},{
			ID: 4,
			Amount: 58_00,
			Category: "Shop",
		},{
			ID: 5,
			Amount: 95_00,
			Category: "Books",
		},
	}

	fmt.Println(TotalInCategory(payments, "Books"))
	//Output:
	// 15000
}