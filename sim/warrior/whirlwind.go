package warrior

import (
	"time"

	"github.com/wowsims/classic/sim/core"
)

func (warrior *Warrior) registerWhirlwindSpell() {
	targets := min(4, warrior.Env.GetNumTargets())
	mhResults := make([]*core.SpellResult, targets)
	ohResults := make([]*core.SpellResult, targets)

	warrior.WhirlwindOH = warrior.RegisterSpell(AnyStance, core.SpellConfig{
		ActionID:    core.ActionID{SpellID: 1680}.WithTag(2),
		SpellSchool: core.SpellSchoolPhysical,
		DefenseType: core.DefenseTypeMelee,
		ProcMask:    core.ProcMaskMeleeOHSpecial,
		Flags:       core.SpellFlagMeleeMetrics | core.SpellFlagAPL | SpellFlagOffensive | core.SpellFlagNoOnCastComplete,

		CritDamageBonus:  warrior.impale(),
		DamageMultiplier: 1,
		ThreatMultiplier: 1.25,
		BonusCoefficient: 1,

		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
			cur := target
			for idx := range ohResults {
				baseDamage := spell.Unit.OHNormalizedWeaponDamage(sim, spell.MeleeAttackPower(cur))
				ohResults[idx] = spell.CalcDamage(sim, cur, baseDamage, spell.OutcomeMeleeSpecialHitAndCrit)
				cur = sim.Environment.NextTargetUnit(cur)
			}
			for _, result := range ohResults {
				spell.DealDamage(sim, result)
			}
		},
	})

	warrior.Whirlwind = warrior.RegisterSpell(BerserkerStance, core.SpellConfig{
		SpellCode:   SpellCode_WarriorWhirlwind,
		ActionID:    core.ActionID{SpellID: 1680},
		SpellSchool: core.SpellSchoolPhysical,
		DefenseType: core.DefenseTypeMelee,
		ProcMask:    core.ProcMaskMeleeMHSpecial,
		Flags:       core.SpellFlagMeleeMetrics | core.SpellFlagAPL | SpellFlagOffensive,

		RageCost: core.RageCostOptions{
			Cost: 25,
		},
		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				GCD: core.GCDDefault,
			},
			IgnoreHaste: true,
			CD: core.Cooldown{
				Timer:    warrior.NewTimer(),
				Duration: time.Second * 10,
			},
		},
		CritDamageBonus:  warrior.impale(),
		DamageMultiplier: 1,
		ThreatMultiplier: 1.25,
		BonusCoefficient: 1,

		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
			cur := target
			for idx := range mhResults {
				baseDamage := spell.Unit.MHNormalizedWeaponDamage(sim, spell.MeleeAttackPower(cur))
				mhResults[idx] = spell.CalcDamage(sim, cur, baseDamage, spell.OutcomeMeleeWeaponSpecialHitAndCrit)
				cur = sim.Environment.NextTargetUnit(cur)
			}
			for _, result := range mhResults {
				spell.DealDamage(sim, result)
			}

			if warrior.AutoAttacks.OH() != nil && warrior.AutoAttacks.OH().SwingSpeed > 0 {
				warrior.WhirlwindOH.Cast(sim, target)
			}
		},
	})
}
