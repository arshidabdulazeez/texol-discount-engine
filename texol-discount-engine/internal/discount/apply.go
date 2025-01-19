package discount

import (
	"sort"
	"fmt"
)

func GetBestDiscount(rules []Rule, order map[string]interface{}) (float64, string) {
	applicableRules := []Rule{}

	for _, rule := range rules {
		if IsRuleApplicable(rule, order) {
			applicableRules = append(applicableRules, rule)
		}
		// fmt.Println(rule, "-------========--------" , applicableRules)

	}

	sort.Slice(applicableRules, func(i, j int) bool {
		return applicableRules[i].Priority < applicableRules[j].Priority
	})
	// fmt.Println("----7777777---========--------" , applicableRules)

	var maxDiscount float64
	var appliedRule string
	for _, rule := range applicableRules {
		// fmt.Println("---------------",order)
		discount := CalculateDiscount(rule, order)
		// fmt.Println("---------------",discount)

		if discount > maxDiscount {
			maxDiscount = discount
			appliedRule = rule.ID
		}
	}

	return maxDiscount, appliedRule
}

func CalculateDiscount(rule Rule, order map[string]interface{}) float64 {
	orderTotal := order["order_total"].(float64)
	// fmt.Println("=========", orderTotal)
	if rule.DiscountPercentage != nil {
		return (*rule.DiscountPercentage / 100.0) * orderTotal
	} else if rule.DiscountFixed != nil {
		return *rule.DiscountFixed
	}
	return 0
}

func IsRuleApplicable(rule Rule, order map[string]interface{}) bool {
		// fmt.Println("pppppppppppppp----", key, "======", value)

	for key, value := range rule.Conditions {
		// fmt.Println("pppppppppppppp----", key, "======", value)
		// fmt.Println("pppppppppppppp--0000000000--", order[key], "======", value)
		if order[key] == value {
			// fmt.Println(key, value)
			return true
		} else if key == "min_order_value" {
			orderTotal, _ := order["order_total"].(float64)
			valueFloat, _ := value.(float64)
			if orderTotal >= valueFloat {
				// fmt.Println(key, value, order["order_total"])
				// fmt.Printf("%T %T", value, order["order_total"])

				return true
			}
			return false
		} else {
			return false
		}  
	}
	return true
}