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
			User: "khofh",
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
	})
})
