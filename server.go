/*
    @author   Chandra Agung Rizky
    @version  1.0
    @since
*/
package main

import (
    "github.com/needkopi/api-product/controllers/products"
    "github.com/needkopi/api-product/router"
)

/*
    initialize
    variabel_name       interface                 init_function
*/
var (
    productController   products.Product        = products.ProductController()
    mux                 router.RouteInterface   = router.NewMuxRouter()
)

const (
    port = ":8000"
)

/*
    main function in this program
*/
func main(){
    mux.GET("/",productController.Index)
    mux.SERVE(port)
}

