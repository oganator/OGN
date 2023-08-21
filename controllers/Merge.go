package controllers

// Ranges over child units and runs unit.Merge() on them
func (e *EntityModel) Merge() {
	for _, v := range e.ChildUnitModels {
		v.Merge()
	}
}

// Copies values from the parent to the child. TODO - pull values from the Override field. Currently only does CostInput
func (u *UnitModel) Merge() {
	for i, v := range u.Parent.GLA.CostInput {
		_, exists := u.CostInput[i]
		if !exists {
			u.CostInput[i] = CostInput{
				Name:              v.Name,
				MasterID:          v.MasterID,
				Type:              v.Type,
				Amount:            v.Amount,
				AmountSigma:       v.AmountSigma,
				COAItemBasis:      v.COAItemBasis,
				COAItemTarget:     v.COAItemTarget,
				Duration:          v.Duration,
				DurationSigma:     v.DurationSigma,
				Start:             v.Start,
				StartEvent:        v.StartEvent,
				End:               v.End,
				EndEvent:          v.EndEvent,
				GrowthItem:        v.GrowthItem,
				GrowthItemOptions: v.GrowthItemOptions,
			}
		}
	}
}
