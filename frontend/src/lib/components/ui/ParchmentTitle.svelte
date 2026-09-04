<script>
    import banner from '$lib/../assets/images/title_banner.png'
    let { title = "Archive", subtitle = "" } = $props()
</script>

<div class="relative w-full max-w-3xl mx-auto flex items-center justify-center select-none pointer-events-none">
    <img
        src={banner}
        alt="Scroll banner"
        class="w-full h-auto object-contain filter drop-shadow-[0_12px_12px_rgba(0,0,0,0.65)]"
    />
    <div class="absolute inset-x-0 top-1/2 -translate-y-1/2 text-center px-16 flex flex-col items-center justify-center">
        <h1 class="font-bold text-lg text-(--wood-dark) tracking-widest uppercase drop-shadow-2xl truncate w-full">{title}</h1>
        {#if subtitle}
            <p class="text-xs sm:text-sm italic font-serif text-(--ink-color) opacity-85 -mt-1 truncate w-full z-5">
                {subtitle}
            </p>
        {/if}
    </div>
</div>
