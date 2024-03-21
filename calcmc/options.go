package calcmc

import "errors"

type Options func(*Params) error

type Params struct {
	altura float64
	base   float64
}

func newParams(opts ...Options) (*Params, error) {
	p := &Params{}
	for _, opt := range opts {
		if err := opt(p); err != nil {
			return nil, err
		}
	}
	return p, nil
}

func WithAltura(altura float64) Options {
	return func(r *Params) error {

		if altura <= 0 {
			return errors.New("altura zerado")
		}

		r.altura = altura

		return nil
	}
}

func WithBase(base float64) Options {
	return func(r *Params) error {

		if base <= 0 {
			return errors.New("base zerado")
		}

		r.base = base

		return nil
	}
}

// setetrs ----
func (p Params) GetAltura() float64 {
	return p.altura
}

func (p Params) GetBase() float64 {
	return p.base
}

// getters ----
func (p *Params) SetAltura(altura float64) {
	p.altura = altura
}

func (p *Params) SetBase(base float64) {
	p.base = base
}
