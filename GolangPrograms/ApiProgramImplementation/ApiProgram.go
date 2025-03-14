package main

import "fmt"

type driverDetails struct {
	driverId   int
	hourlyRate float64
	totalCost  float64
}

var drivers = make(map[int]*driverDetails)

func addDriver(driverId int, hourlyRate float64) {
	drivers[driverId] = &driverDetails{driverId: driverId, hourlyRate: hourlyRate, totalCost: 0}
}

func recordDelivery(driverId int, starttime int, endtime int) {
	driver, exists := drivers[driverId]
	if !exists {
		return
	}

	// start := time.Unix(int64(starttime),0)
	// end := time.Unix(int64(endtime),0)

	// if end.Before(start) {
	// 	end = end.Add(24 * time.Hour)
	// }

	// duration := end.Sub(start).Minutes()/60.0
	// cost := duration * driver.hourlyRate

	// driver.totalCost += cost
	// drivers[driverId] = driver

	durationSeconds := endtime - starttime
	durationHours := float64(durationSeconds) / 3600.0

	payment := durationHours * driver.hourlyRate
	driver.totalCost += payment

	drivers[driverId] = driver
}

func getTotalCost() float64 {
	total := 0.0

	for _, driver := range drivers {
		total += driver.totalCost
	}
	fmt.Println("Total: ", total)

	return total
}

func main() {
	addDriver(101, 10.0)
	addDriver(102, 20.0)
	addDriver(103, 45.0)

	recordDelivery(101, 5460, 9060)
	recordDelivery(102, 5460, 9060)
	recordDelivery(103, 7200, 12600)

	//fmt.Println(*drivers[103])

	fmt.Println("Total Cost: ", getTotalCost())
}
