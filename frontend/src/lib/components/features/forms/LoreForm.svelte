<script lang="ts">
    import DataSelector from "$lib/components/ui/DataSelector.svelte";
    import RichEditor from "$lib/components/ui/RichEditor.svelte";
    import StatusQuip from "$lib/components/ui/StatusQuip.svelte";
    import { subscribe } from "$lib/functions/subscribe";
    import { FetchCategories, FetchLore, SaveLore, SaveLoreCategory } from "$wails/world-builder/app";
    import { Lore, LoreCategory } from "$wails/world-builder/internal/models/models";
    import { onMount } from "svelte";


    let { id } = $props()
    let isLoading = $state(true)
    let lore = $state({
      title: '',
      content: 'The start of a new story!',
      categoryName: '',
      categoryId: 0
    })
    let loreCategories = $state<LoreCategory[]>([])

    async function handleSubmit(e: Event) {
      e.preventDefault()

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
      console.log(result)
    }

    async function handleCategoryCreate(newLabel: string) {
      try {
        const created = await SaveLoreCategory(lore.categoryId, newLabel)

        if (created)
          loreCategories = [...loreCategories, {id: created?.id, label: created?.label}]
        return null
      } catch(err) {
        console.error("Could not create a lore category.", err)
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
        console.error(err)
      } finally {
        isLoading = false
      }
    }

    async function loadLoreCategories() {
      try {
        loreCategories = await FetchCategories()
      } catch(err) {
        console.error(err)
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


<section class="grid grid-cols-1 items-center justify-center p-4">
    <form onsubmit={handleSubmit}>
        <div class="w-full rounded-lg overflow-hidden shadow-2xl border-4 border-[#2b190c]">
            <div class="p-5 rounded-t-md relative">
                <div class="wood-texture-bg"></div>

                <div class="iron-nail nail-tl"></div>
                <div class="iron-nail nail-tr"></div>
                <div class="iron-nail nail-bl"></div>
                <div class="iron-nail nail-br"></div>

                <div class="grid grid-cols-1 md:grid-cols-2 gap-4 relative z-10">
                    <div>
                        <label for="title" class="block text-sm text-left ml-2 text-(--gold-text) mb-1 tracking-wider uppercase">
                            Chronicle Title*
                        </label>
                        <div class="fantasy-input-wrapper flex items-center">
                            <input type="text" name="title" class="fantasy-input" required bind:value={lore.title} placeholder="A brief day for mister Catmancer."  />
                        </div>
                    </div>
                    <div class="relative">
                        <DataSelector items={loreCategories} bind:value={lore.categoryId} onCreate={handleCategoryCreate} placeholder="Select or forge a new category..." labelText="Category*" required={true} />
                    </div>
                </div>
            </div>
            <div class="bg-(--wood-dark) p-2">
                <RichEditor bind:value={lore.content} />
                <div class="mt-2 flex items-center justify-between">
                    <StatusQuip isReady={isFormReady} missingType={missingFieldType} />
                    <!-- <span class="ml-2 text-sm text-[#a38c71] italic font-cinzel">Status: Ready to be inscribed</span> -->
                    <button type="submit" class="fantasy-btn-xl fantasy-bone-n-coper ">
                        Inscribe
                    </button>
                </div>
            </div>


        </div>
    </form>
</section>
