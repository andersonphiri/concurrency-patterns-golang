package pools

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"sync"
	"time"

	"github.com/google/uuid"
)

type IResource interface {
	GetSerializedResource() ([]byte, error)
}

type IResourceService interface {
	GetSerializedResource() ([]byte, error)
}

type ResourceService struct {
	dbContext IResource
}

func NewResourceService(store IResource) *ResourceService {
	service := &ResourceService{dbContext: store}
	// do stuff here
	return service
}

func (service *ResourceService) GetSerializedResource() ([]byte, error) {
	return service.dbContext.GetSerializedResource()
}

type TestResourceModel struct {
	Id          string
	Version     int
	Description string
	Timestamp   int64
}
type TestResourceDbContext struct{}

func NewTestResourceDbContext() IResource {
	ctx := &TestResourceDbContext{}
	return ctx
}
func (resource TestResourceDbContext) GetSerializedResource() ([]byte, error) {
	item := TestResourceModel{Id: uuid.NewString(), Version: rand.Int(),
		Description: fmt.Sprintf("description with random id: %v", rand.Int()), Timestamp: time.Now().UnixNano()}
	ser, err := json.Marshal(item)
	return ser, err
}

// for example, if your service connects to a mongodb store
// then you can happily create a max of 50 000 connections
func createAServiceConnection() interface{} {
	time.Sleep(2 * time.Second)
	return NewResourceService(NewTestResourceDbContext())
}

func warmServiceConnectionCache(n int) *sync.Pool {
	p := &sync.Pool{
		New: createAServiceConnection,
	}
	for i := 0; i < n; i++ {
		p.Put(p.New())
	}
	return p
}

type ResourceController struct {
	connPool *sync.Pool
	poolSize int
}

func NewController(cacheSize int) *ResourceController {
	ctrl := ResourceController{poolSize: cacheSize}
	return &ctrl
}
func (rcv *ResourceController) Init() {
	rcv.connPool = warmServiceConnectionCache(rcv.poolSize)
}
func (rcv *ResourceController) handleGetResource(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	serviceConnection := rcv.connPool.Get().(IResourceService) // could be any
	defer rcv.connPool.Put(serviceConnection)
	// then use the connection or resource
	data, errFetch := serviceConnection.GetSerializedResource()
	if errFetch != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(errFetch.Error()))
		log.Fatalf("error %v", errFetch.Error())
	} else {
		// fmt.Println(conn,string(data))
		w.WriteHeader(http.StatusOK)
		_, errWrite := w.Write(data)
		if errWrite != nil {
			log.Fatalf("error: %v", errWrite.Error())
		}

	}
	
}

func RunWarmCacheExample(host, port string, cacheSize int) {
	log.Printf("listening on port: %v and on host: %v\n", host, port)
	ctrl := NewController(cacheSize)
	log.Printf("initializing controller...\n")
	ctrl.Init()
	log.Printf("initializing controller completed...\n")
	handlerExample := http.HandlerFunc(ctrl.handleGetResource)
	http.Handle("/resource", handlerExample)
	log.Printf("listening on port: %v and on host: %v\n", host, port)
	err := http.ListenAndServe(fmt.Sprintf("%s:%s", host, port), nil)
	if err != nil {
		log.Fatalf(err.Error())
	}
}
