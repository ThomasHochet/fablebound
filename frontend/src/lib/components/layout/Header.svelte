<script lang="ts">
    import Map from '@iconify-svelte/game-icons/components/w/world.svelte'
    import DwarfHelmet from '@iconify-svelte/game-icons/components/d/dwarf-helmet.svelte'
    import Bookmarklet from '@iconify-svelte/game-icons/components/b/bookmarklet.svelte'
    import Castle from '@iconify-svelte/game-icons/components/e/elven-castle.svelte'
    import Banner from '@iconify-svelte/game-icons/components/v/vertical-banner.svelte'

    let { onSelect } = $props();

    let isMenuOpen = $state(false);

    function toggleMenu() { isMenuOpen = !isMenuOpen }

    function handleAction(actionName: string) {
      console.log(`Action clicked: ${actionName}`)
      isMenuOpen = false
    }
</script>

<header class="w-full bg-taupe-800 border-b border-taupe-600 text-taupe-100">
    <div class="max-w-7xl mx-auto px-4 py-3 grid grid-cols-12">
        <!-- Title? -->
        <div class="w-full flex items-center justify-center-safe space-x-3 col-span-full">
            <span class="font-bold text-xl tracking-wider text-[#e6c06f] underline underline-offset-2">World Builder</span>
        </div>
        <div class="mt-2 col-span-full gap-3 flex flex-wrap sm:justify-start justify-center">
            <button onclick={() => onSelect('world')} class="fantasy-btn-xl fantasy-bone-n-coper fantasy-b-gold px-[25px]!">
                <div class="flex items-center justify-center gap-0.5">
                    <Map height="1.25em" />
                    <span>World</span>
                </div>
            </button>
            <button onclick={() => onSelect('character')} class="fantasy-btn-xl fantasy-bone-n-coper fantasy-b-gold px-[25px]!">
                <div class="flex items-center justify-center gap-0.5">
                    <DwarfHelmet height="1.25em" />
                    <span>Characters</span>
                </div>
            </button>
            <button onclick={() => onSelect('lore')} class="fantasy-btn-xl fantasy-bone-n-coper fantasy-b-gold px-[25px]!">
                <div class="flex items-center justify-center gap-0 5">
                    <Bookmarklet height="1.25em" />
                    <span class="pl-1">Lore</span>
                </div>
            </button>
            <button onclick={() => onSelect('location')} class="fantasy-btn-xl fantasy-bone-n-coper fantasy-b-gold px-[25px]!">
                <div class="flex items-center justify-center gap-0 5">
                    <Castle height="1.25em" />
                    <span class="pl-1">City</span>
                </div>
            </button>
            <button onclick={() => onSelect('faction')} class="fantasy-btn-xl fantasy-bone-n-coper fantasy-b-gold px-[25px]!">
                <div class="flex items-center justify-center gap-0 5">
                    <Banner height="1.25em" />
                    <span class="pl-1">Faction</span>
                </div>
            </button>
        </div>
    </div>
</header>
