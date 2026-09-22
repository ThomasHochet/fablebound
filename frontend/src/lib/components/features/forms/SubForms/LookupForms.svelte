<script lang="ts">
  import { onMount } from "svelte";
  import DataSelector from "$lib/components/ui/DataSelector.svelte";
  import { CreateLookup, DeleteLookup, FetchAllLookup, SaveLookup } from "$wails/world-builder/app";
  import { LookupItem } from "$wails/world-builder/internal/services/models";
  import { subscribe } from "$lib/functions/subscribe";
  import DataTableManager from "$lib/components/ui/DataTableManager.svelte";

  import BroomIcon from "@iconify-svelte/game-icons/components/b/broom.svelte"

  let genders = $state<LookupItem[]>([])
  let races = $state<LookupItem[]>([])
  let occupations = $state<LookupItem[]>([])
  let alignments = $state<LookupItem[]>([])
  let status = $state<LookupItem[]>([])

  async function handleCreate(lookupTable: string, label: string): Promise<LookupItem | null> {
    try {
      const newItem = await CreateLookup(lookupTable, label)
      return newItem
    } catch(err) {
      console.error(`Failed to create ${lookupTable}:`, err)
      return null
    }
  }

  async function handleUpdate(lookupTable: string, id: string | number, label: string): Promise<boolean | null> {
    try {
      await SaveLookup(id, label, lookupTable)
      return true
    } catch(err) {
      console.error(err)
      return null
    }
  }

  async function handleDelete(lookupTable: string, id: number) {
    try {
      await DeleteLookup(id, lookupTable)
    } catch(err) {
      console.error(`Failed to delete from ${lookupTable}:`, err)
    }
  }

  async function fetchData() {
    try {
      genders = await FetchAllLookup("gender")
      races = await FetchAllLookup("race")
      occupations = await FetchAllLookup("occupation")
      alignments = await FetchAllLookup("alignment")
      status = await FetchAllLookup("status")
    } catch(err) {
      console.error(err)
    }
  }

  onMount(() => {
    fetchData()

    subscribe('genders', fetchData)
    subscribe('races', fetchData)
    subscribe('occupations', fetchData)
    subscribe('alignments', fetchData)
    subscribe('statuses', fetchData)
  })
</script>

<section class="grid grid-cols-2 items-center justify-center justify-items-center p-4 gap-4">
    <div class="w-full rounded-lg overflow-hidden shadow-2xl border-4 border-[#2b190c]">
        <div class="p-5 rounded-t-md relative">
            <div class="wood-texture-bg"></div>

            <div class="iron-nail nail-tl"></div>
            <div class="iron-nail nail-tr"></div>
            <div class="iron-nail nail-bl"></div>
            <div class="iron-nail nail-br"></div>

            <div class="grid grid-cols-1 md:grid-cols-1 gap-4 relative z-10 my-2">
                <DataTableManager
                    bind:items={genders}
                    labelText="Genders"
                    placeholder="Add gender..."
                    onAdd={(label: string) => handleCreate('gender', label)}
                    onUpdate={(id: string | number, label: string) => handleUpdate('gender', id, label)}
                >
                    {#snippet actionCell(item)}
                        <button
                            type="button"
                            title="Delete  {item.label}"
                            onclick={() => handleDelete('gender', item.id)}
                            class="text-black cursor-pointer opacity-0 group-hover:opacity-100 transition-opacity p-1"
                        >
                            <BroomIcon height="1rem" color="black" class="" />
                        </button>
                    {/snippet}
                </DataTableManager>
            </div>
        </div>
    </div>

    <div class="w-full rounded-lg overflow-hidden shadow-2xl border-4 border-[#2b190c]">
        <div class="p-5 rounded-t-md relative">
            <div class="wood-texture-bg"></div>

            <div class="iron-nail nail-tl"></div>
            <div class="iron-nail nail-tr"></div>
            <div class="iron-nail nail-bl"></div>
            <div class="iron-nail nail-br"></div>

            <div class="grid grid-cols-1 md:grid-cols-1 gap-4 relative z-10 my-2">
                <DataTableManager
                    bind:items={races}
                    labelText="Races"
                    placeholder="Add race..."
                    onAdd={(label: string) => handleCreate('race', label)}
                    onUpdate={(id: string | number, label: string) => handleUpdate('race', id, label)}
                >
                    {#snippet actionCell(item)}
                        <button
                            type="button"
                            title="Delete  {item.label}"
                            onclick={() => handleDelete('race', item.id)}
                            class="text-black cursor-pointer opacity-0 group-hover:opacity-100 transition-opacity p-1"
                        >
                            <BroomIcon height="1rem" color="black" class="" />
                        </button>
                    {/snippet}
                </DataTableManager>
            </div>
        </div>
    </div>

    <div class="w-full rounded-lg overflow-hidden shadow-2xl border-4 border-[#2b190c]">
        <div class="p-5 rounded-t-md relative">
            <div class="wood-texture-bg"></div>

            <div class="iron-nail nail-tl"></div>
            <div class="iron-nail nail-tr"></div>
            <div class="iron-nail nail-bl"></div>
            <div class="iron-nail nail-br"></div>

            <div class="grid grid-cols-1 md:grid-cols-1 gap-4 relative z-10 my-2">
                <DataTableManager
                    bind:items={occupations}
                    labelText="Occupations"
                    placeholder="Add occupation..."
                    onAdd={(label: string) => handleCreate('occupation', label)}
                    onUpdate={(id: string | number, label: string) => handleUpdate('occupation', id, label)}
                >
                    {#snippet actionCell(item)}
                        <button
                            type="button"
                            title="Delete  {item.label}"
                            onclick={() => handleDelete('occupation', item.id)}
                            class="text-black cursor-pointer opacity-0 group-hover:opacity-100 transition-opacity p-1"
                        >
                            <BroomIcon height="1rem" color="black" class="" />
                        </button>
                    {/snippet}
                </DataTableManager>
            </div>
        </div>
    </div>

    <div class="w-full rounded-lg overflow-hidden shadow-2xl border-4 border-[#2b190c]">
        <div class="p-5 rounded-t-md relative">
            <div class="wood-texture-bg"></div>

            <div class="iron-nail nail-tl"></div>
            <div class="iron-nail nail-tr"></div>
            <div class="iron-nail nail-bl"></div>
            <div class="iron-nail nail-br"></div>

            <div class="grid grid-cols-1 md:grid-cols-1 gap-4 relative z-10 my-2">
                <DataTableManager
                    bind:items={alignments}
                    labelText="Alignments"
                    placeholder="Add Alignment..."
                    onAdd={(label: string) => handleCreate('alignment', label)}
                    onUpdate={(id: string | number, label: string) => handleUpdate('alignment', id, label)}
                >
                    {#snippet actionCell(item)}
                        <button
                            type="button"
                            title="Delete  {item.label}"
                            onclick={() => handleDelete('alignment', item.id)}
                            class="text-black cursor-pointer opacity-0 group-hover:opacity-100 transition-opacity p-1"
                        >
                            <BroomIcon height="1rem" color="black" class="" />
                        </button>
                    {/snippet}
                </DataTableManager>
            </div>
        </div>
    </div>

    <div class="w-full rounded-lg overflow-hidden shadow-2xl border-4 border-[#2b190c]">
        <div class="p-5 rounded-t-md relative">
            <div class="wood-texture-bg"></div>

            <div class="iron-nail nail-tl"></div>
            <div class="iron-nail nail-tr"></div>
            <div class="iron-nail nail-bl"></div>
            <div class="iron-nail nail-br"></div>

            <div class="grid grid-cols-1 md:grid-cols-1 gap-4 relative z-10 my-2">
                <DataTableManager
                    bind:items={status}
                    labelText="Status"
                    placeholder="Add status..."
                    onAdd={(label: string) => handleCreate('status', label)}
                    onUpdate={(id: string | number, label: string) => handleUpdate('status', id, label)}
                >
                    {#snippet actionCell(item)}
                        <button
                            type="button"
                            title="Delete  {item.label}"
                            onclick={() => handleDelete('status', item.id)}
                            class="text-black cursor-pointer opacity-0 group-hover:opacity-100 transition-opacity p-1"
                        >
                            <BroomIcon height="1rem" color="black" class="" />
                        </button>
                    {/snippet}
                </DataTableManager>
            </div>
        </div>
    </div>
</section>
