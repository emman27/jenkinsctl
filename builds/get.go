package builds

func List() *Builds {
	return &Builds{
		{JobName: "Hello", BuildID: 1, Description: "test"},
	}
}
