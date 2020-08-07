package main

import (
	"FRSC-Project/api"
	"FRSC-Project/database"
	"FRSC-Project/domain"
	router "FRSC-Project/routes"
	"FRSC-Project/service"
	"fmt"
	"net/http"
	"os"
)

var (
	busDatabase database.Bus       = database.NewBusDB()
	busDomain   domain.BusDomain   = domain.NewBusDomain(busDatabase)
	busService  service.BusService = service.NewBusService(busDomain)
	busApi      api.BusApi         = api.NewBusApi(busService)
	httpRouter  router.Router      = router.NewMuxRouter()
)

func main() {
	port := os.Getenv("PORT")
	defaultPort := ":5001"

	httpRouter.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "up and running")
	})

	httpRouter.POST("/buses/register", busApi.RegisterBus)
	httpRouter.GET("/buses", busApi.GetBuses)

	if port == "" {
		httpRouter.SERVER(defaultPort)
	} else {
		httpRouter.SERVER(":"+port, nil)
	}
}
