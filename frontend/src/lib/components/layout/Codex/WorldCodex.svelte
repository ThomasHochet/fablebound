<script lang="ts">
    import Modal from "$lib/components/features/Modal.svelte";
    import Card from "$lib/components/ui/Card.svelte";
    import { subscribe } from "$lib/functions/subscribe";
    import { DeleteArticle, FetchAllArticles, OpenEditorWindow } from "$wails/world-builder/app";
    import { Article } from "$wails/world-builder/internal/models/models";
    import { onMount } from "svelte";
    import corners from "$lib/../assets/images/corners.png"
    import QuillInk from "@iconify-svelte/game-icons/components/q/quill-ink.svelte";
    import CrossMark from "@iconify-svelte/game-icons/components/c/cross-mark.svelte";

    let isLoading = $state(true)
    let articles = $state<Article[]>([])
    let isModalOpen = $state(false)
    let articleRdyToDel = $state<number>()
    let readyToDeleteData = $derived(articles.find(a => (a.id === articleRdyToDel)))

    async function fetchData() {
      try {
        articles = await FetchAllArticles()
      } catch(err) {
        console.error(err)
      } finally {
        isLoading = false
      }
    }

    function handleEdit(id: number) {
      OpenEditorWindow(`world/form/${id}`).catch(err => {
        console.error("Failed to open window", err)
      })
    }

    function handleDelete(id: number | undefined) {
      const numId = Number(id)
      if (!isNaN(numId)) {
        DeleteArticle(numId)
        isModalOpen = false
      }
    }

    onMount(() => {
      fetchData()

      subscribe('articles', fetchData)
    })
</script>

{#each articles as article(article.id)}
    <Card title={article.title} category={article.category} description={article.description} variant={article.category} onEdit={() => handleEdit(article.id)}
        onDelete={() => {
          (isModalOpen = true)
          articleRdyToDel = article.id
        }}
    />
{/each}

<Modal bind:open={isModalOpen} >
    <div class="border-[#1a0f0f] bg-[#2b190c] p-2 md:p-2 w-[25rem]! h-[22rem]! rounded-md!">
        <div class="wood-texture-bg absolute inset-0 pointer-events-none z-0"></div>

        <img src="{corners}" alt="" class="absolute -top-3 -left-3 w-15 h-15 pointer-events-none z-10">
        <img src="{corners}" alt="" class="absolute -top-3 -right-3 w-15 h-15 pointer-events-none z-10 scale-x-[-1]" />
        <img src="{corners}" alt="" class="absolute -bottom-3 -left-3 w-15 h-15 pointer-events-none z-10 scale-y-[-1]" />
        <img src="{corners}" alt="" class="absolute -bottom-3 -right-3 w-15 h-15 pointer-events-none z-10 rotate-180" />

        <div class="flex flex-col gap-2 parchment-background-inner relative w-full h-full z-10 border border-[#3b2a1e] p-4 rounded-md">
            <span class="mt-4 font-bold">Shall this tale be lost to oblivion?</span>
            <span class="underline mt-2">
                [{readyToDeleteData?.title}]
            </span>
            <span class="italic">Category: {readyToDeleteData?.category}</span>
            <span>Codex : World</span>
            <div class="flex flex-row justify-center px-6 mt-auto gap-4">
                <button class="forge-btn forge-btn-base" onclick={() => {
                  handleDelete(articleRdyToDel)
                }}>
                    <CrossMark height="1.5rem" />
                    <span class="ml-2">Agreed</span>
                </button>
                <button
                    class="forge-btn forge-btn-base"
                    onclick={() => {
                      isModalOpen = false
                      articleRdyToDel = undefined
                    }}
                >
                    <QuillInk height="1.5rem" />
                    <span class="ml-1">Keep</span>
                </button>
            </div>
        </div>

    </div>
</Modal>
