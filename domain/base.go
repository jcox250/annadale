package domain

type Base struct {
	BaseData BaseData
	PageData interface{}
}

type BaseData struct {
	LatestResults []string
	NextFixtures  []string
}
