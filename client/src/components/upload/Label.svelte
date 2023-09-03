<script lang="ts">
    export let label: string
</script>

<div class="column">
    <span>{label}</span>
    <slot />
</div>

<style lang="scss">
    .column {
        display: flex;
        flex-direction: column;
        gap: 6px;
    }

    span {
        font-family: "Amazon Ember";
        font-size: 18px;
        padding-left: 16px;
    }

    @font-face {
        font-family: "Amazon Ember";
        src: local("../../../static")
            url("amazon_ember.ttf") format("truetype")
    }
</style>