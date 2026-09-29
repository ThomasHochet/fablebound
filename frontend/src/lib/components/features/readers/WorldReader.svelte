<script lang="ts">
    import ParchmentTitle from "$lib/components/ui/ParchmentTitle.svelte";
    import corners from "$lib/../assets/images/corners.png"
    import RichEditor from "$lib/components/ui/RichEditor.svelte";
    import { logError } from "$lib/logger";
    import { FetchArticle } from "$wails/world-builder/app";
    import { Article } from "$wails/world-builder/internal/models/models";
    import Mountain from "$lib/../assets/images/moutain_bg.png"


    let { id } = $props()
    let article = $state<Article | null>(null)

    async function fetchWorldArticle(id: any) {
      const numId = Number(id)
      if (!isNaN(numId)) {
        try {
          article = await FetchArticle(numId)
        } catch(err) {
          logError(`Failed to fetch article ${numId}`, err)
        }
      }
    }

    $effect(() => {
      fetchWorldArticle(id)
    })
</script>

{#if article}
    <article class="h-full grid grid-cols-1 items-center justify-center px-2">
        <div class="relative z-20 -mb-10 pointer-events-none">
            <ParchmentTitle title={article.title} subtitle={article.category} />
        </div>
        <div class="relative w-full h-full shadow-2xl rounded-sm border-4 border-[#1a0f0f] bg-[#2b190c] p-1 md:p-1 grid grid-cols-1 grid-rows-1 min-h-56">
            <div class="wood-texture-bg"></div>

            <img src="{corners}" alt="" class="absolute -top-4 -left-4 w-15 h-15 pointer-events-none z-10">
            <img src="{corners}" alt="" class="absolute -top-4 -right-4 w-15 h-15 pointer-events-none z-10 scale-x-[-1]" />
            <img src="{corners}" alt="" class="absolute -bottom-4 -left-4 w-15 h-15 pointer-events-none z-10 scale-y-[-1]" />
            <img src="{corners}" alt="" class="absolute -bottom-4 -right-4 w-15 h-15 pointer-events-none z-10 rotate-180" />

            <div class="parchment-background relative w-full h-full grid grid-rows-[auto_1fr] z-10 border border-[#3b2a1e]">
                <div class="absolute inset-2 border border-[#8b7355]/40 pointer-events-none z-0"></div>

            <RichEditor disabled={true} value={article.description} classes="text-left! py-2" />
            </div>
            <img src={Mountain} alt="" class="opacity-20 absolute bottom-0 right-20 mb-2 mr-2 z-21" />
        </div>
</article>
{:else}
    <p>Loading World article</p>
{/if}
