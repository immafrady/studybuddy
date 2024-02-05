/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"github.com/immafrady/studybuddy/internal/dispatcher"
	"github.com/immafrady/studybuddy/internal/startup"
)

func main() {
	startup.Bootstrap()
	dispatcher.Dispatch()

}
