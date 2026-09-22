<script lang="ts">
  import { onMount } from "svelte";
  import corners from "$lib/../assets/images/corners.png"
  import RichEditor from "$lib/components/ui/RichEditor.svelte";
  import DataSelector from "$lib/components/ui/DataSelector.svelte";
  import StatusQuip from "$lib/components/ui/StatusQuip.svelte";
  import { LookupItem, TraitItem } from "$wails/world-builder/internal/services/models";
  import { Affiliation, Character, Faction, Location } from "$wails/world-builder/internal/models/models";
  import { loadCharacterFormData } from '$lib/functions/services'
    import Divider from "$lib/components/ui/Divider.svelte";
    import PortraitPicker from "$lib/components/ui/PortraitPicker.svelte";
    import MultiDataSelector from "$lib/components/ui/MultiDataSelector.svelte";
    import { FetchCharacter, SaveCharacter } from "$wails/world-builder/app";
    import ParchmentTitle from "$lib/components/ui/ParchmentTitle.svelte";
    import AnvilImpact from "@iconify-svelte/game-icons/components/a/anvil-impact.svelte";

  let { id } = $props()
  let formState = $state({
    firstname: '',
    middlename: '',
    surname: '',
    nickname: '',

    age: 'Unknown',
    description: 'What does he looks like?',
    portrait: '',
    goals: 'What are his objectives?',

    gender_id: null as number | null,
    race_id: null as number | null,
    occupation_id: null as number | null,
    alignment_id: null as number | null,
    status_id: null as number | null,

    affiliation_id: null as number | null,
    faction_id: null as number | null,
    birthplace_location_id: null as number | null,
    current_location_id: null as number | null,

    personality_traits: [] as number[],
    strengths: [] as number[],
    flaws: [] as number[],
    weaknesses: [] as number[],
    fears: [] as number[],

    relationships: [],
    backstories: []
  })

  let genders = $state<LookupItem[]>([])
  let races = $state<LookupItem[]>([])
  let occupations = $state<LookupItem[]>([])
  let alignments = $state<LookupItem[]>([])
  let statuses = $state<LookupItem[]>([])

  let personalityTraits = $state<TraitItem[]>([])
  let strengths = $state<TraitItem[]>([])
  let flaws = $state<TraitItem[]>([])
  let weaknesses = $state<TraitItem[]>([])
  let fears = $state<TraitItem[]>([])

  let locations = $state<Location[]>([])
  let affiliations = $state<Affiliation[]>([])
  let factions = $state<Faction[]>([])

  // For affiliation belonging to a faction
  let isFactionLocked = $state(false)

  const mapToIds = <T>(ids: number[]) => ids.map((id) => ({ id  } as unknown as T));

  async function handleSubmit(e: Event) {
    e.preventDefault()
    const rawBase64 = formState.portrait.split(',')[1]
    try {
      const payload = new Character({
        id: id ? Number(id) : undefined,
        ...formState,
        portrait: rawBase64,

        personality_traits: mapToIds(formState.personality_traits),
        strengths: mapToIds(formState.strengths),
        flaws: mapToIds(formState.flaws),
        weaknesses: mapToIds(formState.weaknesses),
        fears: mapToIds(formState.fears),

        relationships: [],
        backstories: []
      })
      await SaveCharacter(payload)
    } catch(err) {
      console.error(err)
    }
  }

  let isFormReady;
  let missingFieldType;

  async function fetchData() {
    try{
      const data = await loadCharacterFormData()
      genders = data.genders
      races = data.races
      occupations = data.occupations
      alignments = data.alignments
      statuses = data.statuses

      personalityTraits = data.personalityTraits
      strengths = data.strengths
      flaws = data.flaws
      weaknesses = data.weaknesses
      fears = data.fears

      locations = data.locations
      affiliations = data.affiliations
      factions = data.factions

      if(id) {
        const charToEdit = await FetchCharacter(Number(id))
        if (charToEdit) {
          Object.assign(formState, charToEdit)

          const extractIds = (items: any[]) => items?.map(item => item.id) || [];

          formState.personality_traits = extractIds(charToEdit.personality_traits)
          formState.strengths = extractIds(charToEdit.strengths)
          formState.flaws = extractIds(charToEdit.flaws)
          formState.weaknesses = extractIds(charToEdit.weaknesses)
          formState.fears = extractIds(charToEdit.fears)
        }
      }
    } catch(err) {
      console.error(err)
    }
  }

  onMount(() => {
    fetchData()
  })

  $effect(() => {
    if (formState.affiliation_id) {
      const selectedAffiliation = affiliations.find(a => a.id === formState.affiliation_id)

      if (selectedAffiliation && selectedAffiliation.faction_id) {
        // Affiliation belongs to a faction, lock input
        formState.faction_id = selectedAffiliation.faction_id
        isFactionLocked = true
      } else {
        // Independant affiliation, can belong to a faction
        formState.faction_id = null
        isFactionLocked = false
      }

    } else {
      // Nothing selected, just in case
      isFactionLocked = false
    }
  })
</script>

<section class="h-screen grid grid-cols-1 items-center justify-center px-2 overflow-y-auto">
    <div class="relative z-20 -mb-10 pointer-events-none">
        <ParchmentTitle title="Character" subtitle="I don't know what to write for characters, really..." />
    </div>
    <form onsubmit={handleSubmit}>
        <div class="relative w-full h-auto shadow-2xl rounded-sm border-4 border-[#1a0f0f] bg-[#2b190c] p-1 md:p-1 grid grid-cols-1 grid-rows-1">
            <div class="wood-texture-bg"></div>

            <img src="{corners}" alt="" class="absolute -top-4 -left-4 w-15 h-15 pointer-events-none z-10">
            <img src="{corners}" alt="" class="absolute -top-4 -right-4 w-15 h-15 pointer-events-none z-10 scale-x-[-1]" />
            <img src="{corners}" alt="" class="absolute -bottom-4 -left-4 w-15 h-15 pointer-events-none z-10 scale-y-[-1]" />
            <img src="{corners}" alt="" class="absolute -bottom-4 -right-4 w-15 h-15 pointer-events-none z-10 rotate-180" />

            <div class="parchment-background relative w-full h-full grid grid-rows-[auto_1fr] z-10 border border-[#3b2a1e] p-4">
                <div class="absolute inset-2 border border-[#8b7355]/40 pointer-events-none z-0"></div>
                <div class="grid grid-cols-1 md:grid-cols-2 gap-4 relative z-10 px-2 py-1 mb-4">
                    <div class="grid grid-rows-[auto-auto]">
                        <label for="title" class="forge-input-label">
                            Firstname*
                        </label>
                        <input type="text" name="name" class="forge-input" bind:value={formState.firstname} placeholder="Firstname" required  />
                    </div>
                    <div class="grid grid-rows-[auto-auto]">
                        <label for="title" class="forge-input-label">
                            Middlename
                        </label>
                        <input type="text" name="name" class="forge-input" bind:value={formState.middlename} placeholder="Middlename" />
                    </div>

                    <div class="grid grid-rows-[auto-auto]">
                        <label for="title" class="forge-input-label">
                            Surname*
                        </label>
                        <input type="text" name="name" class="forge-input" bind:value={formState.surname} placeholder="Surname" required />
                    </div>
                    <div class="grid grid-rows-[auto-auto]">
                        <label for="title" class="forge-input-label">
                            Nickname
                        </label>
                        <input type="text" name="name" class="forge-input" bind:value={formState.nickname} placeholder="Nickname" />
                    </div>

                    <div class="mb-3 mt-8 col-span-2 border-b border-[#8b7355]/40 pointer-events-none z-0"></div>

                    <div class="grid grid-rows-[auto-auto] col-span-2">
                        <label for="title" class="forge-input-label">
                            Portrait
                        </label>
                        <PortraitPicker characterId={id} bind:currentPortrait={formState.portrait} />
                    </div>

                    <div class="grid grid-rows-[auto-auto]">
                        <label for="title" class="forge-input-label">
                            Age
                        </label>
                        <input type="text" name="name" class="forge-input" bind:value={formState.age} placeholder="Age" />
                    </div>

                    <div class="relative">
                        <DataSelector items={statuses} bind:value={formState.status_id} placeholder="Alive, Dead, It's complicated, ..." labelText="Status" selectOnly={true} variant="forge" />
                    </div>

                    <div class="relative">
                        <DataSelector items={genders} bind:value={formState.gender_id} placeholder="Male, Female, Androgynous,..." labelText="Gender" selectOnly={true} variant="forge" />
                    </div>

                    <div class="relative">
                        <DataSelector items={races} bind:value={formState.race_id} placeholder="Human, Dwarf,..." labelText="Race" selectOnly={true} variant="forge" />
                    </div>


                    <div class="grid grid-rows-[auto-auto]">
                        <label for="title" class="forge-input-label">
                            Goals
                        </label>
                        <RichEditor bind:value={formState.goals} />
                    </div>

                    <div class="grid grid-rows-[auto-auto] relative">
                        <div class="absolute -left-1.5 top-3 bottom-3 w-px bg-[#8b7355]/40 pointer-events-none"></div>
                        <label for="title" class="forge-input-label">
                            Physical Description
                        </label>
                        <RichEditor bind:value={formState.description} />
                    </div>

                    <div class="mb-3 mt-8 col-span-2 border-b border-[#8b7355]/40 pointer-events-none z-0"></div>

                    <div class="relative">
                        <DataSelector items={occupations} bind:value={formState.occupation_id} placeholder="Guard, Warrior, Unknown ..." labelText="Occupation" selectOnly={true} variant="forge" />
                    </div>
                    <div class="relative">
                        <DataSelector items={alignments} bind:value={formState.alignment_id} placeholder="Loyal, Evil, Absolute bullshit, ..." labelText="Alignment" selectOnly={true} variant="forge" />
                    </div>

                    <div class="relative">
                        <DataSelector items={affiliations} bind:value={formState.affiliation_id} placeholder="Mercenary group, King's guard, ..." labelText="Affiliation" selectOnly={true} variant="forge" />
                    </div>
                    <div class="relative">
                        <DataSelector items={factions} bind:value={formState.faction_id} placeholder="A Kingdom, Guild, ..." labelText="Faction" selectOnly={true} disabled={isFactionLocked} variant="forge" />
                    </div>

                    <div class="relative">
                        <DataSelector items={locations} bind:value={formState.birthplace_location_id} placeholder="Where was he born?" labelText="Birthplace" selectOnly={true} variant="forge" />
                    </div>
                    <div class="relative">
                        <DataSelector items={locations} bind:value={formState.current_location_id} placeholder="And where is he now?" labelText="Current Location" selectOnly={true} variant="forge" />
                    </div>


                    <div class="mb-3 mt-8 col-span-2 border-b border-[#8b7355]/40 pointer-events-none z-0"></div>
                    <!-- Traits -->
                    <div class="relative">
                        <MultiDataSelector items={personalityTraits} bind:value={formState.personality_traits} placeholder="Curious, Introvert, ..." labelText="Personality traits" variant="forge" />
                    </div>
                    <div class="relative">
                        <MultiDataSelector items={strengths} bind:value={formState.strengths} placeholder="Courage, Bravery, ..." labelText="Strengths" variant="forge" />
                    </div>
                    <div class="relative">
                        <MultiDataSelector items={flaws} bind:value={formState.flaws} placeholder="Prideful" labelText="Flaws" variant="forge" />
                    </div>
                    <div class="relative">
                        <MultiDataSelector items={weaknesses} bind:value={formState.weaknesses} placeholder="Sick, Grass, ..." labelText="Weaknesses" variant="forge" />
                    </div>
                    <div class="relative">
                        <MultiDataSelector items={fears} bind:value={formState.fears} placeholder="Death, Spiders, ..." labelText="Fears" variant="forge" />
                    </div>
                </div>


                    <div class="border-b border-[#8b7355]/40 pointer-events-none z-0"></div>
                    <div class="mt-2 flex items-center justify-between">
                        <StatusQuip isReady={isFormReady} missingType={missingFieldType} />
                        <button type="submit" class="forge-btn forge-btn-base flex align-middle gap-2">
                            <span><AnvilImpact height="1.6rem" /></span>
                            <span>Inscribe</span>
                        </button>
                    </div>
            </div>
        </div>
    </form>
</section>
