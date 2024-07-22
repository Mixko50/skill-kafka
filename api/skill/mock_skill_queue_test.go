package skill

type mockSkillQueue struct {
	SkillQueue
	errPublish error
}

func (m *mockSkillQueue) PublishSkill(action SkillAction, key *string, skillPayload interface{}) error {
	return m.errPublish
}
