package atm

import "testing"

func Test_ApplicationService_SpendMoney(t *testing.T) {
	card := NewCreditCard(5000.0)
	cardRepository := NewCreditCardRepository()
	cardRepository.Store(card)
	cardService := NewCreditCardService()
	applicationService := NewApplicationService(cardRepository, cardService)
	amount := 2000.0 // Руб

	applicationService.SpendMoney(amount)

	if card.Balance() != 3000.0 {
		t.Errorf("expectred card balance 3000.0 RUB, got: %v\n", card.Balance())
	}
}

//
// ApplicationService
//
func NewApplicationService(cardRepository *creditCardRepository, cardService *creditCardService) *applicationService {
	return &applicationService{cardRepository, cardService}
}

type applicationService struct {
	cardRepository *creditCardRepository
	cardService    *creditCardService
}

func (app *applicationService) SpendMoney(amount float64) {
	card := app.cardRepository.FindCard()
	app.cardService.SpendMoney(card, amount)
	app.cardRepository.Store(card)
}

//
// CardRepository
//
func NewCreditCardRepository() *creditCardRepository {
	return &creditCardRepository{}
}

type creditCardRepository struct {
	card *card
}

func (r *creditCardRepository) Store(card *card) {
	r.card = card
}

func (r *creditCardRepository) FindCard() *card {
	return r.card
}

//
// CreditCardService
//
func NewCreditCardService() *creditCardService {
	return &creditCardService{}
}

type creditCardService struct{}

func (service *creditCardService) SpendMoney(card *card, amount float64) {
	card.Spend(amount)
}

//
// Card
//
func NewCreditCard(amount float64) *card {
	return &card{amount}
}

type card struct {
	balance float64
}

func (c *card) Spend(amount float64) {
	if amount <= c.balance {
		c.balance -= amount
	}
}

func (c *card) Balance() float64 {
	return c.balance
}
