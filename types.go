// This file was auto-generated by Fern from our API Definition.

package petstoreferndefapi

import (
	json "encoding/json"
	fmt "fmt"
	core "github.com/minaelee/ferntest-go-sdk/core"
)

type Category struct {
	Id   *int    `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`

	_rawJSON json.RawMessage
}

func (c *Category) UnmarshalJSON(data []byte) error {
	type unmarshaler Category
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = Category(value)
	c._rawJSON = json.RawMessage(data)
	return nil
}

func (c *Category) String() string {
	if len(c._rawJSON) > 0 {
		if value, err := core.StringifyJSON(c._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(c); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", c)
}

type Tag struct {
	Id   *int    `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`

	_rawJSON json.RawMessage
}

func (t *Tag) UnmarshalJSON(data []byte) error {
	type unmarshaler Tag
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*t = Tag(value)
	t._rawJSON = json.RawMessage(data)
	return nil
}

func (t *Tag) String() string {
	if len(t._rawJSON) > 0 {
		if value, err := core.StringifyJSON(t._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(t); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", t)
}

type OrderStatus string

const (
	OrderStatusPlaced    OrderStatus = "placed"
	OrderStatusApproved  OrderStatus = "approved"
	OrderStatusDelivered OrderStatus = "delivered"
)

func NewOrderStatusFromString(s string) (OrderStatus, error) {
	switch s {
	case "placed":
		return OrderStatusPlaced, nil
	case "approved":
		return OrderStatusApproved, nil
	case "delivered":
		return OrderStatusDelivered, nil
	}
	var t OrderStatus
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (o OrderStatus) Ptr() *OrderStatus {
	return &o
}
