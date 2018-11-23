package behaviortree

// Blackboard ...
type Blackboard struct {
	data map[string]interface{}
}

// NewBlackboard instantiates a new blackboard
func NewBlackboard() *Blackboard {
	bb := &Blackboard{}
	bb.data = map[string]interface{}{}
	return bb
}

func (bb *Blackboard) Read(id string) (interface{}, bool) {
	value, ok := bb.data[id]
	return value, ok
}

func (bb *Blackboard) Write(id string, data interface{}) {
	bb.data[id] = data
}