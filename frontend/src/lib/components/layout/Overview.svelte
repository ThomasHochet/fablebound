<script lang="ts">
    import { onMount } from "svelte";
    import corners from "$lib/../assets/images/corners.png"
    import AnvilImpact from "@iconify-svelte/game-icons/components/a/anvil-impact.svelte"
    import GearHammer from "@iconify-svelte/game-icons/components/g/gear-hammer.svelte"
    import DwarfHelmet from "@iconify-svelte/game-icons/components/d/dwarf-helmet.svelte"
    import WorldIcon from "@iconify-svelte/game-icons/components/w/world.svelte"
    import ScrollUnfurled from "@iconify-svelte/game-icons/components/s/scroll-unfurled.svelte"
    import MedievalTown from "@iconify-svelte/game-icons/components/m/medieval-village-01.svelte"
    import Banner from "@iconify-svelte/game-icons/components/v/vertical-banner.svelte"


    import Dropdown from "../ui/Dropdown.svelte";

    import WorldCodex from "./Codex/WorldCodex.svelte";
    import LocationCodex from "./Codex/LocationCodex.svelte";
    import LoreCodex from "./Codex/LoreCodex.svelte";
    import FactionCodex from "./Codex/FactionCodex.svelte";
    import CharacterCodex from "./Codex/CharacterCodex.svelte";
    import Modal from "../features/Modal.svelte";
    import { OpenEditorWindow } from "$wails/world-builder/app";
    import { logError } from "$lib/logger";

    let selectedOverview = $state(1)
    let isLoading = $state(true)
    let isForgeModalOpen = $state(false)

    const overviewTypes = [
      { id: 1, label: "World" },
      { id: 2, label: "Locations" },
      { id: 3, label: "Lores" },
      { id: 4, label: "Factions" },
      { id: 5, label: "Characters" },
    ]

    let selectedOverviewType = $derived(overviewTypes.find(o => o.id === selectedOverview))

    const viewMap: Record<number, any> = {
      1: WorldCodex,
      2: LocationCodex,
      3: LoreCodex,
      4: FactionCodex,
      5: CharacterCodex
    }

    let ActiveView = $derived(viewMap[selectedOverview])

    function handleForge(type: string) {
      isForgeModalOpen = false
      OpenEditorWindow(type, 0).catch(err => {
        logError(`Failed to open ${type} window, `, err)
      })
    }

    onMount(() => {

    })
</script>

<section class="h-screen w-full p-4 md:p-6">

    <div class="relative w-full h-full shadow-2xl rounded-sm border-4 border-[#1a0f0f] bg-[#2b190c] p-1 md:p-1 grid grid-cols-1 grid-rows-1">

        <div class="wood-texture-bg absolute inset-0 pointer-events-none z-0"></div>

        <img src="{corners}" alt="" class="absolute -top-4 -left-4 w-15 h-15 pointer-events-none z-10">
        <img src="{corners}" alt="" class="absolute -top-4 -right-4 w-15 h-15 pointer-events-none z-10 scale-x-[-1]" />
        <img src="{corners}" alt="" class="absolute -bottom-4 -left-4 w-15 h-15 pointer-events-none z-10 scale-y-[-1]" />
        <img src="{corners}" alt="" class="absolute -bottom-4 -right-4 w-15 h-15 pointer-events-none z-10 rotate-180" />

        <div class="parchment-background relative w-full h-full grid grid-rows-[auto_1fr] z-10 border border-[#3b2a1e]">
            <div class="absolute inset-2 border border-[#8b7355]/40 pointer-events-none z-0"></div>

            <header class="relative z-30 w-full px-6 py-4 border-b border-[#8b7355]/30 grid grid-cols-1 md:grid-cols-[1fr_auto] items-center gap-4">

                <div class="order-1 md:order-2 md:right-0 grid grid-cols-[auto_1fr_auto] items-center gap-3 py-2 md:py-0">
                    <div class="h-[1px] w-8 md:w-12 bg-gradient-to-r from-transparent to-[#7a644f]"></div>
                    <h1 class="font-serif text-lg md:text-2xl font-bold tracking-widest text-(--ink-color) uppercase whitespace-nowrap text-center">
                        The {selectedOverviewType?.label} Codex
                    </h1>
                    <div class="h-[1px] w-8 md:w-12 bg-gradient-to-l from-transparent to-[#7a644f]"></div>
                </div>

                <div class="order-2 md:order-1 grid grid-cols-1 sm:grid-cols-[auto_auto_auto_auto_1fr] items-end gap-3">
                    <div class="grid grid-rows-[auto_auto]">
                        <Dropdown items={overviewTypes} bind:value={selectedOverview} labelText="Overview Type" />
                    </div>

                    <div class="grid grid-rows-[auto_auto]">
                        <label class="text-xs font-serif mb-1 text-left ml-1">Search</label>
                        <div class="relative grid items-center w-fit">
                            <svg class="w-4 h-4 absolute left-2.5 text-[#8c7661] z-10" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"/>
                            </svg>
                            <input type="text" class="border border-[#b7a48d] rounded-md pl-8 pr-3 py-1.5 text-sm font-serif text-(--ink-color) placeholder-[#9c8974] focus:outline-none w-48 sm:w-64 shadow-inner">
                        </div>
                    </div>

                    <button class="forge-btn gap-1" onclick={() => (isForgeModalOpen = true)}>
                        FORGE NEW <AnvilImpact height="1.2rem" />
                    </button>

                    <button class="forge-btn gap-1" onclick={() => handleForge('params')}>
                        <GearHammer height="1.2rem" />
                    </button>
                </div>

            </header>

            <main class="relative z-10 w-full h-full min-h-0 overflow-y-auto! p-6 grid grid-cols-1 md:grid-cols-2! grid-rows-[max-content] gap-4">
                <!-- To use for loading -->
                <!-- <p class="col-span-2 font-serif text-(--ink-color) text-lg text-center">
                    Your entries for the {selectedOverviewType?.label} codex will populate here...
                </p> -->

                {#if ActiveView}
                    <ActiveView />
                {/if}
            </main>
        </div>
    </div>
</section>

<Modal bind:open={isForgeModalOpen}>
    <div class="border-[#1a0f0f] bg-[#2b190c] p-2 md:p-2 w-[25rem]! h-[22rem]! rounded-md!">
        <div class="wood-texture-bg absolute inset-0 pointer-events-none z-0"></div>

        <img src="{corners}" alt="" class="absolute -top-3 -left-3 w-15 h-15 pointer-events-none z-10">
        <img src="{corners}" alt="" class="absolute -top-3 -right-3 w-15 h-15 pointer-events-none z-10 scale-x-[-1]" />
        <img src="{corners}" alt="" class="absolute -bottom-3 -left-3 w-15 h-15 pointer-events-none z-10 scale-y-[-1]" />
        <img src="{corners}" alt="" class="absolute -bottom-3 -right-3 w-15 h-15 pointer-events-none z-10 rotate-180" />

        <div class="flex flex-col gap-2 items-center justify-center parchment-background-inner relative w-full h-full z-10 border border-[#3b2a1e] p-4 rounded-md">
                <button class="forge-btn z-50 w-50 flex items-center justify-start px-5 gap-4" onclick={() => handleForge('character')}>
                    <span>
                        <DwarfHelmet height="2rem" />
                    </span>
                    <span>
                        FORGE CHARACTER
                    </span>
                </button>

                <button class="forge-btn z-50 w-50 flex items-center justify-start px-5 gap-4" onclick={() => handleForge('lore')}>
                    <span>
                        <ScrollUnfurled height="2rem" />
                    </span>
                    <span>
                        FORGE LORE
                    </span>
                </button>

                <button class="forge-btn z-50 w-50 flex items-center justify-start px-5 gap-4" onclick={() => handleForge('location')}>
                    <span>
                        <MedievalTown height="2rem" />
                    </span>
                    <span>
                        FORGE LOCATION
                    </span>
                </button>

                <button class="forge-btn z-50 w-50 flex items-center justify-start px-5 gap-4" onclick={() => handleForge('faction')}>
                    <span>
                        <Banner height="2rem" />
                    </span>
                    <span>
                        FORGE FACTION
                    </span>

                </button>

                <button class="forge-btn z-50 w-50 flex items-center justify-start px-5 gap-4" onclick={() => handleForge('world')}>
                    <span>
                        <WorldIcon height="2rem" />
                    </span>
                    <span>
                        FORGE WORLD
                    </span>

                </button>
        </div>




        <div class="col-span-1"></div>

    </div>
</Modal>
