package productRepo

import (
    "database/sql"

    "github.com/needkopi/api-product/entity"
    _ "github.com/go-sql-driver/mysql"
)

type ProductInterface interface{
    GetAll() ([]entity.Product, error)
}

type productRepo struct{}

func NewProductRepo() ProductInterface {
    return &productRepo{}
}

func connect() (db *sql.DB){
	dbDriver := "mysql"
    dbUser := "root"
    dbPass := "example"
    dbName := "api-products"
    db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
    if err != nil {
        panic(err.Error())
    }
    return db
}

 

func (*productRepo) GetAll() ([]entity.Product, error){
    db := connect()
    defer db.Close()

    data, err := db.Query("SELECT * FROM products")
    if (err != nil){
        return nil,err
    }
    var products []entity.Product
    for data.Next(){

		var id , price int
        var name string

        err = selDB.Scan(&id, &name, &price)
        if err != nil {
            return nil, err
        }
        item := entity.Product{
            Id      : id,
            Name    : name,
            Price   : price,
        }
        products = append(products,item)
    }

    return products, nil
}
