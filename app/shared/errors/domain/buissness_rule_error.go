package domainErr

type BusinessRuleError struct {
	BaseDomainError
	Rule string `json:"rule"`
}

func NewBusinessRuleError(rule, message string) *BusinessRuleError {
	return &BusinessRuleError{
		BaseDomainError: BaseDomainError{
			Code:    "BUSINESS_RULE_VIOLATION",
			Type:    "domain",
			Message: message,
			Data: map[string]interface{}{
				"rule": rule,
			},
		},
		Rule: rule,
	}
}
