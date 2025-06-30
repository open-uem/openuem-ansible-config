package ansible

func NewAnsiblePlaybook() *AnsiblePlaybook {
	pb := AnsiblePlaybook{}
	pb.Hosts = "localhost"
	pb.Connection = "local"
	return &pb
}

func (pb *AnsiblePlaybook) AddAnsibleTask(task any) {
	pb.Tasks = append(pb.Tasks, task)
}
