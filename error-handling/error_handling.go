package erratum

const testVersion = 2

func Use(opener ResourceOpener, input string) (err error) {
	var res Resource

	for {
		res, err = opener()
		if err != nil {
			_, ok := err.(TransientError)
			if !ok {
				return err
			}
			continue
		}
		break
	}
	defer res.Close()

	defer func() {
		if r := recover(); r != nil {
			switch r.(type) {
			case FrobError:
				fe := r.(FrobError)
				err = fe.GetError()
				res.Defrob(fe.GetDeforbTag())
			default:
				err = r.(error)
			}
		}
	}()

	res.Frob(input)

	return err
}
