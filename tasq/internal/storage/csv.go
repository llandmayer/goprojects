package storage

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"reflect"
	"strconv"
	"tasq/internal/entity"
	"time"
)

type csvStorage struct {
	filePath string
}

func newCsvEngine(filePath string) Storage {
	return &csvStorage{
		filePath: filePath,
	}
}

func (s *csvStorage) Load() ([]entity.Task, error) {
	file, err := s.openFile()
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// For learning purpose I'm implementing this manually consider replacing the csv parsing
	// Using "github.com/gocarina/gocsv" gocsv.UnmarshalFile(file, &tasks)
	csvReader := csv.NewReader(file)
	headers, err := s.readHeaders(csvReader)
	if err != nil {
		return nil, err
	}

	return s.readTasks(csvReader, headers)
}

func (s *csvStorage) openFile() (*os.File, error) {
	file, err := os.OpenFile(s.filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return nil, fmt.Errorf("can't open csv file: %v", err)
	}
	return file, nil
}

// func (s *csvStorage) validate(f *csv.Reader) error {
// 	file, err := s.openFile()
// 	if err != nil {
// 		return err
// 	}
// 	_, ok, err := csvlint.Validate(file, '\t', false)
// 	if !ok {
// 		return err
// 	}
// 	return nil
// }

func (s *csvStorage) readHeaders(reader *csv.Reader) (map[string]int, error) {
	headers, err := reader.Read()
	if err != nil {
		return nil, fmt.Errorf("failed to read csv header: %v", err)
	}

	headerIndex := make(map[string]int)
	for i, header := range headers {
		headerIndex[header] = i
	}
	return headerIndex, nil
}

func (s *csvStorage) readTasks(reader *csv.Reader, headers map[string]int) ([]entity.Task, error) {
	var tasks []entity.Task

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("failed to read csv record: %v", err)
		}

		task, err := s.parseTask(record, headers)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (s *csvStorage) parseTask(record []string, headers map[string]int) (entity.Task, error) {
	var task entity.Task
	val := reflect.ValueOf(&task).Elem()
	typ := val.Type()

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := typ.Field(i)
		fieldName := fieldType.Name

		idx, ok := headers[fieldName]
		if !ok {
			return task, fmt.Errorf("record missing field: %v", fieldName)
		}

		fieldValue := record[idx]
		if err := s.setField(field, fieldName, fieldValue); err != nil {
			return task, err
		}
	}

	return task, nil
}

func (s *csvStorage) setField(field reflect.Value, fieldName, fieldValue string) error {
	switch field.Kind() {
	case reflect.Int:
		intValue, err := strconv.ParseInt(fieldValue, 10, 64)
		if err != nil {
			return fmt.Errorf("invalid value for field %s: %v", fieldName, err)
		}
		field.SetInt(intValue)
	case reflect.String:
		field.SetString(fieldValue)
	case reflect.Struct:
		return s.setStructField(field, fieldName, fieldValue)
	default:
		return errors.New("failed to setField, unsupported field type")
	}
	return nil
}

func (s *csvStorage) setStructField(field reflect.Value, fieldName, fieldValue string) error {
	if field.Type() == reflect.TypeOf(time.Time{}) {
		t, err := time.Parse("2006-01-02", fieldValue)
		if err != nil {
			return fmt.Errorf("invalid date for field %s: %v", fieldName, err)
		}
		field.Set(reflect.ValueOf(t))
	}
	return nil
}

func (s *csvStorage) Save(t entity.Task) error {
	file, err := s.openFile()
	if err != nil {
		return fmt.Errorf("can't open csv file: %v", err)
	}
	defer file.Close()

	s.setId(file, &t)
	csvWriter := csv.NewWriter(file)
	if err := csvWriter.Write([]string{strconv.Itoa(t.ID), t.Name, t.Description, t.Status, t.CreatedAt.Format("2006-01-02"), t.UpdatedAt.Format("2006-01-02")}); err != nil {
		return err
	}

	csvWriter.Flush()
	if csvWriter.Error() != nil {
		return errors.New("in Save failed to flush data, missing permissions")
	}

	return nil
}

func (s *csvStorage) Update(t entity.Task) error {
	return nil
}

func (s *csvStorage) setId(f *os.File, t *entity.Task) bool {
	lastId := s.getLastId(f)
	newId := lastId + 1
	t.ID = newId

	return true
}

func (s csvStorage) getLastId(f *os.File) int {
	csvReader := csv.NewReader(f)

	headers, err := s.readHeaders(csvReader)
	if err != nil {
		log.Fatalf("in setId failed to read header: %v", err)
		return 0
	}

	tasks, err := s.readTasks(csvReader, headers)
	if err != nil {
		log.Fatalf("in setId failed to readTasks: %v", err)
		return 0
	}

	if len(tasks) == 0 {
		return 0
	}

	return tasks[len(tasks)-1].ID
}
