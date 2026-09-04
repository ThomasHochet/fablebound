<script lang="ts">
    import RichEditor from "$lib/components/ui/RichEditor.svelte";
    import StatusQuip from "$lib/components/ui/StatusQuip.svelte";
    import DataSelector from "$lib/components/ui/DataSelector.svelte";
    import { onMount } from "svelte";
    import { FetchArticleCategories } from "$wails/world-builder/app";

    let { id } = $props()
    let formState = $state({
      title: '',
      description: 'A fine day for a good story, wouldn\'t you agree?',
      category: '',
      tags: ''
    })
    let categories = $state<[]>([])

    let categoryOptions = $derived(categories.map(c => ({id: c, label: c})))

    let isFormReady = $derived(Boolean(formState.title.trim() && formState.category.trim() && formState.description.trim()))
    let missingFieldType = $derived.by(() => {
      return null
    })

    function handleSubmit(e: SubmitEvent) {
      e.preventDefault()


    }

    async function loadCategories() {
      try {
        categories = await FetchArticleCategories()
      } catch(err){
        console.error(err)
      }
    }

    onMount(() => {
      loadCategories()
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
                            Story Title*
                        </label>
                        <div class="fantasy-input-wrapper flex items-center">
                            <input type="text" name="title" class="fantasy-input" required bind:value={formState.title} placeholder="The world is at your attention."  />
                        </div>
                    </div>
                    <div class="relative">
                        <DataSelector items={categoryOptions} bind:value={formState.category} placeholder="" labelText="Category*" required={true} />
                    </div>
                    <div>
                        <label for="title" class="block text-sm text-left ml-2 text-(--gold-text) mb-1 tracking-wider uppercase">
                            Tags
                        </label>
                        <div class="fantasy-input-wrapper flex items-center">
                            <input type="text" name="title" class="fantasy-input" disabled bind:value={formState.tags} placeholder="Coming soon..."  />
                        </div>
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
