// Filename: main.go
// Project: keepm
// Copyright (C) 2023 RemarQable Software LLC. All Rights Reserved.
//
// Confidentiality Notice: This file contains proprietary information of RemarQable Software LLC.
// It is intended solely for the use within keepm project. Disclosure, copying, distribution, or any other action
// concerning the content of this file is prohibited unless explicitly approved by RemarQable Software LLC.

package main

import (
	_controllers "keepm/_controllers"
	util "keepm/util"
	"log"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

// func init() {
// 	var err error
// 	globals.DB, err = _models.DB()
// 	if err != nil {
// 		log.Fatalf("%v", err)
// 	}
// }

func main() {
	log.Println("Starting keepm...")
	router := gin.Default()

	// Serve static files from the specified directory.
	router.Static("/assets", "./_views/assets")

	// Load HTML templates from the specified glob pattern.
	router.LoadHTMLGlob("_views/templates/**/*.html")

	// Set up cookie-based sessions using a secret from the util package.
	router.Use(sessions.Sessions("session", cookie.NewStore(util.Secret)))

	// Define a group for public routes (no authentication required).
	public := router.Group("/")
	_controllers.PublicRoutes(public)

	// Define a group for private routes (authentication required).
	private := router.Group("/")
	private.Use(_controllers.AuthRequired) // Apply authentication middleware.
	_controllers.PrivateRoutes(private)

	// Start the HTTP server on port 8080.
	router.Run(":8080")
}
