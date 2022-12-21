package a

import "errors"

var err error = errors.New("nilpointer")

func g() *int {
	return nil
}

func h() (*int, error) {
	i1 := 1
	if err != nil {
		return &i1, err
	}
	i1 = 2
	return &i1, nil
}

func i() (*int, error) {
	if err != nil {
		return nil, nil // want "return nil pointer: line:22"
	}

	i1 := 1
	return &i1, nil
}

func j() (*int, *int, error) {
	if err != nil {
		return nil, nil, nil // want "return nil pointer: line:31"
	}
	i1 := 1
	i2 := 2
	return &i1, &i2, nil
}

func k() (*int, *int, error) {
	i1 := 1
	i2 := 2
	if err != nil {
		return &i1, &i2, nil
	}
	return nil, nil, nil // want "return nil pointer: line:44"
}

type Hoge struct{ field string }

func m() *Hoge {
	return nil
}

func n() (*Hoge, error) {
	return nil, nil // want "return nil pointer: line:54"
}

func o() (*Hoge, error) {
	return &Hoge{}, nil
}

func p() (*Hoge, error) {
	return &Hoge{}, err
}

func q() (*Hoge, error) {
	//lint:ignore nilpointer reason
	return nil, nil
}
