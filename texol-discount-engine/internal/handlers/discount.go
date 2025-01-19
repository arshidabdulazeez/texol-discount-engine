package handlers

import (
	"encoding/json"
	"net/http"
	"fmt"

	"texol-discount-engine/texol-discount-engine/internal/discount"
	"texol-discount-engine/texol-discount-engine/pkg/utils"
)

type DiscountRequest struct {
	OrderTotal   float64 `json:"order_total"`
	CustomerType string  `json:"customer_type"`
}

type DiscountResponse struct {
	Discount    float64 `json:"discount"`
	FinalTotal  float64 `json:"final_total"`
	RuleApplied string  `json:"rule_applied"`
}

func ApplyDiscountHandler(w http.ResponseWriter, r *http.Request) {
	var req DiscountRequest
	// customerType := "regular" // Default value
	// if req.CustomerType == "" {
	// 	req.CustomerType = customerType
	// }
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.WriteErrorResponse(w, http.StatusBadRequest, "Invalid input")
		return
	}

	rules, err := discount.LoadRules("config/rules.json")
	fmt.Println(rules)
	fmt.Println(err)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusInternalServerError, "Failed to load rules")
		return
	}

	discountAmount, appliedRule := discount.GetBestDiscount(rules, map[string]interface{}{
		"order_total":   req.OrderTotal,
		"customer_type": req.CustomerType,
	})

	response := DiscountResponse{
		Discount:    discountAmount,
		FinalTotal:  req.OrderTotal - discountAmount,
		RuleApplied: appliedRule,
	}
	utils.WriteJSONResponse(w, http.StatusOK, response)
}
