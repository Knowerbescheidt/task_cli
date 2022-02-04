package db

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/boltdb/bolt"
)

//package level variables normally not such a good practice
var taskBucket = []byte("tasks")
var db *bolt.DB

type Task struct {
	Key    int
	Value  string
	Status string
}

//not an init function jsut an exportet function
func Init(dbPath string) error {
	var err error
	// nächste zeile wird nicht declared da db package level ist
	db, err = bolt.Open(dbPath, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		fmt.Println("Got here")
		return err
	}
	// closure function in function which also returns an error
	return db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(taskBucket)
		return err
	})
}

func CreateTask(task string) (int, error) {

	//durch das aufrufen und definieren der funktion hier können wir direkt die id vergeben vorteil an closure :)
	var id int
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		id64, _ := b.NextSequence()
		id = int(id64)
		key := itob(id)
		t := Task{Key: id, Value: task, Status: "Not Done"}
		taskMarsh, err := json.Marshal(t)
		if err != nil {
			fmt.Println("An error during marshelling occured")
		}
		return b.Put(key, []byte(taskMarsh))
	})
	if err != nil {
		return -1, err
	}
	return id, err
}

func DeleteTask(tid int) error {

	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		return b.Delete(itob(tid))
	})
}

func UpdateTask(t Task) error {
	t.Status = "Done"
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		taskMarsh, err := json.Marshal(t)
		if err != nil {
			fmt.Println("An error during marshelling occured")
		}
		return b.Put(itob(t.Key), []byte(taskMarsh))
	})
}

func AllTasks() ([]Task, error) {
	var tasks []Task
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			var t Task
			err := json.Unmarshal(v, &t)
			if err != nil {
				fmt.Println(err)
			}
			tasks = append(tasks, t)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

//Turning integers to byte slices
func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

// func btoi(b []byte) int {
// 	return int(binary.BigEndian.Uint64(b))
// }

func CloseDB() {
	err := db.Close()
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
}
