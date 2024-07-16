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

	// Buckets
	USERS    = "Users"
	LISTS    = "Lists"
	CHANNELS = "Channels"
	VIDEOS   = "Videos"
)

// Interface for storable resource
type Resource interface {
	GetID() string                // ID is used as key for storage
	GetBucketName() string        // Returns the name of the bucket the resource is stored in
	MarshalData() ([]byte, error) // Returns the resources data in a savable format
}

var masterDBM *DatabaseManager

// Create and initialize master database manager
func Startup() {
	masterDBM = &DatabaseManager{}
	masterDBM.Initialize()
}

// Ensure Master DBM gets shutdown
func Shutdown() {
	masterDBM.Shutdown()
}

// Database Manager
type DatabaseManager struct {
	database     *bolt.DB
	saveFilePath string
}

// Initializes the DBM
func (dbm *DatabaseManager) Initialize() {
	log.Debug("Trying to initialize database...")

	// Load save file path and ensure it exists
	dbm.saveFilePath = path.Join(getCurUserHomeDir(), SAVE_DIR, SAVE_FILE)
	os.MkdirAll(path.Join(getCurUserHomeDir(), SAVE_DIR), os.ModePerm)

	log.Debug("Trying to access save file path", "path", dbm.saveFilePath)

	// Open database. Read/Write for user, Read for group & other
	tempDB, err := bolt.Open(dbm.saveFilePath, 0644, &bolt.Options{Timeout: 10 * time.Second})
	log.Debug("Database opened")

	checkErr(err)
	dbm.database = tempDB

	dbm.initializeBuckets()
	log.Debug("Buckets initialized")
}

// Ensure all buckets exist so they can assuredly be loaded later on
func (dbm *DatabaseManager) initializeBuckets() {
	dbm.database.Update(func(tx *bolt.Tx) error {
		tx.CreateBucketIfNotExists([]byte(USERS))
		tx.CreateBucketIfNotExists([]byte(LISTS))
		tx.CreateBucketIfNotExists([]byte(CHANNELS))
		tx.CreateBucketIfNotExists([]byte(VIDEOS))

		return nil
	})
}

// Ensure database is properly closed
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

// Save resource to database
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

// Load resource from database by ID
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

func ClearBucket(bucketName string) {
	masterDBM.database.Update(func(tx *bolt.Tx) error {
		err := tx.DeleteBucket([]byte(bucketName))
		if err == nil {
			return err
		}
		tx.CreateBucketIfNotExists([]byte(bucketName))
		return nil
	})
}

// Get the home directory of the current user
func getCurUserHomeDir() string {
	curUser, err := user.Current()
	checkErr(err)

	homeDir := curUser.HomeDir
	if homeDir == "" {
		log.Fatal("Couldn't load a user home directory")
	}

	return homeDir
}

// Utility function to check errors
func checkErr(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
