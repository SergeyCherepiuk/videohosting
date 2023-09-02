<script lang="ts">
    import { onMount } from 'svelte';

    export let name: string;
    export let text: string;
    export let limit: number = 0;
    export let singleLine: boolean = true;
    export let maxHeight: string = "300px";
    
    onMount(() => {
        let scroll: HTMLDivElement = document.getElementById("scroll") as HTMLDivElement
        scroll.style.maxHeight = maxHeight

        let input: HTMLTextAreaElement = document.getElementById("input") as HTMLTextAreaElement;
        input.oninput = () => {
            if (singleLine) {
                input.value = input.value.replaceAll("\n", "");
            }
            text = input.value;

            input.style.height = "0px";
            input.style.height = (input.scrollHeight) - 24.5 + "px";
        };
    });
</script>

<div class="input-box">
    <div id="scroll" class="scroll">
        <textarea id="input" {name} rows="1" bind:value={text}/> 
    </div>
    {#if limit > 0}
        <span>{text?.length ?? 0}/{limit}</span>
        {#if (text?.length ?? 0) > limit}
            <!-- Change styles -->
        {/if}
    {/if}
</div>

<style lang="scss">
    $line-height: 175%;
    $border-radius: 10px;
    $border-width: 3px;
    $outline-width: 3px;

    $gray-30: #555555;
    $gray-60: #AAAAAA;
    $gray-70: #CCCCCC;

    $blue-50: #1952E5;
    $blue-90: #E2ECFB;

    * {
        font-family: "Amazon Ember";
        font-size: 18px;
    }

    .input-box {
        display: flex;
        flex-direction: row;
        align-items: start;
        border: $border-width $gray-70 solid;
        border-radius: $border-radius;
    }

    .input-box:hover {
        outline: $outline-width $blue-90 solid;
    }

    .input-box:has(textarea:focus) {
        border: $border-width $blue-50 solid;
        outline: $outline-width $blue-90 solid;
    }
    
    .input-box textarea, .input-box span { 
        padding: 12px;
    }

    .scroll {
        flex: 1;
        display: flex;
        overflow: auto;
        direction: rtl;
    }

    textarea {
        all: unset;
        flex-grow: 1;
        overflow-wrap: break-word;
        direction: ltr;
        line-height: $line-height;
    }

    .input-box span {
        color: $gray-70;
        line-height: $line-height;
    }

    ::-webkit-scrollbar {
        width: 12px;
    }

    ::-webkit-scrollbar-track {
        background: $gray-70; 
    }
    
    ::-webkit-scrollbar-thumb {
        background: $gray-60;
    }

    ::-webkit-scrollbar-thumb:hover {
        background: $gray-30; 
    }

    @font-face {
        font-family: "Amazon Ember";
        src: local("../../../static")
            url("amazon_ember.ttf") format("truetype")
    }
</style>