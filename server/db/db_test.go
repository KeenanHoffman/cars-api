package db_test

import (
	"github.com/go-pg/pg"
	"github.com/keenanhoffman/cars-api/proto"
	. "github.com/keenanhoffman/cars-api/server/db"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)


var _ = Describe("Db", func() {
	var db *pg.DB
	BeforeEach(func() {
		db = pg.Connect(&pg.Options{
			User: "postgres",
			Database: "cars_unit_test",
		})

		err := createSchema(db)
		Expect(err).ToNot(HaveOccurred())
	})

	Context("CreateCar", func() {
		It("Creates a new car", func() {
			defer db.Close()
			postgresDB := Postgres{
				DB: db,
			}

			car := proto.Car{
				Make:  "test-make",
				Model: "test-model",
				Vin:   "test-vin",
			}
			err := postgresDB.CreateCar(car)
			Expect(err).ToNot(HaveOccurred())

			cars := []*proto.Car{}
			err = db.Model(&cars).Select()
			Expect(err).ToNot(HaveOccurred())

			Expect(len(cars)).To(Equal(1))
			Expect(cars[0].Make).To(Equal("test-make"))
			Expect(cars[0].Model).To(Equal("test-model"))
			Expect(cars[0].Vin).To(Equal("test-vin"))
		})
		It("Returns an error when the db errors", func() {
			postgresDB := Postgres{
				DB: db,
			}

			db.Close()
			err := postgresDB.CreateCar(proto.Car{})
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("pg: database is closed"))
		})
	})

	Context("GetCarById", func() {
		It("Gets a new car when given an ID", func() {
			defer db.Close()
			postgresDB := Postgres{
				DB: db,
			}
			err := db.Insert(&proto.Car{
				Make:  "test-make",
				Model: "test-model",
				Vin:   "test-vin",
			})
			Expect(err).ToNot(HaveOccurred())

			car := proto.Car{}
			err = db.Model(&car).
				Where("car.make = ?", "test-make").
				Select()
			Expect(err).ToNot(HaveOccurred())

			retrievedCar, err := postgresDB.GetCarById(car.Id)
			Expect(err).ToNot(HaveOccurred())

			Expect(retrievedCar.Make).To(Equal("test-make"))
			Expect(retrievedCar.Model).To(Equal("test-model"))
			Expect(retrievedCar.Vin).To(Equal("test-vin"))
		})
		It("Fails when no car is found with a given id", func() {
			defer db.Close()
			postgresDB := Postgres{
				DB: db,
			}

			retrievedCar, err := postgresDB.GetCarById(12345)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("pg: no rows in result set"))
			Expect(retrievedCar.Id).To(Equal(int64(0)))
		})
		It("Returns an error when the db errors", func() {
			postgresDB := Postgres{
				DB: db,
			}

			db.Close()
			_, err := postgresDB.GetCarById(1)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("pg: database is closed"))
		})
	})

	Context("GetCars", func() {
		It("Gets all cars", func() {
			defer db.Close()
			postgresDB := Postgres{
				DB: db,
			}
			err := db.Insert(&proto.Car{
				Make:  "test-make",
				Model: "test-model",
				Vin:   "test-vin",
			})
			Expect(err).ToNot(HaveOccurred())

			cars, err := postgresDB.GetCars()
			Expect(err).ToNot(HaveOccurred())

			Expect(len(cars)).To(Equal(1))
			Expect(cars[0].Make).To(Equal("test-make"))
			Expect(cars[0].Model).To(Equal("test-model"))
			Expect(cars[0].Vin).To(Equal("test-vin"))
		})
		It("Returns an error when the db errors", func() {
			postgresDB := Postgres{
				DB: db,
			}

			db.Close()
			_, err := postgresDB.GetCars()
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("pg: database is closed"))
		})
	})

	Context("UpdateCar", func() {
		It("Updates given fields of a car", func() {
			defer db.Close()
			postgresDB := Postgres{
				DB: db,
			}
			err := db.Insert(&proto.Car{
				Make:  "test-make",
				Model: "test-model",
				Vin:   "test-vin",
			})
			Expect(err).ToNot(HaveOccurred())

			car := proto.Car{}
			err = db.Model(&car).
				Where("car.make = ?", "test-make").
				Select()
			Expect(err).ToNot(HaveOccurred())

			newCarDetails := proto.Car{
				Id: car.Id,
				Make: "new-make",
			}
			err = postgresDB.UpdateCar(newCarDetails)
			Expect(err).ToNot(HaveOccurred())

			cars := []*proto.Car{}
			err = db.Model(&cars).Select()
			Expect(err).ToNot(HaveOccurred())

			Expect(len(cars)).To(Equal(1))
			Expect(cars[0].Make).To(Equal("new-make"))
			Expect(cars[0].Model).To(Equal("test-model"))
			Expect(cars[0].Vin).To(Equal("test-vin"))
		})
		It("Updates given fields of a car given all fields", func() {
			defer db.Close()
			postgresDB := Postgres{
				DB: db,
			}
			err := db.Insert(&proto.Car{
				Make:  "test-make",
				Model: "test-model",
				Vin:   "test-vin",
			})
			Expect(err).ToNot(HaveOccurred())

			car := proto.Car{}
			err = db.Model(&car).
				Where("car.make = ?", "test-make").
				Select()
			Expect(err).ToNot(HaveOccurred())

			newCarDetails := proto.Car{
				Id: car.Id,
				Make: "new-make",
				Model: "new-model",
				Vin: "new-vin",
			}
			err = postgresDB.UpdateCar(newCarDetails)
			Expect(err).ToNot(HaveOccurred())

			cars := []*proto.Car{}
			err = db.Model(&cars).Select()
			Expect(err).ToNot(HaveOccurred())

			Expect(len(cars)).To(Equal(1))
			Expect(cars[0].Make).To(Equal("new-make"))
			Expect(cars[0].Model).To(Equal("new-model"))
			Expect(cars[0].Vin).To(Equal("new-vin"))
		})
		It("Fails when no car is found with a given id", func() {
			defer db.Close()
			postgresDB := Postgres{
				DB: db,
			}

			newCarDetails := proto.Car{
				Id: 12345,
			}
			err := postgresDB.UpdateCar(newCarDetails)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("pg: no rows in result set"))
		})
		It("Returns an error when the db errors", func() {
			postgresDB := Postgres{
				DB: db,
			}

			db.Close()
			err := postgresDB.UpdateCar(proto.Car{})
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("pg: database is closed"))
		})
	})
	Context("ReplaceCar", func() {
		It("Replaces a car with a new car with the same ID", func() {
			defer db.Close()
			postgresDB := Postgres{
				DB: db,
			}
			err := db.Insert(&proto.Car{
				Make:  "test-make",
				Model: "test-model",
				Vin:   "test-vin",
			})
			Expect(err).ToNot(HaveOccurred())

			car := proto.Car{}
			err = db.Model(&car).
				Where("car.make = ?", "test-make").
				Select()
			Expect(err).ToNot(HaveOccurred())

			newCarDetails := proto.Car{
				Id:    car.Id,
				Make:  "new-make",
				Model: "new-model",
				Vin:   "new-vin",
			}
			err = postgresDB.ReplaceCar(newCarDetails)
			Expect(err).ToNot(HaveOccurred())

			cars := []*proto.Car{}
			err = db.Model(&cars).Select()
			Expect(err).ToNot(HaveOccurred())

			Expect(len(cars)).To(Equal(1))
			Expect(cars[0].Make).To(Equal("new-make"))
			Expect(cars[0].Model).To(Equal("new-model"))
			Expect(cars[0].Vin).To(Equal("new-vin"))
		})
		It("Fails when no car is found with a given id", func() {
			defer db.Close()
			postgresDB := Postgres{
				DB: db,
			}

			newCarDetails := proto.Car{
				Id: 12345,
			}
			err := postgresDB.ReplaceCar(newCarDetails)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("pg: no rows in result set"))
		})
		It("Returns an error when the db errors", func() {
			postgresDB := Postgres{
				DB: db,
			}

			db.Close()
			err := postgresDB.ReplaceCar(proto.Car{})
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("pg: database is closed"))
		})
	})
	Context("DeleteCar", func() {
		It("Deletes a car with a given ID", func() {
			defer db.Close()
			postgresDB := Postgres{
				DB: db,
			}
			err := db.Insert(&proto.Car{
				Make:  "test-make",
				Model: "test-model",
				Vin:   "test-vin",
			})
			Expect(err).ToNot(HaveOccurred())

			car := proto.Car{}
			err = db.Model(&car).
				Where("car.make = ?", "test-make").
				Select()
			Expect(err).ToNot(HaveOccurred())

			err = postgresDB.DeleteCar(car.Id)
			Expect(err).ToNot(HaveOccurred())

			cars := []*proto.Car{}
			err = db.Model(&cars).Select()
			Expect(err).ToNot(HaveOccurred())

			Expect(len(cars)).To(Equal(0))
		})
		It("Fails when no car is found with a given id", func() {
			defer db.Close()
			postgresDB := Postgres{
				DB: db,
			}

			err := postgresDB.DeleteCar(12345)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("pg: no rows in result set"))
		})
		It("Returns an error when the db errors", func() {
			postgresDB := Postgres{
				DB: db,
			}

			db.Close()
			err := postgresDB.DeleteCar(1)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("pg: database is closed"))
		})
	})
	Context("SearchCars", func() {
		It("Searches for cars with the given info", func() {
			defer db.Close()
			postgresDB := Postgres{
				DB: db,
			}
			err := db.Insert(&proto.Car{
				Make:  "test-make",
				Model: "test-model",
				Vin:   "test-vin",
			})
			Expect(err).ToNot(HaveOccurred())
			err = db.Insert(&proto.Car{
				Make:  "test-make",
				Model: "other-model",
				Vin:   "other-vin",
			})
			Expect(err).ToNot(HaveOccurred())
			err = db.Insert(&proto.Car{
				Make:  "test-make",
				Model: "another-model",
				Vin:   "another-vin",
			})
			Expect(err).ToNot(HaveOccurred())

			cars, err := postgresDB.SearchCars(proto.Car{
				Make: "test-make",
			})
			Expect(err).ToNot(HaveOccurred())

			Expect(len(cars)).To(Equal(3))
			Expect(cars[0].GetMake()).To(Equal("test-make"))
			Expect(cars[0].GetModel()).To(Equal("test-model"))
			Expect(cars[1].GetMake()).To(Equal("test-make"))
			Expect(cars[1].GetModel()).To(Equal("other-model"))
			Expect(cars[2].GetMake()).To(Equal("test-make"))
			Expect(cars[2].GetModel()).To(Equal("another-model"))
		})
		It("Searches for cars with the model field", func() {
			defer db.Close()
			postgresDB := Postgres{
				DB: db,
			}
			err := db.Insert(&proto.Car{
				Make:  "test-make",
				Model: "test-model",
				Vin:   "test-vin",
			})
			Expect(err).ToNot(HaveOccurred())
			err = db.Insert(&proto.Car{
				Make:  "other-make",
				Model: "test-model",
				Vin:   "other-vin",
			})
			Expect(err).ToNot(HaveOccurred())
			err = db.Insert(&proto.Car{
				Make:  "another-make",
				Model: "test-model",
				Vin:   "another-vin",
			})
			Expect(err).ToNot(HaveOccurred())

			cars, err := postgresDB.SearchCars(proto.Car{
				Model: "test-model",
			})
			Expect(err).ToNot(HaveOccurred())

			Expect(len(cars)).To(Equal(3))
			Expect(cars[0].GetModel()).To(Equal("test-model"))
			Expect(cars[0].GetMake()).To(Equal("test-make"))
			Expect(cars[1].GetModel()).To(Equal("test-model"))
			Expect(cars[1].GetMake()).To(Equal("other-make"))
			Expect(cars[2].GetModel()).To(Equal("test-model"))
			Expect(cars[2].GetMake()).To(Equal("another-make"))
		})
		It("Searches for cars with the vin field", func() {
			defer db.Close()
			postgresDB := Postgres{
				DB: db,
			}
			err := db.Insert(&proto.Car{
				Make:  "test-make",
				Model: "test-model",
				Vin:   "test-vin",
			})
			Expect(err).ToNot(HaveOccurred())
			err = db.Insert(&proto.Car{
				Make:  "other-make",
				Model: "test-model",
				Vin:   "other-vin",
			})
			Expect(err).ToNot(HaveOccurred())
			err = db.Insert(&proto.Car{
				Make:  "another-make",
				Model: "test-model",
				Vin:   "another-vin",
			})
			Expect(err).ToNot(HaveOccurred())

			cars, err := postgresDB.SearchCars(proto.Car{
				Model: "test-model",
			})
			Expect(err).ToNot(HaveOccurred())

			Expect(len(cars)).To(Equal(3))
			Expect(cars[0].GetModel()).To(Equal("test-model"))
			Expect(cars[0].GetMake()).To(Equal("test-make"))
			Expect(cars[1].GetModel()).To(Equal("test-model"))
			Expect(cars[1].GetMake()).To(Equal("other-make"))
			Expect(cars[2].GetModel()).To(Equal("test-model"))
			Expect(cars[2].GetMake()).To(Equal("another-make"))
		})
		It("Searches for cars with many fields", func() {
			defer db.Close()
			postgresDB := Postgres{
				DB: db,
			}
			err := db.Insert(&proto.Car{
				Make:  "test-make",
				Model: "test-model",
				Vin:   "test-vin",
			})
			Expect(err).ToNot(HaveOccurred())

			cars, err := postgresDB.SearchCars(proto.Car{
				Model: "test-model",
				Make: "test-make",
				Vin: "test-vin",
			})
			Expect(err).ToNot(HaveOccurred())

			Expect(len(cars)).To(Equal(1))
			Expect(cars[0].GetModel()).To(Equal("test-model"))
			Expect(cars[0].GetMake()).To(Equal("test-make"))
			Expect(cars[0].GetVin()).To(Equal("test-vin"))
		})
		It("Returns empty when no car is found with the given info", func() {
			defer db.Close()
			postgresDB := Postgres{
				DB: db,
			}

			cars, err := postgresDB.SearchCars(proto.Car{
				Model: "test-model",
			})
			Expect(err).ToNot(HaveOccurred())

			Expect(len(cars)).To(Equal(0))
		})
		It("Returns an error when the db errors", func() {
			postgresDB := Postgres{
				DB: db,
			}

			db.Close()
			_, err := postgresDB.SearchCars(proto.Car{
				Model: "test-model",
			})
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("pg: database is closed"))
		})
	})
})
