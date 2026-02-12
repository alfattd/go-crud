package server

import (
	"net/http"

	"github.com/alfattd/crud/internal/handler"
	"github.com/alfattd/crud/internal/platform/config"
	"github.com/alfattd/crud/internal/platform/database"
	"github.com/alfattd/crud/internal/platform/monitor"
	"github.com/alfattd/crud/internal/repository"
	"github.com/alfattd/crud/internal/repository/memory"
	"github.com/alfattd/crud/internal/repository/postgres"
	"github.com/alfattd/crud/internal/service"
)

func New(cfg *config.Config) *http.Server {

	mux := http.NewServeMux()

	var categoryRepo repository.CategoryRepository
	var productRepo repository.ProductRepository

	if cfg.ServiceVersion == "dev" {
		categoryRepo = memory.NewInMemoryCategoryRepo()
		productRepo = memory.NewInMemoryProductRepo()
	} else {
		db := database.NewPostgres(cfg.DBUrl())
		categoryRepo = postgres.NewPostgresCategoryRepo(db)
		productRepo = postgres.NewPostgresProductRepo(db)
	}

	categoryService := service.NewCategoryService(categoryRepo)
	productService := service.NewProductService(productRepo, categoryRepo)

	categoryHandler := handler.NewCategoryHandler(categoryService)
	productHandler := handler.NewProductHandler(productService)

	mux.HandleFunc("/health", monitor.Health)
	mux.HandleFunc("/version", monitor.Version(cfg.ServiceName, cfg.ServiceVersion))
	mux.Handle("/metrics", monitor.MetricsHandler())

	mux.HandleFunc("/categories", func(w http.ResponseWriter, r *http.Request) {

		switch r.Method {

		case http.MethodPost:
			categoryHandler.CreateCategory(w, r)

		case http.MethodGet:
			categoryHandler.ListCategory(w, r)

		case http.MethodPut:
			categoryHandler.UpdateCategory(w, r)

		case http.MethodDelete:
			categoryHandler.DeleteCategory(w, r)

		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/categories/detail", categoryHandler.GetCategoryByID)

	mux.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {

		switch r.Method {

		case http.MethodPost:
			productHandler.CreateProduct(w, r)

		case http.MethodGet:
			productHandler.ListProduct(w, r)

		case http.MethodPut:
			productHandler.UpdateProduct(w, r)

		case http.MethodDelete:
			productHandler.DeleteProduct(w, r)

		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/products/detail", productHandler.GetProductByID)

	handlerWithMetrics := MetricsMiddleware(mux)

	return &http.Server{
		Addr:    ":80",
		Handler: handlerWithMetrics,
	}
}
