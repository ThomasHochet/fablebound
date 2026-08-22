<script>
    import { FetchAllArticles } from "$wails/world-builder/app";
    import { onMount } from "svelte";
    import  coverImage from "$lib/../assets/images/parchments/wooden-floor-background.jpg"

    let articles = $state([])
    let isLoading = $state(true)
    let error = $state(null)

    onMount(async () => {
      try {
        articles = await FetchAllArticles()
      } catch (err) {
        console.error("Failed to load articles", err)
      } finally {
        isLoading = false
      }
    })
</script>

<section class="overflow-y-scroll">
    {#if isLoading}
        <p>Loading articles....</p>
    {:else if error}
        <p class="text-red-500">{error}</p>
    {:else}
        <div class="w-full grid grid-cols-12 justify-items-start mt-1">
            {#each articles as article (article.id)}
                <article class="w-full col-span-6 flex flex-col h-[50vh] px-4 fantasy-border"
                    style="--inlay-bg: url('{coverImage}') center/cover; --inlay-filter: blur(1px) brightness(1);">
                    <div class="grid grid-cols-6">
                        <h4 class="col-span-4 text-start">{article.title}</h4>
                        <p class="col-span-2 italic">{article.category}</p>
                    </div>
                    <div class="horizontal-hr" style="--scale: 1;"></div>
                    <div class="mt-1 p-1 text-start flex-1 line-clamp-6 overflow-hidden">
                        {@html article.description}
                    </div>
                    {#if article.tags != ''}
                        <div class="grid grid-cols-1 justify-start">
                            <p class="opacity-30 text-gray-500">{article.tags}</p>
                        </div>
                    {/if}
                    <div class="col-span-6 justify-center mb-2">
                        <button class="fantasy-btn-md fantasy-bone-n-coper">Edit</button>
                    </div>
                </article>
            {/each}
        </div>
    {/if}
</section>
