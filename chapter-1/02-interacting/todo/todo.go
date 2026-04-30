package todo

import (
	"encoding/json"
	"errors"
	"os"
	"time"
)

type item struct {
	Task        string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

type List []item

func (l *List) Add(task string) {
	i := item{task, false, time.Now(), time.Now()}
	*l = append(*l, i)
}

func (l *List) Complete(i int) error {
	ls := *l
	if i <= 0 || i > len(ls) {
		return errors.New("invalid item")
	}

	ls[i-1].Done = true
	ls[i-1].CompletedAt = time.Now()

	return nil
}
func (l *List) Delete(i int) error {
	ls := *l
	if i <= 0 || i > len(ls) {
		return errors.New("invalid item")
	}

	*l = append(ls[:i-1], ls[i:]...) //Slice trick
	// go nao tem delete para splice, apenas para map

	ls[i-1].Done = true
	ls[i-1].CompletedAt = time.Now()

	return nil
}

func (l *List) Save(fileName string) error {
	marshal, err := json.Marshal(l)
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, marshal, 0644)
}

func (l *List) Get(fileName string) error {
	file, err := os.ReadFile(fileName)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return errors.New("file doesn't exist")
		}
		return err
	}

	if len(file) == 0 {
		return nil
	}
	return json.Unmarshal(file, l)
}
