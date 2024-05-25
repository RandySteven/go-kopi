package enums

type MongoCollection string

const (
	Config             MongoCollection = "config"
	CorexData          MongoCollection = "corexData"
	DepositData        MongoCollection = "depositData"
	FintrustData       MongoCollection = "fintrustData"
	LendingData        MongoCollection = "lendingData"
	LoanChannelingData MongoCollection = "loanChannelingData"
	OnboardingData     MongoCollection = "onboardingData"
	OpsPortalData      MongoCollection = "opsPortalData"
	PaymentData        MongoCollection = "paymentData"
	RewardsData        MongoCollection = "rewardsData"
)

func (coll MongoCollection) ToString() string {
	return string(coll)
}
