<script lang="ts">
    export let text: string
    export let video: FileList
    export let clearVideo: Function

    clearVideo = () => {
        let el: HTMLInputElement | null = document.getElementById("video") as HTMLInputElement
        if (el) { el.value = "" }
    }
</script>

<label for="video"
    class="text-lg text-white font-amazon-ember
        rounded-lg px-4 py-2 bg-blue-50 w-min">
    <p class="truncate">{text}</p>
</label>
<input id="video" type="file" bind:files={video} class="hidden" />