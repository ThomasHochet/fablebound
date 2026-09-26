<script lang="ts">
    import { Location } from "$wails/world-builder/internal/models/models";
    import { FetchLocationTypes, FetchLocationSubtypes, FetchAllLocations, SaveLocation } from "$wails/world-builder/app";
    import RichEditor from "$lib/components/ui/RichEditor.svelte";
    import { onMount } from "svelte";
    import DataSelector from "$lib/components/ui/DataSelector.svelte";
    import StatusQuip from "$lib/components/ui/StatusQuip.svelte";
    import corners from "$lib/../assets/images/corners.png"
    import ParchmentTitle from "$lib/components/ui/ParchmentTitle.svelte";
    import AnvilImpact from "@iconify-svelte/game-icons/components/a/anvil-impact.svelte";
    import  SaveToast, { type SaveStatus } from "$lib/components/ui/SaveToast.svelte";
    import { Window } from "@wailsio/runtime";
    import { logError } from "$lib/logger";

    let { id } = $props()
    let status = $state<SaveStatus>('idle')
    let locations = $state<Location[]>([])
    let types = $state<string[]>([])
    let subtypes = $state<string[]>([])

    let formState = $state({
      name: '',
      type: '',
      subtype: '',
      parentId: null as number | null,
      description: 'Describe your location... Is it a city? A continent? A kingdom?'
    })

    let typeOptions = $derived(types.map(type => ({ id: type, label: type})))
    let subtypeOptions = $derived(subtypes.map(st => ({ id: st, label: st })))
    let locationOptions = $derived(
      locations
        .filter(loc => loc.id !== Number(id)) // Self reference issue
        .map(loc => ({id: loc.id, label: loc.name}))
    )

    async function handleSubmit(e: Event) {
      try {
        e.preventDefault()
        status = 'saving'

        const payload = new Location({
          id: id ? Number(id) : undefined,
          name: formState.name,
          type: formState.type,
          subtype: formState.subtype,
          parent_id: formState.parentId,
          description: formState.description
        })

        const result = await SaveLocation(payload)
        status = 'saved'

        setTimeout(() => {
          if (status === 'saved') status = 'idle'
        }, 3000)

        setTimeout(() => {
          Window.Close()
        }, 1500)
      } catch(err) {
        logError("Failed to save Location", err)
        status = 'error'
      }
    }

    let isFormReady = $derived(Boolean(formState.name.trim() && formState.type === ''))
    let missingFieldType = $derived.by(() => {
      if (formState.type === '') return 'location' as const
      return null
    })

    async function fetchData() {
      try {
        types = await FetchLocationTypes()
        subtypes = await FetchLocationSubtypes()
        locations = await FetchAllLocations()

        if (id) {
          const locToEdit = locations.find(loc => loc.id === Number(id))
          if (locToEdit) {
            formState.name = locToEdit.name || ''
            formState.type = locToEdit.type || ''
            formState.subtype = locToEdit.subtype || ''
            formState.description = locToEdit.description || ''
            formState.parentId = locToEdit.parent_id ?? null
          }
        }
      } catch(err) {
        logError("Failed to fetch Location data", err)
      }
    }

    onMount(() => {
      fetchData()
    })
</script>

<section class="h-full grid grid-cols-1 items-center justify-center px-2 ">
    <div class="relative z-20 -mb-10 pointer-events-none">
        <ParchmentTitle title="Locations" subtitle="Every record, every maps." />
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
                            Location Title*
                        </label>
                        <input type="text" name="name" class="forge-input" bind:value={formState.name} placeholder="Location name" required />
                    </div>
                    <div class="relative">
                        <DataSelector items={typeOptions} bind:value={formState.type} placeholder="Continent, Kingdom, City" labelText="Location Type*" required={true} variant='forge' />
                    </div>
                    <div class="relative">
                        <DataSelector items={subtypeOptions} bind:value={formState.subtype} placeholder="Castle, Inn, Old farm..." labelText="Location Subtype" variant='forge' />
                    </div>
                    <div class="relative">
                        <DataSelector items={locationOptions} bind:value={formState.parentId} placeholder="Faerûn, Baldur's gate..." labelText="Parent location" variant='forge' />
                    </div>
                </div>
                <RichEditor bind:value={formState.description} />
                <div class="border-b border-[#8b7355]/40 pointer-events-none z-0"></div>
                <div class="mt-2 flex items-center justify-between">
                    <StatusQuip isReady={isFormReady} missingType={missingFieldType} />
                    <button type="submit" class="forge-btn forge-btn-base flex align-middle gap-2">
                        <span><AnvilImpact height="1.6rem" /></span>
                        <span>Inscribe</span>
                    </button>
                </div>
            </div>
        </div>
    </form>

    <SaveToast {status} />
</section>
