package composite

import (
	"fmt"
	"testing"
)

func TestComposite(t *testing.T) {
	monitor1 := Monitor{}
	monitor2 := Monitor{}

	mejaKerja := MejaKerja{
		ProductItem: []Product{
			monitor1, monitor2,
		},
	}

	paketWorkingSpace := PaketWorkingSpace{
		ProductItem: []Product{
			mejaKerja,
		},
	}

	fmt.Println(paketWorkingSpace.TotalPriceProduct())

	fmt.Println(monitor1.TotalPriceProduct())
}
