package xerror

func IsError(err error) bool {
	return err != nil
}

func NoError(err error) bool {
	return err == nil
}
