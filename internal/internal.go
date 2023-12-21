package internal

type Internal interface {
	IgnoreMethod(a int)
	ReportMethod(a int)
}

type IgnoreInternal interface {
	Method(a int)
}
