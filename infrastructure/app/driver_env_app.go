package app

import "github.com/joho/godotenv"

type DriverApp struct {
	SecretKey string
	Password  string
}

func readEnv() (DriverApp, error) {
	envApp, err := godotenv.Read()
	if err != nil {
		return DriverApp{}, err
	}
	return DriverApp{
		SecretKey: envApp["SECRET_KEY"],
	}, nil
}

func NewDriverApp() (DriverApp, error) {
	dataEnv, err := readEnv()
	if err != nil {
		return DriverApp{}, err
	}
	return dataEnv, nil
}
