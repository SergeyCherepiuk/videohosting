<script lang="ts">
    export let name: string;
    export let text: string;
    export let limit: number = 0;
</script>

<div class="input-box">
    <input type="text" name={name} bind:value={text}/>
    {#if limit > 0}
        <span>{text?.length ?? 0}/{limit}</span>
    {/if}
</div>

<style lang="scss">
    $padding: 0.75em;

    $border-width: 0.2em;
    $unfocused-border-color: #CCCCCC;
    $focused-border-color: #1952E5;
    
    $outline-width: 0.2em;
    $outline-color: #E2ECFB;

    .input-box {
        display: flex;
        flex-direction: row;
        border: $border-width $unfocused-border-color solid;
        align-items: start;
        border-radius: 0.5em;
    }

    .input-box:hover {
        outline: $outline-width $outline-color solid;
    }

    .input-box:has(input:focus) {
        border: $border-width $focused-border-color solid;
        outline: $outline-width $outline-color solid;
    }
    
    .input-box * {
        font-size: 1.25em;
        font-family: "Amazon Ember";
        padding: $padding;
    }

    input {
        all: unset;
        flex-grow: 1;
    }

    span {
        color: $unfocused-border-color;
    }

    @font-face {
        font-family: "Amazon Ember";
        src: local("../../../static") url("amazon_ember.ttf") format("truetype")
    }
</style>