package gae_echo

// reference our echo instance and create it early
var e = createMux()

func Start() {
	e.Logger.Fatal(e.Start(":8080"))
}
