<script lang='ts'>
    import { TrixEditor } from 'svelte-trix';
    import {CreateArticle} from '$wails/go/main/App.js'

    let { articleId = null,  initialTitle = '', initialDescription = '' } = $props();

    let title = $state(initialTitle)
    let description = $state(initialDescription);

    $effect(() => {
      title = initialTitle
      description = initialDescription
    })

    const handleChange = (html: string) => {
      description = html;
    }

    function handleSave() {
      console.log(title, description)
      CreateArticle(title, description)
    }
</script>

<article class="col-span-12">
    <input autocomplete="off" bind:value={title} class="input" id="title" type="text" placeholder="Title">
    <TrixEditor
        value="Time to write your adventures..."
        onChange={handleChange}
    />
    <button class="fantasy-btn-2xl fantasy-bone-n-coper" onclick={handleSave}>Save</button>
</article>
