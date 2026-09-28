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

    let { character, fullname } = $props()
    let affiliatedText = $derived.by(() => {
      if (character?.affiliation?.id !== undefined) {
        return character?.affiliation
      }
      return "Unaffiliated"
    })
</script>

<aside class="flex flex-col gap-3">
    <div class="flex flex-col">
        <div class="py-2.5 px-3 fantasy-border fantasy-border-brown">
            <img src={character.portrait} alt="Portrait" class="h-full overflow-hidden max-h-80 rounded-lg" />
        </div>

        <span class="mt-1 font-bold text-base">{fullname}</span>
        <span class="italic font-light text-base">{character.nickname}</span>
        <hr class="my-1" />
    </div>
    <div class="flex flex-col gap-2">
        <div class="flex flex-row gap-4 items-center">
            <Hourglass height="1.7rem" class="ml-4" />
            <div class="flex flex-col text-left">
                <span class="font-bold text-sm">Age</span>
                <span class="text-sm">{character.age}</span>
            </div>
        </div>
        <div class="flex flex-row gap-4 items-center">
            <VitruvianMan height="1.7rem" class="ml-4" />
            <div class="flex flex-col text-left">
                <span class="font-bold text-sm">Gender</span>
                <span class="text-sm">{character.gender.label}</span>
            </div>
        </div>
        <div class="flex flex-row gap-4 items-center">
            <DoubleFaceMask height="1.7rem" class="ml-4" />
            <div class="flex flex-col text-left">
                <span class="font-bold text-sm">Race</span>
                <span class="text-sm">{character.race.label}</span>
            </div>
        </div>
        <div class="flex flex-row gap-4 items-center">
            <LightHelm height="1.7rem" class="ml-4" />
            <div class="flex flex-col text-left">
                <span class="font-bold text-sm">Occupation</span>
                <span class="text-sm">{character.occupation.label}</span>
            </div>
        </div>
        <div class="flex flex-row gap-4 items-center">
            <BeveledStar height="1.7rem" class="ml-4" />
            <div class="flex flex-col text-left">
                <span class="font-bold text-sm">Alignment</span>
                <span class="text-sm">{character.alignment.label}</span>
            </div>
        </div>
        <div class="flex flex-row gap-4 items-center">
            <HeartShield height="1.7rem" class="ml-4" />
            <div class="flex flex-col text-left">
                <span class="font-bold text-sm">Status</span>
                <span class="text-sm">{character.status.label}</span>
            </div>
        </div>

        <hr class="my-1" />

        <div class="flex flex-row gap-4 items-center">
            <AffiliationEdgedShield height="1.7rem" class="ml-4" />
            <div class="flex flex-col text-left">
                <span class="font-bold text-sm">Affiliation</span>
                <span class="text-sm">{affiliatedText}</span>
            </div>
        </div>
        <div class="flex flex-row gap-4 items-center">
            <VerticalBanner height="1.7rem" class="ml-4" />
            <div class="flex flex-col text-left">
                <span class="font-bold text-sm">Faction</span>
                <span class="text-sm">{character.faction.name}</span>
            </div>
        </div>
        <div class="flex flex-row gap-4 items-center">
            <MedievalVillage height="1.7rem" class="ml-4" />
            <div class="flex flex-col text-left">
                <span class="font-bold text-sm">Birthplace</span>
                <span class="text-sm">{character.birthplace.name}</span>
            </div>
        </div>

        <div class="flex flex-row gap-4 items-center">
            <World height="1.7rem" class="ml-4" />
            <div class="flex flex-col text-left">
                <span class="font-bold text-sm">Current Location</span>
                <span class="text-sm">{character.current_location.name}</span>
            </div>
        </div>

    </div>
</aside>
