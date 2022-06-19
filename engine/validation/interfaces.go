package validation

type IValidator interface {
	ValidateFilter(filter map[string]interface{}) error
	ValidateMutation(mutation map[string]interface{}) error
	ValidateName(name string) error
	ValidateDocument(document map[string]interface{}) error
	IsTopLevelFilterOp(filterOp string) bool
	IsFilterOp(filterOp string) bool
}
