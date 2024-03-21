package httpTransport

// TODO: возможно для каждой сущности сделать отдельные роуты
func (http Http) routing() {
	user := http.echo.Group("/user")
	user.POST("/signup", http.userHandler.SingUp)
	user.POST("/signin", http.userHandler.SingIn)
	user.GET("/profile", http.userHandler.GetProfile, http.UserAuthorizationCheck)

	jwt := http.echo.Group("/token")
	jwt.POST("refresh", http.jwtHandler.RefreshAccessToken)
}
