package hunter

import (
	"time"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
)

func (hunter *Hunter) getMongooseBiteConfig(rank int) core.SpellConfig {
	spellId := [5]int32{0, 1495, 14269, 14270, 14271}[rank]
	bonus := [5]float64{0, 15, 22, 37, 57}[rank]
	manaCost := [5]float64{0, 30, 40, 50, 65}[rank]
	level := [5]int{0, 16, 30, 44, 58}[rank]

	return core.SpellConfig{
		ActionID:      core.ActionID{SpellID: spellId},
		SpellSchool:   core.SpellSchoolPhysical,
		DefenseType:   core.DefenseTypeMelee,
		ProcMask:      core.ProcMaskMeleeMHSpecial,
		Flags:         core.SpellFlagMeleeMetrics | core.SpellFlagAPL | SpellFlagStrike,
		CastType:      proto.CastType_CastTypeMainHand,
		Rank:          rank,
		RequiredLevel: level,

		ManaCost: core.ManaCostOptions{FlatCost: manaCost},
		Cast: core.CastConfig{
			DefaultCast: core.Cast{GCD: core.GCDDefault},
			CD: core.Cooldown{
				Timer:    hunter.NewTimer(),
				Duration: time.Second * 5,
			},
			IgnoreHaste: true,
		},
		ExtraCastCondition: func(sim *core.Simulation, target *core.Unit) bool {
			return hunter.DefensiveState != nil && hunter.DefensiveState.IsActive()
		},

		CritDamageBonus:  hunter.mortalShots(),
		DamageMultiplier: 1,
		ThreatMultiplier: 1,
		BonusCoefficient: 1,

		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
			base := bonus + spell.Unit.MHNormalizedWeaponDamage(sim, spell.MeleeAttackPower(target))
			result := spell.CalcAndDealDamage(sim, target, base, spell.OutcomeMeleeWeaponSpecialHitAndCrit)
			if result.Landed() && hunter.DefensiveState != nil && hunter.DefensiveState.IsActive() {
				hunter.DefensiveState.Deactivate(sim)
			}
		},
	}
}

func (hunter *Hunter) registerMongooseBiteSpell() {
	hunter.DefensiveState = hunter.RegisterAura(core.Aura{
		Label:    "Defensive State",
		ActionID: core.ActionID{SpellID: 5302},
		Duration: time.Second * 5,
	})

	best := 0
	for rank := 1; rank <= 4; rank++ {
		if [5]int{0, 16, 30, 44, 58}[rank] <= int(hunter.Level) {
			best = rank
		}
	}
	if best == 0 {
		return
	}
	hunter.MongooseBite = hunter.GetOrRegisterSpell(hunter.getMongooseBiteConfig(best))
}
