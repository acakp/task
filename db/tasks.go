package db

import (
	"encoding/binary"
	"fmt"
	"strings"
	"time"

	"github.com/boltdb/bolt"
)

var db *bolt.DB

func Init(dbPath string) error {
	var err error
	db, err = bolt.Open(dbPath, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		panic(err)
	}
	return db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("tasks"))
		return err
	})
}

func AddTask(task string) {
	err := db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("tasks"))
		if err != nil {
			fmt.Println("Failed create a bucket")
			panic(err)
		}
		tasksBucket := tx.Bucket([]byte("tasks"))
		id, _ := tasksBucket.NextSequence()
		task = "[ ] " + task
		return tasksBucket.Put(itob(id), []byte(task))
	})
	if err != nil {
		fmt.Println("Failed update a bucket in AddTask")
		panic(err)
	}
}

func itob(v uint64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, v)
	return b
}

func ShowTasks() {
	db.View(func(tx *bolt.Tx) error {
		tasksBucket := tx.Bucket([]byte("tasks"))
		if tasksBucket == nil {
			fmt.Println("No tasks (tasksBucket doesn't exists)")
			return nil
		}
		count := 1
		tasksBucket.ForEach(func(k, v []byte) error {
			fmt.Printf("%v. %v\n", count, string(v))
			count++
			return nil
		})
		if count == 1 {
			fmt.Println("Task list is empty")
		}
		return nil
	})
}

func Clear() {
	db.Update(func(tx *bolt.Tx) error {
		err := tx.DeleteBucket([]byte("tasks"))
		if err != nil {
			fmt.Println("No tasks (tasksBucket doesn't exists)")
		} else {
			fmt.Println("The task list has been cleared")
		}
		return err
	})
}

func taskAction(taskNumbers []int, action func(b *bolt.Bucket, k, v []byte)) {
	db.Update(func(tx *bolt.Tx) error {
		tasksBucket := tx.Bucket([]byte("tasks"))
		if tasksBucket == nil {
			fmt.Println("No tasks (tasksBucket doesn't exists)")
			return nil
		}
		count := 0
		tasksBucket.ForEach(func(k, v []byte) error {
			count++
			if anyOf(count, taskNumbers) {
				action(tasksBucket, k, v)
			}
			return nil
		})
		return nil
	})
}

func Del(taskNumbers []int) {
	taskAction(taskNumbers, func(b *bolt.Bucket, k, v []byte) {
		b.Delete(k)
	})
}

func Do(taskNumbers []int) {
	taskAction(taskNumbers, func(b *bolt.Bucket, k, v []byte) {
		dog := string(b.Get(k))
		b.Put(k, []byte(replaceAtIndex(dog, 'x', 1)))
	})
}

func replaceAtIndex(input string, replacement byte, index int) string {
	return strings.Join([]string{input[:index], string(replacement), input[index+1:]}, "")
}

func anyOf(num int, nums []int) bool {
	contain := false
	for _, i := range nums {
		if i == num {
			contain = true
		}
	}
	return contain
}
