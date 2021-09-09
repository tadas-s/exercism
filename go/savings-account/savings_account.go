package savings

const FixedInterestRate = 0.05

func GetFixedInterestRate() float32 {
	return FixedInterestRate
}

const DaysPerYear = 365

func GetDaysPerYear() int {
	return DaysPerYear
}

const (
	Jan = iota + 1
	Feb
	Mar
	Apr
	May
	Jun
	Jul
	Aug
	Sep
	Oct
	Nov
	Dec
)

func GetMonth(month int) int {
	return month
}

type AccNo string

const AccountNo AccNo = "XF348IJ"

func GetAccountNumber() AccNo {
	return AccountNo
}
