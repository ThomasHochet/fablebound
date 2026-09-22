<script lang="ts">
    import RichEditor from "$lib/components/ui/RichEditor.svelte";
    import StatusQuip from "$lib/components/ui/StatusQuip.svelte";
    import DataSelector from "$lib/components/ui/DataSelector.svelte";
    import corners from "$lib/../assets/images/corners.png"
    import { onMount } from "svelte";
    import { SaveArticle, FetchArticle, FetchArticleCategories } from "$wails/world-builder/app";
    import ParchmentTitle from "$lib/components/ui/ParchmentTitle.svelte";
    import AnvilImpact from "@iconify-svelte/game-icons/components/a/anvil-impact.svelte"
    import { Article } from "$wails/world-builder/internal/models/models";
    import SaveToast, { type SaveStatus } from "$lib/components/ui/SaveToast.svelte";
    import { Window } from "@wailsio/runtime";


    let { id } = $props()
    let formState = $state({
      title: '',
      description: 'A fine day for a good story, wouldn\'t you agree?',
      category: '',
      tags: ''
    })
    let status = $state<SaveStatus>('idle')
    let categories = $state<string[]>([])

    let categoryOptions = $derived(categories.map(c => ({id: c, label: c})))

    let isFormReady = $derived(Boolean(formState.title.trim() && formState.category.trim() && formState.description.trim()))
    let missingFieldType = $derived.by(() => {
      return null
    })

    async function handleSubmit(e: SubmitEvent) {
      e.preventDefault()

      status = 'saving'

      try {
        const payload = new Article({
          id: id ? Number(id) : undefined,
          title: formState.title,
          description: formState.description,
          category: formState.category,
          tags: formState.tags
        })

        const result = await SaveArticle(payload)
        status = 'saved'

        setTimeout(() => {
          if (status === 'saved') status = 'idle'
        }, 3000)

        setTimeout(() => {
          Window.Close()
        }, 1500)
      } catch(err) {
        console.error("Failed to save world article:", err)
        status = 'error'
      }
    }

    async function loadCategories() {
      try {
        categories = await FetchArticleCategories()
      } catch(err){
        console.error(err)
      }
    }

    async function loadArticle(id: number) {
      try {
        const aData = await FetchArticle(id)
        formState.title = aData?.title ?? ''
        formState.description = aData?.description ?? ''
        formState.category = aData?.category ?? ''
        formState.tags = aData?.tags ?? ''
      } catch(err) {
        console.error(err)
      }
    }

    onMount(() => {
      loadCategories()
      const numId = Number(id)
      if (!isNaN(numId))
        loadArticle(numId)
    })
</script>

<section class="h-full grid grid-cols-1 items-center justify-center px-2 ">
    <div class="relative z-20 -mb-10 pointer-events-none">
        <ParchmentTitle title="World" subtitle="Every little thing matter as long as it's written." />
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
                                Story Title*
                            </label>
                            <input type="text" name="title" class="forge-input" required bind:value={formState.title} placeholder="The world is at your attention."  />
                        </div>
                        <div class="relative">
                            <DataSelector items={categoryOptions} bind:value={formState.category} placeholder="" labelText="Category*" required={true} variant="forge" />
                        </div>
                        <div class="grid grid-rows-[auto-auto]">
                            <label for="tags" class="forge-input-label">
                                Tags
                            </label>
                            <input type="text" name="tags" class="forge-input forge-input-disabled" disabled bind:value={formState.tags} placeholder="Coming soon..."  />
                        </div>
                    </div>

                    <RichEditor bind:value={formState.description} />
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
