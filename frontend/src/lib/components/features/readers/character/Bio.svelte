<script lang="ts">
    let { character, fullname } = $props()
</script>

<aside class="flex flex-col gap-3">
    <div class="flex flex-col">
        <div class="py-2.5 px-3 fantasy-border fantasy-border-brown">
            <img src={character.portrait} alt="Portrait" class="h-full overflow-hidden max-h-80 rounded-lg" />
        </div>

        <span class="mt-1 font-bold text-base">{fullname}</span>
        <span class="italic font-light text-base">{character.nickname}</span>
        <hr class="my-1" />
    </div>
    <div class="flex flex-col gap-1">
        <div class="flex flex-col">
            <span class="font-bold text-sm">Age</span>
            <span class="text-sm">{character.age}</span>
        </div>
        <div class="flex flex-col">
            <span class="font-bold text-sm">Gender</span>
            <span class="text-sm">{character.gender.label}</span>
        </div>
        <div class="flex flex-col">
            <span class="font-bold text-sm">Race</span>
            <span class="text-sm">{character.race.label}</span>
        </div>
        <div class="flex flex-col">
            <span class="font-bold text-sm">Occupation</span>
            <span class="text-sm">{character.occupation.label}</span>
        </div>
        <div class="flex flex-col">
            <span class="font-bold text-sm">Alignment</span>
            <span class="text-sm">{character.alignment.label}</span>
        </div>
        <div class="flex flex-col">
            <span class="font-bold text-sm">Status</span>
            <span class="text-sm">{character.status.label}</span>
        </div>

        <hr class="my-1" />

        <div class="flex flex-col">
            <span class="font-bold text-sm">Affiliation</span>
            <span class="text-sm">{character.affiliation}</span>
        </div>
        <div class="flex flex-col">
            <span class="font-bold text-sm">Faction</span>
            <span class="text-sm">{character.faction.name}</span>
        </div>
        <div class="flex flex-col">
            <span class="font-bold text-sm">Birthplace</span>
            <span class="text-sm">{character.birthplace.name}</span>
        </div>
        <div class="flex flex-col">
            <span class="font-bold text-sm">Current Location</span>
            <span class="text-sm">{character.current_location.name}</span>
        </div>
    </div>
</aside>
