<script lang="ts">
    import { Location } from "$wails/world-builder/internal/models/models";
    import { FetchLocationTypes, FetchLocationSubtypes, FetchAllLocations, SaveLocation } from "$wails/world-builder/app";
    import RichEditor from "$lib/components/ui/RichEditor.svelte";
    import { onMount } from "svelte";
    import { Events } from "@wailsio/runtime";

    let { id } = $props()
    let locations = $state<Location[]>([])
    let types = $state<string[]>([])
    let subtypes = $state<string[]>([])

    let formState = $state({
      name: '',
      type: '',
      subtype: '',
      parentName: '',
      description: 'Describe your location... Is it a city? A continent? A kingdom?'
    })

    async function handleSubmit(e: Event) {
      e.preventDefault()

      const selectedParent = locations.find(loc => loc.name === formState.parentName)
      const parentId = selectedParent ? selectedParent.id : null

      const payload = new Location({
        id: id ? Number(id) : undefined,
        name: formState.name,
        type: formState.type,
        subtype: formState.subtype,
        parent_id: parentId,
        description: formState.description
      })

      const result = await SaveLocation(payload)
      console.info(result)
    }

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
            if (locToEdit.parent_id) {
              const parentLoc = locations.find(loc => loc.id === locToEdit.parent_id)
              console.log(parentLoc)
              if (parentLoc) {
                formState.parentName = parentLoc.name
              }
            }
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

<section class="p-2">
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
</section>
