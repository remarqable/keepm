// Filename: main.go
// Project: keepm
// Copyright (C) 2023 RemarQable Software LLC. All Rights Reserved.
//
// Confidentiality Notice: This file contains proprietary information of RemarQable Software LLC.
// It is intended solely for the use within keepm project. Disclosure, copying, distribution, or any other action
// concerning the content of this file is prohibited unless explicitly approved by RemarQable Software LLC.

package main

import (
	"embed"
	"html/template"
	"io/fs"
	_controllers "keepm/_controllers"
	util "keepm/util"
	"log"
	"net/http"

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

//go:embed _views/templates/**/*
var templatesEmbed embed.FS

//go:embed _views/assets/**/* _views/assets/*
var staticEmbed embed.FS

func main() {
	log.Println("Starting keepm...")
	router := gin.Default()

	// Serve static files from the staticEmbed filesystem.
	staticFS, _ := fs.Sub(staticEmbed, "_views/assets")
	router.StaticFS("/assets", http.FS(staticFS))

	templ := template.Must(template.New("").ParseFS(templatesEmbed, "_views/templates/**/*.html"))
	router.SetHTMLTemplate(templ)

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
