package hunter

import (
	"time"

	"github.com/wowsims/classic/sim/core"
)

func (hunter *Hunter) registerSummonHawkSpell(timer *core.Timer) {
	if !hunter.Talents.SummonHawk {
		return
	}

	hunter.SummonHawk = hunter.GetOrRegisterSpell(core.SpellConfig{
		ActionID:    core.ActionID{SpellID: 1293241},
		SpellSchool: core.SpellSchoolPhysical,
		DefenseType: core.DefenseTypeRanged,
		ProcMask:    core.ProcMaskEmpty,
		Flags:       core.SpellFlagAPL | core.SpellFlagPureDot,
		ManaCost: core.ManaCostOptions{
			FlatCost: 190, // rank 4 @ 60; 80 / 105 / 135 / 190
		},
		DamageMultiplier: 1 + 0.03*float64(hunter.Talents.UnleashedFury),
		BonusCritRating:  2 * float64(hunter.Talents.Ferocity) * core.CritRatingPerCritChance,
		ThreatMultiplier: 1,

		Cast: core.CastConfig{
			DefaultCast: core.Cast{GCD: core.GCDDefault},
			CD: core.Cooldown{
				Timer:    timer,
				Duration: time.Second * 6,
			},
		},

		Dot: core.DotConfig{
			Aura: core.Aura{
				Label:     "Summon Hawk",
				ActionID:  core.ActionID{SpellID: 1293241},
				MaxStacks: 2,
			},
			NumberOfTicks: 6,
			TickLength:    time.Second * 3,
			OnTick: func(sim *core.Simulation, target *core.Unit, dot *core.Dot) {
				dot.CalcAndDealPeriodicSnapshotDamage(sim, target, dot.OutcomeTick)
			},
		},

		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
			tick := 108 + spell.RangedAttackPower(target, false)*0.05

			spell.CalcAndDealDamage(sim, target, tick, spell.OutcomeRangedHitAndCrit)

			dot := spell.Dot(target)
			if dot.IsActive() {
				if dot.GetStacks() < 2 {
					dot.AddStack(sim)
				}
				dot.Refresh(sim)
			} else {
				dot.Apply(sim)
				dot.SetStacks(sim, 1)
			}
			dot.SnapshotBaseDamage = tick * float64(dot.GetStacks())
			dot.SnapshotAttackerMultiplier = spell.AttackerDamageMultiplier(spell.Unit.AttackTables[target.UnitIndex][spell.CastType], true)
		},
	})
}
