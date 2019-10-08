package main

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

var testAccProviders map[string]terraform.ResourceProvider
var testAccProvider *schema.Provider

func init() {
	testAccProvider = Provider()
	testAccProviders = map[string]terraform.ResourceProvider{
		"expense": testAccProvider,
	}
}

func testAccPreCheck(t *testing.T) {
	if v := os.Getenv("EXPENSE_URL"); v == "" {
		t.Fatal("EXPENSE_URL must be set for acceptance tests")
	}
}

func TestAccResourceExpenseCreate(t *testing.T) {

	name := fmt.Sprintf("Unit Testing on %s", time.Now().Format("Jan 02, 2006 at 15:04:05"))
	tripID := "raleigh"
	cost := 100.03
	currency := "US"
	date := "2019-08-22T00:00:00"

	expense := &Expense{
		Name:     name,
		TripID:   tripID,
		Cost:     cost,
		Currency: currency,
		Date:     date,
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckExpenseDelete("expense_item.test", expense),
		Steps: []resource.TestStep{
			{
				Config: testAccCheckExpenseConfig_basic(name, tripID, currency, date, cost),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckExpenseExists("expense_item.test", expense),
					resource.TestCheckResourceAttr("expense_item.test", "name", name),
					resource.TestCheckResourceAttr("expense_item.test", "trip_id", tripID),
					resource.TestCheckResourceAttr("expense_item.test", "currency", currency),
					resource.TestCheckResourceAttr("expense_item.test", "cost", fmt.Sprintf("%.2f", cost)),
					resource.TestCheckResourceAttr("expense_item.test", "date", date),
				),
			},
		},
	})
}

func TestAccResourceExpenseUpdate(t *testing.T) {

	name := fmt.Sprintf("Unit Testing on %s", time.Now().Format("Jan 02, 2006 at 15:04:05"))
	tripID := "raleigh"
	cost := 100.03
	currency := "US"
	date := "2019-08-22T00:00:00"
	newCost := 200.03

	expense := &Expense{
		Name:     name,
		TripID:   tripID,
		Cost:     cost,
		Currency: currency,
		Date:     date,
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckExpenseDelete("expense_item.test", expense),
		Steps: []resource.TestStep{
			{
				Config: testAccCheckExpenseConfig_basic(name, tripID, currency, date, cost),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckExpenseExists("expense_item.test", expense),
					resource.TestCheckResourceAttr("expense_item.test", "name", name),
					resource.TestCheckResourceAttr("expense_item.test", "trip_id", tripID),
					resource.TestCheckResourceAttr("expense_item.test", "currency", currency),
					resource.TestCheckResourceAttr("expense_item.test", "cost", fmt.Sprintf("%.2f", cost)),
					resource.TestCheckResourceAttr("expense_item.test", "date", date),
				),
			},
			{
				Config: testAccCheckExpenseConfig_newcost(name, tripID, currency, date, newCost),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckExpenseExists("expense_item.test", expense),
					resource.TestCheckResourceAttr("expense_item.test", "name", name),
					resource.TestCheckResourceAttr("expense_item.test", "trip_id", tripID),
					resource.TestCheckResourceAttr("expense_item.test", "currency", currency),
					resource.TestCheckResourceAttr("expense_item.test", "cost", fmt.Sprintf("%.2f", newCost)),
					resource.TestCheckResourceAttr("expense_item.test", "date", date),
				),
			},
		},
	})
}

func testAccCheckExpenseConfig_basic(name string, tripId string, currency string, date string, cost float64) string {
	return fmt.Sprintf(`
	resource "expense_item" "test" {
		name          = "%s"
		currency      = "%s"
		cost          = "%.2f"
		date          = "%s"
		trip_id        = "%s"
	}
	`, name, currency, cost, date, tripId)
}

func testAccCheckExpenseConfig_newcost(name string, tripId string, currency string, date string, cost float64) string {
	return fmt.Sprintf(`
	resource "expense_item" "test" {
		name          = "%s"
		currency      = "%s"
		cost          = "%.2f"
		date          = "%s"
		trip_id        = "%s"
	}
	`, name, currency, cost, date, tripId)
}

func testAccCheckExpenseDelete(n string, expense *Expense) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("not found: %s", n)
		}

		var config Config
		if err := config.LoadAndValidate(); err != nil {
			return err
		}

		_, err := config.Client.GetExpenseByID(rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("should be deleted: %s", n)
		}

		return nil
	}
}

func testAccCheckExpenseExists(n string, expense *Expense) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("no expense ID is set")
		}

		var config Config
		if err := config.LoadAndValidate(); err != nil {
			return err
		}

		result, err := config.Client.GetExpenseByID(rs.Primary.ID)
		if err != nil {
			return err
		}

		if expense.Name != result.Name {
			return fmt.Errorf("expense not found")
		}

		*expense = *result

		return nil
	}
}
