package Storage

import (
	// "fmt"
	"os"
	"os/user"
	"path"
	"time"

	"github.com/charmbracelet/log"
	bolt "go.etcd.io/bbolt"
)

// String Constants
const (
	SAVE_DIR  string = ".terminality"
	SAVE_FILE string = "youterm.db"

	USERS    = "Users"
	LISTS    = "Lists"
	CHANNELS = "Channels"
	VIDEOS   = "Videos"
)

type Resource interface {
	GetID() string
	GetBucketName() string
	MarshalData() ([]byte, error)
}

var masterDBM *DatabaseManager

func Startup() {
	masterDBM = &DatabaseManager{}
	masterDBM.Initialize()
}

func Shutdown() {
	masterDBM.Shutdown()
}

type DatabaseManager struct {
	database     *bolt.DB
	saveFilePath string
}

func (dbm *DatabaseManager) Initialize() {
	log.Debug("Trying to initialize database...")

	dbm.saveFilePath = path.Join(getCurUserHomeDir(), SAVE_DIR, SAVE_FILE)
	os.MkdirAll(path.Join(getCurUserHomeDir(), SAVE_DIR), os.ModePerm)

	log.Debug("Trying to access save file path", "path", dbm.saveFilePath)

	tempDB, err := bolt.Open(dbm.saveFilePath, 0644, &bolt.Options{Timeout: 10 * time.Second}) // 0644 indicates user R/W, group and other R
	log.Debug("Database opened")

	checkErr(err)
	dbm.database = tempDB

	dbm.initializeBuckets()
	log.Debug("Buckets initialized")
}

func (dbm *DatabaseManager) initializeBuckets() {
	// Set up buckets
	dbm.database.Update(func(tx *bolt.Tx) error {
		tx.CreateBucketIfNotExists([]byte(USERS))
		tx.CreateBucketIfNotExists([]byte(LISTS))
		tx.CreateBucketIfNotExists([]byte(CHANNELS))
		tx.CreateBucketIfNotExists([]byte(VIDEOS))

		return nil
	})
}

func (dbm *DatabaseManager) Shutdown() {
	masterDBM.database.Close()
	log.Debug("Database closed")
}

func InitializeUser(ID string) {
	masterDBM.database.Update(func(tx *bolt.Tx) error {
		tx.CreateBucketIfNotExists([]byte(USERS))
		//bucket := tx.Bucket([]byte("Users"))
		//userBucket, err := bucket.CreateBucket([]byte(ID))
		return nil
	})
}

func SaveResource(resource Resource) {
	// fmt.Println("Saving resource", resource)
	masterDBM.database.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(resource.GetBucketName()))
		resourceData, err := resource.MarshalData()
		checkErr(err)
		bucket.Put([]byte(resource.GetID()), resourceData)
		return nil
	})

}

func LoadResource(bucketName string, idToLoad string) []byte {
	// fmt.Println("Loading resource", bucketName, idToLoad)
	var output []byte
	masterDBM.database.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketName))
		output = bucket.Get([]byte(idToLoad))
		return nil
	})
	return output
}

func getCurUserHomeDir() string {
	curUser, err := user.Current()
	checkErr(err)

	homeDir := curUser.HomeDir
	if homeDir == "" {
		log.Fatal("Couldn't load a user home directory")
	}

	return homeDir
}

func checkErr(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
