package cake

type Iterator interface {
    Next() MenuItem
    HasNext() bool
}
