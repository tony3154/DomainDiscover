package model

type Domain struct {
	ID                int
	DomainName        string
	Project           string
	Service           string
	CDN               string
	HTTPS             string
	Backend           string
	WhiteList         string
	WhiteListLocation string
	Notes             string
}
