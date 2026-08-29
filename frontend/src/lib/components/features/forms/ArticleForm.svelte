<script lang="ts">
    const handleChange = (html: string) => {
      console.log(html)
    }

    function handleSubmit(e: SubmitEvent) {
      e.preventDefault()


    }
</script>

<form onsubmit={handleSubmit}>
    <article class="w-full grid grid-cols-12 justify-center fantasy-border ">
        <input autocomplete="off" class="input" id="title" type="text" placeholder="Title">
        <div class="w-full col-span-12">
        </div>
        <button class="fantasy-btn-2xl fantasy-bone-n-coper">Save</button>
    </article>
</form>
