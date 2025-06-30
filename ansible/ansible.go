package ansible

func NewAnsiblePlaybook() *AnsiblePlaybook {
	pb := AnsiblePlaybook{}
	return &pb
}

func (pb *AnsiblePlaybook) AddAnsibleTask(task any) {
	pb.Tasks = append(pb.Tasks, task)
}
