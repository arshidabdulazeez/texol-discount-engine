package main

import (
	"fmt"
	"log"
	"net/http"

	"texol-discount-engine/texol-discount-engine/internal/handlers"
)

func main() {
	http.HandleFunc("/apply-discount", handlers.ApplyDiscountHandler)
	fmt.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}









// 000000000000000000000000000000000000000000


// package main

// import "fmt"

// func main() {
//     fmt.Println("Hello, World!")
// }

// type Rule struct {
//     ID          string            `json:"id"`
//     Description string            `json:"description"`
//     Conditions  map[string]string `json:"conditions"`
//     DiscountType string           `json:"discount_type"` // "percentage" or "fixed"
//     Value       float64           `json:"value"`
//     Priority    int               `json:"priority"`
// }

// // Methods
// // =============================================

// func LoadRules(filePath string) ([]Rule, error) {
//     data, err := os.ReadFile(filePath)
//     if err != nil {
//         return nil, err
//     }
//     var rules []Rule
//     err = json.Unmarshal(data, &rules)
//     if err != nil {
//         return nil, err
//     }
//     return rules, nil
// }


// // Validations
// // ==============================================

// func IsRuleApplicable(rule Rule, order map[string]interface{}) bool {
//     for key, value := range rule.Conditions {
//         if order[key] != value {
//             return false
//         }
//     }
//     return true
// }

// // ==================================================

// func GetBestDiscount(rules []Rule, order map[string]interface{}) (float64, string) {
//     var maxDiscount float64
//     var appliedRule string

//     for _, rule := range rules {
//         if IsRuleApplicable(rule, order) {
//             discount := CalculateDiscount(rule, order)
//             if discount > maxDiscount {
//                 maxDiscount = discount
//                 appliedRule = rule.ID
//             }
//         }
//     }
//     return maxDiscount, appliedRule
// }

// func CalculateDiscount(rule Rule, order map[string]interface{}) float64 {
//     orderTotal := order["order_total"].(float64)
//     if rule.DiscountType == "percentage" {
//         return (rule.Value / 100) * orderTotal
//     } else if rule.DiscountType == "fixed" {
//         return rule.Value
//     }
//     return 0
// }


// // ======================================

// var ruleLock sync.RWMutex

// func ConcurrentHandler(order map[string]interface{}) {
//     ruleLock.RLock()
//     defer ruleLock.RUnlock()
//     // Call the logic to calculate discounts here
// }



// func ApplyDiscountHandler(w http.ResponseWriter, r *http.Request) {
//     var order map[string]interface{}
//     if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
//         http.Error(w, "Invalid input", http.StatusBadRequest)
//         return
//     }

//     rules, err := LoadRules("rules.json")
//     if err != nil {
//         http.Error(w, "Failed to load rules", http.StatusInternalServerError)
//         return
//     }

//     discount, appliedRule := GetBestDiscount(rules, order)
//     finalTotal := order["order_total"].(float64) - discount

//     response := map[string]interface{}{
//         "discount":    discount,
//         "final_total": finalTotal,
//         "rule_applied": appliedRule,
//     }

//     w.Header().Set("Content-Type", "application/json")
//     json.NewEncoder(w).Encode(response)
// }




// func TestIsRuleApplicable(t *testing.T) {
//     rule := Rule{
//         Conditions: map[string]string{"customer_type": "premium"},
//     }
//     order := map[string]interface{}{"customer_type": "premium"}

//     if !IsRuleApplicable(rule, order) {
//         t.Error("Expected rule to be applicable")
//     }
// }
