package exile

import (
	"fmt"

	"github.com/genshinsim/gcsim/pkg/core"
	"github.com/genshinsim/gcsim/pkg/core/attributes"
	"github.com/genshinsim/gcsim/pkg/core/event"
	"github.com/genshinsim/gcsim/pkg/core/keys"
	"github.com/genshinsim/gcsim/pkg/core/player/artifact"
	"github.com/genshinsim/gcsim/pkg/core/player/character"
)

func init() {
	core.RegisterSetFunc(keys.TheExile, NewSet)
}

type Set struct {
	Index int
}

func (s *Set) SetIndex(idx int) { s.Index = idx }
func (s *Set) Init() error      { return nil }

// 2-Piece Bonus: Energy Recharge +20%.
// 4-Piece Bonus: Using an Elemental Burst regenerates 2 Energy for all party members (excluding the wearer) every 2s for 6s. This effect cannot stack.
func NewSet(core *core.Core, char *character.CharWrapper, count int, param map[string]int) (artifact.Set, error) {
	s := Set{}

	if count >= 2 {
		m := make([]float64, attributes.EndStatType)
		m[attributes.ER] = 0.20
		char.AddStatMod("exile-2pc", -1, attributes.ER, func() ([]float64, bool) {
			return m, true
		})
	}
	if count >= 4 {
		// TODO: does multiple exile holders extend the duration?
		core.Events.Subscribe(event.PreBurst, func(args ...interface{}) bool {
			if core.Player.Active() != char.Index {
				return false
			}
			if core.Status.Duration("exile") > 0 {
				return false
			}
			core.Status.Add("exile", 6*60)

			for _, c := range core.Player.Chars() {
				this := c
				if char.Index == this.Index {
					continue
				}
				// 3 ticks
				for i := 120; i <= 360; i += 120 {
					core.Tasks.Add(func() {
						this.AddEnergy("exile-4pc", 2)
					}, i)
				}
			}

			return false
		}, fmt.Sprintf("exile-4pc-%v", char.Base.Name))
	}

	return &s, nil
}
