package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"product-service/domain"
	"product-service/service"

	"github.com/gorilla/mux"
)

// ProductHandler handles product-related HTTP requests
type ProductHandler struct {
	productService service.ProductService
	logger         Logger
}

// Logger defines the interface for logging
type Logger interface {
	Info(args ...interface{})
	Error(args ...interface{})
	WithField(key string, value interface{}) Logger
	WithError(err error) Logger
}

// NewProductHandler creates a new product handler
func NewProductHandler(productService service.ProductService, logger Logger) *ProductHandler {
	return &ProductHandler{
		productService: productService,
		logger:         logger,
	}
}

// CreateProductRequest represents the request payload for creating a product
type CreateProductRequest struct {
	Name        string `json:"name" validate:"required,min=2,max=100"`
	Description string `json:"description" validate:"required,min=10,max=500"`
	Price       int64  `json:"price" validate:"required,min=1"`
	Currency    string `json:"currency" validate:"required,len=3"`
	Category    string `json:"category" validate:"required,min=2,max=50"`
}

// UpdateProductRequest represents the request payload for updating a product
type UpdateProductRequest struct {
	Name        *string `json:"name,omitempty" validate:"omitempty,min=2,max=100"`
	Description *string `json:"description,omitempty" validate:"omitempty,min=10,max=500"`
	Price       *int64  `json:"price,omitempty" validate:"omitempty,min=1"`
	Currency    *string `json:"currency,omitempty" validate:"omitempty,len=3"`
	Category    *string `json:"category,omitempty" validate:"omitempty,min=2,max=50"`
}

// ProductResponse represents the response payload for product operations
type ProductResponse struct {
	Success bool           `json:"success"`
	Message string         `json:"message,omitempty"`
	Data    *ProductData   `json:"data,omitempty"`
}

// ProductsResponse represents the response payload for multiple products
type ProductsResponse struct {
	Success bool         `json:"success"`
	Message string       `json:"message,omitempty"`
	Data    []ProductData `json:"data,omitempty"`
	Total   int          `json:"total,omitempty"`
}

// ProductData represents product data in responses
type ProductData struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       MoneyData `json:"price"`
	Category    string `json:"category"`
	Status      string `json:"status"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

// MoneyData represents money data in responses
type MoneyData struct {
	Amount   int64  `json:"amount"`
	Currency string `json:"currency"`
}

// ErrorResponse represents an error response
type ErrorResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Error   string `json:"error,omitempty"`
}

// CreateProduct handles POST /products
func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var req CreateProductRequest
	
	// Parse JSON request body
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.logger.WithError(err).Error("Failed to decode request body")
		h.sendErrorResponse(w, http.StatusBadRequest, "Invalid JSON", err.Error())
		return
	}

	// Validate request
	if err := h.validateCreateRequest(req); err != nil {
		h.logger.WithError(err).Error("Validation failed")
		h.sendErrorResponse(w, http.StatusBadRequest, "Validation failed", err.Error())
		return
	}

	// Create Money value object
	price, err := domain.NewMoney(req.Price, req.Currency)
	if err != nil {
		h.logger.WithError(err).Error("Failed to create money value object")
		h.sendErrorResponse(w, http.StatusBadRequest, "Invalid price", err.Error())
		return
	}

	// Create product
	product, err := h.productService.CreateProduct(req.Name, req.Description, price, req.Category)
	if err != nil {
		h.logger.WithError(err).Error("Failed to create product")
		h.sendErrorResponse(w, http.StatusInternalServerError, "Failed to create product", err.Error())
		return
	}

	h.logger.WithField("product_id", product.ID.String()).Info("Product created successfully")

	// Return success response
	response := ProductResponse{
		Success: true,
		Message: "Product created successfully",
		Data:    h.toProductData(product),
	}

	h.sendJSONResponse(w, http.StatusCreated, response)
}

// GetProduct handles GET /products/{id}
func (h *ProductHandler) GetProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productIDStr := vars["id"]

	if productIDStr == "" {
		h.sendErrorResponse(w, http.StatusBadRequest, "Invalid product ID", "Product ID is required")
		return
	}

	// Create ProductID value object
	productID, err := domain.NewProductID(productIDStr)
	if err != nil {
		h.sendErrorResponse(w, http.StatusBadRequest, "Invalid product ID", err.Error())
		return
	}

	// Get product
	product, err := h.productService.GetProduct(productID)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			h.sendErrorResponse(w, http.StatusNotFound, "Product not found", "Product with given ID does not exist")
			return
		}
		h.logger.WithError(err).Error("Failed to get product")
		h.sendErrorResponse(w, http.StatusInternalServerError, "Failed to get product", err.Error())
		return
	}

	h.logger.WithField("product_id", productIDStr).Info("Product retrieved successfully")

	// Return success response
	response := ProductResponse{
		Success: true,
		Message: "Product retrieved successfully",
		Data:    h.toProductData(product),
	}

	h.sendJSONResponse(w, http.StatusOK, response)
}

// GetAllProducts handles GET /products
func (h *ProductHandler) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	// Parse query parameters
	category := r.URL.Query().Get("category")
	status := r.URL.Query().Get("status")

	var products []domain.Product
	var err error

	// Get products based on filters
	if category != "" {
		products, err = h.productService.GetProductsByCategory(category)
	} else if status == "active" {
		products, err = h.productService.GetActiveProducts()
	} else {
		products, err = h.productService.GetAllProducts()
	}

	if err != nil {
		h.logger.WithError(err).Error("Failed to get products")
		h.sendErrorResponse(w, http.StatusInternalServerError, "Failed to get products", err.Error())
		return
	}

	h.logger.WithField("count", len(products)).Info("Products retrieved successfully")

	// Convert to response format
	productData := make([]ProductData, len(products))
	for i, product := range products {
		productData[i] = *h.toProductData(&product)
	}

	// Return success response
	response := ProductsResponse{
		Success: true,
		Message: "Products retrieved successfully",
		Data:    productData,
		Total:   len(products),
	}

	h.sendJSONResponse(w, http.StatusOK, response)
}

// UpdateProduct handles PUT /products/{id}
func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productIDStr := vars["id"]

	if productIDStr == "" {
		h.sendErrorResponse(w, http.StatusBadRequest, "Invalid product ID", "Product ID is required")
		return
	}

	// Create ProductID value object
	productID, err := domain.NewProductID(productIDStr)
	if err != nil {
		h.sendErrorResponse(w, http.StatusBadRequest, "Invalid product ID", err.Error())
		return
	}

	var req UpdateProductRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.logger.WithError(err).Error("Failed to decode request body")
		h.sendErrorResponse(w, http.StatusBadRequest, "Invalid JSON", err.Error())
		return
	}

	// Validate request
	if err := h.validateUpdateRequest(req); err != nil {
		h.logger.WithError(err).Error("Validation failed")
		h.sendErrorResponse(w, http.StatusBadRequest, "Validation failed", err.Error())
		return
	}

	// Get existing product to get current values
	existingProduct, err := h.productService.GetProduct(productID)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			h.sendErrorResponse(w, http.StatusNotFound, "Product not found", "Product with given ID does not exist")
			return
		}
		h.logger.WithError(err).Error("Failed to get product for update")
		h.sendErrorResponse(w, http.StatusInternalServerError, "Failed to get product", err.Error())
		return
	}

	// Use existing values if not provided
	name := existingProduct.Name
	description := existingProduct.Description
	price := existingProduct.Price
	category := existingProduct.Category

	if req.Name != nil {
		name = *req.Name
	}
	if req.Description != nil {
		description = *req.Description
	}
	if req.Price != nil && req.Currency != nil {
		newPrice, err := domain.NewMoney(*req.Price, *req.Currency)
		if err != nil {
			h.sendErrorResponse(w, http.StatusBadRequest, "Invalid price", err.Error())
			return
		}
		price = newPrice
	}
	if req.Category != nil {
		category = *req.Category
	}

	// Update product
	if err := h.productService.UpdateProduct(productID, name, description, price, category); err != nil {
		h.logger.WithError(err).Error("Failed to update product")
		h.sendErrorResponse(w, http.StatusInternalServerError, "Failed to update product", err.Error())
		return
	}

	h.logger.WithField("product_id", productIDStr).Info("Product updated successfully")

	// Return success response
	response := ProductResponse{
		Success: true,
		Message: "Product updated successfully",
	}

	h.sendJSONResponse(w, http.StatusOK, response)
}

// DeleteProduct handles DELETE /products/{id}
func (h *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productIDStr := vars["id"]

	if productIDStr == "" {
		h.sendErrorResponse(w, http.StatusBadRequest, "Invalid product ID", "Product ID is required")
		return
	}

	// Create ProductID value object
	productID, err := domain.NewProductID(productIDStr)
	if err != nil {
		h.sendErrorResponse(w, http.StatusBadRequest, "Invalid product ID", err.Error())
		return
	}

	// Delete product
	if err := h.productService.DeleteProduct(productID); err != nil {
		if strings.Contains(err.Error(), "not found") {
			h.sendErrorResponse(w, http.StatusNotFound, "Product not found", "Product with given ID does not exist")
			return
		}
		h.logger.WithError(err).Error("Failed to delete product")
		h.sendErrorResponse(w, http.StatusInternalServerError, "Failed to delete product", err.Error())
		return
	}

	h.logger.WithField("product_id", productIDStr).Info("Product deleted successfully")

	// Return success response
	response := ProductResponse{
		Success: true,
		Message: "Product deleted successfully",
	}

	h.sendJSONResponse(w, http.StatusOK, response)
}

// DiscontinueProduct handles POST /products/{id}/discontinue
func (h *ProductHandler) DiscontinueProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productIDStr := vars["id"]

	if productIDStr == "" {
		h.sendErrorResponse(w, http.StatusBadRequest, "Invalid product ID", "Product ID is required")
		return
	}

	// Create ProductID value object
	productID, err := domain.NewProductID(productIDStr)
	if err != nil {
		h.sendErrorResponse(w, http.StatusBadRequest, "Invalid product ID", err.Error())
		return
	}

	// Discontinue product
	if err := h.productService.DiscontinueProduct(productID); err != nil {
		if strings.Contains(err.Error(), "not found") {
			h.sendErrorResponse(w, http.StatusNotFound, "Product not found", "Product with given ID does not exist")
			return
		}
		h.logger.WithError(err).Error("Failed to discontinue product")
		h.sendErrorResponse(w, http.StatusInternalServerError, "Failed to discontinue product", err.Error())
		return
	}

	h.logger.WithField("product_id", productIDStr).Info("Product discontinued successfully")

	// Return success response
	response := ProductResponse{
		Success: true,
		Message: "Product discontinued successfully",
	}

	h.sendJSONResponse(w, http.StatusOK, response)
}

// Helper methods

func (h *ProductHandler) sendJSONResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

func (h *ProductHandler) sendErrorResponse(w http.ResponseWriter, statusCode int, message, error string) {
	response := ErrorResponse{
		Success: false,
		Message: message,
		Error:   error,
	}
	h.sendJSONResponse(w, statusCode, response)
}

func (h *ProductHandler) validateCreateRequest(req CreateProductRequest) error {
	if req.Name == "" {
		return errors.New("name is required")
	}
	if len(req.Name) < 2 || len(req.Name) > 100 {
		return errors.New("name must be between 2 and 100 characters")
	}
	if req.Description == "" {
		return errors.New("description is required")
	}
	if len(req.Description) < 10 || len(req.Description) > 500 {
		return errors.New("description must be between 10 and 500 characters")
	}
	if req.Price <= 0 {
		return errors.New("price must be positive")
	}
	if req.Currency == "" {
		return errors.New("currency is required")
	}
	if len(req.Currency) != 3 {
		return errors.New("currency must be 3 characters")
	}
	if req.Category == "" {
		return errors.New("category is required")
	}
	if len(req.Category) < 2 || len(req.Category) > 50 {
		return errors.New("category must be between 2 and 50 characters")
	}
	return nil
}

func (h *ProductHandler) validateUpdateRequest(req UpdateProductRequest) error {
	if req.Name != nil && (len(*req.Name) < 2 || len(*req.Name) > 100) {
		return errors.New("name must be between 2 and 100 characters")
	}
	if req.Description != nil && (len(*req.Description) < 10 || len(*req.Description) > 500) {
		return errors.New("description must be between 10 and 500 characters")
	}
	if req.Price != nil && *req.Price <= 0 {
		return errors.New("price must be positive")
	}
	if req.Currency != nil && len(*req.Currency) != 3 {
		return errors.New("currency must be 3 characters")
	}
	if req.Category != nil && (len(*req.Category) < 2 || len(*req.Category) > 50) {
		return errors.New("category must be between 2 and 50 characters")
	}
	return nil
}

func (h *ProductHandler) toProductData(product *domain.Product) *ProductData {
	return &ProductData{
		ID:          product.ID.String(),
		Name:        product.Name,
		Description: product.Description,
		Price: MoneyData{
			Amount:   product.Price.Amount(),
			Currency: product.Price.Currency(),
		},
		Category:  product.Category,
		Status:    string(product.Status),
		CreatedAt: product.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		UpdatedAt: product.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
	}
}
