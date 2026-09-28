<script lang="ts">
    import ChestArmor from "@iconify-svelte/game-icons/components/c/chest-armor.svelte"
    import Bullseye from "@iconify-svelte/game-icons/components/b/bullseye.svelte"
    import HighFive from "@iconify-svelte/game-icons/components/h/high-five.svelte"
    import BrokenBone from "@iconify-svelte/game-icons/components/b/broken-bone.svelte"
    import BreakingChain from "@iconify-svelte/game-icons/components/b/breaking-chain.svelte"
    import Eyeball from "@iconify-svelte/game-icons/components/e/eyeball.svelte"
    import Heart from "@iconify-svelte/game-icons/components/h/hearts.svelte"
    import ScrollUnfurled from "@iconify-svelte/game-icons/components/s/scroll-unfurled.svelte";
    import Pawn from "@iconify-svelte/game-icons/components/p/pawn.svelte"

    import DividedSquare from "@iconify-svelte/game-icons/components/d/divided-square.svelte"
    import RichEditor from "$lib/components/ui/RichEditor.svelte";
    import Rafter from "$lib/components/ui/Rafter.svelte";
    import { fly } from 'svelte/transition';

    let { character } = $props()
    let currentBackstoryIndex = $state(0)
    let direction = $state(1) // 1: right, -1: left

    function nextBackstory() {
      if (!character?.backstories?.length) return
      direction = 1
      currentBackstoryIndex = (currentBackstoryIndex + 1) % character.backstories.length
    }

    function prevBackstory() {
      if (!character?.backstories?.length) return
      direction = -1
      currentBackstoryIndex = (currentBackstoryIndex - 1 + character.backstories.length) % character.backstories.length
    }
</script>

<section class="flex flex-col gap-6">

    <!-- Description and goals -->
    <div class="shadow-wrapper">
        <div class="grid grid-cols-1 md:grid-cols-1 corners corners--chamfer-bordered">
            <div class="flex flex-col gap-2">
                <div class="flex flex-row gap-2 text-left items-center text-lg">
                    <ChestArmor height="1.7rem" />
                    <span class="font-bold">Physical Description</span>
                </div>
                <div class="flex flex-col gap-1 text-left text-lg">
                    {@html character.description}
                </div>
            </div>
        </div>
    </div>

    <div class="shadow-wrapper">
        <div class="grid grid-cols-1 md:grid-cols-1 corners corners--chamfer-bordered">
            <div class="flex flex-col gap-2">
                <div class="flex flex-row gap-2 text-left items-center text-lg">
                    <Bullseye height="1.7rem" />
                    <span class="font-bold">Personality & Goals</span>
                </div>
                <div class="flex flex-col gap-1 text-left text-lg">
                    {@html character.goals}
                </div>
            </div>
        </div>
    </div>


    <!-- Traits Grid -->
    <div class="grid grid-cols-2 sm:grid-cols-5 gap-3">
        <div class="shadow-wrapper">
            <div class="p-3 flex flex-col gap-2 corners corners--parchment-bordered light-wood-border min-h-25">
                <div class="flex flex-row items-center text-left text-lg gap-2">
                    <Pawn height="1.7rem" />
                    <span class="font-bold">Personality Traits</span>
                </div>
                <ul class="text-lg text-left">
                    {#if character.personality_traits !== undefined}
                        {#each character.personality_traits.slice(0, 5) as item(item.id)}
                            <li class="flex flex-row items-center gap-1"><DividedSquare height="0.7em" />  {item.label}</li>
                        {/each}
                    {/if}
                </ul>
            </div>
        </div>
        <div class="shadow-wrapper">
            <div class="p-3 flex flex-col gap-2 corners corners--parchment-bordered light-wood-border min-h-25">
                <div class="flex flex-row items-center text-left text-lg gap-2">
                    <HighFive height="1.7rem" />
                    <span class="font-bold">Strengths</span>
                </div>
                <ul class="text-lg text-left">
                    {#if character.strengths !== undefined}
                        {#each character.strengths.slice(0, 5) as item(item.id)}
                            <li class="flex flex-row items-center gap-1"><DividedSquare height="0.7em" />  {item.label}</li>
                        {/each}
                    {/if}
                </ul>
            </div>
        </div>
        <div class="shadow-wrapper">
            <div class="p-3 flex flex-col gap-2 corners corners--parchment-bordered light-wood-border min-h-25">
                <div class="flex flex-row items-center gap-2 text-left text-lg">
                    <BrokenBone height="1.7rem" />
                    <span class="font-bold">Flaws</span>
                </div>
                <ul class="text-lg text-left">
                    {#if character.flaws !== undefined}
                        {#each character.flaws.slice(0, 5) as item(item.id)}
                            <li class="flex flex-row items-center gap-1"><DividedSquare height="0.7em" />  {item.label}</li>
                        {/each}
                    {/if}
                </ul>
            </div>
        </div>
        <div class="shadow-wrapper">
            <div class="p-3 flex flex-col gap-2 corners corners--parchment-bordered light-wood-border min-h-25">
                <div class="flex flex-row items-center gap-2 text-left text-lg">
                    <BreakingChain height="1.7rem" />
                    <span class="font-bold">Weaknesses</span>
                </div>
                <ul class="text-lg text-left">
                    {#if character.weaknesses !== undefined}
                        {#each character.weaknesses.slice(0, 5) as item(item.id)}
                            <li class="flex flex-row items-center gap-1"><DividedSquare height="0.7em" />  {item.label}</li>
                        {/each}
                    {/if}
                </ul>
            </div>
        </div>

        <div class="shadow-wrapper">
            <div class="p-3 flex flex-col gap-2 corners corners--parchment-bordered light-wood-border min-h-25">
                <div class="flex flex-row items-center gap-2 text-left text-lg">
                    <Eyeball height="1.7rem" />
                    <span class="font-bold">Fears</span>
                </div>
                {#if character.fears !== undefined}
                    {#each character.fears.slice(0, 5) as item(item.id)}
                        <li class="flex flex-row items-center gap-1"><DividedSquare height="0.7em" />  {item.label}</li>
                    {/each}
                {/if}
            </div>
        </div>
    </div>

    {console.log(character)}

    <!-- Relationships -->
    <div class="shadow-wrapper">
        <div class="flex flex-col gap-3 corners corners--chamfer-bordered">
            <div class="flex flex-row items-center gap-2 text-left text-lg">
                <Heart height="1.7rem" />
                <span class="font-bold">Relationships</span>
            </div>
            <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-4">
                {#if character?.relationships }
                    {#each character.relationships.slice(0, 4) as relationship(relationship.id)}
                        <!-- Card Item - Each loop (max 4 visible) -->
                        <div class="shadow-wrapper shadow-wrapper-less">
                            <div class="p-2 flex flex-col gap-2 corners corners--parchment-bordered light-wood-border min-h-30 parchment-background-inner!">
                                <div class="flex flex-row text-left">
                                    <span>{relationship.target_char_id}</span>
                                    <span class="text-lg">{relationship.type}</span>
                                </div>


                                <span class="text-base text-center!">{relationship.notes}</span>
                            </div>
                        </div>
                    {/each}
                {:else}
                    <span class="my-4 text-lg col-span-1 sm:col-span-2 lg:col-span-4">This character does not seem to consider to have any relationship.</span>
                {/if}
            </div>
        </div>
    </div>


    <!-- Lore & Backstory -->
    <div class="shadow-wrapper">
        <div class="flex flex-col corners corners--chamfer-bordered">
            <div class="flex flex-row items-center gap-2 p-2 text-left text-lg">
                <ScrollUnfurled height="1.7rem" />
                <span class="font-bold">Lore & Backstories</span>
            </div>
            {#if character?.backstories?.length}
                <div class="flex flex-col gap-3 p-2">
                    <div class="grid grid-cols-1 grid-rows-1 overflow-hidden">
                    {#each character.backstories as backstory, index (backstory.id || index)}
                        {#if index === currentBackstoryIndex}
                            <div
                                class="col-start-1 row-start-1 w-full"
                                in:fly={{ x: direction * 200, duration: 250, delay: 100 }}
                                out:fly={{ x: direction * -200, duration: 250 }}
                            >
                                <RichEditor disabled={true} value={backstory.content} classes="text-left" />
                            </div>
                        {/if}
                    {/each}
                    </div>
                    <div class="flex flex-col items-center w-full">
                        <div class="mt-2 border-b border-[#8b7355]/40 pointer-events-none z-0 w-4/5"></div>

                        <div class="flex flex-row justify-between w-full -mb-4">
                            <!-- Previous Button (Points Left) -->
                            <Rafter direction="right" class="w-18" handleClick={prevBackstory} />

                            <!-- Next Button (Points Right) -->
                            <Rafter direction="left" class="w-18" handleClick={nextBackstory} />
                        </div>
                    </div>
                </div>
            {:else}
                <span class="my-4 text-lg col-span-1 sm:col-span-2 lg:col-span-4">This character lore hasn't been recorded.</span>
            {/if}
        </div>

        <!-- <div class="grid grid-cols-1 grid-rows-[1fr-auto] gap-3">
            {#if character?.backstories?.length}
                {@const currentBackstory = character.backstories[currentBackstoryIndex]}

                <RichEditor disabled={true} value={currentBackstory.content} classes="text-left" />

                <div class="mt-4 col-span-2 border-b border-[#8b7355]/40 pointer-events-none z-0 w-4/5"></div>
                <div class="flex flex-row justify-between -mt-3">
                    <Rafter direction="right" class="w-18" handleClick={() => prevBackstory()} />
                    <Rafter direction="left" class="w-18" handleClick={() => nextBackstory()} />
                </div>
            {/if}
        </div> -->
    </div>

</section>
