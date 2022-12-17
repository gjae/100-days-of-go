package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"
	"github.com/emicklei/go-restful"
	_ "github.com/mattn/go-sqlite3"
	"github.com/gjae/railapi/dbutils"
)

// DB Driver visible to whole program
var DB *sql.DB

// TrainResource is the model for holding rail information
type TrainResource struct {
	ID int
	DriverName string
	OperatingStatus bool
}

// StationResource holds information about locations
type StationResource struct {
	ID int 
	Name string 
	OpeningTime time.Time 
	ClosingTime time.Time 
}

// ScheduleResource links both trains and stations
type ScheduleResource struct {
	ID int 
	TrainID int
	StationID int 
	ArrivalTime time.Time 
}


func (t TrainResource) getTrain(request *restful.Request, response *restful.Response) {
	id := request.PathParameter("train-id")

	err := DB.QueryRow("SELECT ID, DRIVER_NAME, OPERATION_STATUS FROM train WHERE id=?", id).Scan(&t.ID, &t.DriverName, &t.OperatingStatus)
	if err != nil {
		log.Println(err)
		response.AddHeader("Content-Type", "text/plain")
		response.WriteErrorString(http.StatusNotFound, "Train could not be found.")
	} else {
		response.WriteEntity(t)
	}
}

func (t TrainResource) createTrain(request *restful.Request, response *restful.Response) {
	log.Println("Body ", request.Request.Body)

	decoder := json.NewDecoder(request.Request.Body)
	var b TrainResource

	err := decoder.Decode(&b)

	log.Println(b.DriverName, b.OperatingStatus)

	log.Println("DB ", DB)
	// Error handling is obvious here, so ommiting

	statement, _ := DB.Prepare("insert into train(DRIVER_NAME, OPERATION_STATUS) values (?, ?)")
	result, err := statement.Exec(b.DriverName, b.OperatingStatus)

	if err == nil {
		newID, _ := result.LastInsertId()
		b.ID = int(newID)
		log.Println(b)
		response.WriteHeaderAndEntity(http.StatusCreated, b)
	} else {
		response.AddHeader("Content-Type", "text/plain")
		response.WriteErrorString(http.StatusInternalServerError, err.Error())
	}

}

func (t TrainResource) RemoveTrain(request *restful.Request, response *restful.Response) {
	id := request.PathParameter("train-id")
	statement, err := DB.Prepare("DELETE FROM train WHERE id=?")
	_, err = statement.Exec(id)

	if err != nil {
		log.Println(err)
		response.AddHeader("Content-Type", "text/plain")
		response.WriteErrorString(http.StatusInternalServerError, err.Error())
	} else {
		response.WriteHeader(http.StatusOK)
	}
}

func (t *TrainResource) Register(container *restful.Container) {
	ws := new(restful.WebService)

	ws.Path("/v1/trains").Consumes(restful.MIME_JSON).Produces(restful.MIME_JSON)
	ws.Route(ws.GET("/{train-id}").To(t.getTrain))
	ws.Route(ws.POST("").To(t.createTrain))
	ws.Route(ws.DELETE("/{train-id}").To(t.RemoveTrain))

	container.Add(ws)
}

func main() {
	var err error
	DB, err = sql.Open("sqlite3", "./railapi.db")

	if err != nil {
		log.Println(err)
		return
	}

	dbutils.Initialize(DB)
	wsContainer := restful.NewContainer()
	wsContainer.Router(restful.CurlyRouter{})
	t := TrainResource{}
	t.Register(wsContainer)
	log.Println("Start listening on localhost:8000")
	server := &http.Server{Addr: ":8000", Handler: wsContainer}
	log.Fatal(server.ListenAndServe())
}