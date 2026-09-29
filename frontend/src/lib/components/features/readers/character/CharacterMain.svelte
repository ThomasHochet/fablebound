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
    import BadgeBG from "$lib/../assets/images/badge-bg.png"
    import CornerGrey from "$lib/../assets/images/corners_grey.png"
    import Mountain from "$lib/../assets/images/moutain_bg.png"

    let { character } = $props()
    let currentBackstoryIndex = $state(0)
    let currentRelPageIndex = $state(0)
    let direction = $state(1) // 1: right, -1: left
    let relDirection = $state(1)

    const ITEMS_PER_PAGE = 3 // Number of cards per slide

    let totalRelPages = $derived(
      character?.relationships ? Math.ceil(character.relationships.length / ITEMS_PER_PAGE) : 0
    )

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

    function nextRelationships() {
      if (totalRelPages <= 1) return
      relDirection = 1
      currentRelPageIndex = (currentRelPageIndex + 1) % totalRelPages
    }

    function prevRelationships() {
      if (totalRelPages <= 1) return
      relDirection = -1
      currentRelPageIndex = (currentRelPageIndex - 1 + totalRelPages) % totalRelPages
    }
</script>



<section class="flex flex-col gap-6">

    <!-- Description and goals -->
    <div class="grid grid-cols-1 md:grid-cols-1 mt-4">
        <fieldset class="border border-[#8b7355]/40 rounded-md p-4 pt-0">
            <!-- Legend sits natively on top of the border and breaks it transparently -->
            <legend class="text-left flex flex-row gap-2.5 items-center text-lg px-2" >
                <div class="grid place-items-center relative shrink-0 w-8 h-8">
                    <img src={BadgeBG} alt="" class="col-start-1 row-start-1 w-full h-full object-contain pointer-events-none" />
                    <ChestArmor height="1.2rem" class="col-start-1 row-start-1 text-[#2c251b] z-10" />
                </div>
                <span class="font-bold text-[#2c251b]">Physical Description</span>
            </legend>
            <div class="flex flex-col gap-1 text-left text-lg pt-1">
                {@html character.description}
            </div>
        </fieldset>
    </div>

    <div class="grid grid-cols-1">
        <fieldset class="border border-[#8b7355]/40 rounded-md p-4 pt-0">
            <legend class="text-left flex flex-row gap-2 items-center text-lg px-2">
                <div class="grid place-items-center relative shrink-0 w-8 h-8">
                    <img src={BadgeBG} alt="" class="col-start-1 row-start-1 w-full h-full object-contain pointer-events-none" />
                    <Bullseye height="1.2rem" class="col-start-1 row-start-1 text-[#2c251b] z-10" />
                </div>
                <span class="font-bold text-[#2c251b]">Personality & Goals</span>
            </legend>
            <div class="flex flex-col gap-1 text-left text-lg pt-1">
                {@html character.goals}
            </div>
        </fieldset>
    </div>


    <!-- Traits Grid -->
    <div class="grid grid-cols-2 sm:grid-cols-5 gap-3">
        <div class="p-2 flex flex-col gap-2 min-h-35 border border-[#8b7355]/40 rounded-sm shadow-2xl">
            <div class="flex flex-row items-center text-left text-base gap-2">
                <div class="grid place-items-center relative shrink-0 w-7 h-7 ml-1">
                    <img src={BadgeBG} alt="" class="col-start-1 row-start-1 w-full h-full object-contain pointer-events-none" />
                    <Pawn height="1.2rem" class="col-start-1 row-start-1 text-[#2c251b] z-10" />
                </div>
                <span class="font-bold text-[#2c251b]">Personality Traits</span>
            </div>
            <div class="border-b border-[#8b7355]/40 pointer-events-none z-0 w-full"></div>
            <ul class="text-base text-left">
                {#if character.personality_traits !== undefined}
                    {#each character.personality_traits.slice(0, 5) as item(item.id)}
                        <li class="flex flex-row items-center gap-2.5"><DividedSquare height="0.7em" />  {item.label}</li>
                    {/each}
                {/if}
            </ul>
        </div>
        <div class="p-2 flex flex-col gap-2 min-h-35 border border-[#8b7355]/40 rounded-sm shadow-xl">
            <div class="flex flex-row items-center text-left text-base gap-2">
                <div class="grid place-items-center relative shrink-0 w-7 h-7 ml-1">
                    <img src={BadgeBG} alt="" class="col-start-1 row-start-1 w-full h-full object-contain pointer-events-none" />
                    <HighFive height="1.2rem" class="col-start-1 row-start-1 text-[#2c251b] z-10" />
                </div>
                <span class="font-bold text-[#2c251b]">Strengths</span>
            </div>
            <div class="border-b border-[#8b7355]/40 pointer-events-none z-0 w-full"></div>
            <ul class="text-base text-left">
                {#if character.strengths !== undefined}
                    {#each character.strengths.slice(0, 5) as item(item.id)}
                        <li class="flex flex-row items-center gap-2.5"><DividedSquare height="0.7em" />  {item.label}</li>
                    {/each}
                {/if}
            </ul>
        </div>
        <div class="p-2 flex flex-col gap-2 min-h-35 border border-[#8b7355]/40 rounded-sm shadow-xl">
            <div class="flex flex-row items-center text-left text-base gap-2">
                <div class="grid place-items-center relative shrink-0 w-7 h-7 ml-1">
                    <img src={BadgeBG} alt="" class="col-start-1 row-start-1 w-full h-full object-contain pointer-events-none" />
                    <BrokenBone height="1.2rem" class="col-start-1 row-start-1 text-[#2c251b] z-10" />
                </div>
                <span class="font-bold text-[#2c251b]">Flaws</span>
            </div>
            <div class="border-b border-[#8b7355]/40 pointer-events-none z-0 w-full"></div>
            <ul class="text-lg text-left">
                {#if character.flaws !== undefined}
                    {#each character.flaws.slice(0, 5) as item(item.id)}
                        <li class="flex flex-row items-center gap-2.5"><DividedSquare height="0.7em" />  {item.label}</li>
                    {/each}
                {/if}
            </ul>
        </div>
        <div class="p-2 flex flex-col gap-2 min-h-35 border border-[#8b7355]/40 rounded-sm shadow-xl">
            <div class="flex flex-row items-center text-left text-base gap-2">
                <div class="grid place-items-center relative shrink-0 w-7 h-7 ml-1">
                    <img src={BadgeBG} alt="" class="col-start-1 row-start-1 w-full h-full object-contain pointer-events-none" />
                    <BreakingChain height="1.1rem" class="col-start-1 row-start-1 text-[#2c251b] z-10" />
                </div>
                <span class="font-bold text-[#2c251b]">Weaknesses</span>
            </div>
            <div class="border-b border-[#8b7355]/40 pointer-events-none z-0 w-full"></div>

            <ul class="text-lg text-left">
                {#if character.weaknesses !== undefined}
                    {#each character.weaknesses.slice(0, 5) as item(item.id)}
                        <li class="flex flex-row items-center gap-2.5"><DividedSquare height="0.7em" />  {item.label}</li>
                    {/each}
                {/if}
            </ul>
        </div>

        <div class="p-2 flex flex-col gap-2 min-h-35 border border-[#8b7355]/40 rounded-sm shadow-xl">
            <div class="flex flex-row items-center text-left text-base gap-2">
                <div class="grid place-items-center relative shrink-0 w-7 h-7 ml-1">
                    <img src={BadgeBG} alt="" class="col-start-1 row-start-1 w-full h-full object-contain pointer-events-none" />
                    <Eyeball height="1.1rem" class="col-start-1 row-start-1 text-[#2c251b] z-10" />
                </div>
                <span class="font-bold text-[#2c251b]">Fears</span>
            </div>
            <div class="border-b border-[#8b7355]/40 pointer-events-none z-0 w-full"></div>
            <ul class="text-lg text-left">
                {#if character.fears !== undefined}
                    {#each character.fears.slice(0, 5) as item(item.id)}
                        <li class="flex flex-row items-center gap-2.5"><DividedSquare height="0.7em" />  {item.label}</li>
                    {/each}
                {/if}
            </ul>
        </div>
    </div>

    {console.log(character.relationships)}

    <!-- Relationships -->
    <div class="grid grid-cols-1">
        <fieldset class="border border-[#8b7355]/40 rounded-md p-4 pt-0">
            <legend class="text-left flex flex-row items-center gap-2 text-lg px-2">
                <div class="grid place-items-center relative shrink-0 w-8 h-8">
                    <img src={BadgeBG} alt="" class="col-start-1 row-start-1 w-full h-full object-contain pointer-events-none" />
                    <Heart height="1.2rem" class="col-start-1 row-start-1 text-[#2c251b] z-10" />
                </div>
                <span class="font-bold text-[#2c251b]">Relationships</span>
            </legend>

            {#if character?.relationships && character.relationships.length > 0}
                <div class="flex flex-col gap-3 pt-1">
                    <!-- Stacked grid container for fly transitions -->
                    <div class="grid grid-cols-1 grid-rows-1 overflow-hidden min-h-62">
                        {#each Array(totalRelPages) as _, pageIndex (pageIndex)}
                            {#if pageIndex === currentRelPageIndex}
                                <div
                                    class="col-start-1 row-start-1 w-full grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-4"
                                    in:fly={{ x: relDirection * 200, duration: 250, delay: 100 }}
                                    out:fly={{ x: relDirection * -200, duration: 250 }}
                                >
                                    {#each character.relationships.slice(pageIndex * ITEMS_PER_PAGE, (pageIndex + 1) * ITEMS_PER_PAGE) as relationship (relationship.id)}
                                        <!-- Card Item -->
                                        <div class="relative z-10 w-full max-w-2xl parchment-background-inner font-serif rounded-lg px-4 py-3 overflow-hidden select-none shadow-4xl min-h-45 max-h-60">
                                            <div class="absolute inset-1 border border-[#a8977d]/60 rounded-md pointer-events-none"></div>

                                            <img src={CornerGrey} alt="" class="absolute -top-2 -left-2 w-9 h-9 pointer-events-none z-10" />
                                            <img src={CornerGrey} alt="" class="absolute -top-2 -right-2 w-9 h-9 pointer-events-none z-10 scale-x-[-1]" />
                                            <img src={CornerGrey} alt="" class="absolute -bottom-2 -left-2 w-9 h-9 pointer-events-none z-10 scale-y-[-1]" />
                                            <img src={CornerGrey} alt="" class="absolute -bottom-2 -right-2 w-9 h-9 pointer-events-none z-10 rotate-180" />

                                            <div class="relative z-10 grid grid-cols-1 grid-rows-[auto_1fr] h-full">
                                                <div class="flex items-center justify-between pb-2 border-b border-[#a8977d]/70">
                                                    <div class="flex items-center gap-2.5">
                                                        <h2 class="text-base font-bold tracking-wider text-[#1c120c] font-serif text-left">
                                                            {relationship.target_character.firstname + " " + relationship.target_character.surname}
                                                        </h2>
                                                    </div>
                                                    <span class="italic text-base text-[#2c251b]/60">
                                                        {relationship.type}
                                                    </span>
                                                </div>
                                                <div class="pt-1 pb-2">
                                                    <p class="text-[#3b2a1e] text-left text-lg font-serif tracking-wide line-clamp-4">
                                                        {@html relationship.notes}
                                                    </p>
                                                </div>
                                            </div>
                                        </div>
                                    {/each}
                                </div>
                            {/if}
                        {/each}
                    </div>

                    {#if totalRelPages > 1}
                        <div class="flex flex-col items-center w-full">
                            <div class="mt-2 border-b border-[#8b7355]/40 pointer-events-none z-0 w-4/5"></div>

                            <div class="flex flex-row justify-between w-full -mb-4">
                                <!-- Previous Button -->
                                <Rafter direction="right" class="w-18" handleClick={prevRelationships} />

                                <!-- Next Button -->
                                <Rafter direction="left" class="w-18" handleClick={nextRelationships} />
                            </div>
                        </div>
                    {/if}
                </div>
            {:else}
                <span class="my-4 text-lg text-left block">
                    This character does not seem to have any relationships.
                </span>
            {/if}
        </fieldset>
    </div>



    <!-- Lore & Backstory -->
    <div class="grid grid-cols-1">
        <fieldset class="border border-[#8b7355]/40 rounded-md p-4 pt-0">
            <legend class="flex flex-row items-center gap-2 p-2 text-left text-lg px-2">
                <div class="grid place-items-center relative shrink-0 w-8 h-8">
                    <img src={BadgeBG} alt="" class="col-start-1 row-start-1 w-full h-full object-contain pointer-events-none" />
                    <ScrollUnfurled height="1.2rem" class="col-start-1 row-start-1 text-[#2c251b] z-10" />
                </div>
                <span class="font-bold">Lore & Backstories</span>
            </legend>
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
        </fieldset>
    </div>

    <img src={Mountain} alt="" class="opacity-20 absolute bottom-0 right-20 mb-2 mr-2" />

</section>
