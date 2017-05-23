package captcha

type captcha struct {
	pattern       int
	firstOperand  int
	seconfOperand int
	operator      int
}

func (c *captcha) String() string {
	if c.pattern == 2 {
		return "one " + c.oper() + " 1"
	}
	return "1 " + c.oper() + " one"
}

func (c *captcha) oper() string {
	m := map[int]string{
		1: "+",
		2: "-",
		3: "*",
	}
	return m[c.operator]
}
