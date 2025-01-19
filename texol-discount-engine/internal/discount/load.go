package discount

import (
	"encoding/json"
	"os"
	"fmt"
)

type Rule struct {
	ID          string            `json:"id"`
	Description string            `json:"description"`
	Conditions  map[string]interface{} `json:"condition"`
	DiscountFixed *float64           `json:"discount_fixed,omitempty"`
	DiscountPercentage *float64           `json:"discount_percentage,omitempty"`
	// Value       float64           `json:"value"`
	Priority    int               `json:"priority"`
}

func (i *Rule) Validate() error {
	if i.DiscountFixed != nil && i.DiscountPercentage != nil {
		return fmt.Errorf("only one of 'discount_fixed' or 'discount_percentage' can be set")
	}
	return nil
}

func LoadRules(filePath string) ([]Rule, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	var rules []Rule
	err = json.Unmarshal(data, &rules)
	if err != nil {
		return nil, err
	}


	for _, item := range rules {
		if err := item.Validate(); err != nil {
			return nil, err
		}
	}

	return rules, nil
}