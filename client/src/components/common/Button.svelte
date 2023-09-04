<script lang="ts">
    import type { MouseEventHandler } from "svelte/elements";

    export let text: string
    export let onClick: MouseEventHandler<HTMLButtonElement>
</script>

<button
    class="
        text-lg text-white font-amazon-ember
        rounded-lg px-4 py-2 bg-blue-50 w-min"
    on:click={onClick}>
    {text}
</button>