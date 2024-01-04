package _controllers

import (
	"github.com/gin-gonic/gin"
)

// API Design Guide
/* ___________________________________________________________________________*\
 * URL 					 HTTP Verb	|   Action	| Used for			           |
 * ____________________________________________________________________________|
 * /employees 		   |	POST	| (C)reate   | Create a new employee       |
 * /employees/ 		   |	GET		| (R)etrieve | Retrieve all employees      |
 * /employees/:id 	   |	GET		| (R)etrieve | Retrieve a specific employee|
 * /employees/:id 	   |	PUT		| (U)pdate	 | Update a specific employee  |
 * /employees/:id 	   |	DELETE	| (D)elete	 | Delete a specific employee  |
 * /employees/new 	   |	GET		| New		 | Return a new HTML form	   |
 * /employees/:id/edit |	GET		| Edit	     | Return an edit HTML form	   |
 \*___________________________________________________________________________*/

// PublicRoutes registers routes that are accessible without authentication.
// These include routes for user login.
func PublicRoutes(g *gin.RouterGroup) {
	g.GET("/login", LoginGet)
	g.POST("/login", LoginPost)
}

// PrivateRoutes registers routes that require user authentication.
// These include routes for the dashboard, settings, and logout.
func PrivateRoutes(g *gin.RouterGroup) {
	// Routes
	g.GET("/", IndexGet)

	// Employees
	// Create
	//TODO g.POST("/employees", employeesPOST)

	// // Retieve
	// g.GET("/employees", employeesGET)
	// g.GET("/employees/:ID", employeesGET)

	// // Update
	// g.PUT("/employees/:ID", employeesPUT)

	// // Delete
	// // TODO g.DELETE("/employees/:id", employeesDelete)

	// // Forms
	// g.GET("/employees/new", employeesFormGET)
	// g.GET("/employees/:ID/formedit", employeesFormGET)
	// g.GET("/employees/:ID/inlineedit", employeesInlineEditGET)

	// // Hiring
	// g.GET("/hiring", hiringGet)

	// // Onboarding
	// g.GET("/onboarding", onboardingGet)

	// // Time Tracking
	// g.GET("/timetracking", timetrackingGet)

	// // Documents
	// g.GET("/templates", templatesGet)

	// Logout
	g.GET("/logout", LogoutGet)

}
