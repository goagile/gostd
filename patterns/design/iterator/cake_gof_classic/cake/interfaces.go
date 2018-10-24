package cake

type Iterabler interface {
  Iterator() Iterator
}

type Iterator interface {
    Next() MenuItem
    HasNext() bool
}
