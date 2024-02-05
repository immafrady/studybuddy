/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	homescreen "github.com/immafrady/studybuddy/internal/screens/home"
	"github.com/immafrady/studybuddy/internal/startup"
)

func main() {
	startup.Bootstrap()
	homescreen.Run()

}
