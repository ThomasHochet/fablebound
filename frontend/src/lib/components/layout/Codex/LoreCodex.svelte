<script lang="ts">
    import { FetchAllLore, DeleteLore, DeleteLoreCategory, OpenEditorWindow, OpenReaderWindow, FetchGeneralLore } from "$wails/world-builder/app";
    import BottomButton from "$lib/components/ui/BottomButton.svelte";
    import { onMount } from "svelte";
    import { Lore, LoreCategory } from "$wails/world-builder/internal/models/models";
    import RichEditor from "$lib/components/ui/RichEditor.svelte";
    import { subscribe } from "$lib/functions/subscribe";
    import Card from "$lib/components/ui/Card.svelte";
    import CrossMark from "@iconify-svelte/game-icons/components/c/cross-mark.svelte";
    import corners from "$lib/../assets/images/corners.png"
    import QuillInk from "@iconify-svelte/game-icons/components/q/quill-ink.svelte";
    import Modal from "$lib/components/features/Modal.svelte";
    import { logError } from "$lib/logger";

    let isLoading = $state(true)
    let lores = $state<Lore[]>([])
    let isModalOpen = $state(false)
    let loreRdyToDel = $state<number>()
    let readyToDeleteData = $derived(lores.find(l => (l.id === loreRdyToDel)))

    function handleOpen(id: number, title: string) {
      OpenReaderWindow('lore', id, title).catch(err => {
        logError(`Failed to open lore reader: ${id}:${title}`, err)
      })
    }

    function handleEdit(id: number) {
      OpenEditorWindow('lore', id).catch(err => {
        logError("Failed to open window", err)
      })
    }

    async function handleDelete(id: number | undefined) {
      const numId = Number(id)
      if (!isNaN(numId)) {
        DeleteLore(numId)
        isModalOpen = false
      }
    }

    async function fetchData() {
      try {
        lores = await FetchGeneralLore()
      } catch(err) {
        logError("Failed to load lores.", err)
      } finally {
        isLoading = false
      }
    }

    onMount(() => {
      fetchData()

      return subscribe('lores', fetchData)
    })
</script>

{#each lores as lore(lore.id)}
    <Card
        title={lore.title}
        category={lore.lore_category?.label}
        description={lore.content}
        variant={lore.lore_category?.label.toLowerCase()}
        onEdit={() => handleEdit(lore.id)}
        onDelete={() => {
          loreRdyToDel = lore?.id
          isModalOpen = true
        }}
        dblClickEvent={() => handleOpen(lore.id, lore.title)}
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
                [{readyToDeleteData?.title}]
            </span>
            <span class="italic">Category: {readyToDeleteData?.lore_category?.label}</span>
            <span>Codex : Lore</span>
            <div class="flex flex-row justify-center px-6 mt-auto gap-4">
                <button class="forge-btn forge-btn-base" onclick={() => {
                  handleDelete(loreRdyToDel)
                }}>
                    <CrossMark height="1.5rem" />
                    <span class="ml-2">Agreed</span>
                </button>
                <button
                    class="forge-btn forge-btn-base"
                    onclick={() => {
                      isModalOpen = false
                      loreRdyToDel = undefined
                    }}
                >
                    <QuillInk height="1.5rem" />
                    <span class="ml-1">Keep</span>
                </button>
            </div>
        </div>

    </div>
</Modal>
