package selectbox

// Option represents a selectable option in a selectbox
type Option struct {
	Label    string
	Value    string
	Selected bool
}

// TransformToOptions converts a slice of base models to selectbox options
func TransformToOptions(items []interface{ GetLabel() string; GetValue() string }) []Option {
	var options []Option
	for _, item := range items {
		options = append(options, Option{
			Label: item.GetLabel(),
			Value: item.GetValue(),
		})
	}
	return options
}
