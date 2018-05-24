package router

// appendRoutes registers routes in the router
func (r *Router) appendRoutes() {
	// API
	api := r.e.Group("/api")
	api.Use(r.mwBasicAuth())

	// API, Version 1
	v1 := api.Group("/v1")

	// Endpoints
	v1.POST("/datastore/store_entry", r.c.StoreEntry)
	v1.GET("/datastore/get_entry", r.c.GetEntry)
}
