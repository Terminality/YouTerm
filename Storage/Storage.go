package Storage

import (
	"log"
	"os/user"
	"path"

	bolt "go.etcd.io/bbolt"
)

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

	dbm.saveFilePath = path.Join(getCurUserHomeDir(), ".terminalize", "youterm.db")

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

		return nil
	})
}

func (dbm *DatabaseManager) Shutdown() {
	masterDBM.database.Close()
}

func SaveChannelInfo() {

}

func LoadChannelInfo(idToLoad string) {

}

func SaveVideoInfo() {

}

func LoadVideoInfo(idToLoad string) {

}

func SaveUserInfo() {

}

func LoadUserInfo(idToLoad string) {

}

func SaveGeneric(bucketName string, idToSave string) {

}

func LoadGeneric(bucketName string, idToLoad string) {

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
