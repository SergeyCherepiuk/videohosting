<script lang="ts">
    import type { MouseEventHandler } from "svelte/elements";

    export let text: string
    export let onClick: MouseEventHandler<HTMLButtonElement>
</script>

<div class="flex flex-row w-full justify-between items-center">
    <span class="text-3xl font-amazon-ember">{text}</span>
    <button on:click={onClick}><span class="text-4xl">&times;</span></button>
</div>