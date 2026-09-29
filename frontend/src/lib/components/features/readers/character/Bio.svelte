<script lang="ts">
    import VerticalBanner from "@iconify-svelte/game-icons/components/v/vertical-banner.svelte";
    import AffiliationEdgedShield from "@iconify-svelte/game-icons/components/e/edged-shield.svelte";
    import MedievalVillage from "@iconify-svelte/game-icons/components/m/medieval-village-01.svelte"
    import World from "@iconify-svelte/game-icons/components/w/world.svelte";
    import LightHelm from "@iconify-svelte/game-icons/components/l/light-helm.svelte"
    import Hourglass from "@iconify-svelte/game-icons/components/h/hourglass.svelte"
    import HeartShield from "@iconify-svelte/game-icons/components/h/heart-shield.svelte"
    import BeveledStar from "@iconify-svelte/game-icons/components/b/beveled-star.svelte"
    import DoubleFaceMask from "@iconify-svelte/game-icons/components/d/double-face-mask.svelte"
    import VitruvianMan from "@iconify-svelte/game-icons/components/v/vitruvian-man.svelte"
    import DividedSquare from "@iconify-svelte/game-icons/components/d/divided-square.svelte"
    import BadgeBG from "$lib/../assets/images/badge-bg.png"
    import CastleTown from "$lib/../assets/images/castle-town.png"

    let { character, fullname } = $props()
    let affiliatedText = $derived.by(() => {
      if (character?.affiliation?.id !== undefined) {
        return character?.affiliation
      }
      return "Unaffiliated"
    })
</script>

<aside class="flex flex-col">

    <div class="flex flex-col">
        <div class="py-2.5 px-3 fantasy-border fantasy-border-brown">
            <img src={character.portrait} alt="Portrait" class="h-full overflow-hidden max-h-80 rounded-lg" />
        </div>

        <span class="mt-1 font-bold text-base">{fullname}</span>
        <span class="italic font-light text-base">{character.nickname}</span>
        <div class="mb-3 mt-4 col-span-2 flex items-center justify-center pointer-events-none z-0">
            <div class="flex-grow border-b border-[#8b7355]/40"></div>
            <DividedSquare height="1em" class="mx-2 shrink-0 text-[#8b7355]" />
            <div class="flex-grow border-b border-[#8b7355]/40"></div>
        </div>
    </div>
    <div class="flex flex-col gap-2">
        <div class="flex flex-row gap-4 items-center">
            <div class="grid place-items-center relative shrink-0 ml-2">
                <img src={BadgeBG} alt="" class="col-start-1 row-start-1 w-11 h-10 object-contain" />
                <Hourglass height="1.7rem" class="col-start-1 row-start-1 text-[#2c251b] z-10" />
            </div>
            <div class="flex flex-col text-left">
                <span class="font-bold text-sm">Age</span>
                <span class="text-sm">{character.age}</span>
            </div>
        </div>
        <div class="flex flex-row gap-4 items-center">
            <div class="grid place-items-center relative shrink-0 ml-2">
                <img src={BadgeBG} alt="" class="col-start-1 row-start-1 w-11 h-10 object-contain" />
                <VitruvianMan height="1.7rem" class="col-start-1 row-start-1 text-[#2c251b] z-10 mb-0.5" />
            </div>
            <div class="flex flex-col text-left">
                <span class="font-bold text-sm">Gender</span>
                <span class="text-sm">{character.gender.label}</span>
            </div>
        </div>
        <div class="flex flex-row gap-4 items-center">
            <div class="grid place-items-center relative shrink-0 ml-2">
                <img src={BadgeBG} alt="" class="col-start-1 row-start-1 w-11 h-10 object-contain" />
                <DoubleFaceMask height="1.7rem" class="col-start-1 row-start-1 text-[#2c251b] z-10" />
            </div>
            <div class="flex flex-col text-left">
                <span class="font-bold text-sm">Race</span>
                <span class="text-sm">{character.race.label}</span>
            </div>
        </div>
        <div class="flex flex-row gap-4 items-center">
            <div class="grid place-items-center relative shrink-0 ml-2">
                <img src={BadgeBG} alt="" class="col-start-1 row-start-1 w-11 h-10 object-contain" />
                <LightHelm height="1.7rem" class="col-start-1 row-start-1 text-[#2c251b] z-10" />
            </div>
            <div class="flex flex-col text-left">
                <span class="font-bold text-sm">Occupation</span>
                <span class="text-sm">{character.occupation.label}</span>
            </div>
        </div>
        <div class="flex flex-row gap-4 items-center">
            <div class="grid place-items-center relative shrink-0 ml-2">
                <img src={BadgeBG} alt="" class="col-start-1 row-start-1 w-11 h-10 object-contain" />
                <BeveledStar height="1.7rem" class="col-start-1 row-start-1 text-[#2c251b] z-10" />
            </div>
            <div class="flex flex-col text-left">
                <span class="font-bold text-sm">Alignment</span>
                <span class="text-sm">{character.alignment.label}</span>
            </div>
        </div>
        <div class="flex flex-row gap-4 items-center">
            <div class="grid place-items-center relative shrink-0 ml-2">
                <img src={BadgeBG} alt="" class="col-start-1 row-start-1 w-11 h-10 object-contain" />
                <HeartShield height="1.7rem" class="col-start-1 row-start-1 text-[#2c251b] z-10" />
            </div>
            <div class="flex flex-col text-left">
                <span class="font-bold text-sm">Status</span>
                <span class="text-sm">{character.status.label}</span>
            </div>
        </div>

        <div class="mb-3 mt-4 col-span-2 flex items-center justify-center pointer-events-none z-0">
            <div class="flex-grow border-b border-[#8b7355]/40"></div>
            <span class="ml-1 text-sm text-[#8b7355]">Faction</span>
            <DividedSquare height="1em" class="mx-2 shrink-0 text-[#8b7355]" />
            <span class="mr-1 text-sm text-[#8b7355]">World</span>
            <div class="flex-grow border-b border-[#8b7355]/40"></div>
        </div>

        <div class="flex flex-row gap-4 items-center">
            <div class="grid place-items-center relative shrink-0 ml-2">
                <img src={BadgeBG} alt="" class="col-start-1 row-start-1 w-11 h-10 object-contain" />
                <AffiliationEdgedShield height="1.7rem" class="col-start-1 row-start-1 text-[#2c251b] z-10 mt-0.5" />
            </div>
            <div class="flex flex-col text-left">
                <span class="font-bold text-sm">Affiliation</span>
                <span class="text-sm">{affiliatedText}</span>
            </div>
        </div>
        <div class="flex flex-row gap-4 items-center">
            <div class="grid place-items-center relative shrink-0 ml-2">
                <img src={BadgeBG} alt="" class="col-start-1 row-start-1 w-11 h-10 object-contain" />
                <VerticalBanner height="1.7rem" class="col-start-1 row-start-1 text-[#2c251b] z-10 mt-0.5" />
            </div>
            <div class="flex flex-col text-left">
                <span class="font-bold text-sm">Faction</span>
                <span class="text-sm">{character.faction.name}</span>
            </div>
        </div>
        <div class="flex flex-row gap-4 items-center">
            <div class="grid place-items-center relative shrink-0 ml-2">
                <img src={BadgeBG} alt="" class="col-start-1 row-start-1 w-11 h-10 object-contain" />
                <MedievalVillage height="1.7rem" class="col-start-1 row-start-1 text-[#2c251b] z-10" />
            </div>
            <div class="flex flex-col text-left">
                <span class="font-bold text-sm">Birthplace</span>
                <span class="text-sm">{character.birthplace.name}</span>
            </div>
        </div>

        <div class="flex flex-row gap-4 items-center">
            <div class="grid place-items-center relative shrink-0 ml-2">
                <img src={BadgeBG} alt="" class="col-start-1 row-start-1 w-11 h-10 object-contain" />
                <World height="1.7rem" class="col-start-1 row-start-1 text-[#2c251b] z-10" />
            </div>
            <div class="flex flex-col text-left">
                <span class="font-bold text-sm">Current Location</span>
                <span class="text-sm">{character.current_location.name}</span>
            </div>
        </div>

        <img src={CastleTown} alt="" class="opacity-50">

    </div>

</aside>
