package tasks

import (
	"encoding/json"
	"errors"
	"github.com/alexeyco/simpletable"
	"github.com/mergestat/timediff"
	"os"
	"strconv"
	"time"
)

func YELLOW(msg string) string {
	return "\u001b[33m" + msg + "\u001b[0m"
}
func CYAN(msg string) string {
	return "\u001b[36m" + msg + "\u001b[0m"
}
func GREEN(msg string) string {
	return "\u001b[32m" + msg + "\u001b[0m"
}

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

	table := simpletable.New()
	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "Done?"},
			{Align: simpletable.AlignCenter, Text: "ID"},
			{Align: simpletable.AlignCenter, Text: "Task"},
			{Align: simpletable.AlignCenter, Text: "Time Added"},
		},
	}
	for i, item := range *t {
		i++
		done := "❌"
		if item.Done {
			done = "✅"
		}
		if item.Done {
			table.Body.Cells = append(table.Body.Cells, []*simpletable.Cell{
				{Align: simpletable.AlignCenter, Text: done},
				{Align: simpletable.AlignCenter, Text: strconv.Itoa(i)},
				{Align: simpletable.AlignLeft, Text: GREEN(item.Task)},
				{Align: simpletable.AlignCenter, Text: GREEN(timediff.TimeDiff(item.CreatedAt))},
			})

		} else {
			table.Body.Cells = append(table.Body.Cells, []*simpletable.Cell{
				{Align: simpletable.AlignCenter, Text: done},
				{Align: simpletable.AlignCenter, Text: strconv.Itoa(i)},
				{Align: simpletable.AlignLeft, Text: CYAN(item.Task)},
				{Align: simpletable.AlignCenter, Text: YELLOW(timediff.TimeDiff(item.CreatedAt))},
			})
		}
		// ("%v\t%d\t%s\t%s\n", done, i, item.Task, timediff.TimeDiff(item.CreatedAt))
	}
	table.SetStyle(simpletable.StyleCompactLite)
	table.Println()

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
