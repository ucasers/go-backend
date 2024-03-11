package test

import (
	"fmt"
	"go-backend/app/utils/config"
	"testing"
)

func TestCreateYamlFactory(t *testing.T) {
	factory1 := config.CreateYamlFactory("config")
	factory2 := config.CreateYamlFactory("gorm")
	fmt.Print(factory1)
	fmt.Print(factory2)
}
