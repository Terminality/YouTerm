package Storage

import (
	"fmt"
	"log"
	"os"
	"os/user"
	"path"

	bolt "go.etcd.io/bbolt"
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

	dbm.saveFilePath = path.Join(getCurUserHomeDir(), ".terminality", "youterm.db")
	os.MkdirAll(path.Join(getCurUserHomeDir(), ".terminality"), os.ModePerm)

	tempDB, err := bolt.Open(dbm.saveFilePath, 0644, nil) // 0644 indicates user R/W, group and other R
	checkErr(err)
	dbm.database = tempDB

	dbm.initializeBuckets()
}

func (dbm *DatabaseManager) initializeBuckets() {
	// Set up buckets
	dbm.database.Update(func(tx *bolt.Tx) error {
		tx.CreateBucketIfNotExists([]byte("Channels"))
		tx.CreateBucketIfNotExists([]byte("Videos"))
		tx.CreateBucketIfNotExists([]byte("Lists"))
		tx.CreateBucketIfNotExists([]byte("Users"))

		return nil
	})
}

func (dbm *DatabaseManager) Shutdown() {
	masterDBM.database.Close()
}

func SaveResource(resource Resource) {
	fmt.Println("Saving resource", resource)
	masterDBM.database.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(resource.GetBucketName()))
		resourceData, err := resource.MarshalData()
		checkErr(err)
		bucket.Put([]byte(resource.GetID()), resourceData)
		return nil
	})

}

func LoadResource(bucketName string, idToLoad string) []byte {
	fmt.Println("Loading resource", bucketName, idToLoad)
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