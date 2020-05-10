package db_test

import (
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	"github.com/keenanhoffman/cars-api/proto"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func createSchema(postgresDB *pg.DB) error {
	for _, model := range []interface{}{(*proto.Car)(nil)} {
		err := postgresDB.CreateTable(model, &orm.CreateTableOptions{
			Temp: true,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func TestDb(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Db Suite")
}
