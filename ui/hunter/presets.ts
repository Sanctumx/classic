import { Phase } from '../core/constants/other.js';
import * as PresetUtils from '../core/preset_utils.js';
import {
	AgilityElixir,
	Alcohol,
	AttackPowerBuff,
	Conjured,
	Consumes,
	Debuffs,
	Food,
	HealthElixir,
	IndividualBuffs,
	ManaRegenElixir,
	Potions,
	Profession,
	Race,
	RaidBuffs,
	SaygesFortune,
	SpellPowerBuff,
	StrengthBuff,
	TristateEffect,
	WeaponImbue,
	ZanzaBuff,
} from '../core/proto/common.js';
import {
	Hunter_Options as HunterOptions,
	Hunter_Options_Ammo as Ammo,
	Hunter_Options_PetAttackSpeed as PetAttackSpeed,
	Hunter_Options_PetType as PetType,
	Hunter_Options_QuiverBonus,
} from '../core/proto/hunter.js';
import { SavedTalents } from '../core/proto/ui.js';
import MMWeaveAPL from './apls/MMWeave.apl.json';
import SurvDWAPL from './apls/SurvDW.apl.json';
import SurvDWGear from './gear_sets/surv_dw.gear.json';
import MMWeaveGear from './gear_sets/mm_weave.gear.json';

export const GearMMWeave = PresetUtils.makePresetGear('MM Weave', MMWeaveGear);
export const GearSurvDW = PresetUtils.makePresetGear('Surv DW', SurvDWGear);


export const GearPresets = {
	[Phase.Phase1]: [GearSurvDW, GearMMWeave],
};

export const APLSurvDW = PresetUtils.makePresetAPLRotation('Surv DW', SurvDWAPL);
export const APLMMWeave = PresetUtils.makePresetAPLRotation('MM Weave', MMWeaveAPL);

export const DefaultGear = GearSurvDW;

export const APLPresets = {
	[Phase.Phase1]: [APLMMWeave, APLSurvDW],
};

export const DefaultAPL = APLSurvDW;

export const SurvDwHawk = PresetUtils.makePresetTalents(
	'Surv DW Hawk',
	SavedTalents.create({ talentsString: '53200005021-005005-5002302300500201' }),
);
export const SurvDw = PresetUtils.makePresetTalents(
	'Surv DW',
	SavedTalents.create({ talentsString: '5-005005201-500230230050222151' }),
);
export const MmWeave = PresetUtils.makePresetTalents(
	'MM Weave',
	SavedTalents.create({ talentsString: '5-005005201-500230230050222151' }),
);
export const TalentPresets = {
	[Phase.Phase1]: [SurvDwHawk, SurvDw, MmWeave],
};

export const DefaultTalents = SurvDwHawk;

export const DefaultOptions = HunterOptions.create({
	ammo: Ammo.ThoriumHeadedArrow,
	quiverBonus: Hunter_Options_QuiverBonus.Speed15,
	petAttackSpeed: PetAttackSpeed.OneTwo,
	petType: PetType.Cat,
	petUptime: 1,
});

export const DefaultConsumes = Consumes.create({
	agilityElixir: AgilityElixir.ElixirOfTheMongoose,
	alcohol: Alcohol.AlcoholRumseyRumBlackLabel,
	attackPowerBuff: AttackPowerBuff.WinterfallFirewater,
	defaultConjured: Conjured.ConjuredDemonicRune,
	defaultPotion: Potions.MajorFrenzyPotion,
	dragonBreathChili: true,
	food: Food.FoodSmokedDesertDumpling,
	healthElixir: HealthElixir.ElixirOfFortitude,
	mainHandImbue: WeaponImbue.ElementalSharpeningStone,
	offHandImbue: WeaponImbue.ElementalSharpeningStone,
	manaRegenElixir: ManaRegenElixir.MagebloodPotion,
	spellPowerBuff: SpellPowerBuff.GreaterArcaneElixir,
	strengthBuff: StrengthBuff.ElixirOfGiants,
	zanzaBuff: ZanzaBuff.GroundScorpokAssay,
});

export const DefaultRaidBuffs = RaidBuffs.create({
	arcaneBrilliance: true,
	battleShout: TristateEffect.TristateEffectImproved,
	divineSpirit: true,
	fireResistanceAura: true,
	fireResistanceTotem: true,
	giftOfTheWild: TristateEffect.TristateEffectImproved,
	leaderOfThePack: true,
	manaSpringTotem: TristateEffect.TristateEffectRegular,
	strengthOfEarthTotem: TristateEffect.TristateEffectImproved,
	trueshotAura: true,
	windfuryTotem: true,
});

export const DefaultIndividualBuffs = IndividualBuffs.create({
	blessingOfKings: true,
	blessingOfMight: TristateEffect.TristateEffectImproved,
	blessingOfWisdom: TristateEffect.TristateEffectRegular,
	fengusFerocity: false,
	moldarsMoxie: false,
	rallyingCryOfTheDragonslayer: false,
	saygesFortune: SaygesFortune.SaygesUnknown, // or omit
	slipkiksSavvy: false,
	songflowerSerenade: false,
	spiritOfZandalar: false,
	warchiefsBlessing: false,
});

export const DefaultDebuffs = Debuffs.create({
	curseOfRecklessness: false,
	exposeArmor: TristateEffect.TristateEffectImproved,
	faerieFire: true,
	huntersMark: TristateEffect.TristateEffectRegular,
	judgementOfWisdom: true,
	sunderArmor: true,
	giftOfArthas: true,
});

export const OtherDefaults = {
	distanceFromTarget: 5,
	profession1: Profession.Enchanting,
	profession2: Profession.Engineering,
	race: Race.RaceOrc,
};