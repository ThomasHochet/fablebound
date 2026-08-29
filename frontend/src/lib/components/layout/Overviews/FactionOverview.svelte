<script lang="ts">
    import CharacterForm from '$lib/components/features/forms/CharacterForm.svelte';
    import BottomButton from '$lib/components/ui/BottomButton.svelte';
    import RichEditor from '$lib/components/ui/RichEditor.svelte';
    import { DeleteAffiliation, FetchAllAffiliations, FetchAllFactions, OpenEditorWindow } from '$wails/world-builder/app';
    import { Faction, Affiliation } from '$wails/world-builder/internal/models/models'
    import { onMount } from 'svelte';

    let isLoading = $state(true)
    let isOpenEditor = $state(false)
    let factions = $state<Faction[]>()
    let affiliations = $state<Affiliation[]>()

    async function fetchData() {
      try {
        factions = await FetchAllFactions()
        affiliations = await FetchAllAffiliations()
      } catch(err) {
        console.error("Failed to load data", err)
      } finally {
        isLoading = false
      }
    }

    function handleOpen() {
      OpenEditorWindow('faction/form').catch(err => {
        console.error("Failed to open window", err)
      })
    }

    function handleEdit(id: number) {
      OpenEditorWindow(`faction/form/${id}`).catch(err => {
        console.error("Failed to open window,", err)
      })
    }

    async function handleDelete(id: number) {
      try {
        DeleteAffiliation(id)
      } catch(err) {
        console.error("Failed to delete affiliation", err)
      } finally {
        isLoading = true
      }

      fetchData()
    }

    onMount(() =>
      fetchData()
    )
</script>

<section id="overview-factions" class="w-full grid grid-cols-12 p-1 gap-1">
    {#each affiliations as affiliation (affiliation.id)}
        <RichEditor disabled bind:value={affiliation.description} classes="col-span-6 min-h-52 p-0.5">
            {#snippet children(prose)}
                <div class="grid grid-cols-6 h-full" style="grid-template-rows: auto auto 1fr auto;">
                    <div class="col-span-6 grid grid-cols-6 text-left">
                        <p class="col-span-4 font-bold underline">{affiliation.name}</p>
                        <p class="col-span-2">{affiliation?.faction?.name ?? ''}</p>
                    </div>
                    <div class="col-span-6">
                        <hr/>
                        {#if affiliation?.type}
                            <p class="italic text-gray-500 text-left ml-8">{affiliation.type}</p>
                        {/if}
                    </div>
                    <div class="col-span-6 text-left self-start">
                        {@render prose()}
                    </div>
                    <div class="col-span-6 justify-items-center mb-1">
                        <button class="fantasy-btn-md fantasy-bone-n-coper" onclick={() => handleEdit(affiliation.id)}>Edit</button>
                        <button class="fantasy-btn-md fantasy-bone-n-coper" onclick={() => handleDelete(affiliation.id)}>Delete</button>
                    </div>
                </div>

            {/snippet}
        </RichEditor>
    {/each}
    <BottomButton onclick={handleOpen}>Add</BottomButton>
</section>
