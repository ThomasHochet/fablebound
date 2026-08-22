<script lang='ts'>
    import { TrixEditor } from 'svelte-trix';
    import { CreateArticle } from '$wails/world-builder/app.js';

    let { section } = $props();

    // let title = $state(initialTitle)
    // let description = $state(initialDescription);

    // console.log("EditorWindow")

    // $effect(() => {
    //   title = initialTitle
    //   description = initialDescription
    // })

    const handleChange = (html: string) => {
      console.log(html)
    }

    // function handleSave() {
    //   console.log(title, description)
    //   CreateArticle(title, description)
    // }
</script>

<article class="w-full grid grid-cols-12 justify-center fantasy-border ">
    <input autocomplete="off" bind:value={section} class="input" id="title" type="text" placeholder="Title">
    <div class="w-full col-span-12">
        <TrixEditor
            value="Time to write your adventures..."
            onChange={handleChange}
        />
    </div>
    <button class="fantasy-btn-2xl fantasy-bone-n-coper">Save</button>
</article>
