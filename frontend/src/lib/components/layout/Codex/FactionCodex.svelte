<script lang="ts">
    import CharacterForm from '$lib/components/features/forms/CharacterForm.svelte';
    import BottomButton from '$lib/components/ui/BottomButton.svelte';
    import Card from '$lib/components/ui/Card.svelte';
    import RichEditor from '$lib/components/ui/RichEditor.svelte';
    import { subscribe } from '$lib/functions/subscribe';
    import { DeleteAffiliation, DeleteFaction, FetchAllAffiliations, FetchAllFactions, OpenEditorWindow } from '$wails/world-builder/app';
    import { Faction, Affiliation } from '$wails/world-builder/internal/models/models'
    import { onMount } from 'svelte';
    import corners from "$lib/../assets/images/corners.png"
    import QuillInk from "@iconify-svelte/game-icons/components/q/quill-ink.svelte";
    import CrossMark from "@iconify-svelte/game-icons/components/c/cross-mark.svelte";
    import Modal from "$lib/components/features/Modal.svelte";

    let isLoading = $state(true)
    let factions = $state<Faction[]>([])
    let affiliations = $state<Affiliation[]>([])
    let isModalOpen = $state(false)

    let combinedList = $derived.by(() => {
      const tFactions = factions.map(f => ({ ...f, _kind: 'faction' as const}))
      const tAffiliations = affiliations.map(a => ({ ...a, _kind: 'affiliation' as const}))

      return [...tFactions, ...tAffiliations].sort(() => Math.random() - 0.5)
    })

    let facRdyToDel = $state<any>()
    let readyToDeleteData = $derived(combinedList.find(c => (c.id === facRdyToDel.id && c._kind == facRdyToDel._kind)))

    async function fetchData() {
      try {
        const [fData, aData] = await Promise.all([
          FetchAllFactions(),
          FetchAllAffiliations()
        ])
        factions = fData ?? []
        affiliations = aData ?? []
      } catch(err) {
        console.error("Failed to load data", err)
      } finally {
        isLoading = false
      }
    }



    function handleEdit(kind: 'faction' | 'affiliation', id: number) {
      OpenEditorWindow(`faction/form/${id}`).catch(err => {
        console.error("Failed to open window,", err)
      })
    }

    async function handleDelete(kind: 'faction' | 'affiliation' | undefined, id: number) {
      console.log(kind, id)
      try {
        if (kind === 'affiliation')
          DeleteAffiliation(id)
        else
          DeleteFaction(id)
        isModalOpen = false
      } catch(err) {
        console.error("Failed to delete affiliation", err)
      } finally {
        isLoading = true
      }

      fetchData()
    }

    onMount(() => {
      fetchData()

      return subscribe('affiliations', fetchData)
    })
</script>

{#if isLoading}

{:else}
    {#each combinedList as item(`${item._kind}-${item.id}`)}
        {#if item._kind === 'affiliation'}
            {@render affiliationCard(item)}
        {:else}
            {@render factionCard(item)}
        {/if}
    {/each}
{/if}
{#snippet affiliationCard(item: Affiliation)}
    <Card
        title={item.name}
        category={item.type || 'Affiliation'}
        description={item.description}
        variant="affiliation"
        onEdit={() => handleEdit('affiliation', item.id)}
        onDelete={() => {
          isModalOpen = true
          facRdyToDel = item
        }}
    />
{/snippet}

{#snippet factionCard(item: Faction)}
	<Card
	    title={item.name}
		category="faction"
		description={item.description}
		variant="faction"
		onEdit={() => handleEdit('faction', item.id)}
		onDelete={() => {
          isModalOpen = true
          facRdyToDel = item
        }}
	/>
{/snippet}

<Modal bind:open={isModalOpen}>
    {console.log(readyToDeleteData)}
    <div class="border-[#1a0f0f] bg-[#2b190c] p-2 md:p-2 w-[25rem]! h-[22rem]! rounded-md!">
        <div class="wood-texture-bg absolute inset-0 pointer-events-none z-0"></div>

        <img src="{corners}" alt="" class="absolute -top-3 -left-3 w-15 h-15 pointer-events-none z-10">
        <img src="{corners}" alt="" class="absolute -top-3 -right-3 w-15 h-15 pointer-events-none z-10 scale-x-[-1]" />
        <img src="{corners}" alt="" class="absolute -bottom-3 -left-3 w-15 h-15 pointer-events-none z-10 scale-y-[-1]" />
        <img src="{corners}" alt="" class="absolute -bottom-3 -right-3 w-15 h-15 pointer-events-none z-10 rotate-180" />

        <div class="flex flex-col gap-2 parchment-background-inner relative w-full h-full z-10 border border-[#3b2a1e] p-4 rounded-md">
            <span class="mt-4 font-bold">Shall this tale be lost to oblivion?</span>
            <span class="underline mt-2">
                [{readyToDeleteData?.name}]
            </span>
            <span class="italic">Category: {readyToDeleteData?._kind}</span>
            <span>Codex : Factions</span>
            <div class="flex flex-row justify-center px-6 mt-auto gap-4">
                <button class="forge-btn forge-btn-base" onclick={() => {
                  handleDelete(readyToDeleteData?._kind, readyToDeleteData?.id)
                }}>
                    <CrossMark height="1.5rem" />
                    <span class="ml-2">Agreed</span>
                </button>
                <button
                    class="forge-btn forge-btn-base"
                    onclick={() => {
                      isModalOpen = false
                      facRdyToDel = undefined
                    }}
                >
                    <QuillInk height="1.5rem" />
                    <span class="ml-1">Keep</span>
                </button>
            </div>
        </div>

    </div>
</Modal>
