package main

import (
	"path/filepath"

	"acakp.task/cmd"
	"acakp.task/db"
	"github.com/mitchellh/go-homedir"
)

// - [x] add
// - [x] clear
// - [x] list
// - [x] del
// - [x] do
// - [ ] REFACTOR CODE IT'S NEED TO BE PRETTY
// чтобы реализовать do, мне нужно будет сильно поменять то, как
// программа взаимодейсвует с заметками.
// сейчас моя бд - это просто id в key и строка с задачей в value.
// я собираюсь сделать, чтобы tasksList bucket в key хранил тот же id,
// а в value хранил другой bucket, который будет состоять из Task struct.
// см. draw.io
// ... или просто добавлять string в начало с "- [ ]", и добавлять х туда,
// когда нужно

func main() {
	home, _ := homedir.Dir()
	dbPath := filepath.Join(home, "tasks.db")
	db.Init(dbPath)
	cmd.Execute()
}
