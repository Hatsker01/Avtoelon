package repo

import str "github.com/Avtoelon/pkg/structs"

type CarsRepoInterface interface{
	CreateCar(car *str.Car)(*str.Car,error)
}