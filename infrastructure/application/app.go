package application

func NewApp() {
	dependencies := loadAppDependencies()
	setUpRoutes(dependencies)
}
