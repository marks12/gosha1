package generator

import (
	"gosha/webapp/types"
	"math/rand"
)

func GenSelfUpdate() types.SelfUpdate {

	return types.SelfUpdate{
		Id:   rand.Intn(100500),
		IsSuccess: (rand.Intn(100500) % 2 > 0),
		//SelfUpdate remove this line for disable generator functionality
	}
}

func GenListSelfUpdate() (list []types.SelfUpdate) {

	for i:=0; i<rand.Intn(5) + 2; i++{
		list = append(list, GenSelfUpdate())
	}

	return
}
