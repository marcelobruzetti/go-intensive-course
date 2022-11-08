package usecase

import (
	"database/sql"
	"go-intensivo/internal/order/entity"
	"go-intensivo/internal/order/infra/database"
	"testing"

	"github.com/stretchr/testify/suite"

	_ "github.com/mattn/go-sqlite3"
)

type CalculatePriceUseCaseTestSuite struct {
	suite.Suite
	OrderRepository database.OrderRepository
	Db *sql.DB
}

func (suite *CalculatePriceUseCaseTestSuite) SetupTest() {
	db, err := sql.Open("sqlite3", ":memory:")
	suite.NoError(err)

	// create table orders
	_, err = db.Exec("CREATE TABLE orders (id VARCHAR(255) NOT NULL, price FLOAT NOT NULL, tax FLOAT NOT NULL, final_price FLOAT NOT NULL)")
	suite.Db = db
	suite.OrderRepository = *database.NewOrderRepository(db)
}

func (suite *CalculatePriceUseCaseTestSuite) TearDownTest() {
	suite.Db.Close()
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(CalculatePriceUseCaseTestSuite))
}

func (suite *CalculatePriceUseCaseTestSuite) TestCalculateFinalPrice() {
	order, err := entity.NewOrder("1", 10.0, 2.0)
	suite.NoError(err)
	order.CalculateFinalPrice()
	

	calculateFinalPriceInput := OrderInputDTO{
		ID: order.ID,
		Price: order.Price,
		Tax: order.Tax,
	}
	calculateFinalPriceUseCase := NewCalculateFinalPriceUseCase(suite.OrderRepository)
	output, err := calculateFinalPriceUseCase.Execute(calculateFinalPriceInput)
	suite.NoError(err)
	suite.Equal(order.ID, output.ID)
	suite.Equal(order.Price, output.Price)
	suite.Equal(order.Tax, output.Tax)
	suite.Equal(order.FinalPrice, output.FinalPrice)
}