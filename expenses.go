package splitwise

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Expenses contains method to work with expense resource
type Expenses interface {
	// Expenses returns current user's expenses
	Expenses(ctx context.Context) ([]Expense, error)

	// ExpenseByID returns information of an expense identified by id argument
	ExpenseByID(ctx context.Context, id uint64) (*Expense, error)

	// CreateExpense Creates an expense. You may either split an expense equally (only with group_id provided), or
	// supply a list of shares.
	//If providing a list of shares, each share must include paid_share and owed_share, and must be identified by one
	// of the following:
	//email, first_name, and last_name
	//user_id
	//Note: 200 OK does not indicate a successful response. The operation was successful only if errors is empty.
	CreateExpense(ctx context.Context, dto *CreateExpenseDTO) ([]Expense, error)
}

type Expense struct {
	Cost                   string      `json:"cost"`
	Description            string      `json:"description"`
	Details                string      `json:"details"`
	Date                   time.Time   `json:"date"`
	RepeatInterval         string      `json:"repeat_interval"`
	CurrencyCode           string      `json:"currency_code"`
	CategoryID             int         `json:"category_id"`
	ID                     int         `json:"id"`
	GroupID                int         `json:"group_id"`
	FriendshipID           int         `json:"friendship_id"`
	ExpenseBundleID        int         `json:"expense_bundle_id"`
	Repeats                bool        `json:"repeats"`
	EmailReminder          bool        `json:"email_reminder"`
	EmailReminderInAdvance interface{} `json:"email_reminder_in_advance"`
	NextRepeat             string      `json:"next_repeat"`
	CommentsCount          int         `json:"comments_count"`
	Payment                bool        `json:"payment"`
	TransactionConfirmed   bool        `json:"transaction_confirmed"`
	Repayments             []struct {
		From   int    `json:"from"`
		To     int    `json:"to"`
		Amount string `json:"amount"`
	} `json:"repayments"`
	CreatedAt time.Time `json:"created_at"`
	CreatedBy struct {
		ID                 int    `json:"id"`
		FirstName          string `json:"first_name"`
		LastName           string `json:"last_name"`
		Email              string `json:"email"`
		RegistrationStatus string `json:"registration_status"`
		Picture            struct {
			Small  string `json:"small"`
			Medium string `json:"medium"`
			Large  string `json:"large"`
		} `json:"picture"`
	} `json:"created_by"`
	UpdatedAt time.Time `json:"updated_at"`
	UpdatedBy struct {
		ID                 int    `json:"id"`
		FirstName          string `json:"first_name"`
		LastName           string `json:"last_name"`
		Email              string `json:"email"`
		RegistrationStatus string `json:"registration_status"`
		Picture            struct {
			Small  string `json:"small"`
			Medium string `json:"medium"`
			Large  string `json:"large"`
		} `json:"picture"`
	} `json:"updated_by"`
	DeletedAt time.Time `json:"deleted_at"`
	DeletedBy struct {
		ID                 int    `json:"id"`
		FirstName          string `json:"first_name"`
		LastName           string `json:"last_name"`
		Email              string `json:"email"`
		RegistrationStatus string `json:"registration_status"`
		Picture            struct {
			Small  string `json:"small"`
			Medium string `json:"medium"`
			Large  string `json:"large"`
		} `json:"picture"`
	} `json:"deleted_by"`
	Category struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"category"`
	Receipt struct {
		Large    string `json:"large"`
		Original string `json:"original"`
	} `json:"receipt"`
	Users []struct {
		User struct {
			ID        int    `json:"id"`
			FirstName string `json:"first_name"`
			LastName  string `json:"last_name"`
			Picture   struct {
				Medium string `json:"medium"`
			} `json:"picture"`
		} `json:"user"`
		UserID     int    `json:"user_id"`
		PaidShare  string `json:"paid_share"`
		OwedShare  string `json:"owed_share"`
		NetBalance string `json:"net_balance"`
	} `json:"users"`
	Comments []struct {
		ID           int       `json:"id"`
		Content      string    `json:"content"`
		CommentType  string    `json:"comment_type"`
		RelationType string    `json:"relation_type"`
		RelationID   int       `json:"relation_id"`
		CreatedAt    time.Time `json:"created_at"`
		DeletedAt    time.Time `json:"deleted_at"`
		User         struct {
			ID        int    `json:"id"`
			FirstName string `json:"first_name"`
			LastName  string `json:"last_name"`
			Picture   struct {
				Medium string `json:"medium"`
			} `json:"picture"`
		} `json:"user"`
	} `json:"comments"`
}

type CreateExpenseDTO struct {
	Cost           string    `json:"cost"`
	Description    string    `json:"description"`
	Details        string    `json:"details"`
	Date           time.Time `json:"date"`
	RepeatInterval string    `json:"repeat_interval"`
	CurrencyCode   string    `json:"currency_code"`
	CategoryID     int       `json:"category_id"`
	GroupID        int       `json:"group_id"`
	SplitEqually   bool      `json:"split_equally"`
}

type CreateExpenseBySharesDTO struct {
	Cost                string    `json:"cost"`
	Description         string    `json:"description"`
	Details             string    `json:"details"`
	Date                time.Time `json:"date"`
	RepeatInterval      string    `json:"repeat_interval"`
	CurrencyCode        string    `json:"currency_code"`
	CategoryID          int       `json:"category_id"`
	GroupID             int       `json:"group_id"`
	Users0UserID        int       `json:"users__0__user_id"`
	Users0PaidShare     string    `json:"users__0__paid_share"`
	Users0OwedShare     string    `json:"users__0__owed_share"`
	Users1FirstName     string    `json:"users__1__first_name"`
	Users1LastName      string    `json:"users__1__last_name"`
	Users1Email         string    `json:"users__1__email"`
	Users1PaidShare     string    `json:"users__1__paid_share"`
	Users1OwedShare     string    `json:"users__1__owed_share"`
	UsersIndexProperty1 string    `json:"users__{index}__{property}1"`
	UsersIndexProperty2 string    `json:"users__{index}__{property}2"`
}

type createResponse struct {
	Expenses []Expense `json:"expenses"`
	Errors   struct {
		Base []string `json:"base"`
	} `json:"errors"`
}

func (c client) Expenses(ctx context.Context) ([]Expense, error) {
	return nil, nil
}

func (c client) ExpenseByID(ctx context.Context, id uint64) (*Expense, error) {
	return nil, nil
}

func (c client) CreateExpense(ctx context.Context, payload *CreateExpenseDTO) ([]Expense, error) {
	url := c.baseURL + "/api/v3.0/create_expense"

	buf, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(buf))
	if err != nil {
		return nil, err
	}

	token, err := c.AuthProvider.Auth()
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer "+token)
	req.Header.Add("Content-Type", "application/json")
	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("failed request: code=%d, status=%s", res.StatusCode, res.Status)
	}
	defer func() {
		_ = res.Body.Close()
	}()

	var response createResponse
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	if len(response.Errors.Base) > 0 {
		return nil, fmt.Errorf("request error: %s", response.Errors.Base[0])
	}

	return response.Expenses, nil
}
