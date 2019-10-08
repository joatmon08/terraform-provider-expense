package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceExpense() *schema.Resource {
	return &schema.Resource{
		Create: resourceExpenseCreate,
		Read:   resourceExpenseRead,
		Update: resourceExpenseUpdate,
		Delete: resourceExpenseDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"trip_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"cost": &schema.Schema{
				Type:     schema.TypeFloat,
				Required: true,
			},
			"currency": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"date": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceExpenseCreate(d *schema.ResourceData, m interface{}) error {
	config := m.(*Config)
	expense := &Expense{
		Name:     d.Get("name").(string),
		TripID:   d.Get("trip_id").(string),
		Cost:     d.Get("cost").(float64),
		Currency: d.Get("currency").(string),
		Date:     d.Get("date").(string),
	}
	id, err := config.Client.CreateExpense(expense)
	if err != nil {
		d.SetId("")
		return nil
	}
	d.SetId(id)
	return resourceExpenseRead(d, m)
}

func resourceExpenseRead(d *schema.ResourceData, m interface{}) error {
	config := m.(*Config)

	// Attempt to read from an upstream API
	obj, err := config.Client.GetExpenseByID(d.Id())

	if err != nil {
		d.SetId("")
		return err
	}

	d.Set("name", obj.Name)
	d.Set("trip_id", obj.TripID)
	d.Set("cost", obj.Cost)
	d.Set("currency", obj.Currency)
	d.Set("date", obj.Date)
	return nil
}

func resourceExpenseUpdate(d *schema.ResourceData, m interface{}) error {
	config := m.(*Config)
	expense := &Expense{
		ID:       d.Id(),
		Name:     d.Get("name").(string),
		TripID:   d.Get("trip_id").(string),
		Cost:     d.Get("cost").(float64),
		Currency: d.Get("currency").(string),
		Date:     d.Get("date").(string),
	}
	err := config.Client.UpdateExpense(expense)
	if err != nil {
		return err
	}
	return resourceExpenseRead(d, m)
}

func resourceExpenseDelete(d *schema.ResourceData, m interface{}) error {
	config := m.(*Config)
	err := config.Client.DeleteExpense(d.Id())
	if err != nil {
		return err
	}
	d.SetId("")
	return nil
}
