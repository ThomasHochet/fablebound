<script lang="ts">
    import FactionModal from "./SubForms/FactionModal.svelte";
    import { DeleteFaction, FetchAffiliation, FetchAllFactions, FetchFaction, FetchFactionAffiliations, SaveAffiliation, SaveFaction } from "$wails/world-builder/app";
    import { Faction, Affiliation } from "$wails/world-builder/internal/models/models";
    import { onMount } from "svelte";
    import RichEditor from "$lib/components/ui/RichEditor.svelte";
    import corners from "$lib/../assets/images/corners.png"
    import ParchmentTitle from "$lib/components/ui/ParchmentTitle.svelte";
    import DataSelector from "$lib/components/ui/DataSelector.svelte";
    import StatusQuip from "$lib/components/ui/StatusQuip.svelte";
    import AnvilImpact from "@iconify-svelte/game-icons/components/a/anvil-impact.svelte";
    import QuillInk from "@iconify-svelte/game-icons/components/q/quill-ink.svelte";
    import CrossMark from "@iconify-svelte/game-icons/components/c/cross-mark.svelte";
    import Modal from "../Modal.svelte";
    import { subscribe } from "$lib/functions/subscribe";

    let { id } = $props();
    let isModalOpen = $state(false)
    let modalType = $state('edit')
    let isLoading = $state(true)
    let affiliation = $state({
      name: '',
      description: 'Describe your affiliation...',
      type: '',
      factionId: null as number | null
    })

    // on mount value
    let factions = $state<Faction[]>([])
    let affiliations = $state<Affiliation[]>([])

    let faction = $state({
      id: null as number | null,
      name: '',
      description: 'What is your purpose?'
    })

    $effect(() => {
      const id = affiliation.factionId

      if (!id) {
        faction.id = null
        faction.name = '';
        faction.description = 'What is your purpose?';
        return;
      }

      FetchFaction(id).then(fFaction => {
        faction.id = fFaction?.id ?? null
        faction.name = fFaction?.name ?? ''
        faction.description = fFaction?.description ?? ''
      })

      FetchFactionAffiliations(id).then(fAffiliations => {
        affiliations = fAffiliations
      })
    })

    async function handleSubmit(e: SubmitEvent) {
        e.preventDefault();
        try {
          const payload = new Affiliation({
            id: id ? Number(id) : undefined,
            name: affiliation.name,
            description: affiliation.description,
            type: affiliation.type,
            faction_id: affiliation.factionId
          })
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
        const affiliationData = await FetchAffiliation(id)

        affiliation.name = affiliationData?.name ?? ''
        affiliation.description = affiliationData?.description ?? ''
        affiliation.type = affiliationData?.type ?? ''
        affiliation.factionId = affiliationData?.faction?.id ?? null
      } catch(err) {
        console.error("Failed to load Faction.", err)
      } finally {
        isLoading = false;
      }
    }


    async function handleFactionSave() {
        try {
          const payload = new Faction({
            id: faction.id ? Number(faction.id) : undefined,
            name: faction.name,
            description: faction.description
          })

          const createdFaction = await SaveFaction(payload)
          await loadFactions()

          isModalOpen = false

          if (createdFaction?.name) {
            affiliation.factionId = createdFaction.id
          }
        } catch(err) {
          console.error(err)
        }
    }

    function handleFactionDelete() {
      // Conditions based on affiliations and maybe a fetch to lookup if selected faction is linked to anything.
      // Nope, handled in front and control.
      if (affiliation.factionId !== null) {
        DeleteFaction(affiliation.factionId)
      }
      affiliation.factionId = null
      faction.id = null
      faction.name = ''
      faction.description = ''

      isModalOpen = false
      loadFactions()
    }

    let isFormReady = $derived(Boolean(affiliation.name.trim()))
    let missingFieldType = $derived.by(() => {
      return null
    })

    onMount(() => {
      loadFactions()

      const numericId = Number(id)
      if (!isNaN(numericId)) {
        loadAffiliation(numericId)
      } else {
        console.error("Invalid provided ID", id)
      }

      return subscribe('factions', loadFactions)
    })
</script>

<section class="h-full grid grid-cols-1 items-center justify-center px-2 ">
    <div class="relative z-20 -mb-10 pointer-events-none">
        <ParchmentTitle title="Factions" subtitle="Divisions. Unions. A world built through connections." />
    </div>
    <form onsubmit={handleSubmit}>
        <div class="relative w-full h-full shadow-2xl rounded-sm border-4 border-[#1a0f0f] bg-[#2b190c] p-1 md:p-1 grid grid-cols-1 grid-rows-1">
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
                                Affiliation Name*
                            </label>
                            <input type="text" name="title" class="forge-input" required bind:value={affiliation.name} placeholder="The world is at your attention."  />
                        </div>
                        <div class="flex items-end gap-2">
                            <DataSelector items={factions} bind:value={affiliation.factionId} placeholder="" labelText="Faction" variant="forge" selectOnly />
                            <button
                                type="button"
                                class="forge-btn forge-btn-xs h-9"
                                onclick={() => {
                                  modalType = 'edit'
                                  isModalOpen = true
                                }}
                            ><QuillInk height="1rem" /></button>
                            <button
                                type="button"
                                class="forge-btn forge-btn-xs h-9"
                                onclick={() => {
                                  modalType = 'delete'
                                  if (faction.id !== null) {
                                    isModalOpen = true
                                  }
                                }}
                            ><CrossMark height="1rem" /></button>
                        </div>
                        <div class="grid grid-rows-[auto-auto]">
                            <label for="title" class="forge-input-label">
                                Affiliation Type
                            </label>
                            <input type="text" name="type" class="forge-input" required bind:value={affiliation.type} placeholder="Mercenary group, guards, travelers..."  />
                        </div>

                    </div>
                    <RichEditor bind:value={affiliation.description} />
                    <div class="border-b border-[#8b7355]/40 pointer-events-none z-0"></div>
                    <div class="mt-2 flex items-center justify-between">
                        <StatusQuip isReady={isFormReady} missingType={missingFieldType} />
                        <button type="submit" class="forge-btn forge-btn-base flex align-middle gap-2">
                            <span class="mb-1"><AnvilImpact height="1.6rem" /></span>
                            <span>Inscribe</span>
                        </button>
                    </div>
                </div>
        </div>
    </form>
</section>

<Modal bind:open={isModalOpen}>
    {#if modalType === 'edit'}
        <div class="border-[#1a0f0f] bg-[#2b190c] p-2 md:p-2 w-[70rem]! h-[40rem]! rounded-md!">
            <div class="wood-texture-bg absolute inset-0 pointer-events-none z-0"></div>

            <img src="{corners}" alt="" class="absolute -top-3 -left-3 w-15 h-15 pointer-events-none z-10">
            <img src="{corners}" alt="" class="absolute -top-3 -right-3 w-15 h-15 pointer-events-none z-10 scale-x-[-1]" />
            <img src="{corners}" alt="" class="absolute -bottom-3 -left-3 w-15 h-15 pointer-events-none z-10 scale-y-[-1]" />
            <img src="{corners}" alt="" class="absolute -bottom-3 -right-3 w-15 h-15 pointer-events-none z-10 rotate-180" />

            <div class="parchment-background relative w-full h-full grid grid-rows-[auto_1fr] z-10 border border-[#3b2a1e] p-4">
                <div class="absolute inset-2 border border-[#8b7355]/40 pointer-events-none z-0"></div>
                <div class="grid grid-cols-1 md:grid-cols-1 gap-4 relative z-10 px-2 py-1 mb-4">
                    <div class="grid grid-rows-[auto-auto]">
                        <label for="title" class="forge-input-label">
                            Faction Title*
                        </label>
                        <input type="text" name="title" class="forge-input" required bind:value={faction.name} placeholder="Seat of power, how thou shalt be named?"  />
                    </div>
                </div>
                <RichEditor bind:value={faction.description} />
                <div class="border-b border-[#8b7355]/40 pointer-events-none z-0"></div>
                <div class="mt-2 flex items-center justify-center">
                <button type="button" class="forge-btn forge-btn-base flex align-middle gap-2" onclick={() => handleFactionSave()}>
                    <span class="mb-1"><AnvilImpact height="1.6rem" /></span>
                    <span>Inscribe</span>
                </button>
                </div>
            </div>
        </div>
    {:else if modalType === 'delete' && faction.id !== null}
        <div class="border-[#1a0f0f] bg-[#2b190c] p-2 md:p-2 w-[25rem]! h-[22rem]! rounded-md!">
            <div class="wood-texture-bg absolute inset-0 pointer-events-none z-0"></div>

            <img src="{corners}" alt="" class="absolute -top-3 -left-3 w-15 h-15 pointer-events-none z-10">
            <img src="{corners}" alt="" class="absolute -top-3 -right-3 w-15 h-15 pointer-events-none z-10 scale-x-[-1]" />
            <img src="{corners}" alt="" class="absolute -bottom-3 -left-3 w-15 h-15 pointer-events-none z-10 scale-y-[-1]" />
            <img src="{corners}" alt="" class="absolute -bottom-3 -right-3 w-15 h-15 pointer-events-none z-10 rotate-180" />

            <div class="flex flex-col gap-2 parchment-background-inner relative w-full h-full z-10 border border-[#3b2a1e] p-4 rounded-md">
                {#if affiliations.length !== 0}
                    <span class="mt-4 font-bold">Beware, this faction cannot cease to be. Some affiliations belongs to it.</span>
                {:else}
                    <span class="mt-4 font-bold">Shall this faction fall to oblivion?</span>
                {/if}
                {#each affiliations as affiliation(affiliation.id)}
                    <span class="underline mt-2">
                        {affiliation.name + " | " + affiliation.type}
                    </span>
                {/each}
                <div class="flex flex-row justify-center px-6 mt-auto gap-4">
                    {#if affiliations.length === 0}
                        <button class="forge-btn forge-btn-base" onclick={() => handleFactionDelete()}>
                            <CrossMark height="1.5rem" />
                            <span class="ml-2">Agreed</span>
                        </button>
                    {/if}
                    <button class="forge-btn forge-btn-base" onclick={() => isModalOpen = false}>
                        <QuillInk height="1.5rem" />
                        <span class="ml-1">Keep</span>
                    </button>
                </div>
            </div>
        </div>
    {/if}

</Modal>
