package calcmc

type CalMetroQuadrado struct {
	params *Params
}

func NewCalcMQ(opts ...Options) (*CalMetroQuadrado, error) {
	p, err := newParams(opts...)

	if err != nil {
		return nil, err
	}

	return &CalMetroQuadrado{
		params: p,
	}, nil
}

func (c *CalMetroQuadrado) CalculaMQ() float64 {
	return c.params.GetAltura() * c.params.GetBase()
}

func (c *CalMetroQuadrado) GetAltura() float64 {
	return c.params.GetAltura()
}

func (c *CalMetroQuadrado) GetBase() float64 {
	return c.params.GetBase()
}

func (c *CalMetroQuadrado) SetAltura(altura float64) {
	c.params.SetAltura(altura)
}

func (c *CalMetroQuadrado) SetBase(base float64) {
	c.params.SetBase(base)
}
