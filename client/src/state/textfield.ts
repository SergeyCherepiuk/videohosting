import { writable, type Writable } from "svelte/store"

export enum TFState {
    Unfocused,
    Focused,
    Error,
}

export enum TFEvent {
    Focus,
    Unfocus,
    ExceedLimit,
    BackToLimit,
}

function textFieldMachine(state: TFState, event: TFEvent): TFState {
    switch (state) {
        case TFState.Unfocused:
            if (event == TFEvent.Focus) return TFState.Focused
        case TFState.Focused:
            if (event == TFEvent.Unfocus) return TFState.Unfocused
            if (event == TFEvent.ExceedLimit) return TFState.Error
        case TFState.Error:
            if (event == TFEvent.BackToLimit) return TFState.Focused
    }
    return state
}

export function useTextFieldMachine(
    initial: TFState,
): { state: Writable<TFState>, send: (event: TFEvent) => void } {
    let state = writable(initial)

    function send(event: TFEvent) {
        state.update(state => textFieldMachine(state, event))
    }

    return { state, send }
}