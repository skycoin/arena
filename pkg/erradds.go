package pkg

type Error struct {
	Val int
}

func AddendWrong(osArg []string) Error {
	if len(osArg) == 0 {
		return Error{
			Val: 1,
		}
	}

	if len(osArg) < 3 {
		return Error{
			Val: 2,
		}
	}

	if len(osArg) > 3 {
		return Error{
			Val: 3,
		}
	}

	return Error{
		Val: 0,
	}
}
