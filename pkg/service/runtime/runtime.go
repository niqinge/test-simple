package runtime

type Item interface{

}

type Runtime struct {
	
}

func (r *Runtime) Del(ctx BaseContext) error {
	panic("implement me")
}

func (r *Runtime) Create(ctx BaseContext) error {
	panic("implement me")
}

func (r *Runtime) Update(ctx BaseContext) error {
	panic("implement me")
}

func (r *Runtime) Query(ctx BaseContext) error {
	panic("implement me")
}

func (r *Runtime) QueryPage(ctx PageContext) error {
	panic("implement me")
}

