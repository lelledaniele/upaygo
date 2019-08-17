package apperror

type RESTError struct {
	M string `json:"error"`
}

func (e *RESTError) Error() string {
	return e.M
}
