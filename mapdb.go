package mapdb

import "sort"
import "os"
import "encoding/gob"
import "log"

type MapDB struct {
  data map[string]interface{}
  filePath string
}

func Register(value interface{}) {
  gob.Register(value)
}

// Open the MapDB
func OpenDB (dbFile string, regFunction func()) *MapDB {

  if regFunction != nil {
    regFunction()
  }

  reader, err := os.OpenFile(dbFile, os.O_RDONLY | os.O_CREATE, 0644)
  if err != nil {
    log.Fatal(err)
  }
  defer reader.Close()

  decoder := gob.NewDecoder(reader)
  var data map[string]interface{}
  err = decoder.Decode(&data)
  if err != nil {
    data = make(map[string]interface{})
  }

  return &MapDB{
    data: data,
    filePath: dbFile,
  }
}

func (db *MapDB) Close() (err error) {
  return
}

func (db *MapDB) GetFileName() string {
  return db.filePath
}

func (db *MapDB) Set(key string, value interface{}) bool {
  db.data[key] = value
  _, ok := db.data[key]
  return ok
}

func (db *MapDB) Get(key string) interface{} {
  return db.data[key];
}

func (db *MapDB) Del(key string) {
  delete (db.data, key)
}

func (db *MapDB) Keys() []string {
  var keys []string
  for k := range db.data {
    keys = append(keys, k)
  }
  sort.Strings(keys)
  return keys
}

func (db *MapDB) Flush() error {
  return nil
}

func (db *MapDB) Save() (err error) {
  writer, err := os.Create(db.filePath)
  if (err != nil) {
    return
  }
  defer writer.Close()

  encoder := gob.NewEncoder(writer)
  err = encoder.Encode(db.data)
  if err != nil {
    return
  }
  return
}
