package main

func main()  {
	stationManager := &stationManager{
		isPlatformFree: true,
	}
	passengerTrain := &passengerTrain{
		mediator: stationManager,
	}
	freightTrain := freightTrain{
		mediator: stationManager,
	}
	passengerTrain.arrive()
	freightTrain.arrive()
	passengerTrain.depart()
}
