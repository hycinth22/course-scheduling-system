package controllers

func getOffset(pageIndex int, pageSize int) int {
	return (pageIndex - 1) * pageSize
}
