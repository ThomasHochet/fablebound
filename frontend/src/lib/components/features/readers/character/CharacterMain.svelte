<script lang="ts">
    let { character } = $props()
</script>

<section class="flex flex-col gap-6">

    <!-- Description and goals -->
    <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
        <div class="flex flex-col gap-2 border p-3">
            <span class="font-bold text-base">Physical Description</span>
            <div class="flex flex-col gap-1 text-left">
                {@html character.description}
            </div>
        </div>

        <div class="flex flex-col gap-2 border p-3">
            <span class="font-bold text-base">Goals</span>
            <div class="flex flex-col gap-1 text-left">
                {@html character.goals}
            </div>
        </div>
    </div>

    <!-- Traits Grid -->
    <div class="grid grid-cols-2 sm:grid-cols-4 gap-3">
        <div class="border p-2 flex flex-col gap-2">
            <div>
                <span class="font-bold text-base">Strengths</span>
            </div>
            <ul>
                <li></li>
                <li></li>
            </ul>
        </div>
        <div class="border p-2 flex flex-col gap-2">
            <div>
                <span class="font-bold text-base">Flaws</span>
            </div>
            <ul>
                <li></li>
                <li></li>
            </ul>
        </div>
        <div class="border p-2 flex flex-col gap-2">
            <div>
                <span class="font-bold text-base">Weaknesses</span>
            </div>
            <ul>
                <li></li>
                <li></li>
            </ul>
        </div>
        <div class="border p-2 flex flex-col gap-2">
            <div>
                <span class="font-bold text-base">Fears</span>
            </div>
            <ul>
                <li></li>
                <li></li>
            </ul>
        </div>
    </div>

</section>
