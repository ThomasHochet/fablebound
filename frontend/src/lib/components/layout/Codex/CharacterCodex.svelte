<script lang="ts">
    import Modal from "$lib/components/features/Modal.svelte";
    import Card from "$lib/components/ui/Card.svelte";
    import { subscribe } from "$lib/functions/subscribe";
    import { DeleteCharacter, FetchAllCharacters, OpenEditorWindow } from "$wails/world-builder/app";
    import { Character } from "$wails/world-builder/internal/models/models";
    import { onMount } from "svelte";
    import corners from "$lib/../assets/images/corners.png"
    import QuillInk from "@iconify-svelte/game-icons/components/q/quill-ink.svelte";
    import CrossMark from "@iconify-svelte/game-icons/components/c/cross-mark.svelte";

    let characters = $state<Character[]>([])
    let isLoading = $state(true)
    let isModalOpen = $state(false)
    let charRdyToDel = $state<number>()
    let readyToDeleteData = $derived(characters.find(c => (c.id === charRdyToDel)))

    function handleEdit(id: number) {
      OpenEditorWindow(`character/form/${id}`).catch(err => {
        console.error("Failed to open window", err)
      })
    }

    async function handleDelete(id: number | undefined) {
      const numId = Number(id)
      if (!isNaN(numId)) {
        DeleteCharacter(numId)
        isModalOpen = false
      }
    }

    async function fetchData() {
      try {
        characters = await FetchAllCharacters(100)
      } catch (err) {
        console.error(err)
      } finally {
        isLoading = false
      }
    }

    onMount(() => {
      fetchData()
      subscribe('characters', fetchData)
    })
</script>

{#each characters as character(character.id)}
    <Card
        title={character.firstname + " " + character.surname}
        category={character.occupation?.label}
        description={character.description}
        variant="character"
        onEdit={() => handleEdit(character.id)}
        onDelete={() => {
          isModalOpen = true
          charRdyToDel = character.id
        }}
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
                [{readyToDeleteData?.firstname + " " + readyToDeleteData?.surname}]
            </span>
            <!-- <span class="italic">Category: Character</span> -->
            <span>Codex : Characters</span>
            <div class="flex flex-row justify-center px-6 mt-auto gap-4">
                <button class="forge-btn forge-btn-base" onclick={() => {
                  handleDelete(charRdyToDel)
                }}>
                    <CrossMark height="1.5rem" />
                    <span class="ml-2">Agreed</span>
                </button>
                <button
                    class="forge-btn forge-btn-base"
                    onclick={() => {
                      isModalOpen = false
                      charRdyToDel = undefined
                    }}
                >
                    <QuillInk height="1.5rem" />
                    <span class="ml-1">Keep</span>
                </button>
            </div>
        </div>

    </div>
</Modal>
