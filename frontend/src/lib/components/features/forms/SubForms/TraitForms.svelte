<script lang="ts">
  import { onMount } from "svelte";
  import { CreateLookup, CreateTrait, DeleteLookup, DeleteTrait, FetchAllTraits, SaveLookup, SaveTrait } from "$wails/world-builder/app";
  import { TraitItem } from "$wails/world-builder/internal/services/models";
  import { subscribe } from "$lib/functions/subscribe";
  import DataTableManager from "$lib/components/ui/DataTableManager.svelte";

  import BroomIcon from "@iconify-svelte/game-icons/components/b/broom.svelte"
    import { logError } from "$lib/logger";

  let personalityTraits = $state<TraitItem[]>([])
  let strengths = $state<TraitItem[]>([])
  let flaws = $state<TraitItem[]>([])
  let weaknesses = $state<TraitItem[]>([])
  let fears = $state<TraitItem[]>([])

  async function handleCreate(traitTable: string, label: string): Promise<TraitItem | null> {
    try {
      const newItem = await CreateTrait(traitTable, label)
      return newItem
    } catch(err) {
      logError(`Failed to create ${traitTable}:`, err)
      return null
    }
  }

  async function handleUpdate(traitTable: string, id: string | number, label: string): Promise<boolean | null> {
    try {
      await SaveTrait(id, label, traitTable)
      return true
    } catch(err) {
      logError(`Failed to update ${traitTable}`, err)
      return null
    }
  }

  async function handleDelete(traitTable: string, id: number) {
    try {
      await DeleteTrait(traitTable, id)
    } catch(err) {
      logError(`Failed to delete from ${traitTable}:`, err)
    }
  }

  async function fetchData() {
    try {
      personalityTraits = await FetchAllTraits("personality")
      strengths = await FetchAllTraits("strength")
      flaws = await FetchAllTraits("flaw")
      weaknesses = await FetchAllTraits("weakness")
      fears = await FetchAllTraits("fear")
    } catch(err) {
      logError("Failed to fetch Traits data", err)
    }
  }

  onMount(() => {
    fetchData()

    subscribe('personality_traits', fetchData)
    subscribe('strengths', fetchData)
    subscribe('flaws', fetchData)
    subscribe('weaknesses', fetchData)
    subscribe('fears', fetchData)
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
                    bind:items={personalityTraits}
                    labelText="Personality Traits"
                    placeholder="Add personality trait..."
                    onAdd={(label: string) => handleCreate('personality', label)}
                    onUpdate={(id: string | number, label: string) => handleUpdate('personality', id, label)}
                >
                    {#snippet actionCell(item)}
                        <button
                            type="button"
                            title="Delete  {item.label}"
                            onclick={() => handleDelete('personality', item.id)}
                            class="text-black cursor-pointer opacity-0 group-hover:opacity-100 transition-opacity p-1 opacity-0 group-hover:opacity-100 transition-opacity p-1"
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
                    bind:items={strengths}
                    labelText="Strengths"
                    placeholder="Add strength..."
                    onAdd={(label: string) => handleCreate('strength', label)}
                    onUpdate={(id: string | number, label: string) => handleUpdate('strength', id, label)}
                >
                    {#snippet actionCell(item)}
                        <button
                            type="button"
                            title="Delete  {item.label}"
                            onclick={() => handleDelete('strength', item.id)}
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
                    bind:items={flaws}
                    labelText="Flaws"
                    placeholder="Add flaw..."
                    onAdd={(label: string) => handleCreate('flaw', label)}
                    onUpdate={(id: string | number, label: string) => handleUpdate('flaw', id, label)}
                >
                    {#snippet actionCell(item)}
                        <button
                            type="button"
                            title="Delete  {item.label}"
                            onclick={() => handleDelete('flaw', item.id)}
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
                    bind:items={weaknesses}
                    labelText="Weaknesses"
                    placeholder="Add weakness..."
                    onAdd={(label: string) => handleCreate('weakness', label)}
                    onUpdate={(id: string | number, label: string) => handleUpdate('weakness', id, label)}
                >
                    {#snippet actionCell(item)}
                        <button
                            type="button"
                            title="Delete  {item.label}"
                            onclick={() => handleDelete('weakness', item.id)}
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
                    bind:items={fears}
                    labelText="Fears"
                    placeholder="Add fear..."
                    onAdd={(label: string) => handleCreate('fear', label)}
                    onUpdate={(id: string | number, label: string) => handleUpdate('fear', id, label)}
                >
                    {#snippet actionCell(item)}
                        <button
                            type="button"
                            title="Delete  {item.label}"
                            onclick={() => handleDelete('fear', item.id)}
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
