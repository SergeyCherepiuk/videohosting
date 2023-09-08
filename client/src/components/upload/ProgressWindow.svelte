<script lang="ts">
    import type { MouseEventHandler } from "svelte/elements";
    import Header from "./Header.svelte";
    import ProgressBar from "./ProgressBar.svelte";

    export let fileName: string
    export let fileSize: string
    export let progress: number
    export let onClose: MouseEventHandler<HTMLButtonElement>
</script>

<div class="flex flex-col p-10 rounded-2xl bg-white shadow-md w-full sm:w-3/4 xl:w-[960px]">
    <div class="mb-8">
        <Header text="Upload new video" onClick={onClose} />
    </div>
    <span class="font-amazon-ember text-xl">{fileName}</span>
    <div class="flex flex-row justify-between w-full">
        <span class="font-amazon-ember text-md text-gray-50">{fileSize}</span>
        <span class="font-amazon-ember text-md text-gray-50">{progress}%</span>
    </div>
    <ProgressBar {progress} />
</div>