package tasks

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/mergestat/timediff"
	"os"
	"text/tabwriter"
	"time"
)

type item struct {
	Task      string
	Done      bool
	CreatedAt time.Time
}

type TODOs []item

func (t *TODOs) Add(task string) {
	todo := item{
		Task:      task,
		Done:      false,
		CreatedAt: time.Now(),
	}
	*t = append(*t, todo)
}

func (t *TODOs) Complete(id int) error {
	ls := *t
	
	id--
	if id < 0 || id >= len(ls) {
		return errors.New("invalid index")
	}
	ls[id].Done = true

	return nil
}

func (t *TODOs) List() {
	ls := *t
	if len(ls) == 0 {
		return
	}
	// display
	writer := tabwriter.NewWriter(os.Stdout, 0, 4, 2, ' ', tabwriter.TabIndent)
	fmt.Fprintln(writer, "Done\tID\tTask\tTime Added")
	for i, item := range *t {
		i++
		done := "❌"
		if item.Done {
			done = "✅"
		}
		fmt.Fprintf(writer, "%v\t%d\t%s\t%s\n", done, i, item.Task, timediff.TimeDiff(item.CreatedAt))
	}
	writer.Flush()
}

func (t *TODOs) Delete(id int) error {
	ls := *t
	id--
	if id < 0 || id >= len(ls) {
		return errors.New("invalid index")
	}

	*t = append(ls[:id], ls[id+1:]...) // ... are for unpacking the items

	return nil
}

func (t *TODOs) Load(filename string) error {
	file, err := os.ReadFile(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}

	if len(file) == 0 {
		return err
	}

	err = json.Unmarshal(file, t)
	if err != nil {
		return nil
	}

	return nil
}

func (t *TODOs) Store(filename string) error {
	data, err := json.Marshal(t)
	if err != nil {
		return err
	}

	return os.WriteFile(filename, data, 0644)
}
