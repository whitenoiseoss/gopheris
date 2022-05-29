package internal

func NewBitmask() *Bitmask {
	return &Bitmask{
		Flags: 0,
	}
}

type Bitmask struct {
	Flags int
}

func (b *Bitmask) AddFlag(flag int) {
	b.Flags |= flag
}

func (b *Bitmask) AddFlags(flags ...int) {
	for _, flag := range flags {
		b.Flags |= flag
	}
}

func (b *Bitmask) CheckFlag(flag int) bool {
	return b.Flags&flag != 0
}

func (b *Bitmask) CheckFlags(flags ...int) bool {
	for _, flag := range flags {
		if b.Flags&flag == 0 {
			return false
		}
	}

	return true
}

func (b *Bitmask) RemoveFlag(flag int) {
	b.Flags &= ^flag
}

func (b *Bitmask) RemoveFlags(flags ...int) {
	for _, flag := range flags {
		b.Flags &= ^flag
	}
}

func (b *Bitmask) ToggleFlag(flag int) {
	b.Flags ^= flag
}

func (b *Bitmask) ToggleFlags(flags ...int) {
	for _, flag := range flags {
		b.Flags ^= flag
	}
}
