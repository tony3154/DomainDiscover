package model

type Domain struct {
	ID                int
	DomainName        string
	Project           string
	Service           string
	CDN               string
	SslState          bool
	WhiteList         string
	WhiteListLocation string
}
