package processors

type Updater struct {
	Column   string
	NewValue string
	IProcessor
}

func NewUpdater(column, newValue string) *Updater {
	return &Updater{
		Column:   column,
		NewValue: newValue,
	}
}
