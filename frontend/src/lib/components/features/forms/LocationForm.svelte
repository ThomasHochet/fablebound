<script lang="ts">
    import { Location } from "$wails/world-builder/internal/models/models";
    import { FetchLocationTypes, FetchLocationSubtypes, FetchAllLocations, SaveLocation } from "$wails/world-builder/app";
    import RichEditor from "$lib/components/ui/RichEditor.svelte";
    import { onMount } from "svelte";
    import DataSelector from "$lib/components/ui/DataSelector.svelte";
    import StatusQuip from "$lib/components/ui/StatusQuip.svelte";

    let { id } = $props()
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
      e.preventDefault()

      const payload = new Location({
        id: id ? Number(id) : undefined,
        name: formState.name,
        type: formState.type,
        subtype: formState.subtype,
        parent_id: formState.parentId,
        description: formState.description
      })

      const result = await SaveLocation(payload)
      console.info(result)
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
        console.error(err)
      }
    }

    onMount(() => {
      fetchData()
    })
</script>

<section class="grid grid-cols-1 items-center justify-center p-4">
    <form onsubmit={handleSubmit}>
        <div class="w-full rounded-lg overflow-hidden shadow-2xl border-4 border-[#2b190c]">
            <div class="p-5 rounded-t-md relative">
                <div class="wood-texture-bg"></div>

                <div class="iron-nail nail-tl"></div>
                <div class="iron-nail nail-tr"></div>
                <div class="iron-nail nail-bl"></div>
                <div class="iron-nail nail-br"></div>

                <div class="grid grid-cols-1 md:grid-cols-2 gap-4 relative z-10 my-2">
                    <div>
                        <label for="title" class="block text-sm text-left ml-2 text-(--gold-text) mb-1 tracking-wider uppercase">
                            Location Title*
                        </label>
                        <div class="fantasy-input-wrapper flex items-center">
                            <input type="text" name="name" class="fantasy-input" bind:value={formState.name} placeholder="Location name" required />
                        </div>
                    </div>
                    <div class="relative">
                        <DataSelector items={typeOptions} bind:value={formState.type} placeholder="Continent, Kingdom, City" labelText="Location Type*" required={true} />
                    </div>
                    <div class="relative">
                        <DataSelector items={subtypeOptions} bind:value={formState.subtype} placeholder="Castle, Inn, Old farm..." labelText="Location Subtype" />
                    </div>
                    <div class="relative">
                        <DataSelector items={locationOptions} bind:value={formState.parentId} placeholder="Faerûn, Baldur's gate..." labelText="Parent location" />
                    </div>
                </div>
            </div>
            <div class="bg-(--wood-dark) p-2">
                <RichEditor bind:value={formState.description} />
                <div class="mt-2 flex items-center justify-between">
                    <StatusQuip isReady={isFormReady} missingType={missingFieldType} />
                    <button type="submit" class="fantasy-btn-xl fantasy-bone-n-coper ">
                        Inscribe
                    </button>
                </div>
            </div>


        </div>
    </form>
</section>

<!-- <section class="p-2">
    <form onsubmit={handleSubmit}>
        <div class="grid grid-cols-12 gap-1">
            <div class="fantasy-border fantasy-border-brown fantasy-input-wrapper fantasy-input-inlay col-span-4 col-start-3">
                <input type="text" name="name" bind:value={formState.name} placeholder="Location name" required />
            </div>
            <div class="fantasy-border fantasy-border-brown fantasy-input-wrapper fantasy-input-inlay col-span-4 col-start-7">
                <input type="text" name="type" id="type" list="types" bind:value={formState.type} placeholder="World, Continent, Kingdom..." required title="Specifies what kind of location it is. A world, a continent, a kingdom...">
                <datalist id="types">
                    {#each types as type}
                        <option value={type}></option>
                    {/each}
                </datalist>
            </div>
            <div class="fantasy-border fantasy-border-brown fantasy-input-wrapper fantasy-input-inlay col-span-4 col-start-3">
                <input type="text" name="subtype" id="subtype" list="subtypes" bind:value={formState.subtype} placeholder="Castle, Inn, Old Farm..." title="A bigger specification for clarification, a Lord's dwelling, a dungeon keep...">
                <datalist id="subtypes">
                    {#each subtypes as subtype}
                        <option value={subtype}></option>
                    {/each}
                </datalist>
            </div>
            <div class="fantasy-border fantasy-border-brown fantasy-input-wrapper fantasy-input-inlay col-span-4 col-start-7">
                <input type="text" name="parent-location" id="parent-location" list="locations" bind:value={formState.parentName} placeholder="Faerûn, Baldur's gate..." title="If the location is within another. A Kingdom within a Continent, nor a Town inside a Kingdom.">
                <datalist id="locations">
                    {#each locations as location(location.id)}
                        <option value={location.name}></option>
                    {/each}
                </datalist>
            </div>
        </div>

        <RichEditor bind:value={formState.description} />
        <div class="flex">
            <button type="submit" class="fantasy-btn fantasy-bone-n-coper">
                Save
            </button>
        </div>
    </form>
</section> -->
