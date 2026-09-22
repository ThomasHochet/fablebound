<script lang="ts">
    import CornerGrey from "$lib/../assets/images/corners_grey.png"
    import QuillInk from "@iconify-svelte/game-icons/components/q/quill-ink.svelte";
    import CrossMark from "@iconify-svelte/game-icons/components/c/cross-mark.svelte";

    // Icon for each types
    import LocationMarker from "@iconify-svelte/game-icons/components/v/virtual-marker.svelte";
    import ElvenCastle from "@iconify-svelte/game-icons/components/e/elven-castle.svelte";
    import EarthSpit from "@iconify-svelte/game-icons/components/e/earth-spit.svelte"
    import MedievalVillage from "@iconify-svelte/game-icons/components/m/medieval-village-01.svelte"

    import WorldIcon from "@iconify-svelte/game-icons/components/w/world.svelte";

    import FactionBanner from "@iconify-svelte/game-icons/components/v/vertical-banner.svelte";
    import AffiliationEdgedShield from "@iconify-svelte/game-icons/components/e/edged-shield.svelte";

    import BookCover from "@iconify-svelte/game-icons/components/b/book-cover.svelte";
    import ScrollUnfurled from "@iconify-svelte/game-icons/components/s/scroll-unfurled.svelte";

    import LightHelm from "@iconify-svelte/game-icons/components/l/light-helm.svelte"

    type CardVariant = 'general' | 'world' | 'character' | 'continent' | 'town' | 'castle' | 'miscellaneous' | string | undefined;

    let {
      title = "Unnamed entry",
      category = "General",
      description = "The content has been sent to the void...",
      variant = "general",
      onEdit,
      onDelete
    }: {
      title: string,
      category: string | undefined,
      description: string | undefined,
      variant?: CardVariant,
      onEdit: (() => void) | null,
      onDelete: (() => void) | null
    } = $props()


    const themes = {
      general: {
        pill: "bg-[#5c656b] text-[#f2ede4] border border-[#485055]",
        icon: LocationMarker
      },
      lore: {
        pill: "bg-yellow-950 text-[#f2ede4]",
        icon: BookCover
      },
      miscellaneous: {
        pill: "bg-yellow-100 text-(--ink-color)",
        icon: ScrollUnfurled
      },
      world: {
        pill: "bg-blue-800 text-[#f2ede4]",
        icon: WorldIcon
      },
      character: {
        pill: "bg-(--wood-light) text-(--gold-text)",
        icon: LightHelm
      },
      castle: {
        pill: "bg-(--wood-dark) text-(--parchment--light)",
        icon: ElvenCastle
      },
      continent: {
        pill: "bg-taupe-900 text-(--gold-text)",
        icon: EarthSpit
      },
      town: {
        pill: "bg-green-950/80 text-(--parchment-light)",
        icon: MedievalVillage
      },
      faction: {
        pill: "bg-stone-800 text-(--parchment-light)",
        icon: FactionBanner
      },
      affiliation: {
        pill: "bg-emerald-950 text-(--gold-text)",
        icon: AffiliationEdgedShield
      }
    }

    let normalizedVariant = $derived(variant.toLowerCase() ?? "general")
    let currentTheme = $derived(
      themes[normalizedVariant as keyof typeof themes] ?? themes.general
    )
    let Icon = $derived(currentTheme.icon)


</script>

<div class="relative z-10 w-full max-w-2xl parchment-background-inner font-serif rounded-lg p-4 overflow-hidden select-none shadow-4xl min-h-60 max-h-60">
    <div class="absolute inset-1 border border-[#a8977d]/60 rounded-md pointer-events-none"></div>

    <img src="{CornerGrey}" alt="" class="absolute -top-2 -left-2 w-9 h-9 pointer-events-none z-10">
    <img src="{CornerGrey}" alt="" class="absolute -top-2 -right-2 w-9 h-9 pointer-events-none z-10 scale-x-[-1]" />
    <img src="{CornerGrey}" alt="" class="absolute -bottom-2 -left-2 w-9 h-9 pointer-events-none z-10 scale-y-[-1]" />
    <img src="{CornerGrey}" alt="" class="absolute -bottom-2 -right-2 w-9 h-9 pointer-events-none z-10 rotate-180" />

    <div class="relative z-10 px-2 pt-1 grid grid-cols-1 grid-rows-[auto_1fr_auto] h-full">
        <!-- Header Row -->
        <div class="flex items-center justify-between pb-3 border-b border-[#a8977d]/70">
          <div class="flex items-center gap-2.5">
            <!-- <svg class="w-7 h-7 text-[#231710]" fill="currentColor" viewBox="0 0 24 24">
              <path d="M12 2C8.13 2 5 5.13 5 9c0 5.25 7 13 7 13s7-7.75 7-13c0-3.87-3.13-7-7-7zm0 9.5c-1.38 0-2.5-1.12-2.5-2.5s1.12-2.5 2.5-2.5 2.5 1.12 2.5 2.5-1.12 2.5-2.5 2.5z"/>
            </svg> -->
            <Icon height="1.5rem" />
            <h2 class="text-[20px] font-extrabold tracking-wider uppercase text-[#1c120c] font-serif text-left">
              {title.toUpperCase()}
            </h2>
          </div>

          <!-- General Pill Badge -->
          <span class="{currentTheme.pill} ml-2 px-5 py-1 rounded-full text-xs font-sans font-bold tracking-widest uppercase shadow-sm">
            {category.toUpperCase()}
          </span>
        </div>

        <!-- Card Description -->
        <div class="pt-4 pb-14">
          <p class="text-[#3b2a1e] text-left text-lg font-serif tracking-wide line-clamp-3">
            {@html description}
          </p>
        </div>

        <!-- Bottom Action Buttons -->
        <div class="absolute -bottom-1 -right-1 flex items-center gap-2.5">
          <!-- Quill / Edit Button -->
          <button type="button" class="forge-btn forge-btn-xs" onclick={onEdit}>
            <QuillInk height="1.5rem" />
          </button>

          <!-- Skull / Danger Action Button -->
          <button type="button" class="forge-btn forge-btn-xs" onclick={onDelete}>
            <CrossMark height="1.5rem" />
          </button>
        </div>
    </div>

</div>
