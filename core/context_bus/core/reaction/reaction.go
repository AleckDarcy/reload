package reaction

func (c *Configure) InitializeSnapshot() *PrerequisiteSnapshot {
	return c.PreTree.InitializeSnapshot()
}
