package main

import (
	"bufio"
	"database/sql"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var products []Products

type Products struct {
	Stt         int
	Name        string
	Price       string
	Brand       string
	Inventory   int
	Description string
	Url         string
}

type Message struct {
	Name         string `json:"name"`
	Email        string `json:"email"`
	Phone_number string `json:"phone"`
	Message      string `json:"message"`
}

type Products_DB struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Summary string `json:"summary"`
	Image   string `json:"image"`
	Price   int    `json:"price"`
	Number  int    `json:"number"`
}

var (
	queryProductSQL     = "SELECT * FROM products WHERE id='%s'"
	createProductSQL    = "INSERT INTO products VALUES ('%d', '%s', '%s', '%s', '%d', '%d')"
	queryAllProductsSQL = "SELECT * FROM products"
	updateProductSQL    = "UPDATE products SET title='%s', summary='%s', image='%s', price='%d', number='%d' WHERE id='%d'"
)

func updateProduct(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			template := template.Must(template.ParseFiles("updateProduct.html"))

			template.Execute(w, nil)
			return
		}

		body, err := ioutil.ReadAll(r.Body)
		log.Println(string(body))

		var product Products_DB
		er := json.Unmarshal(body, &product)
		if er != nil {
			log.Println(er.Error())
			return
		}

		db.Exec(fmt.Sprintf(updateProductSQL,
			product.Title,
			product.Summary,
			product.Image,
			product.Price,
			product.Number,
			product.Id,
		))
		log.Println(product)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(400)
		}
		w.WriteHeader(200)
	}
}

func createProduct(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			template := template.Must(template.ParseFiles("addProduct.html"))

			template.Execute(w, nil)
			return
		}

		body, err := ioutil.ReadAll(r.Body)

		var product Products_DB
		er := json.Unmarshal(body, &product)
		if er != nil {
			log.Println(er.Error())
			return
		}
		product.Image = "/Products/" + product.Image
		log.Println(product)

		db.Exec(fmt.Sprintf(createProductSQL,
			product.Id,
			product.Title,
			product.Summary,
			product.Image,
			product.Price,
			product.Number,
		))
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(400)
		}
		w.WriteHeader(200)
	}
}

func queryProduct(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var product Products_DB
		var id = r.FormValue("id")
		err := db.QueryRow(fmt.Sprintf(queryProductSQL, id)).Scan(
			&product.Id,
			&product.Title,
			&product.Summary,
			&product.Image,
			&product.Price,
			&product.Number,
		)
		if err != nil {
			log.Println(err.Error())
			return
		}
		log.Println(product)

		json, _ := json.Marshal(product)

		w.WriteHeader(200)
		w.Write(json)
	}
}

func printProductsDB(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var products_DB []Products_DB

		rows, err := db.Query(queryAllProductsSQL)

		if err != nil {
			log.Println("Select products from database fail: ", err.Error())
			return
		}

		for rows.Next() {
			var product Products_DB
			err := rows.Scan(&product.Id, &product.Title, &product.Summary, &product.Image, &product.Price, &product.Number)

			if err != nil {
				log.Println("Khong the scan san pham!", err.Error())
				return
			}
			products_DB = append(products_DB, product)
		}

		tmp, err := template.ParseFiles("listproduct.html")
		if err != nil {
			log.Println("Loi: ", err.Error())
			return
		}
		template := template.Must(tmp, err)

		if e := template.Execute(w, products_DB); e != nil {
			log.Println("Khong the doc file! ", e.Error())
		}
	}
}

func printProducts(w http.ResponseWriter, r *http.Request) {
	template := template.Must(template.ParseFiles("product.html"))

	template.Execute(w, products)
}

func printInformation(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		addInformation(w, r)
		return
	}

	template := template.Must(template.ParseFiles("contact.html"))

	template.Execute(w, nil)
}

func addInformation(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)

	var ms Message
	er := json.Unmarshal(body, &ms)
	if er != nil {
		log.Println(er.Error())
		return
	}
	log.Println(ms)

	now := time.Now()

	data := ms.Name + "|" + ms.Email + "|" + ms.Phone_number + "|" + ms.Message + "|" + now.Format("03:04:05PM")

	fmt.Println(data)
	if err := WriteFile("Contacts.txt", data); err != nil {
		return
	}
	w.WriteHeader(http.StatusOK)
}

func FormatPrice(price string) string {
	count := 0

	for i := len(price) - 1; i >= 0; i-- {
		count++
		if count == 3 {
			price = price[:i] + "," + price[i:]
			count = 0
		}
	}
	return price
}

func WriteFile(path string, data string) error {
	file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	defer file.Close()

	fmt.Fprintln(file, data)
	return nil
}

func ReadFile(path string) error {
	file, err := os.Open(path)

	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		pr := strings.Split(scanner.Text(), "|")

		number, _ := strconv.Atoi(pr[0])
		//price, _ := strconv.Atoi(pr[2])
		count, _ := strconv.Atoi(pr[4])

		product := Products{
			Stt:         number,
			Name:        pr[1],
			Price:       FormatPrice(pr[2]),
			Brand:       pr[3],
			Inventory:   count,
			Description: pr[5],
			Url:         pr[6],
		}
		products = append(products, product)
	}

	return nil
}

func main() {
	err := ReadFile("Products.txt")

	if err != nil {
		log.Fatal(err)
	}

	// Ket noi co so du lieu
	db, err := sql.Open("mysql", "kiettran:KietTran@2003@tcp(127.0.0.1:3306)/baitapweb")
	if err != nil {
		log.Println("Ket noi co so du lieu that bai!")
		panic(err.Error())
	}
	log.Println("Ket noi co so du lieu thanh cong!")
	defer db.Close()
	//-----------------------

	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir("assets"))
	file_image := http.FileServer(http.Dir("Products"))
	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))
	mux.Handle("/Products/", http.StripPrefix("/Products/", file_image))
	mux.HandleFunc("/bai2_contact", printInformation)
	mux.HandleFunc("/bai1_products", printProducts)
	mux.HandleFunc("/bai3_products", printProductsDB(db))
	mux.HandleFunc("/create_products", createProduct(db))
	mux.HandleFunc("/update_products", updateProduct(db))
	mux.HandleFunc("/product", queryProduct(db))

	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: mux,
	}

	fmt.Println("\nServer is running on port 8080")
	server.ListenAndServe()
}
