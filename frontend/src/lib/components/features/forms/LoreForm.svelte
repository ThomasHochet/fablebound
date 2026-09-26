<script lang="ts">
    import DataSelector from "$lib/components/ui/DataSelector.svelte";
    import ParchmentTitle from "$lib/components/ui/ParchmentTitle.svelte";
    import RichEditor from "$lib/components/ui/RichEditor.svelte";
    import StatusQuip from "$lib/components/ui/StatusQuip.svelte";
    import { subscribe } from "$lib/functions/subscribe";
    import { FetchCategories, FetchLore, SaveLore, SaveLoreCategory } from "$wails/world-builder/app";
    import { Lore, LoreCategory } from "$wails/world-builder/internal/models/models";
    import corners from "$lib/../assets/images/corners.png"

    import { onMount } from "svelte";
    import AnvilImpact from "@iconify-svelte/game-icons/components/a/anvil-impact.svelte";
    import  SaveToast, { type SaveStatus } from "$lib/components/ui/SaveToast.svelte";
    import { Window } from "@wailsio/runtime";
    import { logError } from "$lib/logger";

    let { id } = $props()
    let status = $state<SaveStatus>('idle')
    let isLoading = $state(true)
    let lore = $state({
      title: '',
      content: 'The start of a new story!',
      categoryName: '',
      categoryId: null as unknown | number
    })
    let loreCategories = $state<LoreCategory[]>([])

    async function handleSubmit(e: Event) {
      e.preventDefault()
      status = 'saving'
      try {
        if (!lore.title.trim() || !lore.categoryId) {
          console.warn("Please select a valid category and title before saving.")
          return;
        }

        const payload = new Lore({
          id: id ? Number(id) : undefined,
          title: lore.title,
          content: lore.content,
          category_id: lore.categoryId
        })
        const result = await SaveLore(payload)
        status = 'saved'

        setTimeout(() => {
          if (status === 'saved') status = 'idle'
        }, 3000)

        setTimeout(() => {
          Window.Close()
        }, 1500)
      } catch(err) {
        logError("Failed to save Lore", err)
        status = 'error'
      }
    }

    async function handleCategoryCreate(newLabel: string) {
      try {
        const created = await SaveLoreCategory(lore.categoryId, newLabel)

        if (created)
          loreCategories = [...loreCategories, {id: created?.id, label: created?.label}]
        return null
      } catch(err) {
        logError("Could not create a lore category.", err)
        return null
      }
    }

    async function loadLore(id: number) {
      try {
        const loreData = await FetchLore(id)
        lore.title = loreData?.title ?? ''
        lore.content = loreData?.content ?? ''
        lore.categoryName = loreData?.lore_category?.label ?? ''
        lore.categoryId = loreData?.category_id ?? 0
      } catch(err) {
        logError(`Failed to load Lore ID:${id}.`, err)
      } finally {
        isLoading = false
      }
    }

    async function loadLoreCategories() {
      try {
        loreCategories = await FetchCategories()
      } catch(err) {
        logError("Failed to load Lore Categories", err)
      }
    }

    let isFormReady = $derived(Boolean(lore.title.trim() && lore.categoryId));

    let missingFieldType = $derived.by(() => {
        if (!lore.categoryId) return 'lore' as const;
        return null;
    });

    onMount(() => {
      loadLoreCategories()
      const numId = Number(id)
      if (!isNaN(numId)) {
        loadLore(numId)
      }
    })
</script>


<section class="h-full grid grid-cols-1 items-center justify-center px-2 ">
    <div class="relative z-20 -mb-10 pointer-events-none">
        <ParchmentTitle title="Lore" subtitle="A mundane record still holds experiences worth sharing." />
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
                            Chronicle Title*
                        </label>
                        <input type="text" name="title" class="forge-input" required bind:value={lore.title} placeholder="A brief day for mister Catmancer."  />
                    </div>
                    <div class="relative">
                        <DataSelector items={loreCategories} bind:value={lore.categoryId} onCreate={handleCategoryCreate} placeholder="Select or forge a new category..." labelText="Category*" required={true} variant="forge" />
                    </div>
                </div>
                <RichEditor bind:value={lore.content} />
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

    <SaveToast {status} />
</section>
