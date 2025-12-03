package main

import "fmt"

type Saver interface {
	Save() error
}

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
}

func (c Config) Save() error {
	fmt.Printf("Saving config: %s,  %s\n", c.Host, c.Port)
	return nil
}

func main() {
	var dbConfig Saver = &Config{
		Host:     "localhost",
		Port:     "5432",
		User:     "admin",
		Password: "secret",
	}

	err := dbConfig.Save()
	if err != nil {
		fmt.Println("Error saving config:", err)
	}

	fmt.Println(dbConfig)
}
