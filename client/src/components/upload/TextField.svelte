<script lang="ts">
    import { onMount } from 'svelte';
    import { TFState, TFEvent, useTextFieldMachine } from '../../state/textfield';

    export let name: string;
    export let text: string;
    export let limit: number | null = null;
    export let singleLine: boolean = true;

    let { state, send } = useTextFieldMachine(TFState.Unfocused);
    function onFocus() { send(TFEvent.Focus) }
    function onBlur() { send(TFEvent.Unfocus) }

    onMount(() => {
        let input = document.getElementById("input") as HTMLTextAreaElement;
        input.oninput = () => {
            if (singleLine) {
                input.value = input.value.replaceAll("\n", "");
            }
            text = input.value;

            if (limit && text.length > limit) {
                send(TFEvent.ExceedLimit);
            } else if (limit) {
                send(TFEvent.BackToLimit);
            }

            input.style.height = "0px";
            input.style.height = (input.scrollHeight) + "px";
        };
    });

    let preloadColors = [
        "border-gray-70", "border-blue-50", "border-red-50",
        "outline-transparent", "outline-blue-90", "outline-red-90",
        "text-red-50", "text-gray-70",
    ]
</script>

<div class="
    flex flex-row items-start
    rounded-xl border-3 border-{
        ($state == TFState.Focused) ? "blue-50"
        : ($state == TFState.Unfocused) ? "gray-70"
        : "red-50"
    }
    outline {
        ($state == TFState.Error) ? "outline-3 outline-red-90"
        : ($state == TFState.Focused) ? "outline-3 outline-blue-90"
        : "outline-transparent hover:outline-3 hover:outline-blue-90"
    }">
    
    <div id="scroll" class="flex flex-1 overflow-auto max-h-80 direction-rtl">
        <textarea
            id="input" {name} rows="1"
            class="
                resize-none outline-none
                flex-grow p-3
                font-amazon-ember text-lg break-words
                rounded-xl direction-ltr"
            on:focus={onFocus} on:blur={onBlur} /> 
    </div>

    {#if limit}
        <span class="font-amazon-ember text-lg p-3 text-{
            ($state == TFState.Error) ? "red-50" : "gray-70"
        }">
            {text?.length ?? 0}/{limit}
        </span>
    {/if}
</div>

<style>
    ::-webkit-scrollbar {
        width: 12px;
    }

    ::-webkit-scrollbar-track {
        background: theme("colors.gray-70"); 
        border-radius: 10px;
    }
    
    ::-webkit-scrollbar-thumb {
        background: theme("colors.gray-60");
        border-radius: 10px;
    }

    ::-webkit-scrollbar-thumb:hover {
        background: theme("colors.gray-30"); 
    }
</style>