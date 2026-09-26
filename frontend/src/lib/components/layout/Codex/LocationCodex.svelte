<script lang="ts">
    import { DeleteLocation, FetchAllLocations, OpenEditorWindow, OpenReaderWindow } from "$wails/world-builder/app";
    import BottomButton from "$lib/components/ui/BottomButton.svelte";
    import { onDestroy, onMount } from "svelte";
    import { Location } from "$wails/world-builder/internal/models/models";
    import RichEditor from "$lib/components/ui/RichEditor.svelte";
    import { Events } from "@wailsio/runtime";
    import { subscribe } from "$lib/functions/subscribe";
    import Card from "$lib/components/ui/Card.svelte";
    import corners from "$lib/../assets/images/corners.png"
    import QuillInk from "@iconify-svelte/game-icons/components/q/quill-ink.svelte";
    import CrossMark from "@iconify-svelte/game-icons/components/c/cross-mark.svelte";
    import Modal from "$lib/components/features/Modal.svelte";
    import { logError } from "$lib/logger";

    let isLoading = $state(true)
    let locations = $state<Location[]>([])
    let isModalOpen = $state(false)
    let locRdyToDel = $state<number>()
    let readyToDeleteData = $derived(locations.find(l => (l.id === locRdyToDel)))


    function handleEdit(id: number) {
      OpenEditorWindow('location', id).catch(err => {
        logError("Failed to open window", err)
      })
    }

    async function handleDelete(id: number | undefined) {
      const numId = Number(id)
      if (!isNaN(numId)) {
        DeleteLocation(numId)
        isModalOpen = false
      }
    }

    function openLocationReader(id: number, title: string) {
      OpenReaderWindow('location', id, title).catch(err => {
        logError(`Failed to open 'location' ${id}`, err)
      })
    }

    async function fetchData() {
      try {
        locations = await FetchAllLocations()
      } catch(err) {
        logError("Failed to load locations.", err)
      } finally {
        isLoading = false
      }
    }

    onMount(() => {
      fetchData()

      return subscribe('locations', fetchData)
    })
</script>

{#each locations as location(location.id)}
    <Card
        title={location.name}
        category={location.type}
        description={location.description}
        variant={location.type}
        onEdit={() => handleEdit(location.id)}
        onDelete={() => {
          (isModalOpen = true)
          locRdyToDel = location.id
        }}
        dblClickEvent={() => openLocationReader(location.id, location.name)}
    />
{/each}

<Modal bind:open={isModalOpen} >
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
            <span class="italic">Category: {readyToDeleteData?.type}</span>
            <span>Codex : Locations</span>
            <div class="flex flex-row justify-center px-6 mt-auto gap-4">
                <button class="forge-btn forge-btn-base" onclick={() => {
                  handleDelete(locRdyToDel)
                }}>
                    <CrossMark height="1.5rem" />
                    <span class="ml-2">Agreed</span>
                </button>
                <button
                    class="forge-btn forge-btn-base"
                    onclick={() => {
                      isModalOpen = false
                      locRdyToDel = undefined
                    }}
                >
                    <QuillInk height="1.5rem" />
                    <span class="ml-1">Keep</span>
                </button>
            </div>
        </div>

    </div>
</Modal>
