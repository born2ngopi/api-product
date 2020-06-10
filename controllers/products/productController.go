package products

import (
    "net/http"
    "encoding/json"

    //"github.com/needkopi/api-product/services/productService"
    "github.com/needkopi/api-product/repositories/productRepo"
)

type Product interface{
    Index(w http.ResponseWriter, r *http.Request)
}

type productController struct{}

func ProductController() Product {
    return &productController{}
}

/*
    constructor
    name_variable   interface                           init_finction
*/
var (
    //productService  productService.ProductInterface     = productService.NewProductService()
    productRepo     productRepo.RepoInterface           = productRepo.NewProductRepo();
}

func (*productController) Index(w http.ResponseWriter, r *http.Request){
    w.Write([]byte("muantap bos"))
}

/*
    function for get all product
    @param  response
    @param  request
*/
func (*productController) FetchAll(w http.ResponseWriter, r *http.Request){
    w.Header().Set("Content-type","application/json")
    data, err := productRepo.GetAll()
    if( err != nil ){
        w.WriteHeader(http.StatusInternalServerError)
        w.Write([]byte(`{"error" : "error get all product"}`))
    }
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(data)

}
