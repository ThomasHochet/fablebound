<script lang="ts">
  import { onMount } from "svelte";
  import RichEditor from "$lib/components/ui/RichEditor.svelte";
  import DataSelector from "$lib/components/ui/DataSelector.svelte";
  import StatusQuip from "$lib/components/ui/StatusQuip.svelte";

  let { id } = $props()
  let formState = $state({
      firstname: '',
      middlename: '',
      surname: '',
      nickname: '',

      age: 'Unknown',
      description: '',
      portrait: '',
      goals: '',

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

  async function handleSubmit(e: Event) {
    e.preventDefault()
  }

  async function fetchData() {
    try{

    } catch(err) {
      console.error(err)
    }
  }

  onMount(() => {
    fetchData()
  })
</script>

<section class="grid grid-cols-1 items-center justify-center p-4">
    <form onsubmit={handleSubmit}>
        <div class="w-full rounded-lg overflow-hidden shadow-2xl border-4 border-[#2b190c]">
            <div class="p-5 rounded-t-md relative">
                <div class="wood-texture-bg"></div>

                <div class="iron-nail nail-tl"></div>
                <div class="iron-nail nail-tr"></div>
                <div class="iron-nail nail-bl"></div>
                <div class="iron-nail nail-br"></div>

                <div class="grid grid-cols-1 md:grid-cols-2 gap-4 relative z-10 my-2">
                    <div>
                        <label for="title" class="block text-sm text-left ml-2 text-(--gold-text) mb-1 tracking-wider uppercase">
                            Location Title*
                        </label>
                        <div class="fantasy-input-wrapper flex items-center">
                            <input type="text" name="name" class="fantasy-input" bind:value={formState.name} placeholder="Location name" required />
                        </div>
                    </div>
                    <div class="relative">
                        <DataSelector items={typeOptions} bind:value={formState.type} placeholder="Continent, Kingdom, City" labelText="Location Type*" required={true} />
                    </div>
                    <div class="relative">
                        <DataSelector items={subtypeOptions} bind:value={formState.subtype} placeholder="Castle, Inn, Old farm..." labelText="Location Subtype" />
                    </div>
                    <div class="relative">
                        <DataSelector items={locationOptions} bind:value={formState.parentId} placeholder="Faerûn, Baldur's gate..." labelText="Parent location" />
                    </div>
                </div>
            </div>
            <div class="bg-(--wood-dark) p-2">
                <RichEditor bind:value={formState.description} />
                <div class="mt-2 flex items-center justify-between">
                    <StatusQuip isReady={isFormReady} missingType={missingFieldType} />
                    <button type="submit" class="fantasy-btn-xl fantasy-bone-n-coper ">
                        Inscribe
                    </button>
                </div>
            </div>


        </div>
    </form>
</section>
