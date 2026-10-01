package hunter

import (
	"time"

	"github.com/wowsims/classic/sim/core"
)

const RaptorStrikeRanks = 8

var RaptorStrikeSpellId = [RaptorStrikeRanks + 1]int32{0, 2973, 14260, 14261, 14262, 14263, 14264, 14265, 14266}
var RaptorStrikeSpellIdMeleeSpecialist = [RaptorStrikeRanks + 1]int32{0, 415335, 415336, 415337, 415338, 415340, 415341, 415342, 415343}
var RaptorStrikeBaseDamage = [RaptorStrikeRanks + 1]float64{0, 5, 11, 21, 34, 50, 80, 110, 140}
var RaptorStrikeManaCost = [RaptorStrikeRanks + 1]float64{0, 15, 25, 35, 45, 55, 70, 80, 100}
var RaptorStrikeLevel = [RaptorStrikeRanks + 1]int{0, 1, 8, 16, 24, 32, 40, 48, 56}

func (hunter *Hunter) TryRaptorStrike(sim *core.Simulation, mhSwingSpell *core.Spell) *core.Spell {
	if hunter.curQueueAura != nil &&
		hunter.RaptorStrikeHit != nil &&
		hunter.RaptorStrikeHit.CanCast(sim, hunter.CurrentTarget) {
		return hunter.RaptorStrikeHit
	}
	return mhSwingSpell
}

func (hunter *Hunter) getRaptorStrikeConfig(rank int) core.SpellConfig {
	spellID := RaptorStrikeSpellId[rank]
	level := RaptorStrikeLevel[rank]

	hunter.RaptorStrikeHit = hunter.newRaptorStrikeHitSpell(rank)

	return core.SpellConfig{
		SpellCode:     SpellCode_HunterRaptorStrike,
		ActionID:      core.ActionID{SpellID: spellID},
		SpellSchool:   core.SpellSchoolPhysical,
		DefenseType:   core.DefenseTypeMelee,
		ProcMask:      core.ProcMaskEmpty,
		Flags:         core.SpellFlagAPL | SpellFlagStrike,
		Rank:          rank,
		RequiredLevel: level,

		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				GCD: core.GCDDefault,
			},
		},
		ExtraCastCondition: func(sim *core.Simulation, target *core.Unit) bool {
			return hunter.curQueueAura == nil &&
				hunter.DistanceFromTarget <= core.MaxMeleeAttackDistance &&
				hunter.RaptorStrikeHit.IsReady(sim)
		},

		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
			if hunter.raptorQueueAura != nil {
				hunter.raptorQueueAura.Activate(sim)
			}
		},
	}
}

func (hunter *Hunter) newRaptorStrikeHitSpell(rank int) *core.Spell {
	spellID := RaptorStrikeSpellId[rank]
	baseDamage := RaptorStrikeBaseDamage[rank]
	manaCost := RaptorStrikeManaCost[rank]

	return hunter.RegisterSpell(core.SpellConfig{
		SpellCode:        SpellCode_HunterRaptorStrikeHit,
		ActionID:         core.ActionID{SpellID: spellID}.WithTag(1),
		SpellSchool:      core.SpellSchoolPhysical,
		DefenseType:      core.DefenseTypeMelee,
		ProcMask:         core.ProcMaskMeleeMHSpecial,
		Flags:            core.SpellFlagMeleeMetrics | core.SpellFlagNoOnCastComplete,
		CritDamageBonus:  hunter.mortalShots(),
		DamageMultiplier: 1,
		BonusCoefficient: 1,

		ManaCost: core.ManaCostOptions{FlatCost: manaCost},
		Cast: core.CastConfig{
			CD: core.Cooldown{
				Timer:    hunter.NewTimer(),
				Duration: time.Second * 6,
			},
		},

		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
			damage := baseDamage + hunter.MHWeaponDamage(sim, spell.MeleeAttackPower(target))
			result := spell.CalcAndDealDamage(sim, target, damage, spell.OutcomeMeleeWeaponSpecialHitAndCrit)
			if hunter.curQueueAura != nil {
				hunter.curQueueAura.Deactivate(sim)
			}
			if result.Landed() {
				hunter.AutoAttacks.RestartRangedSwing(sim)
			}
		},
	})
}

func (hunter *Hunter) makeQueueSpellsAndAura() *core.Spell {
	if hunter.raptorQueueAura != nil {
		return nil
	}

	queueAura := hunter.RegisterAura(core.Aura{
		Label:    "Raptor Strike Queued",
		ActionID: core.ActionID{SpellID: 14266},
		Duration: core.NeverExpires,
		OnGain: func(aura *core.Aura, sim *core.Simulation) {
			hunter.curQueueAura = aura
			hunter.curQueuedAutoSpell = hunter.RaptorStrike
		},
		OnExpire: func(aura *core.Aura, sim *core.Simulation) {
			if hunter.curQueueAura == aura {
				hunter.curQueueAura = nil
				hunter.curQueuedAutoSpell = nil
			}
		},
	})
	hunter.raptorQueueAura = queueAura
	return nil
}

func (hunter *Hunter) registerRaptorStrikeSpell() {
	rank := map[int32]int{
		25: 4,
		40: 6,
		50: 7,
		60: 8,
	}[hunter.Level]
	if rank == 0 {
		rank = 8
	}

	hunter.RaptorStrike = hunter.GetOrRegisterSpell(hunter.getRaptorStrikeConfig(rank))
	hunter.makeQueueSpellsAndAura()
}
