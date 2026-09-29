<script lang="ts">
    import ParchmentTitle from "$lib/components/ui/ParchmentTitle.svelte";
    import corners from "$lib/../assets/images/corners.png"
    import RichEditor from "$lib/components/ui/RichEditor.svelte";
    import { logError } from "$lib/logger";
    import { FetchCharacter } from "$wails/world-builder/app";
    import { Character } from "$wails/world-builder/internal/models/models";
    import Bio from "./character/Bio.svelte";
    import CharacterMain from "./character/CharacterMain.svelte";
    import Compass from "$lib/../assets/images/compass_bg.png"

    let { id } = $props()
    let character = $state<Character | null>(null)
    let fullname = $derived.by(() => {
      if (character?.middlename !== undefined) {
        return character.firstname + " " + character.middlename + " " + character.surname
      }
      return character?.firstname + " " + character?.surname
    })

    async function getCharacter(id: any) {
      const numId = Number(id)
      if (!isNaN(numId)) {
        try {
          character = await FetchCharacter(numId)
          if(character?.portrait) {
            if(!character.portrait.startsWith('data:image')){
              character.portrait = `data:image/png;base64,${character.portrait}`
            }
          }
        } catch(err) {
          logError(`Failed to fetch article ${numId}`, err)
        }
      }
    }

    $effect(() => {
      getCharacter(id)
    })
</script>

{#if character}
    <article class="h-screen grid grid-cols-1 justify-items-center items-start px-2 overflow-y-auto">
        <div class="w-full h-auto max-w-7xl grid grid-cols-1 place-items-center">
            <div class="relative z-20 -mb-10 pointer-events-none">
                <ParchmentTitle title={fullname} subtitle={character.nickname} />
            </div>
            <div class="relative w-full h-auto shadow-2xl rounded-sm border-4 border-[#1a0f0f] bg-[#2b190c] p-1 md:p-1 grid grid-cols-1 grid-rows-1 min-h-56">
                <div class="wood-texture-bg"></div>

                <img src="{corners}" alt="" class="absolute -top-4 -left-4 w-15 h-15 pointer-events-none z-10">
                <img src="{corners}" alt="" class="absolute -top-4 -right-4 w-15 h-15 pointer-events-none z-10 scale-x-[-1]" />
                <img src="{corners}" alt="" class="absolute -bottom-4 -left-4 w-15 h-15 pointer-events-none z-10 scale-y-[-1]" />
                <img src="{corners}" alt="" class="absolute -bottom-4 -right-4 w-15 h-15 pointer-events-none z-10 rotate-180" />

                <div class="parchment-background relative w-full h-auto grid grid-rows-[auto_1fr] z-10 border border-[#3b2a1e]">
                    <div class="absolute inset-2 border border-[#8b7355]/40 pointer-events-none z-0"></div>
                    <div class="grid grid-cols-1 md:grid-cols-[260px_auto_1fr] gap-6 p-4">
                        <Bio character={character} fullname={fullname} />
                        <div class="mt-2 border-l border-[#8b7355]/40 pointer-events-none z-0 w-4/5"></div>
                        <CharacterMain {character} />
                    </div>
                </div>
                <img src={Compass} alt="" class="opacity-20 absolute right-5 z-11 w-50 rotate-40" />
            </div>
        </div>
</article>
{:else}
    <p>Loading Character article</p>
{/if}
