package main

func calculateFinalBill(costPerMessage float64, numMessages int) float64 {
	// ?

	baseCost := calculateBaseBill(costPerMessage, numMessages)
	finalCost := baseCost - (calculateDiscountRate(numMessages) * baseCost)
	return finalCost

}

func calculateDiscountRate(messagesSent int) float64 {
	// ?
	if messagesSent > 5000 {
		return 0.20
	}else if messagesSent > 1000{
		return 0.10
	}
	return 0.0
}

// don't touch below this line

func calculateBaseBill(costPerMessage float64, messagesSent int) float64 {
	return costPerMessage * float64(messagesSent)
}
