package sip

import (
	"fmt"
	"math"
)

// SipCalculator provides SIP calculation methods.
type SipCalculator struct{}

// SipReturn holds the results of a SIP calculation.
type SipReturn struct {
	FutureValue    float64
	TotalInvested  float64
	Gains          float64
	ReturnsPercent float64
}

func (s SipCalculator) CalculateSipReturn(monthlyAmount float64, annualReturnRate float64, years float64) (SipReturn, error) {
	if monthlyAmount <= 0 {
		return SipReturn{}, fmt.Errorf("monthly amount must be greater than 0")
	}
	if monthlyAmount > 1e9 {
		return SipReturn{}, fmt.Errorf("monthly amount is unrealistically high")
	}
	if annualReturnRate <= 0 {
		return SipReturn{}, fmt.Errorf("annual return rate must be greater than 0")
	}
	if annualReturnRate > 100 {
		return SipReturn{}, fmt.Errorf("annual return rate must not exceed 100%%")
	}
	if years <= 0 {
		return SipReturn{}, fmt.Errorf("years must be greater than 0")
	}
	if years > 100 {
		return SipReturn{}, fmt.Errorf("years must not exceed 100")
	}

	monthlyRate := annualReturnRate / (12 * 100)
	months := years * 12

	futureValue := monthlyAmount * ((math.Pow(1+monthlyRate, months) - 1) / monthlyRate) * (1 + monthlyRate)

	totalInvested := monthlyAmount * months
	gains := futureValue - totalInvested
	returnsPercent := (gains / totalInvested) * 100

	return SipReturn{
		FutureValue:    futureValue,
		TotalInvested:  totalInvested,
		Gains:          gains,
		ReturnsPercent: returnsPercent,
	}, nil
}
