<script lang="ts">
    import { circIn, circOut } from "svelte/easing";
    import { fade, fly } from "svelte/transition";
    import Upload from "../../components/upload/Upload.svelte";
    import VideoPicker from "../../components/upload/VideoPicker.svelte";

    let video: FileList
    let clearVideo: Function

    let isWindowOpened: boolean = false;
    $: {
        isWindowOpened = video != undefined
            && video.item(0) != null
            && video?.item(0)?.type.startsWith("video/") == true
    }
</script>

<div class="flex items-center justify-center w-screen h-screen">
    <VideoPicker text="Upload video" bind:video={video} bind:clearVideo={clearVideo} />
    {#if isWindowOpened}
        <div
            class="flex absolute w-full h-full items-center justify-center bg-black bg-opacity-30"
            in:fade={{ duration: 200 }}
            out:fade={{ duration: 200 }}>
            <div 
                class="flex w-full h-full items-center justify-center"
                in:fly={{ y: -100, duration: 150, easing: circOut }}
                out:fly={{ y: -100, duration: 150, easing: circIn }}>
                <Upload
                    bind:video={video} clearVideo={clearVideo}
                    onClose={() => { isWindowOpened = false }} />
            </div>
        </div>
    {/if}
</div>