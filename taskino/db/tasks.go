package db

import (
	"encoding/json"
	"errors"
	"strconv"
	"time"

	bolt "go.etcd.io/bbolt"
)

var (
	dbPath = "db.db"
	bucket = []byte("tasks")
	db     *bolt.DB
)

type Task struct {
	ID      int
	Content string
}

func Init() {
	var err error

	db, err = bolt.Open(dbPath, 0600, &bolt.Options{Timeout: 1 * time.Second})

	if err != nil {
		panic(err)
	}

	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(bucket)

		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		panic(err)
	}
}

func GetTasks() ([]Task, error) {
	if db == nil {
		return nil, errors.New("database is not initialized")
	}

	var tasks []Task

	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucket)

		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			t := &Task{}
			err := json.Unmarshal(v, t)

			if err != nil {
				return err
			}

			tasks = append(tasks, *t)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func CreateTask(task string) error {
	if db == nil {
		return errors.New("database is not initialized")
	}

	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucket)

		id, _ := b.NextSequence()

		t := Task{
			ID:      int(id),
			Content: task,
		}

		buf, err := json.Marshal(t)

		if err != nil {
			return err
		}

		return b.Put([]byte(strconv.Itoa(t.ID)), buf)
	})
}

func DeleteTask(k string) error {
	if db == nil {
		return errors.New("database is not initialized")
	}

	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucket)

		return b.Delete([]byte(k))
	})
}
