<script>
    let content = $props();
</script>

<article class="grid grid-cols-12 p-4 gap-4">
    <div class="grid content-end">
        <button><i class="ra-scroll-unfurled"></i>Edit</button>
    </div>
    <h1>{content.title}</h1>
    <details>{@html content.description}</details>
</article>
