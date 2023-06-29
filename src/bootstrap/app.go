package bootstrap

type Application struct {
	CamundaConfig *CamundaConfig
}

func App() Application {
	app := &Application{}
	app.CamundaConfig = newCamundaConfig()
	return *app
}
