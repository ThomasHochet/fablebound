<script module lang="ts">
    export type SaveStatus = 'idle' | 'unsaved' | 'saving' | 'saved' | 'error';
</script>
<script lang="ts">
    import { fade, scale } from 'svelte/transition';
    import { cubicInOut } from 'svelte/easing';
    import corners from "$lib/../assets/images/gCorner.png"
    import circle from "$lib/../assets/images/gCircle.png"
    import Feather from "@iconify-svelte/game-icons/components/f/feather.svelte"
    import CheckMark from "@iconify-svelte/game-icons/components/c/check-mark.svelte"
    import OpenBook from "@iconify-svelte/game-icons/components/o/open-book.svelte"
    import DeathSkull from "@iconify-svelte/game-icons/components/d/death-skull.svelte"


    let {
        status = 'idle',
        title = "STORY INSCRIBED",
        subtitle = "Saved to the Database Chronicles."
    }: {
        status?: SaveStatus;
        title?: string;
        subtitle?: string;
    } = $props();

    function cardFlip(node: HTMLElement, { duration = 250, delay = 0}) {
      return {
        duration,
        delay,
        easing: cubicInOut,
        css: (t: number, u: number) => `
          transform: rotateY(${u * 90}deg);
          backface-visibility: hidden;
        `
      }
    }
</script>

{#if status === 'saved' || status === 'saving' || status === 'error'}
    <div
        in:scale={{ duration: 500, start: 0.9 }}
        out:fade={{ duration: 450 }}
        class="fixed top-1/3 left-1/2 -translate-x-1/2 z-50 flex flex-col items-center pointer-events-none drop-shadow-[0_20px_25px_rgba(0,0,0,0.9)]"
    >
        <!-- Top Emblem Crest -->
        <!-- <div class="relative -mb-7 z-100 flex items-center justify-center w-16 h-16 rounded-full bg-[#2b190c] border-2 border-[#b7a48d] shadow-md"> -->
        <div class="relative w-25 h-25 -mb-10 z-100 flex items-center justify-center">
            <img src="{circle}" alt="" class="absolute inset-0 w-full h-full object-contain drop-shadow-[0_8px_6px_rgba(0,0,0,0.8)] z-100 [perspective:600px]" />
            {#if status === 'saving'}
                <div
                    in:cardFlip={{ duration: 200, delay: 200}}
                    out:cardFlip={{ duration: 200 }}
                    class=""
                >
                    <Feather class="feather"  />
                </div>
            {:else if status === 'saved'}
                <div
                    in:cardFlip={{ duration: 200, delay: 200}}
                    out:cardFlip={{ duration: 200 }}
                    class=""
                >
                    <OpenBook class="open-book" />
                    <CheckMark class="check-mark" />
                </div>
            {:else if status === 'error'}
            <div
                in:cardFlip={{ duration: 200, delay: 200}}
                out:cardFlip={{ duration: 200 }}
                class=""
            >
                <DeathSkull class="death-skull" />
            </div>
            {/if}
        </div>

        <div class="relative min-w-[20rem] p-2 border-[#1a0f0f] bg-[#2b190c] rounded-lg text-center font-serif shadow-inner">
            <div class="wood-texture-bg absolute inset-0 pointer-events-none z-0"></div>

            <img src="{corners}" alt="" class="absolute -top-2 -left-2.5 w-8 h-8 pointer-events-none -z-10">
            <img src="{corners}" alt="" class="absolute -top-2 -right-2.5 w-8 h-8 pointer-events-none -z-10 scale-x-[-1]" />
            <img src="{corners}" alt="" class="absolute -bottom-2 -left-2.5 w-8 h-8 pointer-events-none -z-10 scale-y-[-1]" />
            <img src="{corners}" alt="" class="absolute -bottom-2 -right-2.5 w-8 h-8 pointer-events-none -z-10 rotate-180" />

            <div class="relative parchment-background-inner z-50 w-full h-full">
                <div class="absolute inset-1 border border-[#b7a48d]/40 rounded pointer-events-none"></div>
                {#if status === 'saving'}
                    <h3 class="mt-5 text-base font-bold tracking-widest text-[#8c7661] uppercase">
                        Scribing History...
                    </h3>
                    <p class="mt-1 mb-5 text-xs text-[#523d2b]/80 italic">
                        The chronicler is at work.
                    </p>
                {:else if status === 'saved'}
                    <h3 class="mt-5 text-lg font-bold tracking-widest text-[#1c5028] uppercase drop-shadow-sm">
                        {title}
                    </h3>
                    <p class="mt-0.5 mb-5 text-xs text-[#523d2b] font-serif">
                        {subtitle}
                    </p>
                {:else if status === 'error'}
                <h3 class="mt-5 text-lg font-bold tracking-widest text-red-900 uppercase drop-shadow-sm">
                    Void connection
                </h3>
                <p class="mt-0.5 mb-5 text-xs text-[#523d2b] font-serif">
                    The database chronicles connection seems erratic.
                </p>
                {/if}
            </div>


        </div>
    </div>
{/if}
