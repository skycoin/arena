package pkg

type ErrorAdds struct {
	Val int
}

func AddendWrong() ErrorAdds {
	return ErrorAdds{
		Val: 1,
	}
}
