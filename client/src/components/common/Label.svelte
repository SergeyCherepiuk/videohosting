<script lang="ts">
    export let label: string
</script>

<div class="flex flex-col gap-2">
    <span class="font-amazon-ember text-lg ps-4">{label}</span>
    <slot />
</div>