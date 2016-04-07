package data

import (
	"database/sql"
	"fmt"
	"log"
	"upper.io/db"
	_ "upper.io/db/postgresql"
)

type Access struct {
	Table      string
	Collection db.Collection
	Condition  map[string]interface{}
	Result     db.Result
	Sess       db.Database
	Rows       *sql.Rows
}

var adapter = `postgresql`

func NewAccess(tableName string) *Access {
	var err error
	var access = Access{Table: tableName}
	var settings = NewConfig().ConnectionString()
	if access.Sess, err = db.Open(adapter, settings); err != nil {
		panic(fmt.Sprintf(`Data.Access() Error: %v`, err))
	}

	if access.Collection, err = access.Sess.Collection(access.Table); err != nil {
		panic(fmt.Sprintf(`data.Init(-%v-) Error: %v`, access.Table, err))
	}
	return &access
}

func (self *Access) Save(entity interface{}) (iface interface{}, err error) {
	if iface, err = self.Collection.Append(entity); err != nil {
		log.Printf("Access.Save() Error: %v", err)
	}
	return iface, err
}

func (self *Access) Seek() {
	self.Result = self.Collection.Find(self.Condition)
}

func (self *Access) Search(cond ...interface{}) {
	self.Result = self.Collection.Find(cond)
}
func (self *Access) Update(model interface{}, cond db.Cond) (err error) {
	result := self.Collection.Find(cond)
	if err = result.Update(model); err != nil {
		log.Printf("Access.Update() Error: %v", err.Error())
	}
	return err
}

func (self *Access) Delete(cond db.Cond) (err error) {
	result := self.Collection.Find(cond)
	if err := result.Remove(); err != nil {
		log.Printf("Access.Remove() Error: %v", err.Error())
	}
	return err
}

func (self *Access) Remove(cond ...interface{}) (err error) {
	result := self.Collection.Find(cond)
	if err := result.Remove(); err != nil {
		log.Printf("Access.Delete() Error: %v", err.Error())
	}
	return err
}
