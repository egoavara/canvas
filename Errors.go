package canvas

import "github.com/pkg/errors"

var (
	ErrorTooBusy     = errors.New("Currently, Query queue is too busy to handle this")
	ErrorUnsupported = errors.New("There is unsupported feature in requested query, please check that")
	ErrorNoQuery = errors.New("'query' must be a real pointer with data, not nil")
	ErrorInvalidData = errors.New("'query' data is invalid, please check this")
)
