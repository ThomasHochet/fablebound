<script lang="ts">
    import { DeleteLocation, FetchAllLocations, OpenEditorWindow } from "$wails/world-builder/app";
    import BottomButton from "$lib/components/ui/BottomButton.svelte";
    import { onDestroy, onMount } from "svelte";
    import { Location } from "$wails/world-builder/internal/models/models";
    import RichEditor from "$lib/components/ui/RichEditor.svelte";
    import { Events } from "@wailsio/runtime";
    import { subscribe } from "$lib/functions/subscribe";

    let isLoading = $state(true)
    let locations = $state<Location[]>()

    function handleOpen() {
      OpenEditorWindow('location/form').catch(err => {
        console.error("Failed to open window", err)
      })
    }

    function handleEdit(id: number) {
      OpenEditorWindow(`location/form/${id}`).catch(err => {
        console.error("Failed to open window", err)
      })
    }

    async function handleDelete(id: number) {
      try {
        DeleteLocation(id)
      } catch(err) {
        console.error(err)
      } finally {
        isLoading = false
      }
    }

    async function fetchData() {
      try {
        locations = await FetchAllLocations()
      } catch(err) {
        console.error("Failed to load locations.", err)
      } finally {
        isLoading = false
      }
    }

    onMount(() => {
      fetchData()

      return subscribe('locations', fetchData)
    })
</script>

<section id="location-overview" class="w-full min-h-0 h-screen grid grid-cols-12 content-start p-1 gap-1 overflow-y-auto!">
    {#each locations as location (location.id)}
        <RichEditor disabled bind:value={location.description} classes="col-span-6 min-h-36 max-h-52 p-0.5">
            {#snippet children(prose)}
                <div class="grid grid-cols-6 h-full" style="">
                    <p class="col-span-3">{location.name}</p>
                    {#if (location?.subtype)}
                        <p class="col-span-3 text-gray-500">{location.type + " " + location.subtype}</p>
                    {:else}
                        <p class="col-span-3 text-gray-500">{location.type}</p>
                    {/if}
                    <hr class="col-span-6">
                    <div class="col-span-6 flex min-w-0 text-left h-full self-start truncate! overflow-hidden!">
                        {@render prose()}
                    </div>
                    <div class="col-span-6 justify-items-center mb-1">
                        <button class="fantasy-btn-md fantasy-bone-n-coper" onclick={() => handleEdit(location.id)}>Edit</button>
                        <button class="fantasy-btn-md fantasy-bone-n-coper" onclick={() => handleDelete(location.id)}>Delete</button>
                    </div>
                </div>
            {/snippet}
        </RichEditor>
    {/each}

    <BottomButton onclick={handleOpen}>Add</BottomButton>
</section>
