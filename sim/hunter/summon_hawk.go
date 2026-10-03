package hunter

import (
	"time"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
)

func (hunter *Hunter) registerSummonHawkSpell(sharedTimer *core.Timer) {
	if !hunter.Talents.SummonHawk {
		return
	}
	if sharedTimer == nil {
		sharedTimer = hunter.NewTimer()
	}

	baseDamage := 32.0
	if hunter.Level >= 60 {
		baseDamage = 108
	} else if hunter.Level >= 48 {
		baseDamage = 80
	} else if hunter.Level >= 36 {
		baseDamage = 55
	}

	hawkMult := 1 + 0.03*float64(hunter.Talents.UnleashedFury)
	hawkCrit := 2 * float64(hunter.Talents.Ferocity) * core.CritRatingPerCritChance

	var swingBase float64
	var swingTarget *core.Unit
	expireAt := []time.Duration{0, 0}

	hawkSwing := hunter.GetOrRegisterSpell(core.SpellConfig{
		ActionID:    core.ActionID{SpellID: 1293241}.WithTag(1),
		SpellSchool: core.SpellSchoolPhysical,
		DefenseType: core.DefenseTypeRanged,
		ProcMask:    core.ProcMaskRangedSpecial,
		Flags:       core.SpellFlagMeleeMetrics | core.SpellFlagPassiveSpell,
		CastType:    proto.CastType_CastTypeRanged,

		BonusCritRating:  hawkCrit,
		DamageMultiplier: hawkMult,
		ThreatMultiplier: 1,
		BonusCoefficient: 1,

		ApplyEffects: func(sim *core.Simulation, _ *core.Unit, spell *core.Spell) {
			if swingTarget == nil {
				return
			}
			spell.CalcAndDealDamage(sim, swingTarget, swingBase, spell.OutcomeRangedHitAndCrit)
		},
	})

	hunter.SummonHawk = hunter.GetOrRegisterSpell(core.SpellConfig{
		ActionID:     core.ActionID{SpellID: 1293241},
		SpellSchool:  core.SpellSchoolPhysical,
		DefenseType:  core.DefenseTypeRanged,
		ProcMask:     core.ProcMaskRangedSpecial,
		Flags:        core.SpellFlagMeleeMetrics | core.SpellFlagAPL | SpellFlagShot,
		CastType:     proto.CastType_CastTypeRanged,
		MissileSpeed: 24,

		ManaCost: core.ManaCostOptions{FlatCost: 80},
		Cast: core.CastConfig{
			DefaultCast: core.Cast{GCD: core.GCDDefault},
			IgnoreHaste: true,
			CD: core.Cooldown{
				Timer:    sharedTimer,
				Duration: time.Second * 6,
			},
		},

		DamageMultiplier: hawkMult,
		ThreatMultiplier: 1,
		BonusCoefficient: 1,

		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
			initial := baseDamage + 0.05*spell.RangedAttackPower(target, false)
			spell.CalcAndDealDamage(sim, target, initial, spell.OutcomeRangedHitAndCrit)

			swingBase = 20 + 0.01*spell.RangedAttackPower(target, false)
			swingTarget = target

			for i := 0; i < 2; i++ {
				if expireAt[i] > sim.CurrentTime {
					continue
				}
				slot := i
				expireAt[slot] = sim.CurrentTime + 18*time.Second

				var pa *core.PendingAction
				pa = &core.PendingAction{
					NextActionAt: sim.CurrentTime + 2*time.Second,
					Priority:     core.ActionPriorityDOT,
					OnAction: func(sim *core.Simulation) {
						if sim.CurrentTime > expireAt[slot] || swingTarget == nil {
							expireAt[slot] = 0
							return
						}
						hawkSwing.Cast(sim, swingTarget)
						pa.NextActionAt = sim.CurrentTime + 2*time.Second
						sim.AddPendingAction(pa)
					},
				}
				sim.AddPendingAction(pa)
				break
			}
		},
	})
}
