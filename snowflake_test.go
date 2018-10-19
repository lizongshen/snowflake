package snowflake

import (
	"log"
	"testing"
)

func TestSnowflake(t *testing.T) {
	w, err := NewIdWorker(0)
	if err != nil {
		t.Error(err)
	}

	id, err := w.NextId()
	if err != nil {
		t.Error(err)
	}

	log.Println(id)
}

func TestSnowflakes(t *testing.T) {

	w, err := NewIdWorker(0)
	if err != nil {
		t.Error(err)
	}

	m := make(map[int64]int, 1000)
	for i := 0; i < 1000000; i++ {
		id, err := w.NextId()
		if err != nil {
			t.Error(err)
		}

		m[id] = 0
	}

	log.Println(len(m))

	if len(m) != 1000000 {
		t.Errorf("生成的ID有重复的！")
	}

}
