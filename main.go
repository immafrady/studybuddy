/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"github.com/immafrady/studybuddy/cmd"
	"github.com/immafrady/studybuddy/internal/startup"
)

func main() {
	startup.Bootstrap()
	cmd.Execute()
}
