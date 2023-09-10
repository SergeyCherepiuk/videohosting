<script lang="ts">
    import type { MouseEventHandler } from "svelte/elements";    
    
    export let text: string
    export let isPrimary: boolean
    export let onClick: MouseEventHandler<HTMLButtonElement>
</script>

<button class="text-lg { isPrimary ? "text-white" : "text-black"} font-amazon-ember
    rounded-lg px-4 py-2 { isPrimary ? "bg-blue-50" : "bg-gray-70" } w-min"
    on:click={onClick}>
    <p class="truncate">{text}</p>
</button>