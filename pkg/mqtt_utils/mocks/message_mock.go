package mocks

type MockMessage struct {
	Topic   string
	Payload []byte
}

func (m *MockMessage) GetTopic() string {
	return m.Topic
}
