<script lang="ts">
    import { FetchAllLore, DeleteLore, DeleteLoreCategory, OpenEditorWindow } from "$wails/world-builder/app";
    import BottomButton from "$lib/components/ui/BottomButton.svelte";
    import { onMount } from "svelte";
    import { Lore, LoreCategory } from "$wails/world-builder/internal/models/models";
    import RichEditor from "$lib/components/ui/RichEditor.svelte";
    import { subscribe } from "$lib/functions/subscribe";

    let isLoading = $state(true)
    let lores = $state<Lore[]>()

    function handleOpen() {
      OpenEditorWindow('lore/form').catch(err => {
        console.error("Failed to open window", err)
      })
    }

    function handleEdit(id: number) {
      OpenEditorWindow(`lore/form/${id}`).catch(err => {
        console.error("Failed to open window", err)
      })
    }

    async function handleDelete(id: number) {
      try {
        DeleteLore(id)
      } catch(err) {
        console.error(err)
      } finally {
        isLoading = false
      }
    }

    async function fetchData() {
      try {
        lores = await FetchAllLore()
      } catch(err) {
        console.error("Failed to load lores.", err)
      } finally {
        isLoading = false
      }
    }

    onMount(() => {
      fetchData()

      return subscribe('lores', fetchData)
    })
</script>

<section id="lore-overview65" class="w-full min-h-0 h-screen grid grid-cols-12 content-start p-1 gap-1 overflow-y-auto!">
    {#each lores as lore (lore.id)}
        <RichEditor disabled bind:value={lore.content} classes="col-span-6 min-h-36 max-h-52 p-0.5">
            {#snippet children(prose)}
                <div class="grid grid-cols-6 h-full" style="">
                    <p class="col-span-3">{@html lore.title}</p>
                    <p class="col-span-3 text-gray-500">{lore.lore_category?.label}</p>
                    <hr class="col-span-6">
                    <div class="col-span-6 flex min-w-0 text-left h-full self-start truncate! overflow-hidden!">
                        {@render prose()}
                    </div>
                    <div class="col-span-6 justify-items-center mb-1">
                        <button class="fantasy-btn-md fantasy-bone-n-coper" onclick={() => handleEdit(lore.id)}>Edit</button>
                        <button class="fantasy-btn-md fantasy-bone-n-coper" onclick={() => handleDelete(lore.id)}>Delete</button>
                    </div>
                </div>
            {/snippet}
        </RichEditor>
    {/each}

    <BottomButton onclick={handleOpen}>Add</BottomButton>
</section>
