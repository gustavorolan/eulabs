package impl

import (
	"eulabs/src/main/core/dto"
	"eulabs/src/main/core/product"
	"eulabs/src/main/core/utils"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
	"time"
)

type ServiceImpl struct {
	Repository product.Repository
}

var loggerInfo = log.New(os.Stdout, "[INFO] ", log.Ldate|log.Ltime|log.Lshortfile)

var loggerError = log.New(os.Stdout, "[ERROR] ", log.Ldate|log.Ltime|log.Lshortfile)

func (s ServiceImpl) Create(request *dto.NewProductRequest) dto.Response {

	productEntity := product.NewProductRequestToEntity(request)

	productEntitySaved, err := s.Repository.Create(&productEntity)

	if err != nil {
		return dto.NewErrorResponseInternalServerError(err.Error())
	}

	loggerInfo.Println(`Product created successfully id: ` + strconv.Itoa(productEntitySaved.ID))

	return dto.NewSuccessResponseCreated(productEntitySaved.ToResponse())
}

func (s ServiceImpl) GetById(id string) dto.Response {
	products, err := s.Repository.FindById(id)

	loggerInfo.Println(`Consulting product by id: ` + id)

	if err != nil {
		loggerError.Println(err.Error())
		return dto.NewErrorResponseInternalServerError(err.Error())
	} else if len(products) > 1 {
		message := "More than one entity has returned for this id: " + id
		loggerError.Println(message)
		return dto.NewErrorResponseInternalServerError(message)
	} else if len(products) == 0 {
		message := "Product was not found id: " + id
		loggerError.Println(message)
		return dto.NewErrorResponseNotFound(message)
	}

	return dto.NewSuccessResponseOk(products[0].ToResponse())
}

func (s ServiceImpl) Update(request *dto.UpdateProductRequest) dto.Response {

	productEntity := product.UpdateProductRequestToEntity(request)

	productEntityUpdated, err := s.Repository.Update(&productEntity)

	if err != nil {
		loggerError.Println(err.Error())
		return dto.NewErrorResponseInternalServerError(err.Error())
	}

	loggerInfo.Println(`Product updated successfully id: ` + strconv.Itoa(productEntityUpdated.ID))

	return dto.NewSuccessResponseOk(productEntityUpdated.ToResponse())
}

func (s ServiceImpl) Delete(id string) dto.Response {

	_, err := s.Repository.DeleteById(id)

	if err != nil {
		loggerError.Println(err.Error())
		return dto.NewErrorResponseInternalServerError(err.Error())
	}

	response := map[string]interface{}{
		"message": "A Product Product has been deleted id: " + id,
	}

	loggerInfo.Println(response)

	return dto.NewSuccessResponseOk(response)
}

func (s ServiceImpl) GetAll(pageable *dto.Pageable) dto.Response {

	loggerInfo.Println("Consulting all products")

	page, err := s.Repository.FindAll(pageable)

	if err != nil {
		loggerError.Println(err.Error())
		return dto.NewErrorResponseInternalServerError(err.Error())
	}

	pageResponse := product.MapToProductResponseList(page)

	return dto.NewSuccessResponseOk(pageResponse)
}

func (s ServiceImpl) CreateManyWithChanel() dto.Response {
	saveDbChanel := make(chan *product.Product)
	printChanel := make(chan *product.Product)

	wg := sync.WaitGroup{}

	go sendData(saveDbChanel, &wg)
	go sendData(printChanel, &wg)

	for i := 0; i < 20; i++ {
		wg.Add(i)
		go s.receiveData(saveDbChanel, &wg)
	}

	go s.receiveDataSout(printChanel, &wg)

	wg.Wait()

	return dto.NewSuccessResponseCreated(nil)
}

func (s ServiceImpl) receiveDataSout(chanel chan *product.Product, wg *sync.WaitGroup) {
	defer wg.Done()
	wg.Add(1)
	for {
		data, ok := <-chanel
		if !ok {
			fmt.Println("Channel closed")
			return
		}
		fmt.Println(data)
	}

}

func (s ServiceImpl) receiveData(chanel chan *product.Product, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		data, ok := <-chanel
		if !ok {
			fmt.Println("Channel closed")
			return
		}
		_, _ = s.Repository.Create(data)
	}

}

func sendData(productChanel chan *product.Product, wg *sync.WaitGroup) {
	defer wg.Done()
	wg.Add(1)
	i := 0
	for i < 10000 {
		p := product.Product{
			Description: strconv.Itoa(i),
			Name:        strconv.Itoa(i),
		}

		productChanel <- &p
		i++
	}
	defer close(productChanel)
}

func (s ServiceImpl) CreateMany() dto.Response {
	products := make([]product.Product, 0)
	i := 0
	for i < 10000 {
		p := product.Product{
			Description: strconv.Itoa(i),
			Name:        strconv.Itoa(i),
		}

		products = append(products, p)
		i++
	}

	s.createMany(products)

	response := utils.Map(products,
		func(p product.Product) int {
			return p.ID
		})

	return dto.NewSuccessResponseOk(response)
}

func (s ServiceImpl) createMany(products []product.Product) {

	var wg sync.WaitGroup

	chunks := utils.ChunkBy(products, 30)

	for index, ps := range chunks {
		wg.Add(1)
		go s.createList(ps, index, &wg)
	}
	wg.Wait()
}

func (s ServiceImpl) createList(ps []*product.Product, index int, wg *sync.WaitGroup) {

	fmt.Printf(time.Now().String()+" Iniciando processamento do chunck: %v\n", index)

	err := s.Repository.CreateBatch(ps)
	if err != nil {
		panic(err)
	}

	fmt.Printf(time.Now().String()+" Finalizando processamento do chunck: %v\n", index)

	defer wg.Done()
}
