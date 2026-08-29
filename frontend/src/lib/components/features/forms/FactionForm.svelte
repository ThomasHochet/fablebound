<script lang="ts">
    import FactionModal from "./SubForms/FactionModal.svelte";
    import { DeleteFaction, FetchAffiliation, FetchAllFactions, FetchFaction, SaveAffiliation, SaveFaction } from "$wails/world-builder/app";
    import { Faction, Affiliation } from "$wails/world-builder/internal/models/models";
    import { onMount } from "svelte";
    import RichEditor from "$lib/components/ui/RichEditor.svelte";

    let { id } = $props();
    let affiliation = $state<Affiliation | null>()

    // on mount value
    let factions = $state<Faction[]>([])


    // booleans
    let isModalOpen = $state(false)
    let isLoading = $state(true)

    // misc
    let affiliationName = $state('')
    let affiliationDescription = $state('Describe your affiliation...')
    let affiliationType = $state('')
    let selectedFactionName = $state('')
    let selectedFactionId = $derived(
      factions.find((faction) => faction.name === selectedFactionName)?.id ?? null
    )
    let selectedFaction = $state<Faction | null>()

    async function handleSubmit(e: SubmitEvent) {
        e.preventDefault();
        try {
          const payload = new Affiliation({
            name: affiliationName,
            description: affiliationDescription,
            type: affiliationType,
            faction_id: selectedFactionId
          })
          // console.log(payload)
          const createdAffiliation = await SaveAffiliation(payload)
        } catch(err) {
          console.error(err)
        }
    }

    async function loadFactions() {
      try {
          factions = await FetchAllFactions()
      } catch(err) {
          console.error("Failed to load factions.", err)
      } finally {
          isLoading = false;
      }
    }

    async function loadAffiliation(id: number) {
      try {
        affiliation = await FetchAffiliation(id)
      } catch(err) {
        console.error("Failed to load Faction.", err)
      } finally {
        isLoading = false;
      }

      affiliationName = affiliation?.name
      affiliationDescription = affiliation?.description
      affiliationType = affiliation?.type
      selectedFactionId = affiliation?.faction?.id
      selectedFactionName = affiliation?.faction?.name
    }

    async function handleFactionSave(data: { title: string; description: string }) {
        try {
          const payload = new Faction({
              name: data.title,
              description: data.description
          })

          const createdFaction = await SaveFaction(payload)
          await loadFactions()

          if (createdFaction?.name) {
            selectedFactionName = createdFaction.name
          }
        } catch(err) {
          console.error(err)
        }
    }

    async function handleFactionEdit() {
      try {
        if (selectedFactionId != null){
          selectedFaction = await FetchFaction(selectedFactionId)
          isModalOpen = true
        }
      } catch(err) {
        console.error("Failed to fetch faction.", err)
      }
    }

    function handleFactionDelete() {
      // Conditions based on affiliations and maybe a fetch to lookup if selected faction is linked to anything.
      if (selectedFactionId != null) {
        DeleteFaction(selectedFactionId)
      }
      loadFactions()
      selectedFactionName = ''
    }

    onMount(() => {
      loadFactions()

      const numericId = Number(id)
      if (!isNaN(numericId)) {
        loadAffiliation(numericId)
      } else {
        console.error("Invalid provided ID", id)
      }
    })
</script>

<article>
    <div>
        <button onclick={() => isModalOpen = true}>Add Faction</button>
    </div>

    <form onsubmit={handleSubmit}>
        <input type="hidden" name="faction-id" value={selectedFactionId} />
        <input type="text" name="title" placeholder="Affiliation" bind:value={affiliationName} />
        <div class="flex">
            <input type="text" name="factions-choice" id="factions-choice" list="factions" placeholder="Faction ?" bind:value={selectedFactionName}>
            <datalist id="factions">
                {#each factions as faction}
                    <option value={faction.name}></option>
                {/each}
            </datalist>
            <button class="fantasy-btn-md fantasy-bone-n-coper" onclick={handleFactionEdit}>
                Edit
            </button>
            <button class="fantasy-btn-md fantasy-bone-n-coper" onclick={handleFactionDelete}>
                Delete
            </button>
        </div>
        <input type="text" name="type" id="affiliation-type" placeholder="Type" bind:value={affiliationType} />
        <RichEditor bind:value={affiliationDescription} />
        <div class="flex">
            <button type="submit" class="fantasy-btn fantasy-bone-n-coper">
                Save
            </button>
        </div>
    </form>

    <FactionModal
        bind:open={isModalOpen}
        onSave={handleFactionSave}
        factionData={selectedFaction ?? null}
    />
</article>
