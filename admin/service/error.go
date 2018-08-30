package service

// ResourceNotFoundError is resource does not exists error
type ResourceNotFoundError struct {
	message string
}

// NewResourceNotFoundError returns a new ResourceNotFoundError instance
func NewResourceNotFoundError(message string) *ResourceNotFoundError {
	return &ResourceNotFoundError{
		message: message,
	}
}

// Error returns error message
func (rn *ResourceNotFoundError) Error() string {
	return rn.message
}

// StoreSystemError caused by store error
type StoreSystemError struct {
	message string
}

// NewStoreSystemError returns a new StoreSystemError instance
func NewStoreSystemError(message string) *StoreSystemError {
	return &StoreSystemError{
		message: message,
	}
}

// Error returns error message
func (ss *StoreSystemError) Error() string {
	return ss.message
}
